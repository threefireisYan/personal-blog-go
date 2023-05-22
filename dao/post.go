package dao

import (
	"log"
	"personal-blog-go/models"
)

func GetAllPost() ([]models.Category, error) {
	rows, err := DB.Query("select * from test.blog_post")
	if err != nil {
		log.Println("GetAllCategory查询出错：", err)
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
