package views

import (
	"errors"
	"log"
	"net/http"
	"personal-blog-go/common"
	"personal-blog-go/service"
	"strconv"
	"strings"
)

func (receiver *HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category

	//获取文章参数
	path := r.URL.Path

	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别此请求路径"))
		return
	}

	//数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
