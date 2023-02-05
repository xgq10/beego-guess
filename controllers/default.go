package controllers

import (
	"beego-guess/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var subject models.S
	err := func() error {
		id, err := c.GetInt("id")
		if err != nil {
			id = 1
		}
		subject, err = models.GetSubject(id)
		if err != nil {
			return err
		}
		return nil
	}()

	if err != nil {
		c.Ctx.WriteString("wrong params get\n")
	}

	var option map[string]string
	if err = json.Unmarshal([]byte(subject.Option), &option); err != nil {
		c.Ctx.WriteString("Sorry, there is no subject!")
	}
	c.Data["ID"] = subject.Id
	c.Data["Option"] = option
	c.Data["Img"] = "/static" + subject.Img
	c.TplName = "guess.tpl"
}

func (c *MainController) Post() {
	var subject models.S
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println(err)
	}
	subject, _ = models.GetSubject(id)
	answer := c.GetString("key")
	right := models.Answer(subject.Id, answer)

	c.Data["Right"] = right
	c.Data["Next"] = subject.Id + 1
	c.Data["ID"] = subject.Id
	c.TplName = "guess.tpl"
}
