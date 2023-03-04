package sentryclient

import (
	"github.com/getsentry/sentry-go"
	"github.com/zhwei820/gconv/deepcopy"
)

func FromSentryEvent(e *sentry.Event) *SendEventReq {
	req := &SendEventReq{}
	_ = deepcopy.CopyExported(req, e)
	return req
}
