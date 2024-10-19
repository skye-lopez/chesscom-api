package chesscom

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
