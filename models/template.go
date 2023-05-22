package models

import (
	"io"
	"log"
	"text/template"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (t TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func InitTemplate(templateDir string) (HtmlTemplate, error) {

	tp, err := readTemplate(
		[]string{"index", "category", "custom", "detail", "login", "Pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}

func readTemplate(templates []string, tmeplaterDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(view + ".html")
		//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板都进行解析
		home := tmeplaterDir + `home.html`
		header := tmeplaterDir + `layout\header.html`
		footer := tmeplaterDir + `layout\footer.html`
		personal := tmeplaterDir + `layout\personal.html`
		post := tmeplaterDir + `layout\post-list.html`
		pagination := tmeplaterDir + `layout\pagination.html`
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		c, err := t.ParseFiles(tmeplaterDir+viewName, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println("解析模板错误：", err)
			return nil, err
		}
		var tb TemplateBlog
		tb.Template = c
		tbs = append(tbs, tb)
	}
	return tbs, nil
}

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}