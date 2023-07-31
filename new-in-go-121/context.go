package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// START OMIT
func main() {
	ctx := context.Background()
	// Useful for async work that needs the request values
	justValuesCtx := context.WithoutCancel(ctx)
	_ = justValuesCtx
	causeCtx, _ := context.WithTimeoutCause(
		ctx,
		time.Nanosecond,
		errors.New("request timeout exceeded"),
	)
	time.Sleep(2 * time.Nanosecond)
	// No more "context cancelled" errors without knowing the cause!
	// Note: calling the cancel func does not set the cause.
	fmt.Println(context.Cause(causeCtx))
}

// END OMIT
