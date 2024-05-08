package models

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"string"`
	Content string `json:"content"`
}

var articleList = []article{
	{
		ID:      0,
		Title:   "Effective Go",
		Content: `Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.`,
	}, {
		ID:      1,
		Title:   "The Secret",
		Content: `The Secret’s principles for manifestation – visualization, gratitude, intention, and mastering your thoughts and feelings – allow you to easily use the law of attraction to create anything you desire. You can manifest abundance of every kind, beauty all around you, better health, better relationships, and a life filled with gratitude and happiness. Your manifesting journey begins now, with deciding what you want.`,
	},
}

// Return a list of all the articles
func GetAllArticles() []article {
	return articleList
}

func GetArticle(id int) (article, string) {
	for _, article := range articleList {
		if article.ID == id {
			return article, ""
		}
	}
	return article{}, "No article by that ID"
}
