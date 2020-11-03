package dao

import (
	"database/sql"
	"fmt"
	"productManagement/model"
)

func QueryAllScores(db *sql.DB, limit string) ([]model.UserScore, error) {
	var au []model.UserScore

	rows, err := db.Query(fmt.Sprintf("SELECT id, quizz_id, username, score FROM scores LIMIT  %v", limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var s model.UserScore
		err := rows.Scan(&s.Id, &s.QuizId, &s.Score, &s.Username)
		if err != nil {
			return nil, err
		}

		au = append(au, s)
	}

	return au, nil
}

func QueryUserScore(db *sql.DB, username string) (model.UserScore, error) {
	var user model.UserScore

	rows, err := db.Query(fmt.Sprintf("SELECT id, quizz_id, username, score FROM scores WHERE username='%v'", username))
	if err != nil {
		return model.UserScore{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.QuizId, &user.Score, &user.Username)
		if err != nil {
			return model.UserScore{}, err
		}
	}

	return user, nil
}