package service

import (
	"html/template"
	"log"
	"personal-blog-go/config"
	"personal-blog-go/dao"
	"personal-blog-go/models"
)

func GetAllIndexInfo(page int, pageSize int, slug string) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var total int

	var posts []models.Post
	if slug == "" {
		posts, err = dao.GetAllPost(page, pageSize)
		total = dao.CountGetAllPost()
	} else {
		posts, err = dao.GetPostPageBySlug(page, pageSize, slug)
		total = dao.CountGetAllPostBySlug(slug)
	}

	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	//获取总页数
	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}

	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}
	return hr, nil
}
