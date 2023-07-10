package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	KEY = "trace_id"
)

func NewRequestID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func PrecessEnter(ctx context.Context) {
	PrintLog(ctx, "Go cici")
}

func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
}

func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

func main() {
	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
	PrecessEnter(ctx)
}
