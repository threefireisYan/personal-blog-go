package dao

import (
	"log"
	"personal-blog-go/models"
)

func GetAllPost(page int, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from test.blog_post limit ?,?", page, pageSize)
	if err != nil {
		log.Println("GetAllCategory查询出错：", err)
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostAll() ([]models.Post, error) {
	rows, err := DB.Query("select * from test.blog_post ")
	if err != nil {
		log.Println("GetAllCategory查询出错：", err)
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CountGetAllPost() int {
	rows := DB.QueryRow("select count(1) from test.blog_post")
	var count int
	_ = rows.Scan(&count)
	return count
}
func CountGetAllPostBySlug(slug string) int {
	rows := DB.QueryRow("select count(1) from test.blog_post where slug = ?", slug)
	var count int
	_ = rows.Scan(&count)
	return count
}

func CountGetAllPostByCategoryId(cId int) int {
	rows := DB.QueryRow("select count(1) from test.blog_post where category_id = ?", cId)
	var count int
	_ = rows.Scan(&count)
	return count
}

func GetPostPageBySlug(page int, pageSize int, slug string) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from test.blog_post where slug = ? limit ?,?", slug, page, pageSize)
	if err != nil {
		log.Println("GetAllCategory查询出错：", err)
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageByCategoryId(cId int, page int, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from test.blog_post where category_id = ? limit ?,?", cId, page, pageSize)
	if err != nil {
		log.Println("GetAllCategory查询出错：", err)
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostById(pId int) (*models.Post, error) {
	rows := DB.QueryRow("select * from blog_post where pid = ?", pId)
	var post = &models.Post{}
	err := rows.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt)
	if err != nil {
		return nil, err
	}

	return post, nil

}

func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into blog_post (title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values (?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt)

	if err != nil {
		log.Println(err)
	}

	pid, err := ret.LastInsertId()
	post.Pid = int(pid)

}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title = ?,content=?,markdown = ?,category_id = ?,type = ?,slug = ?,update_at = ? where pid = ?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid)
	if err != nil {
		log.Println(err)
	}
}

func GetPostSearch(condition string) ([]models.Post, error) {
	rows, err := DB.Query("select * from test.blog_post where title like ?", "%"+condition+"%")
	if err != nil {
		log.Println("GetAllCategory查询出错：", err)
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
