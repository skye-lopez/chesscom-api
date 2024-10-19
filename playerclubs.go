package chesscom

type PlayerClubsResp struct {
	Clubs []PlayerClub `json:"clubs"`
}

type PlayerClub struct {
	Id           string `json:"@id"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	Url          string `json:"url"`
	LastActivity int    `json:"last_activity"`
	Joined       int    `json:"joined"`
}

func PlayerClubs(username string) (*PlayerClubsResp, error) {
	resp := &PlayerClubsResp{}
	err := request("https://api.chess.com/pub/player/"+username+"/clubs", resp)
	return resp, err
}
