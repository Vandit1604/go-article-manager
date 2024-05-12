package models

import (
	"database/sql"
	"strings"
)

var DB *sql.DB

type article struct {
	ID      int64  `json:"id"`
	Title   string `json:"string"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

// TODO: Functionality to add articles to database via the app
func RegisterArticle() {
}

func GetAllArticles() ([]article, error) {
	rows, err := DB.Query("SELECT * FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []article

	for rows.Next() {
		var at article

		err := rows.Scan(&at.ID, &at.Title, &at.Content, &at.Date)
		if err != nil {
			return nil, err
		}
		// lil hack to get the date in correct format
		var temp []string = strings.Split(at.Date, "T")
		at.Date = temp[0]

		articles = append(articles, at)
	}

	return articles, nil
}

func GetArticle(ID int64) (article, error) {
	rows, err := DB.Query("SELECT * FROM articles WHERE ID=$1;", ID)
	if err != nil {
		return article{}, err
	}
	defer rows.Close()
	var at article
	for rows.Next() {
		err := rows.Scan(&at.ID, &at.Title, &at.Content, &at.Date)
		if err != nil {
			return article{}, err
		}
	}
	// lil hack to get the date in correct format
	var temp []string = strings.Split(at.Date, "T")
	at.Date = temp[0]

	return at, nil
}
