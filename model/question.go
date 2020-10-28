package model

type Question struct {
	Id string
	Question string
	Responses []string
	CorrectResponse string
}
