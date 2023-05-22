package views

import (
	"errors"
	"log"
	"net/http"
	"personal-blog-go/common"
	"personal-blog-go/service"
)

func (receiver *HTMLApi) Index(w http.ResponseWriter, r *http.Request) {

	//页面上涉及到的所有数据，必须有定义
	info, err := service.GetAllIndexInfo()
	index := common.Template.Index
	if err != nil {
		log.Println("Index获取数据出错", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
	}
	hr := info
	
	index.WriteData(w, hr)

}
