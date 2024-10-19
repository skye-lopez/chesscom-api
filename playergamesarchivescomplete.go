package chesscom

type PlayerGamesArchivesCompleteResp struct {
	Games []ArchiveGame `json:"games"`
}

type ArchivePlayer struct {
	Username string `json:"username"`
	Result   string `json:"result"`
	Id       string `json:"@id"`
	Rating   int    `json:"rating"`
}

type ArchiveGame struct {
	Accuracies struct {
		White float64 `json:"white"`
		Black float64 `json:"black"`
	} `json:"accuracies"`
	White       ArchivePlayer `json:"white"`
	Black       ArchivePlayer `json:"black"`
	Url         string        `json:"url"`
	Fen         string        `json:"fen"`
	Pgn         string        `json:"pgn"`
	TimeControl string        `json:"time_control"`
	Rules       string        `json:"rules"`
	Eco         string        `json:"eco"`
	Tournament  string        `json:"tournament"`
	Match       string        `json:"match"`
	StartTime   int           `json:"start_time"`
	EndTime     int           `json:"end_time"`
}

func PlayerGamesArchivesComplete(username string, year string, month string) (*PlayerGamesArchivesCompleteResp, error) {
	resp := &PlayerGamesArchivesCompleteResp{}
	err := request("https://api.chess.com/pub/player/"+username+"/games/"+year+"/"+month, resp)
	return resp, err
}
