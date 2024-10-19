package chesscom

type PlayerMatchesResp struct {
	Finished   []PlayerMatch `json:"finished"`
	InProgress []PlayerMatch `json:"in_progress"`
	Registered []PlayerMatch `json:"registered"`
}

type PlayerMatch struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Id      string `json:"@id"`
	Club    string `json:"club"`
	Results struct {
		PlayedAsWhite string `json:"played_as_white"`
		PlayedAsBlack string `json:"played_as_black"`
	} `json:"results,omitempty"`
	Board string `json:"board"`
}

func PlayerMatches(username string) (*PlayerMatchesResp, error) {
	resp := &PlayerMatchesResp{}
	err := request("https://api.chess.com/pub/player/"+username+"/matches", resp)
	return resp, err
}
