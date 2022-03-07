package wireDemo

import (
	"context"
	"fmt"
)

func ExampleApp() {

	app, err := InitializeService("")
	if err != nil {
		panic(err)
	}

	user, err := app.storage.GetUserByID(context.Background(), 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user.ID, user.Name)

	// Output:
}

func ExampleMock() {

	app, err := InitializeMockService("/test/path")
	if err != nil {
		panic(err)
	}

	user, err := app.storage.GetUserByID(context.Background(), 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user.ID, user.Name)

	// Output:
}
