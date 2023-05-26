package service

import (
	"log"
	"personal-blog-go/config"
	"personal-blog-go/dao"
	"personal-blog-go/models"
)

func FindPostPigeonhole() (*models.PigenoholeRes, error) {
	//	查询所有文章，进行月份整理
	//	查询所有分类
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}

	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
	}
	return &models.PigenoholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pigeonholeMap,
	}, nil
}
