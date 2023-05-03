package conc

import (
	"context"
	"fmt"
	"testing"
)

func TestMainCtx(t *testing.T) {
	ProcessRequest("jane", "abc123")
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), "userID", userID)
	ctx = context.WithValue(ctx, "authToken", authToken)
	HandleResponse(ctx)
}
func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response for %v (%v)",
		ctx.Value("userID"),
		ctx.Value("authToken"))

}
