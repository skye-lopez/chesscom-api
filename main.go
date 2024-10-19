package chesscom

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func request(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)
}

type PlayerProfileResp struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	Username   string `json:"username"`
	Title      string `json:"title"`
	Status     string `json:"status"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Location   string `json:"location"`
	Country    string `json:"country"`
	TwitchUrl  string `json:"twitch_url"`
	PlayerId   int    `json:"player_id"`
	Joined     int    `json:"joined"`      // TODO: This is a timestamp
	LastOnline int    `json:"last_online"` // TODO: This is a timestamp
	Followers  int    `json:"followers"`
	Fide       int    `json:"fide"`
	IsStreamer bool   `json:"is_streamer"`
}

// Returns a players general player information. See PlayerProfileResp. URL = https://api.chess.com/pub/player/{username}
func PlayerProfile(username string) (*PlayerProfileResp, error) {
	resp := &PlayerProfileResp{}
	err := request("https://api.chess.com/pub/player/"+username, resp)
	return resp, err
}

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
