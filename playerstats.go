package chesscom

type DailySectionStats struct {
	Best struct {
		Game   string `json:"game"`
		Rating int    `json:"rating"`
		Date   int    `json:"date"`
	} `json:"best"`
	Last struct {
		Rating int `json:"rating"`
		Date   int `json:"date"`
		Rd     int `json:"rd"`
	} `json:"last"`
	Record struct {
		Win            int `json:"win"`
		Loss           int `json:"loss"`
		Draw           int `json:"draw"`
		TimePerMove    int `json:"time_per_move"`
		TimeoutPercent int `json:"timeout_percent"`
	} `json:"record"`
	Tournament struct {
		Points        int `json:"points"`
		Withdraw      int `json:"withdraw"`
		Count         int `json:"count"`
		HighestFinish int `json:"highest_finish"`
	} `json:"tournament"`
}

type SectionStats struct {
	Best struct {
		Game   string `json:"game"`
		Rating int    `json:"rating"`
		Date   int    `json:"date"`
	} `json:"best"`
	Last struct {
		Rating int `json:"rating"`
		Date   int `json:"date"`
		Rd     int `json:"rd"`
	} `json:"last"`
	Record struct {
		Win  int `json:"win"`
		Loss int `json:"loss"`
		Draw int `json:"draw"`
	} `json:"record"`
}

type TacticsStats struct {
	Highest struct {
		Rating int `json:"rating"`
		Date   int `json:"date"`
	} `json:"highest"`
	Lowest struct {
		Rating int `json:"rating"`
		Date   int `json:"date"`
	} `json:"lowest"`
}

type PuzzleRushStats struct {
	Best struct {
		TotalAttempts int `json:"total_attempts"`
		Score         int `json:"score"`
	} `json:"best"`
}

type PlayerStatsResp struct {
	ChessDaily    DailySectionStats `json:"chess_daily"`
	Chess960Daily DailySectionStats `json:"chess960_daily"`
	ChessRapid    SectionStats      `json:"chess_rapid"`
	ChessBullet   SectionStats      `json:"chess_bullet"`
	ChessBlitz    SectionStats      `json:"chess_blitz"`
	Tactics       TacticsStats      `json:"tactics"`
	PuzzleRush    PuzzleRushStats   `json:"puzzle_rush"`
	Fide          int               `json:"fide"`
}

func PlayerStats(username string) (*PlayerStatsResp, error) {
	resp := &PlayerStatsResp{}
	err := request("https://api.chess.com/pub/player/"+username+"/stats", resp)
	return resp, err
}
