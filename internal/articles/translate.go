package articles

import (
	"log"
	"news-summarizer/internal/llm"
	"news-summarizer/internal/models"
)

func TransalateToPortuguese(article *models.Article) {

	request := llm.BuildTranslationToBrazilianRequest(article)
	llmResponse, err := llm.SendRequest(request)
	if err != nil {
		log.Printf("Error translating to Portuguese %s: %v", article.Title, err)
		return
	}

	if len(llmResponse.Choices) > 0 {
		article.NewSumamryPT = llmResponse.Choices[0].Message.Content
	}
}
