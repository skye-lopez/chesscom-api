package chesscom

type PlayerTournamentsResp struct {
	Finished   []FinishedTournament   `json:"finished"`
	InProgress []InProgressTournament `json:"in_progress"`
	Registered []InProgressTournament `json:"registered"`
}

type FinishedTournament struct {
	Url           string `json:"url"`
	Id            string `json:"@id"`
	Status        string `json:"status"`
	TotalPlayers  int    `json:"total_players"`
	Wins          int    `json:"wins"`
	Losses        int    `json:"losses"`
	Draws         int    `json:"draws"`
	PointsAwarded int    `json:"points_awarded"`
	Placement     int    `json:"placement"`
}

type InProgressTournament struct {
	Url    string `json:"url"`
	Id     string `json:"@id"`
	Status string `json:"status"`
}

func PlayerTournaments(username string) (*PlayerTournamentsResp, error) {
	resp := &PlayerTournamentsResp{}
	err := request("https://api.chess.com/pub/player/"+username+"/tournaments", resp)
	return resp, err
}
