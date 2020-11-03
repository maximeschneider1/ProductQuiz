package dao

import (
	"database/sql"
	"fmt"
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
			Responses:       []string{"", "", "", ""},
		}
		err := rows.Scan(&q.Id, &q.Question, &q.Responses[0], &q.Responses[1], &q.Responses[2], &q.Responses[3], &q.CorrectResponse)
		if err != nil {
			return nil, err
		}
		aq = append(aq, q)
	}

	return aq, nil
}
