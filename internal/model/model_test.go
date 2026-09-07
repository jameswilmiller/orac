package model

import "testing"

func TestHasFailure(t *testing.T) {

	// 1. make a report card where test "t5" failed with the error boom
	r := Result{Failures: []FailureKey{{Test: "t5", Signature: "boom"}}}

	// 2. Ask: did t5 fail with boom? the answer should be yes
	if r.HasFailure(FailureKey{Test: "t5", Signature: "boom"}) == false {
		t.Error("Expected HasFailure to say yes, but it said no")
	}
}
