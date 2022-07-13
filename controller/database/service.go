package database

type NewArticle struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	PublishedAt string `json:"published_at" binding:"required"`
	Source      string `json:"source" binding:"required"`
	RawContent  string `json:"content" binding:"required"`
}

func ListArticles() ([]Article, error) {
	var articles []Article
	res := DB.Raw(getArticles).Scan(&articles)
	if res.Error != nil {
		return nil, res.Error
	}
	return articles, nil
}

func GetArticle(id string) (any, error) {
	var article struct {
		Article     string
		WordsInLine string
		CharsInWord string
		PagesNum    string
	}
	res := DB.Raw(getArticleByID, id).Scan(&article)
	if res.Error != nil {
		return "", res.Error
	}
	return article, nil
}

func GetArticleWords(articleID string, wordGroupID string) (any, error) {
	var articleWords []struct {
		Word  string
		Count int
		Index string
	}
	res := DB.Raw(getArticleWords, articleID).Scan(&articleWords)
	if res.Error != nil {
		return nil, res.Error
	}
	return articleWords, nil
}

func CreateArticle(newArticle NewArticle) (any, error) {
	articleToInsert, err := parseArticle(newArticle)
	if err != nil {
		return nil, err
	}

	res := DB.Create(&articleToInsert)
	if res.Error != nil {
		return nil, res.Error
	}
	return articleToInsert, nil
}

// get article with statistics
// chars in word/line/page/article, words in line/page, how many times each word appears
//params: author, publishedAt, title, source

// get list of words in article, position for each word
// params: article, word

// get a list of words in article, position for each word, only if in word group

// get word by position

// get appearances of an expression in an article
