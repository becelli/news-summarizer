package articles

import (
	"log"
	"news-summarizer/internal/llm"
	"news-summarizer/internal/models"
)

// GenerateSummary creates a headline and summary for the article in Portuguese using the LLM server.
func GenerateSummary(article *models.Article) {

	request := llm.BuildSummaryRequest(article)
	llmResponse, err := llm.SendRequest(request)
	if err != nil {
		log.Printf("Error generating summary for %s: %v", article.Title, err)
		return
	}

	if len(llmResponse.Choices) > 0 {
		article.NewSummary = llmResponse.Choices[0].Message.Content
	}
}
