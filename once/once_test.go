package once

import (
	"context"
	"fmt"
)

var myData = map[int64]string{}
var myOnce Once

func GetData(ctx context.Context) (map[int64]string, error) {
	if err := myOnce.Do(func() error {
		fmt.Println("Do once")
		myData[1] = "a"
		myData[2] = "b"
		return nil
	}); err != nil {
		return nil, err
	}
	return myData, nil
}
