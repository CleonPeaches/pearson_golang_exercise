package types

type Question struct {
	Text          string   `json:"text"`
	Options       []string `json:"options"`
	CorrectAnswer string   `json:"correctAnswer"`
	Distractors   []string `json:"distractors"`
}

type Quiz struct {
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}
