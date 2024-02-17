package acres

import (
	"bytes"
	"strings"
	"testing"
)

func TestResult(t *testing.T) {
	successMsg := "Done with test instance."
	var b bytes.Buffer
	DefaultOutputSettings.Writer = &b
	r := NewEmitter().Done().Print(successMsg)
	rIsSuccessful := r.Success()
	if !rIsSuccessful {
		t.Fail()
	}

	t.Log("output buffer:", b.String())
	if !strings.Contains(b.String(), successMsg) {
		t.Fail()
	}
}
