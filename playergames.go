package chesscom

type PlayerGamesResp struct {
	Games []PlayerGame `json:"games"`
}

type PlayerGame struct {
	White        string `json:"white"`
	Black        string `json:"black"`
	Url          string `json:"url"`
	Fen          string `json:"fen"`
	Pgn          string `json:"pgn"`
	Turn         string `json:"turn"`
	TimeControl  string `json:"time_control"`
	TimeClass    string `json:"time_class"`
	Rules        string `json:"rules"`
	Tournament   string `json:"tournament"`
	Match        string `json:"match"`
	DrawOffer    string `json:"draw_offer"`
	MoveBy       int    `json:"move_by"`
	LastActivity int    `json:"last_activity"`
	StartTime    int    `json:"start_time"`
}

func PlayerGames(username string) (*PlayerGamesResp, error) {
	resp := &PlayerGamesResp{}
	err := request("https://api.chess.com/pub/player/"+username+"/games", resp)
	return resp, err
}
