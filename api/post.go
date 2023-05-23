package api

import (
	"errors"
	"net/http"
	"personal-blog-go/common"
	"personal-blog-go/dao"
	"personal-blog-go/models"
	"personal-blog-go/service"
	"personal-blog-go/utils"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}

	post, err := dao.GetPostById(pId)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)

}

func (*Api) SvarAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//	Post 是save操作
	//获取用户id，判断用户是否登陆
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid

	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)

		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		//update
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(float64)
		categoryId := int(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		pidFloat := params["pid"].(float64)
		pid := int(pidFloat)

		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}

}
