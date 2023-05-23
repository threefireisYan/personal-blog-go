package views

import (
	"net/http"
	"personal-blog-go/common"
	"personal-blog-go/service"
)

func (receiver *HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()

	writing.WriteData(w, wr)
}
