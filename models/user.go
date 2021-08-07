package models

type User struct {
	ID         uint    `json:"id" bson:"_id"`
	Username   string  `json:"username" bson:"username"`
	Password   string  `json:"password" bson:"password"`
	FullName   string  `json:"full_name" bson:"full_name"`
	Scores     []Score `json:"scores" bson:"scores"`
	TotalScore uint    `json:"total_score" bson:"total_score"`
	Claim      []uint  `json:"claim" bson:"claim"`
}
