package basic

import "time"
import "fmt"

// Comment struct
type Comment struct {
	ID      int
	Message string
}

// Article struct
type Article struct {
	ID        int
	Title     string
	Published time.Time
	Body      string
	Comments  []Comment
}

// Publish date
func (a *Article) Publish() {
	a.Published = time.Now()
}

func createArticle() {
	article := Article{
		1,
		"TDD",
		time.Now(),
		"Test driven development",
		make([]Comment, 0, 3),
	}
	fmt.Printf(article.Title)
}
