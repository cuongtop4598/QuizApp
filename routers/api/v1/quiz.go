package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/cuongtop4598/QuizWithGo/QuizApp/db"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/models"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/pkg/util"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/repository"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

//@Summary Get a single article
//@Produce json
//@Param id path int true "ID"
//@Success 200 {object} app.Response
//@Failure 500 {object} app.Response
//@Router /api/v1/articles/{id} [get]
func GetQuizz(c *gin.Context) {
	//appG := app.Gin{C: c}
	// id := com.StrTo(c.Param("id")).MustInt()
	// valid := validation.Validation{}
	// valid.Min(id, 1, "id")

	// if valid.HasErrors() {
	// 	app.MarkErrors(valid.Errors)
	// 	appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	// 	return
	// }

	if pusher := c.Writer.Pusher(); pusher != nil {
		// Push is supported
		options := &http.PushOptions{
			Header: http.Header{
				"Accept-Encoding": c.Request.Header["Accept-Encoding"],
			},
		}
		err := pusher.Push("/public/img/ic-logo.png", options)
		if err != nil {
			log.Fatal(err)
		}
		resource, err := db.InitResource()
		if err != nil {
			log.Fatal(err)
		}
		quiz := repository.NewQuizEntity(resource)
		data, status, err := quiz.GetListQuestionByID(c.Param("id"))
		if err != nil {
			log.Fatal(err)
		}
		gintemplate.HTML(c, status, "index", gin.H{
			"title":  "Quiz",
			"model":  data,
			"number": 1,
		})
	} else {
		resource, err := db.InitResource()
		if err != nil {
			log.Fatal(err)
		}
		question := repository.NewQuestionEntity(resource)
		id := c.Request.FormValue("id")
		data, status, err := question.GetOneByID(id)
		if err != nil {
			log.Fatal(err)
		}
		gintemplate.HTML(c, status, "index", gin.H{
			"title":  "Quiz",
			"model":  data,
			"number": 1,
		})
		// c.JSON(http.StatusOK, gin.H{
		// 	"status": "not support server push",
		// })

	}
}

func GetQuizByID(c *gin.Context) {
	resource, err := db.InitResource()
	if err != nil {
		log.Fatal(err)
	}
	optionEntity := repository.NewOptionEntity(resource)
	question := repository.NewQuestionEntity(resource)
	lenCollection := question.GetLenght()

	err = c.Request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	if c.Param("id") != "" {
		id, _ := strconv.Atoi(c.Param("id"))
		if int64(id) <= lenCollection {
			options := optionEntity.GetOptionsByIdQuestion(c.Param("id"))
			data, status, err := question.GetOneByID(c.Param("id"))
			if err != nil {
				c.JSON(status, err)
				return
			}
			gintemplate.HTML(c, status, "index", gin.H{
				"title":    "Quiz",
				"question": data,
				"options":  options,
				"end":      false,
			})
		} else {
			gintemplate.HTML(c, 200, "index", gin.H{
				"title":    "Quiz",
				"question": &models.Question{},
				"options":  &models.Option{},
				"end":      true,
			})
		}
	} else {
		idStr := c.Request.FormValue("id")
		idInt, _ := strconv.Atoi(idStr)
		if int64(idInt) > lenCollection {
			gintemplate.HTML(c, 200, "index", gin.H{
				"title":    "Quiz",
				"question": &models.Question{},
				"options":  &models.Option{},
				"end":      true,
			})
		} else {
			options := optionEntity.GetOptionsByIdQuestion(idStr)
			data, status, err := question.GetOneByID(idStr)
			if err != nil {
				c.JSON(status, err)
				return
			}
			gintemplate.HTML(c, status, "index", gin.H{
				"title":    "Quiz",
				"question": data,
				"options":  options,
				"end":      false,
			})
		}

	}

}

func SaveAnswer(c *gin.Context) {
	c.Request.ParseForm()
	var options []string
	// using for get option from checkbox
	for k, v := range c.Request.Form {
		if strings.Contains(k, "answer") {
			i := strings.Split(k, "answer")[1]
			log.Println("options:", i)
			log.Println(v)
			options = append(options, i)
		}

	}
	// using for get option from radio
	if c.Request.FormValue("answer") != "" {
		options = append(options, c.Request.FormValue("answer"))
	}
	answers, err := c.Cookie("answers")
	if err != nil {
		answers = c.Request.FormValue("currentID_s") + strings.Join(options, "")
		c.SetCookie("answers", "", 1000, "", "", false, false)
	}
	answers = answers + ":" + c.Request.FormValue("currentID_s") + strings.Join(options, "")
	c.SetCookie("answers", answers, 1000, "", "", false, false)

	id, err := strconv.Atoi(c.Request.FormValue("currentID_s"))
	if err != nil {

		log.Fatal(err)
	}
	log.Println("id:", id)
	id = id + 1
	id_str := strconv.Itoa(id)
	c.Redirect(302, "/api/v1/quizzes/"+id_str)

}

func GetResult(c *gin.Context) {
	var score = 0
	cookie, err := c.Cookie("answers")
	if err != nil {
		fmt.Println("coookie not found")
	}
	resource, err := db.InitResource()
	if err != nil {
		log.Fatal(err)
	}
	optionEntity := repository.NewOptionEntity(resource)

	question := repository.NewQuestionEntity(resource)

	lenCollection := question.GetLenght()

	options_user := util.GetOptionsFromCookie(cookie)
	fmt.Println(options_user)
	var questionNum = 1
	for questionNum <= int(lenCollection) {
		right := true
		options := optionEntity.GetOptionsByIdQuestion(strconv.Itoa(questionNum))
		fmt.Println(options)
		check := options_user[strconv.Itoa(questionNum)]
		fmt.Println("check:", check)
		if check != nil {
			for _, v := range options {
				id := fmt.Sprintf("%v", v.ID)
				if v.Correct {
					if strings.Contains(strings.Join(check, ""), id) {
						continue
					} else {
						right = false
						break
					}
				} else {
					if strings.Contains(strings.Join(check, ""), id) {
						right = false
						break
					} else {
						continue
					}
				}

			}
			if right {
				score += 1
			}
		}
		questionNum++
	}
	gintemplate.HTML(c, 200, "result", gin.H{
		"score": score,
		"total": lenCollection,
	})
}
