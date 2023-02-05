package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strings"
)

type S struct {
	Id        int
	Option    string
	AnswerKey string
	Status    int
	Img       string
}

func init() {
	var s *S
	s.TableName()
	orm.RegisterModel(new(S))
}

func GetSubject(id int) (s S, err error) {
	o := orm.NewOrm()
	s = S{Id: id}
	err = o.Read(&s)
	if err != nil {
		return s, err
	}
	return
}

func Answer(sid int, answerKey string) bool {
	subject, err := GetSubject(sid)
	if err != nil {
		return false
	}
	return strings.Compare(strings.ToUpper(answerKey), subject.AnswerKey) == 0
}

func (o *S) TableName() string {
	return "subject"
}
