package controllers

import (
	"github.com/robfig/revel"
	"revelapp/app/models"
	"revelapp/app/routes"
)

type Topic struct {
	Application
}

// 帖子列表
func (c *Topic) Index(page int) revel.Result {
	title := "最近发表"

	categories := models.GetCategories()
	topics, pagination := models.GetTopics(page, "", "", "created", routes.Topic.Index(page))

	return c.Render(title, topics, pagination, categories)
}

func (c *Topic) Hot(page int) revel.Result {
	title := "最多点击"

	categories := models.GetCategories()
	topics, pagination := models.GetTopics(page, "", "", "hits", routes.Topic.Hot(page))

	c.Render(title, topics, pagination, categories)
	return c.RenderTemplate("Topic/Index.html")
}

func (c *Topic) Good(page int) revel.Result {
	title := "好帖推荐"

	categories := models.GetCategories()
	topics, pagination := models.GetTopics(page, "good", true, "created", routes.Topic.Good(page))

	c.Render(title, topics, pagination, categories)
	return c.RenderTemplate("Topic/Index.html")
}

// 帖子分类查询，帖子列表按时间排序
func (c *Topic) Category(id int64, page int) revel.Result {
	title := "最近发表"

	categories := models.GetCategories()
	topics, pagination := models.GetTopics(page, "category_id", id, "created", routes.Topic.Category(id, page))

	c.Render(title, topics, pagination, categories)
	return c.RenderTemplate("Topic/Index.html")
}

func (c *Topic) New() revel.Result {
	title := "发表新帖"
	categories := models.GetCategories()

	return c.Render(title, categories)
}

func (c *Topic) NewPost(topic models.Topic) revel.Result {
	topic.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Topic.New())
	}

	topic.UserId = c.RenderArgs["user"].(*models.User).Id
	if topic.Save() {
		c.Flash.Success("发表新帖成功")
	} else {
		c.Flash.Error("发表新帖失败")
	}

	return c.Redirect(routes.Topic.New())
}

// 帖子详细
func (c *Topic) Show(id int64) revel.Result {
	topic := models.FindTopicById(id)

	if topic.Id == 0 {
		return c.NotFound("帖子不存在")
	}

	topic.Hits += 1
	topic.Save()

	replies := models.GetReplies(id)
	categories := models.GetCategories()

	title := topic.Title
	return c.Render(title, topic, replies, categories)
}

// 回复帖子
func (c *Topic) Reply(id int64, reply_content string) revel.Result {
	c.Validation.Required(reply_content).Message("请填写回复内容")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Topic.Show(id))
	}

	reply := new(models.Reply)
	reply.TopicId = id
	reply.UserId = c.RenderArgs["user"].(*models.User).Id
	reply.Content = reply_content

	if !reply.Save() {
		c.Flash.Error("发表回复失败")
	}

	return c.Redirect(routes.Topic.Show(id))
}

func (c *Topic) Edit(id int64) revel.Result {
	title := "编辑帖子"

	topic := models.FindTopicById(id)

	if topic.Id == 0 {
		return c.NotFound("帖子不存在")
	}

	categories := models.GetCategories()

	c.Render(title, topic, categories)
	return c.RenderTemplate("Topic/New.html")
}

func (c *Topic) EditPost(id int64, topic models.Topic) revel.Result {
	topic.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Topic.Edit(id))
	}

	topic_ := models.FindTopicById(id)

	if topic_.Id == 0 {
		return c.NotFound("帖子不存在")
	}

	topic_.Title = topic.Title
	topic_.CategoryId = topic.CategoryId
	topic_.Content = topic.Content

	if topic_.Save() {
		c.Flash.Success("编辑帖子成功")
	} else {
		c.Flash.Error("编辑帖子失败")
	}

	return c.Redirect(routes.Topic.Edit(id))
}
