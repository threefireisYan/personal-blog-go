package dao

import (
	"log"
	"personal-blog-go/models"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from test.blog_category")
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

func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("select name from test.blog_category where cid = ?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}
