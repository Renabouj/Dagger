package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		fmt.Printf("Error connecting to Dagger Engine: %s", err)
		os.Exit(1)
	}

	defer client.Close()
}
