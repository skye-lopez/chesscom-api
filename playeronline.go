package chesscom

type PlayerOnlineResp struct {
	Online bool `json:"online"`
}

func PlayerOnline(username string) (*PlayerOnlineResp, error) {
	resp := &PlayerOnlineResp{}
	err := request("https://api.chess.com/pub/player"+username+"/is-online", resp)
	return resp, err
}
