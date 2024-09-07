package articles

import (
	"fmt"
	"log"
	"news-summarizer/internal/llm"
	"news-summarizer/internal/models"
)

// EvaluateImportance uses the LLM to score the importance of the article summary.
func EvaluateImportance(article *models.Article) {

	request := llm.BuildImportanceRequest(article)
	llmResponse, err := llm.SendRequest(request)
	if err != nil {
		log.Printf("Error evaluating importance of %s: %v", article.Title, err)
		return
	}

	if len(llmResponse.Choices) > 0 {
		log.Printf("Score for %s: %s", article.Title, llmResponse.Choices[0].Message.Content)
		fmt.Sscanf(llmResponse.Choices[0].Message.Content, "%d", &article.Score)
	}
}
