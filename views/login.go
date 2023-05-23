package views

import (
	"net/http"
	"personal-blog-go/common"
	"personal-blog-go/config"
)

func (receiver *HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}
