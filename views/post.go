package views

import (
	"errors"
	"net/http"
	"personal-blog-go/common"
	"personal-blog-go/service"
	"strconv"
	"strings"
)

func (receiver HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	//	获取路径参数
	//获取文章参数
	path := r.URL.Path

	pIdStr := strings.TrimPrefix(path, "/p/")
	// 获取到 7.html
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错！"))
		return
	}
	detail.WriteData(w, postRes)
}
