package articles

import (
	"log"
	"news-summarizer/internal/models"
	"time"

	"github.com/mmcdole/gofeed"
)

func FetchArticles(feedURL string) ([]*models.Article, error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(feedURL)
	if err != nil {
		return nil, err
	}

	var articles []*models.Article
	for _, item := range feed.Items {
		publishedDate := GetPublishedDate(item)
		isRecent, err := VerifyPublishedDate(publishedDate)
		if err != nil {
			log.Printf("Error verifying published date: %v", err)
			continue
		}

		if !isRecent {
			continue
		}

		// log the article
		articles = append(articles, &models.Article{
			Title:       item.Title,
			Description: item.Description + "\n" + item.Content + "\n",
			Link:        item.Link,
			Published:   publishedDate,
		})
	}

	return articles, nil
}
func GetPublishedDate(item *gofeed.Item) string {
	if item.PublishedParsed != nil {
		return item.PublishedParsed.Format("2006-01-02 15:04:05")
	}

	if item.PublishedParsed == nil && item.UpdatedParsed != nil {
		return item.UpdatedParsed.Format("2006-01-02 15:04:05")
	}

	return ""
}

func VerifyPublishedDate(publishedDate string) (bool, error) {
	currentDate := time.Now()
	// 24 hours
	thresholdDate := currentDate.Add(-24 * time.Hour)

	// parse date and time
	parsedPublishedDate, err := time.Parse("2006-01-02 15:04:05", publishedDate)
	if err != nil {
		log.Printf("Error parsing date: %v", err)
		return false, err
	}

	// check if the article is older than 24 hours
	return parsedPublishedDate.After(thresholdDate), nil
}
