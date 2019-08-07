package model

type AnswerField struct {
	IsCorrect bool `json:"isCorrect"`
	Content string `json:"content"`
}
