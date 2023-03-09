package memory

import (
	"context"
	"fmt"
	go_gin "golang.source-fellows.com/seminar/go-gin/v2"
	"time"
)

type AutoRepository struct {
	autos []go_gin.Auto
}

func (cs *AutoRepository) AddAuto(ctx context.Context, auto go_gin.Auto) error {
	cs.autos = append(cs.autos, auto)
	return nil
}

func (cs *AutoRepository) GetAllAuto(ctx context.Context) ([]go_gin.Auto, error) {
	logger := go_gin.Logger(ctx)
	logger.Info("GetAllAuto aufgerufen")
	fmt.Printf("context:header:x-trace-id:%v", ctx.Value("x-trace-id"))
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("weiter.................")
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return cs.autos, nil
}
