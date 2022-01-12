package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"net/http"
	"toki/code-golang/oauth2/dingtalk"
)

type Config struct {
	ClientID     string `env:"DINGTALK_CLIENT_ID,required"`
	ClientSecret string `env:"DINGTALK_CLIENT_SECRET,required"`
	RedirectURI  string `env:"OAUTH2_REDIRECT_URI" envDefault:"http://localhost:8080/callback"`
}

func main() {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	if cfg.ClientID == "" {
		panic("error")
	}

	dingtalkConn := dingtalk.Open(cfg.ClientID, cfg.ClientSecret, cfg.RedirectURI)

	r := gin.Default()
	r.GET("/dingtalk", func(c *gin.Context) {
		callbackURL, err := dingtalkConn.LoginURL(cfg.RedirectURI, "xx")
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusFound, callbackURL)
	})
	r.GET("/callback", func(c *gin.Context) {
		ident, err := dingtalkConn.HandleCallback(c.Request)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ident)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
