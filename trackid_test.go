package trackid

import (
	"context"
	"testing"
)

func TestGetSetClear(t *testing.T) {
	ctx := context.Background()
	ctx = Set(ctx, "")
	val := Get(ctx)
	if val == "" {
		t.Fatal("fatal 1")
	}
	ctx = Set(ctx, "1")
	val = Get(ctx)
	if val != "1" {
		t.Fatal("fatal 2")
	}
	Clear(ctx)
	val = Get(ctx)
	if val != "" {
		t.Fatal("fatal 3")
	}
	ctx = context.Background()
	val = Get(ctx)
	if val != "" {
		t.Fatal("fatal 4")
	}
}
