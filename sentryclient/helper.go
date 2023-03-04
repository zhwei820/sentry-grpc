package sentryclient

import (
	"github.com/getsentry/sentry-go"
	"github.com/zhwei820/gconv"
	"github.com/zhwei820/gconv/deepcopy"
)

const TraceID = "x_request_id"
const errorVerbose = "errorVerbose"

type ErrorRecordBase struct {
	Msg        string `json:"msg" bson:"msg"`
	XRequestID string `json:"x_request_id" bson:"x_request_id"`

	Detail       string `json:"detail" bson:"detail"`
	ErrorVerbose string `json:"error_verbose" bson:"error_verbose"`

	Biz       string `json:"biz" bson:"biz"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}

/*
业务线： biz
告警时间： 2023-03-04 14:46:53.288 CST
告警消息: biz.failed
告警详情： errorVerbose
x_request_id： pTriEEmtQBgwInaMQHxyyywyCXhVJhZj
*/

func FromSentryEvent(e *sentry.Event, biz string) *SendEventReq {
	return &SendEventReq{
		Msg:          e.Message,
		Detail:       gconv.String(e.Extra),
		XRequestId:   gconv.String(e.Extra[TraceID]),
		ErrorVerbose: gconv.String(e.Extra[errorVerbose]),
		Biz:          biz,
		Timestamp:    e.Timestamp.UnixMilli(),
	}
}

func ToErrorRecordBase(e *SendEventReq) *ErrorRecordBase {
	req := &ErrorRecordBase{}
	_ = deepcopy.CopyExported(req, e)
	return req
}
