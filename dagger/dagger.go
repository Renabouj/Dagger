package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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

	absPath, err := filepath.Abs("../")
	if err != nil {
		fmt.Printf("Error getting the absolute path of path passed in: %s", err)
		os.Exit(1)
	}

	src := client.Host().Directory(absPath)

	golang := client.Container().From("golang:latest")
	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src").WithEnvVariable("CGO_ENABLED", "0")

	path := "build/"
	golang = golang.WithExec([]string{"sh", "-c", "go build -o " + path})
	// _, err = golang.WithExec([]string{"go", "build", "-o", path}).ExitCode(ctx)
	// if err != nil {
	// 	fmt.Printf("Error executing command in container: %s", err)
	// 	os.Exit(1)
	// }

	buildDir := golang.Directory(path)
	// _, err = buildDir.Export(ctx, path)
	// if err != nil {
	// 	fmt.Printf("Error writing content to host: %s", err)
	// 	os.Exit(1)
	// }

	// workDir := client.Host().Directory(".")
	_, err = buildDir.Export(ctx, path)
	if err != nil {
		fmt.Printf("Error writing directory: %s", err)
		os.Exit(1)
	}

	// _, err = client.Container().Build(src).ExitCode(ctx)
	// if err != nil {
	// 	fmt.Printf("Error building docker container: %s", err)
	// 	os.Exit(1)
	// }

}
