package router

import (
	"net/http"
	"personal-blog-go/api"
	"personal-blog-go/views"
)

func Router() {
	//	1.页面	views
	//	2.数据（json）
	//	3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/v1/post", api.API.SvarAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
