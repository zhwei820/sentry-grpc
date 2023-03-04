package sentryclient

import (
	"fmt"
	"testing"

	"github.com/zhwei820/gconv"
)

func TestFromSentryEvent(t *testing.T) {
	fmt.Println("ErrorRecordBase", gconv.Export(ErrorRecordBase{}))
}
