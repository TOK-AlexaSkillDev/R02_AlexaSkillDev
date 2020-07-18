package main

import (
	"net/http"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
)

type Enemy struct {
	NameId int `form:"name"`
	Winner int
}

var enemy []Enemy

func GetDataSkill(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(2)
	var b Enemy
	c.Bind(&b)
	b.Winner = randInt
	enemy = append(enemy, b)
	c.HTML(http.StatusOK, "skillList.html", map[string]interface{}{
		"enemy": enemy,
	})
}

func GetDataSkillHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", map[string]interface{}{})
}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("../../templates/*.html")
	r.Static("/css", "./../../css")
	r.Static("/images", "./../../images")
	r.GET("/skillHome", GetDataSkillHome)
	r.GET("/snafkinSkill", GetDataSkill)
	r.Run()
}
