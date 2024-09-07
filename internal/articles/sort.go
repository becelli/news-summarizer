package articles

import (
	"news-summarizer/internal/models"
	"sort"
)

func SelectTopArticles(articles []*models.Article, quantity int, threshold int) []*models.Article {
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Score > articles[j].Score
	})

	var topArticles []*models.Article
	for _, article := range articles {
		if article.Score >= threshold {
			topArticles = append(topArticles, article)
		}
	}

	if len(topArticles) > quantity {
		topArticles = topArticles[:quantity]
	}

	return topArticles
}
