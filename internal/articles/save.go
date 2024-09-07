package articles

import (
	"fmt"
	"news-summarizer/internal/models"
	"os"
	"time"
)

func SaveArticlesToMarkdown(articles []*models.Article) error {
	filename := fmt.Sprintf("news_summary_%s.md", time.Now().Format("2006_01_02"))
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, article := range articles {
		text := fmt.Sprintf(
			"**Original title:** %s\n\n**Summary to be announced:** %s\n\n**Portuguese::** %s\n**Score:** %d\n\n**Published Date:** %s\n\n**Link:** %s\n\n---\n\n",
			article.Title,
			article.NewSummary,
			article.NewSumamryPT,
			article.Score,
			article.Published,
			article.Link,
		)

		_, err := file.WriteString(text)
		if err != nil {
			return err
		}
	}

	return nil
}
