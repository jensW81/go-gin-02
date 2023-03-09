package main

import (
	"context"
	"fmt"
	go_gin "golang.source-fellows.com/seminar/go-gin/v2"
	"golang.source-fellows.com/seminar/go-gin/v2/http"
	"golang.source-fellows.com/seminar/go-gin/v2/memory"
)

func main() {
	repo := memory.AutoRepository{}
	auto := go_gin.Audi{}
	auto.Kennzeichen = "S-WW 4711"
	ctx := context.Background()
	repo.AddAuto(ctx, auto)
	fmt.Println(http.StartServer(&repo))
}
