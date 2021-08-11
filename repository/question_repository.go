package repository

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cuongtop4598/QuizWithGo/QuizApp/db"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var QuestionEntity IQuestion

type IQuestion interface {
	GetAll() ([]models.Question, int, error)
	GetOneByID(id string) (*models.Question, int, error)
	GetOneByObjID(id string) (*models.Question, int, error)
	GetLenght() int64
}

type questionEntity struct {
	resource *db.Resource
	repo     *mongo.Collection
}

func NewQuestionEntity(resource *db.Resource) IQuestion {
	questionRepo := resource.DB.Collection("question")
	QuestionEntity = &questionEntity{resource: resource, repo: questionRepo}
	return QuestionEntity
}

func (entity *questionEntity) GetAll() ([]models.Question, int, error) {
	questionList := []models.Question{}
	ctx, cancel := initContext()
	defer cancel()
	cursor, err := entity.repo.Find(ctx, bson.M{})

	if err != nil {
		return []models.Question{}, 400, err
	}

	for cursor.Next(ctx) {
		var question models.Question
		err = cursor.Decode(&question)
		if err != nil {
			logrus.Print(err)
		}
		questionList = append(questionList, question)
	}
	return questionList, http.StatusOK, nil
}

func (entity *questionEntity) GetOneByObjID(id string) (*models.Question, int, error) {
	var question models.Question
	ctx, cancel := initContext()
	defer cancel()
	logrus.Print(id)
	objID, _ := primitive.ObjectIDFromHex(id)

	err := entity.repo.FindOne(ctx, bson.M{"_id": objID}).Decode(&question)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	return &question, http.StatusOK, nil
}

func (entity *questionEntity) GetOneByID(id string) (*models.Question, int, error) {
	var question models.Question
	ctx, cancel := initContext()
	defer cancel()
	logrus.Print(id)
	i, _ := strconv.Atoi(id)
	err := entity.repo.FindOne(ctx, bson.M{"id": i}).Decode(&question)
	if err != nil {
		log.Fatal(err)
		//return nil, http.StatusNotFound, err
	}
	return &question, http.StatusOK, nil
}

func (entity *questionEntity) GetLenght() int64 {
	ctx, cancel := initContext()
	defer cancel()
	len, err := entity.repo.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return len
}
