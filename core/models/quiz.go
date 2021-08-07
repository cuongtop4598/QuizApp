package models

type Question struct {
	ID             uint     `json:"id" bson:"_id"`
	QuizID         uint     `json:"quizid" bson:"quizid"`
	Question       string   `json:"question" bson:"question"`
	Options        []Option `json:"options" bson:"options"`
	CorrectOptions []uint   `json:"cor_options" bson:"cor_options"`
}

type Option struct {
	ID         uint   `json:"id" bson:"_id"`
	Content    string `json:"content" bson:"content"`
	Correct    bool   `json:"correct" bson:"correct"`
	QuestionID uint   `json:"question_id" bson:"question_id"`
}

type Quiz struct {
	ID         uint       `json:"id" bson:"_id"`
	Name       string     `json:"name" bson:"name"`
	Questions  []Question `json:"questions" bson:"questions"`
	CategoryID uint       `json:"category_id" bson:"category_id"`
	LevelID    uint       `json:"level_id" bson:"level_id"`
}

type Category struct {
	ID      uint   `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Quizzes []Quiz `json:"quizzes" bson:"quizzes"`
}

type Level struct {
	ID    uint   `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Value uint   `json:"value" bson:"value"`
}

type Score struct {
	ID         uint   `json:"id" bson:"_id"`
	UserID     uint   `json:"user_id" bson:"user_id"`
	Username   string `json:"username" bson:"username"`
	CategoryID uint   `json:"category_id" bson:"category_id"`
	QuizID     uint   `json:"quiz_id" bson:"quiz_id"`
	LevelID    uint   `json:"level_id" bson:"level_id"`
	Value      uint   `json:"value" bson:"value"`
}
