package views

import (
	"log"
	"net/http"
	"personal-blog-go/common"
	"personal-blog-go/service"
)

func (receiver *HTMLApi) Pigenohole(w http.ResponseWriter, r *http.Request) {
	//获取首页的渲染
	pigenohole := common.Template.Pigeonhole

	pigenoholeRes, err := service.FindPostPigeonhole()
	if err != nil {
		log.Println(err)
	}
	pigenohole.WriteData(w, pigenoholeRes)
}
