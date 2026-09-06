package model

import (
	"crypto/sha256"
	"encoding/hex"
)

type TestId string
type Order []TestId
type Status string
type FailureKey struct {
	Test   TestId
	Digest string
}
type Result struct {
	Statuses map[TestId]Status
	Failures []FailureKey
}

type Verdict struct {
	Victim          TestId
	Culprits        []TestId
	MinimalOrder    Order
	ConfirmRuns     int
	ConfirmFails    int
	IsolationRuns   int
	IsolationPasses int

	ReproCommand string
}

func (r Result) HasFailure(k FailureKey) bool {
	for _, f := range r.Failures {
		if f.Test == k.Test && (f.Digest == "" || f.Digest == k.Digest) {
			return true
		}
	}
	return false
}

func (o Order) Key() string {
	h := sha256.New()

	for _, id := range o {
		h.Write([]byte(id))
		h.Write([]byte{0})
	}

	return hex.EncodeToString(h.Sum(nil))[:16]
}
