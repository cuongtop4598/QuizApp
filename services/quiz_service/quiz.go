package quiz_service

import "github.com/cuongtop4598/QuizWithGo/QuizApp/models"

type Quiz struct {
	ID         uint
	Name       string
	Questions  []models.Question
	CategoryID uint
	LevelID    uint
}

// func (q *Quiz) Add() error {
// 	quiz := map[string]interface{}{
// 		"name":        q.Name,
// 		"questions":   q.Questions,
// 		"category_id": q.CategoryID,
// 		"level_id":    q.LevelID,
// 	}

// 	if err := models.AddQuiz(quiz); err != nil {
// 		return err
// 	}
// 	return nil
// }
