package chesscom

type PlayerGamesToMoveResp struct {
	Games []PlayerGameToMove `json:"games"`
}

type PlayerGameToMove struct {
	Url          string `json:"url"`
	MoveBy       int    `json:"move_by"`
	DrawOffer    bool   `json:"draw_offer"`
	LastActivity int    `json:"last_activity"`
}

func PlayerGamesToMove(username string) (*PlayerGamesToMoveResp, error) {
	resp := &PlayerGamesToMoveResp{}
	err := request("https://api.chess.com/pub/player/"+username+"/games/to-move", resp)
	return resp, err
}
