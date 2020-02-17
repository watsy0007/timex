package timex

import (
	"testing"
	"time"
)

func TestAtEndOfWeek(t *testing.T) {

	d := time.Date(2020, 2, 17, 0, 0, 0, 0, time.UTC)
	expect := time.Date(2020, 2, 23, 23, 59, 59, NSec, time.UTC)
	actual := AtEndOfWeek(d)
	if actual != expect {
		t.Errorf("the value is %v , expect %v", actual, expect)
	}
}
