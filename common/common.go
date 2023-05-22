package common

import (
	"personal-blog-go/config"
	"personal-blog-go/models"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + `\templates\`)
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}
