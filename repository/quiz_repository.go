package repository

import (
	"net/http"

	"github.com/cuongtop4598/QuizWithGo/QuizApp/db"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var QuizEntity IQuiz

type quizEntity struct {
	resource *db.Resource
	repo     *mongo.Collection
	question *mongo.Collection
}

type IQuiz interface {
	// CreateOne(quiz form.QuizForm)(models.Quiz, int , error)
	GetListQuestionByID(id string) ([]models.Question, int, error)
}

// func NewQuizEntity
func NewQuizEntity(resource *db.Resource) IQuiz {
	quizRepo := resource.DB.Collection("quizz")
	question := resource.DB.Collection("question")
	QuizEntity = &quizEntity{resource: resource, repo: quizRepo, question: question}
	return QuizEntity
}

func (entity *quizEntity) GetListQuestionByID(id string) ([]models.Question, int, error) {
	var question []models.Question
	ctx, cancel := initContext()
	defer cancel()
	logrus.Print(id)
	objID, _ := primitive.ObjectIDFromHex(id)

	listquest, err := entity.question.Find(ctx, bson.M{"_id": objID})

	listquest.Decode(&question)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return question, http.StatusOK, nil
}
