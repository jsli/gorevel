package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"time"
)

type Reply struct {
	Id      int64
	TopicId int64
	Topic   *Topic
	UserId  int64
	User    *User
	Content string
	Created time.Time
}

func (r *Reply) Save() bool {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer q.Close()

	_, err = q.Save(r)
	if err != nil {
		fmt.Println(err)
		return false
	}

	topic := new(Topic)
	topic.Id = r.TopicId
	q.Find(topic)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if topic.Id > 0 {
		topic.Replies += 1
		topic.Save()
	}
	return true
}

func GetReplies(id int64) []*Reply {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	defer q.Close()

	var replies []*Reply
	err = q.WhereEqual("topic_id", id).FindAll(&replies)
	if err != nil {
		fmt.Println(err)
	}

	return replies
}
