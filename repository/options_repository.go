package repository

import (
	"log"
	"strconv"

	"github.com/cuongtop4598/QuizWithGo/QuizApp/db"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var OptionsEntity IOption

type IOption interface {
	GetOptionsByIdQuestion(id string) []models.Option
}

type optionEntity struct {
	resource *db.Resource
	repo     *mongo.Collection
}

func NewOptionEntity(resource *db.Resource) IOption {
	optionsRepo := resource.DB.Collection("option")
	OptionsEntity := &optionEntity{resource: resource, repo: optionsRepo}
	return OptionsEntity
}

func (entity *optionEntity) GetOptionsByIdQuestion(id string) []models.Option {
	var options []models.Option
	ctx, cancel := initContext()
	defer cancel()
	logrus.Print(id)
	i, _ := strconv.Atoi(id)
	cursor, err := entity.repo.Find(ctx, bson.M{"question_id": i})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		temp := &models.Option{}
		cursor.Decode(temp)
		options = append(options, *temp)
	}
	return options
}
