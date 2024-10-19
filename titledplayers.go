package chesscom

import (
	"errors"
	"fmt"
)

type TitledPlayersResp struct {
	Players []string `json:"players"`
}

// Returns a list of all titled players usernames. Valid abbrev's = [GM, WGM, IM, WIM, FM, WFM, NM, WNM, CM, WCM]. URL = https://api.chess.com/pub/titled/{title-abbrev}
func TitledPlayers(abbrev string) (*TitledPlayersResp, error) {
	validAbbrevs := map[string]bool{
		"GM":  true,
		"WGM": true,
		"IM":  true,
		"WIM": true,
		"FM":  true,
		"WFM": true,
		"NM":  true,
		"WNM": true,
		"CM":  true,
		"WCM": true,
	}

	resp := &TitledPlayersResp{}

	_, ok := validAbbrevs[abbrev]
	if !ok {
		errMsg := fmt.Sprintf("Provided abbrev: %s is invalid. Valid options are: [GM, WGM, IM, WIM, FM, WFM, NM, WNM, CM, WCM]", abbrev)
		return resp, errors.New(errMsg)
	}

	err := request("https://api.chess.com/pub/titled/"+abbrev, resp)

	return resp, err
}
