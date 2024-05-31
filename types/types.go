package types

type Question struct {
	Text          string   `json:"text"`
	Topic         string   `json:"topic"`
	Options       []string `json:"options"`
	CorrectAnswer string   `json:"correctAnswer"`
	Distractors   []string `json:"distractors"`
}

type Quiz struct {
	Title     string     `json:"text"`
	Questions []Question `json:"questions"`
}
