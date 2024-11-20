package main

import (
	"context"
	"log"

	"github.com/acronix0/song-libary-api/internal/app"
)

func main() {
		ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}