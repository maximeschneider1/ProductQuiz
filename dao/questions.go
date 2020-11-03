package dao

import (
	"database/sql"
	"fmt"
	"math/rand"
	"productManagement/model"
)

func QueryAllQuestion(db *sql.DB, limit string) ([]model.Question, error) {
	var aq []model.Question

	rows, err := db.Query(fmt.Sprintf("SELECT id, question, response_1, response_2, response_3, response_4, correct FROM questions LIMIT %v", limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		q := model.Question{
			Responses:       []model.Response{{}, {}, {}, {}},
		}

		err := rows.Scan(&q.Id, &q.Question, &q.Responses[0].Answer, &q.Responses[1].Answer, &q.Responses[2].Answer, &q.Responses[3].Answer, &q.CorrectResponse)
		if err != nil {
			return nil, err
		}

		for i, response := range q.Responses {
			if response.Answer == q.CorrectResponse {
				q.Responses[i].IsCorrect = true
			} else {
				q.Responses[i].IsCorrect = false
			}
		}

		aq = append(aq, q)
	}

	return aq, nil
}


func QueryRandomQuestions(db *sql.DB) (model.Question, error) {
	q := model.Question{
		Responses:       []model.Response{{}, {}, {}, {}},
	}
	var count int
	rows, err := db.Query("SELECT COUNT(*) from questions")
	if err != nil {
		return model.Question{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return model.Question{}, err
		}
	}

	// generate random number
	random := rand.Intn(count - 1) + 1

	rowsTwo, err := db.Query(fmt.Sprintf("SELECT id, question, response_1, response_2, response_3, response_4, correct FROM questions WHERE id= %v", random))
	if err != nil {
		return model.Question{}, err
	}
	defer rowsTwo.Close()
	for rowsTwo.Next() {
		err := rowsTwo.Scan(&q.Id, &q.Question, &q.Responses[0].Answer, &q.Responses[1].Answer, &q.Responses[2].Answer, &q.Responses[3].Answer, &q.CorrectResponse)
		if err != nil {
			return model.Question{}, err
		}
	}

	return q, nil
}