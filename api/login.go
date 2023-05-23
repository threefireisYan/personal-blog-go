package api

import (
	"net/http"
	"personal-blog-go/common"
	"personal-blog-go/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//	接受用户名与密码
	//	返回对应json数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}

	common.Success(w, loginRes)

}
