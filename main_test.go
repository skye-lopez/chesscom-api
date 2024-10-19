package chesscom

// NOTE: These will be better written out; for now they just make sure we can parse an endpoint properly.

import (
	"testing"
)

func TestPlayerProfile(t *testing.T) {
	got, err := PlayerProfile("erik")
	if err != nil {
		t.Errorf("PlayerProfile error occured; %e", err)
	}

	t.Log(got)
}
