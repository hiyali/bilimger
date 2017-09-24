package bilimger

import (
	"fmt"
	"testing"
)

func TestVoice(t *testing.T) {
	v := fmt.Sprintf("city No.%d", 5)
	expectedVal := "city No.5"
	if v != expectedVal {
		t.Error("Expected: "+expectedVal+", but got", v)
	}
}
