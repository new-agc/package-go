package agcmsg

type AnswerForm struct {
	// Equation string
	Answer string
}

type ScoreLogForm struct {
	No     int     `json:"no"`
	Weight int     `json:"weight"`
	Status string  `json:"status"`
	Score  float64 `json:"score"`
}

type ScoringData struct {
	Id              string  `json:"id"`
	Status          string  `json:"status"`
	Score           float64 `json:"score"`
	ScoreLog        string  `json:"score_log,omitempty"`
}

type AGCTestMessage struct {
	TeamName		string			`json:"team"`
	AnswerSheet		[]string		`json:"answer_sheet"`
	Status			string			`json:"status"`
}

type CommonAnswerMessage struct {
	TeamName		string			`json:"team"`
	Secret			string			`json:"secret"`
	Track			string			`json:"track"`
	AnswerSheet		[]string		`json:"answer_sheet"`
}