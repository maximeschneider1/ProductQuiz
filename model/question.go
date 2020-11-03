package model

type Question struct {
	Id string
	Question string
	Responses []Response
	CorrectResponse string
}

type Response struct {
	Answer string
	IsCorrect bool
}
