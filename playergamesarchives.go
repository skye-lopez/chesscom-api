package chesscom

type PlayerGamesArchivesResp struct {
	Archives []string `json:"archives"`
}

func PlayerGamesArchives(username string) (*PlayerGamesArchivesResp, error) {
	resp := &PlayerGamesArchivesResp{}
	err := request("https://api.chess.com/pub/player/"+username+"/games/archives", resp)
	return resp, err
}
