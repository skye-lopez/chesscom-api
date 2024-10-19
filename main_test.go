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

func TestPlayerStats(t *testing.T) {
	got, err := PlayerStats("erik")
	if err != nil {
		t.Errorf("PlayerStats error occured; %e", err)
	}

	t.Log(got)
}

func TestPlayerOnline(t *testing.T) {
	got, err := PlayerOnline("erik")
	if err != nil {
		t.Errorf("PlayerOnline error occured; %e", err)
	}

	t.Log(got)
}

func TestPlayerGames(t *testing.T) {
	got, err := PlayerGames("erik")
	if err != nil {
		t.Errorf("PlayerGames error occured; %e", err)
	}

	if len(got.Games) <= 0 {
		t.Errorf("PlayerGames error occured; no games where returned for erik")
	}

	t.Log(got)
}

func TestPlayerGamesToMove(t *testing.T) {
	got, err := PlayerGamesToMove("erik")
	if err != nil {
		t.Errorf("PlayerGamesToMove error occured; %e", err)
	}

	if len(got.Games) <= 0 {
		t.Errorf("PlayerGamesToMove error occured; no games where returned for erik")
	}

	t.Log(got)
}
