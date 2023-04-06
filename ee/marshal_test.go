package ee

import (
	"encoding/json"
	"fmt"
	"github.com/ImSingee/tt"
	"strconv"
	"testing"
)

func TestMarshalJson(t *testing.T) {
	errs := []error{
		New("err1"),                   // fundamental
		Errorf("err2"),                // fundamental
		WithStack(fmt.Errorf("err3")), // withStack
		WithMessage(fmt.Errorf("underlying"), "err4"), // withMessage
		Wrap(fmt.Errorf("underlying"), "err5"),        // withStack
	}

	for _, err := range errs {
		t.Run("marshal "+err.Error(), func(t *testing.T) {
			p, e := json.Marshal(err)
			tt.AssertIsNotError(t, e)

			tt.AssertEqual(t, strconv.Quote(err.Error()), string(p))
		})
	}
}
