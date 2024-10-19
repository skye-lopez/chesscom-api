package chesscom

// NOTE: These will be better written out; for now they just make sure we can parse an endpoint properly.

import (
	"testing"
)

func TestPlayerProfile(t *testing.T) {
	_, err := PlayerProfile("erik")
	if err != nil {
		t.Errorf("PlayerProfile error occured; %e", err)
	}
}

func TestTitledPlayers(t *testing.T) {
	got, err := TitledPlayers("GM")
	if err != nil {
		t.Errorf("TitledPlayers error occured; %e", err)
	}

	if len(got.Players) == 0 {
		t.Errorf("TitledPlayers error occured; no players were returned in the GM section")
	}
}

func TestPlayerStats(t *testing.T) {
	_, err := PlayerStats("erik")
	if err != nil {
		t.Errorf("PlayerStats error occured; %e", err)
	}
}

func TestPlayerOnline(t *testing.T) {
	_, err := PlayerOnline("erik")
	if err != nil {
		t.Errorf("PlayerOnline error occured; %e", err)
	}
}

func TestPlayerGames(t *testing.T) {
	got, err := PlayerGames("erik")
	if err != nil {
		t.Errorf("PlayerGames error occured; %e", err)
	}

	if len(got.Games) <= 0 {
		t.Errorf("PlayerGames error occured; no games were returned for erik")
	}
}

func TestPlayerGamesToMove(t *testing.T) {
	got, err := PlayerGamesToMove("erik")
	if err != nil {
		t.Errorf("PlayerGamesToMove error occured; %e", err)
	}

	if len(got.Games) <= 0 {
		t.Errorf("PlayerGamesToMove error occured; no games were returned for erik")
	}
}

func TestPlayerGamesArchives(t *testing.T) {
	got, err := PlayerGamesArchives("erik")
	if err != nil {
		t.Errorf("PlayerGamesArchives error occured; %e", err)
	}

	if len(got.Archives) <= 0 {
		t.Errorf("PlayerGamesArchives error occured; no games were returned for erik")
	}
}

func TestPlayerGamesArchivesComplete(t *testing.T) {
	got, err := PlayerGamesArchivesComplete("erik", "2009", "10")
	if err != nil {
		t.Errorf("PlayerGamesArchivesComplete error occured; %e", err)
	}

	if len(got.Games) <= 0 {
		t.Errorf("PlayerGamesArchivesComplete error occured; no games were returned for erik")
	}
}

func TestPlayerClubs(t *testing.T) {
	got, err := PlayerClubs("erik")
	if err != nil {
		t.Errorf("PlayerClubs error occured; %e", err)
	}

	if len(got.Clubs) <= 0 {
		t.Errorf("PlayerClubs error occured; no clubs were returned for erik")
	}
}

func TestPlayerMatches(t *testing.T) {
	got, err := PlayerMatches("erik")
	if err != nil {
		t.Errorf("PlayerMatches error occured; %e", err)
	}

	if len(got.Finished) <= 0 {
		t.Errorf("PlayerMatches error occured; no matches were returned for erik")
	}
}

func TestPlayerTournaments(t *testing.T) {
	got, err := PlayerTournaments("erik")
	if err != nil {
		t.Errorf("PlayerTournaments error occured; %e", err)
	}

	if len(got.Finished) <= 0 {
		t.Errorf("PlayerTournaments error occured; no tournaments were returned for erik")
	}
}
