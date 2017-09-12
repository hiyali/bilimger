package main

import (
  "testing"
  "github.com/hiyali/bilimger/world"
)

func TestVoice(t *testing.T) {
  v := world.Voice()
  expectedVal := "city No.5"
	if v != expectedVal {
    t.Error("Expected: "+ expectedVal +", but got", v)
	}
}

