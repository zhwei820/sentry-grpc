package sentryclient

import (
	context "context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewSentryCli(ctx context.Context, sentryAddr string) SentryClient {
	conn, err := grpc.Dial(sentryAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return NewSentryClient(conn)
}
