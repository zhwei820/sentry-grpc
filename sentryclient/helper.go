package sentryclient

import (
	"github.com/getsentry/sentry-go"
	"github.com/zhwei820/gconv"
	"github.com/zhwei820/gconv/deepcopy"
)

const TraceID = "x-request-id"

type ErrorRecordBase struct {
	Msg        string `json:"msg" bson:"msg"`
	XRequestID string `json:"x_request_id" bson:"x_request_id"`

	Detail string `json:"detail" bson:"detail"`

	Collection string `json:"collection" bson:"collection"`
}

func FromSentryEvent(e *sentry.Event, biz string) *SendEventReq {
	return &SendEventReq{
		Msg:        e.Message,
		Detail:     gconv.String(e.Extra),
		XRequestId: gconv.String(e.Extra[TraceID]),
		Collection: biz,
	}
}

func ToErrorRecordBase(e *SendEventReq) *ErrorRecordBase {
	req := &ErrorRecordBase{}
	_ = deepcopy.CopyExported(req, e)
	return req
}
