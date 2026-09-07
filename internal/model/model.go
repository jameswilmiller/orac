package model

import (
	"crypto/sha256"
	"encoding/hex"
)

type TestID string
type Order []TestID
type Status string

const (
	StatusPass  Status = "pass"
	StatusFail  Status = "fail"
	StatusSkip  Status = "skip"
	StatusError Status = "error"
)

type FailureKey struct {
	Test      TestID
	Signature string
}
type Result struct {
	Statuses map[TestID]Status
	Failures []FailureKey
}

type Verdict struct {
	Victim          TestID
	Culprits        []TestID
	MinimalOrder    Order
	ConfirmRuns     int
	ConfirmFails    int
	IsolationRuns   int
	IsolationPasses int

	ReproCommand string
}

// HasFailure reports whether the given FailureKey exists within the results failures
//
// If the provided FailureKey has an empty Signature, it returns true if any failure
// matches the key's Test ID, ignoring the signature check
func (r Result) HasFailure(k FailureKey) bool {
	for _, f := range r.Failures {
		if f.Test == k.Test && (k.Signature == "" || f.Signature == k.Signature) {
			return true
		}
	}
	return false
}

// Key returns a deterministic 16 character hexadecimal hash representing the Order.
//
// It computes a SHA-256 hash across all IDs in the order, so we have fixed size keys for large orders
// separates each ID with a null byte to prevent collisions, and returns the first 16 characters.
func (o Order) Key() string {
	h := sha256.New()

	for _, id := range o {
		h.Write([]byte(id))
		h.Write([]byte{0}) // Prevents ["a", "bc"] and ["ab", "c"] from producing identical hashes
	}

	return hex.EncodeToString(h.Sum(nil))[:16]
}
