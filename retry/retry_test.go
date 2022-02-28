package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func TestBackoff(t *testing.T) {
	retryTimes := 3
	retryInterval := time.Second

	tryErr := 0

	do := func() error {
		if tryErr != 3 {
			tryErr++
			fmt.Println("error...")
			return errors.New("error")
		}
		fmt.Println("retry...")
		return nil
	}

	boff := backoff.NewExponentialBackOff()
	boff.InitialInterval = retryInterval
	err := backoff.Retry(do, backoff.WithMaxRetries(boff, uint64(retryTimes)))
	if err != nil {
		panic(err)
	}
}
