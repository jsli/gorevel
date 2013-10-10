package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"github.com/robfig/revel"
	"strings"
	"time"
)

type Topic struct {
	Id         int64
	Title      string `qbs:"size:255"`
	Content    string
	CategoryId int64
	Category   *Category
	UserId     int64
	User       *User
	Hits       int
	Replies    int
	Good       bool
	Created    time.Time
	Updated    time.Time
}

func (topic *Topic) Validate(v *revel.Validation) {
	v.Required(topic.Title).Message("请输入标题")
	v.MaxSize(topic.Title, 105).Message("最多35个字")
	v.Required(topic.Category).Message("请选择分类")
	v.Required(topic.Content).Message("帖子内容不能为空")
}

func (t *Topic) Save() bool {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer q.Close()

	_, err = q.Save(t)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func FindTopicById(id int64) *Topic {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	topic := new(Topic)
	err = q.WhereEqual("topic.id", id).Find(topic)

	if err != nil {
		fmt.Println(err)
	}

	return topic
}

func GetTopics(page int, where string, value interface{}, order string, url string) ([]*Topic, *Pagination) {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	page -= 1
	if page < 0 {
		page = 0
	}

	var topics []*Topic
	var total int64
	if where == "" {
		total = q.Count("topic")
		err = q.OmitFields("Content").OrderByDesc(order).
			Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&topics)
	} else {
		total = q.WhereEqual(where, value).Count("topic")
		err = q.WhereEqual(where, value).
			OmitFields("Content").OrderByDesc(order).
			Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&topics)
	}

	if err != nil {
		fmt.Println(err)
	}

	url = url[:strings.Index(url, "=")+1]
	pagination := NewPagination(page, int(total), url)

	return topics, pagination
}
