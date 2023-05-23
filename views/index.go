package views

import (
	"errors"
	"log"
	"net/http"
	"personal-blog-go/common"
	"personal-blog-go/service"
	"strconv"
)

func (receiver *HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	//获取首页的渲染
	index := common.Template.Index

	//数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10

	//页面上涉及到的所有数据，必须有定义
	info, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
	}

	hr := info

	index.WriteData(w, hr)

}
