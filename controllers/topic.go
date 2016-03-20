package controllers

import (
	"fmt"
	"strconv"

	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
	// "github.com/deepzz0/go-common/log"
)

const (
	DUOSHUO_COMMENT_NUM_URL = "http://api.duoshuo.com/threads/counts.json?short_name=%s&threads=%d"
)

type TopicController struct {
	Common
}

func (this *TopicController) Get() {
	this.Layout = "homelayout.html"
	this.TplName = "topicTemplate.html"
	this.Leftbar("")
	this.Topic()
}

func (this *TopicController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	resp.Data = "文章索引错误."
	id := this.Ctx.Input.Param(":id")
	ID, err := strconv.Atoi(id)
	if err == nil {
		topic := models.TMgr.GetTopic(int32(ID))
		if topic != nil {
			resp.Data = string(topic.Content)
		}
	}
}

func (this *TopicController) Topic() {
	id := this.Ctx.Input.Param(":id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		this.Data["IsFalse"] = true
		return
	}
	topic := models.TMgr.GetTopic(int32(ID))
	if topic == nil {
		this.Data["IsFalse"] = true
		return
	}
	this.Data["IsFalse"] = false
	this.Data["Title"] = topic.Title + " - " + models.Blogger.BlogName
	this.Data["Url"] = fmt.Sprintf("%s/%s/%d.html", this.domain, topic.CreateTime.Format(helper.Layout_y_m_d), topic.ID)
	this.Data["Title"] = topic.Title
	this.Data["Time"] = topic.CreateTime.Format(helper.Layout_y_m_d2)
	this.Data["PCategory"] = topic.PCategory
	this.Data["PTags"] = topic.PTags
	this.Data["ID"] = fmt.Sprint(topic.ID)
	this.Data["Content"] = string(topic.Content)
}
