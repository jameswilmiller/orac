package model

import "testing"

func TestHasFailure(t *testing.T) {
	stored := Result{Failures: []FailureKey{{Test: "t5", Signature: "boom"}}}
	// 1. make a report card where test "t5" failed with the error boom
	cases := []struct {
		name  string
		query FailureKey
		want  bool
	}{
		{"exact match", FailureKey{Test: "t5", Signature: "boom"}, true},
		{"empty signature is wildcard", FailureKey{Test: "t5"}, true},
		{"wrong signature", FailureKey{Test: "t5", Signature: "other"}, false},
		{"test didn't fail", FailureKey{Test: "t1"}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := stored.HasFailure(tc.query); got != tc.want {
				t.Errorf("HasFailure(%+v) = %v; want %v", tc.query, got, tc.want)
			}
		})
	}
}

func TestKeySameOrderSameKey(t *testing.T) {
	a := Order{"t1", "t2", "t3"}
	b := Order{"t1", "t2", "t3"}

	if a.Key() != b.Key() {
		t.Errorf("same order gave different keys: %s != %s ", a.Key(), b.Key())
	}
}

func TestKeyDifferentOrderDifferentKey(t *testing.T) {
	a := Order{"t1", "t2", "t3"}
	b := Order{"t1", "t3", "t2"}

	if a.Key() == b.Key() {
		t.Errorf("different order gave same keys: %s != %s ", a.Key(), b.Key())
	}
}

func TestKeySeparataborPreventsCollision(t *testing.T) {
	a := Order{"ab", "c"}
	b := Order{"a", "bc"}

	if a.Key() == b.Key() {
		t.Errorf("different order gave same keys: %s != %s ", a.Key(), b.Key())
	}
}
