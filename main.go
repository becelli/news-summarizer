package main

import (
	"log"
	"news-summarizer/internal/articles"
	"news-summarizer/internal/models"
	"sync"
	"time"
)

func fetchRSSFeeds() []*models.Article {
	rssFeeds := []string{
		"https://blog.becelli.com.br/rss.xml",
	}

	var (
		allArticles []*models.Article
		mu          sync.Mutex
		wg          sync.WaitGroup
	)

	// Fetch articles concurrently
	for i, feedURL := range rssFeeds {
		wg.Add(1)
		go func(i int, feedURL string) {
			defer wg.Done()
			articles, err := articles.FetchArticles(feedURL)
			log.Printf("Fetched %d articles from %s", len(articles), feedURL)
			if err != nil {
				log.Printf("Error fetching articles from %s: %v", feedURL, err)
				return
			}

			// Safely append articles to the shared slice
			mu.Lock()
			allArticles = append(allArticles, articles...)
			mu.Unlock()
		}(i, feedURL)
	}

	// Wait for all Goroutines to finish
	wg.Wait()
	return allArticles
}

func main() {
	allArticles := fetchRSSFeeds()
	log.Printf("Fetched %d articles", len(allArticles))

	// Set up wait group and mutex

	totalSteps := len(allArticles)
	currentStep := 0
	startTime := time.Now()

	// Use goroutines to process articles concurrently
	for _, article := range allArticles {
		// Generate summary
		articles.GenerateSummary(article)
		// Evaluate importance
		articles.EvaluateImportance(article)

		articles.TransalateToPortuguese(article)

		currentStep++
		logProgress(currentStep, totalSteps, startTime)
	}

	// Select top 7 articles
	log.Println("Selecting top articles...")
	topArticles := articles.SelectTopArticles(allArticles, 40, 0)

	// Save to markdown
	log.Println("Saving articles to markdown...")
	err := articles.SaveArticlesToMarkdown(topArticles)
	if err != nil {
		log.Fatalf("Error saving markdown file: %v", err)
	}

	log.Println("News summary saved successfully!")

}

func logProgress(currentStep, totalSteps int, startTime time.Time) {
	percentage := float64(currentStep) / float64(totalSteps) * 100
	elapsed := time.Since(startTime)

	// Use a weighted average of the previous and current estimated time per step
	// This helps to stabilize the ETA calculation early on.
	averageTimePerStep := elapsed / time.Duration(currentStep)
	eta := averageTimePerStep * time.Duration(totalSteps-currentStep)
	log.Printf("%d of %d articles, %.2f%% complete. ETA: %s", currentStep, totalSteps, percentage, eta.Truncate(time.Second))
}
