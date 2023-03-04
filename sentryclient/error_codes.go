package sentryclient

import (
	"fmt"

	"gitee.com/spwx/errors"
)

const baseCode = 12110100
const (
	//  ErrCodeCantCancelAssetOrder cant cancel asset order
	ErrCodeCantCancelAssetOrder = baseCode + iota + 1
	ErrCodeClearlyCheckSettlementFailed
	// transfer liability order can be roughly split into 3 parts.
	// 1. transfer amount from transferee to sentry, and freeze the amount.
	// 2. create new order.
	// 3. settle old order, payback to transferer.
	// when we have problem in (2,3) and succeed in (1), we tend to retry on failure steps, until it all become success.
	ErrCodeNeedRetryInTransferLiabilityOrder
)

func RegisterTranslate() {
	errors.RegisterTranslate([]errors.TransInfo{
		{
			Tag: errors.EnUs,
			Key: fmt.Sprintf("%d", ErrCodeCantCancelAssetOrder),
			Msg: "can't cancel asset order",
		},
		{
			Tag: errors.EnUs,
			Key: fmt.Sprintf("%d", ErrCodeClearlyCheckSettlementFailed),
			Msg: "clearly check settlement failed",
		},
	})
}
