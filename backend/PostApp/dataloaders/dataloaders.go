package handlers

import (
	"database/sql"

	"PostApp/models"
)

func LoadPostsFromPostgreSQL(db *sql.DB) ([]models.Post, error) {
	var posts []models.Post

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.ImageURI, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
