package model

type Answer struct {
	Type string            `json:"type"`
	Choices []*AnswerField `json:"choices"`
}
