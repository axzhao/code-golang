package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"

	"go.uber.org/zap"
)

const (
	apiURL   = "https://api.dingtalk.com"
	authURL  = "https://login.dingtalk.com/oauth2/auth"
	tokenURL = "https://api.dingtalk.com/v1.0/oauth2/userAccessToken"
)

type Identity struct {
	UserID   string
	Username string
	Email    string
	Mobile   string
}

type Config struct {
	ClientID     string `env:"DINGTALK_CLIENT_ID"`
	ClientSecret string `env:"DINGTALK_CLIENT_SECRET"`
	RedirectURI  string `env:"OAUTH2_REDIRECT_URI" envDefault:"http://localhost:8080/callback"`
}

func Open(cID, cSecret, cRedirectURI string) *Connector {
	return &Connector{
		oauth2Config: &oauth2.Config{
			ClientID:     cID,
			ClientSecret: cSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  authURL,
				TokenURL: tokenURL,
			},
			Scopes:      []string{"openid"},
			RedirectURL: cRedirectURI,
		},
		logger: zap.NewExample(),
	}
}

type Connector struct {
	oauth2Config *oauth2.Config
	logger       *zap.Logger
}

func (c *Connector) LoginURL(callbackURL, state string) (string, error) {
	if c.oauth2Config.RedirectURL != callbackURL {
		return "", fmt.Errorf("expected callback URL %q did not match the URL in the config %q",
			callbackURL, c.oauth2Config.RedirectURL)
	}
	return c.oauth2Config.AuthCodeURL(state, oauth2.ApprovalForce), nil
}

type tokenPayload struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Code         string `json:"code"`
	GrantType    string `json:"grantType"`
}

type connectorData struct {
	AccessToken string `json:"accessToken"`
}

func (c *Connector) HandleCallback(r *http.Request) (identity Identity, err error) {
	q := r.URL.Query()
	if errMsg := q.Get("error"); errMsg != "" {
		return identity, errors.New(errMsg)
	}

	//token, err := c.oauth2Config.Exchange(ctx, q.Get("authCode"))
	//if err != nil {
	//	return identity, fmt.Errorf("dingtalk: get token: %v", err)
	//}
	//client := c.oauth2Config.Client(ctx, token)

	ctx := r.Context()
	var payload bytes.Buffer
	err = json.NewEncoder(&payload).Encode(&tokenPayload{
		ClientID:     c.oauth2Config.ClientID,
		ClientSecret: c.oauth2Config.ClientSecret,
		Code:         q.Get("authCode"),
		GrantType:    "authorization_code",
	})
	if err != nil {
		return identity, err
	}
	req, err := http.NewRequest("POST", tokenURL, &payload)
	if err != nil {
		return identity, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return identity, err
	}
	if resp.StatusCode != http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return identity, err
		}
		return identity, errors.New(fmt.Sprintf("status: %d, resp: %s", resp.StatusCode, string(data)))
	}
	var data connectorData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return identity, err
	}
	client := c.oauth2Config.Client(ctx, &oauth2.Token{AccessToken: data.AccessToken})
	p, err := c.profile(ctx, client, data.AccessToken)
	if err != nil {
		return identity, fmt.Errorf("dingtalk: get profile: %v", err)
	}

	identity.Username = p.Name
	identity.Email = p.Email
	identity.Mobile = fmt.Sprintf(`+%s%s`, p.StateCode, p.Mobile)
	identity.UserID = p.UnionID

	return identity, nil
}

type profile struct {
	UnionID   string `json:"unionId"`
	Name      string `json:"nick"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	StateCode string `json:"stateCode"`
}

func (c *Connector) profile(ctx context.Context, client *http.Client, accessToken string) (profile, error) {
	req, err := http.NewRequest("GET", apiURL+"/v1.0/contact/users/me", nil)
	if err != nil {
		return profile{}, fmt.Errorf("new req: %v", err)
	}
	req.Header.Set("x-acs-dingtalk-access-token", accessToken)

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return profile{}, fmt.Errorf("get URL %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return profile{}, fmt.Errorf("read body: %v", err)
		}
		return profile{}, fmt.Errorf("%s: %s", resp.Status, body)
	}

	var d profile
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return profile{}, fmt.Errorf("JSON decode: %v", err)
	}

	return d, err
}
