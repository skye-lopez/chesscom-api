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

func TestTitledPlayers(t *testing.T) {
	got, err := TitledPlayers("GM")
	if err != nil {
		t.Errorf("TitledPlayers error occured; %e", err)
	}

	if len(got.Players) == 0 {
		t.Errorf("TitledPlayers error occured; no players where returned in the GM section")
	}

	t.Log(got)
}
