package sentryclient

import (
	context "context"

	"github.com/zhwei820/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewSentryCli(ctx context.Context, sentryAddr string) SentryClient {
	conn, err := grpc.Dial(sentryAddr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(log.UnaryClientLogTraceInterceptorWrapper(log.GetComponentName())))
	if err != nil {
		panic(err)
	}
	return NewSentryClient(conn)
}
