package llm

import (
	"news-summarizer/internal/models"
)

func BuildSummaryRequest(article *models.Article) models.LLMRequest {
	return models.LLMRequest{
		Model: "lmstudio-community/gemma-2-9b-it-GGUF",
		Messages: []models.LLMMessage{
			{
				Role: "system",
				Content: `You're an AI news specialist for a podcast targeting tech professionals and developers. Your task is to provide a title and a summary that should be concise, engaging, and informative, capturing the main points of the article.

Format:
Title: A brief, engaging title that captures the article's essence.
Summary: A 1-2 paragraph summary narrating the article’s main points, clear and accessible.

Example of a response:
Title: Hacker Creates Fake Death Record to Evade Child Support
Summary: A hacker manipulated state death records to falsely declare his own death using a forged certification. The scheme, which involved system damage and unpaid child support, led to financial losses exceeding $190,000. The hacker was ultimately sentenced to 6 years and 9 months in prison.

You always provide a title and summary for the articles.`,
			},
			{
				Role:    "user",
				Content: "Title: " + article.Title + "\n\nContent: " + article.Description,
			},
		},
		Temperature: 0.6,
		MaxTokens:   300,
		Stream:      false,
	}
}

func BuildImportanceRequest(article *models.Article) models.LLMRequest {
	return models.LLMRequest{
		Model: "lmstudio-community/gemma-2-9b-it-GGUF",
		Messages: []models.LLMMessage{
			{
				Role:    "system",
				Content: `You are an expert in evaluating news content for a podcast targeting technology professionals and software developers. You always reply with a number only, no explanations. Your task is to assess the significance of news summaries on a scale from 0 to 100, focusing on relevance, quality, content type, and impact. Consider the importance of the news to the target audience, giving higher scores to summaries that are highly relevant to current tech trends, industry challenges, or significant advancements. Evaluate the clarity, informativeness, and engagement level of the summary, penalizing those that are vague, unclear, or fail to captivate the reader. Assign lower scores to content that is not a genuine news article (e.g., blog posts, opinion pieces, sponsored content) or lacks valuable information.`,
			},
			{
				Role:    "user",
				Content: "Summary: " + article.NewSummary + "\n\nOriginal Content: " + article.Description,
			},
		},
		Temperature: 0.1,
		MaxTokens:   10,
		Stream:      false,
	}
}

func BuildTranslationToBrazilianRequest(article *models.Article) models.LLMRequest {
	return models.LLMRequest{
		Model: "lmstudio-community/gemma-2-9b-it-GGUF",
		Messages: []models.LLMMessage{
			{
				Role:    "system",
				Content: `You are an AI translation specialist for a news podcast aimed at the Brazilian audience. Your task is to translate the following news summary into Brazilian Portuguese. The translation should maintain the original meaning and tone of the summary while ensuring that the language is appropriate for a Brazilian audience. If the summary contains cultural references or idiomatic expressions, adapt or explain them as needed.`,
			},
			{
				Role:    "user",
				Content: "Summary: " + article.NewSummary,
			},
		},
		Temperature: 0.3,
		MaxTokens:   400,
		Stream:      false,
	}
}
