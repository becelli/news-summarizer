package llm

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"news-summarizer/internal/models"
)

// SendRequest sends an LLM request and returns the response.
func SendRequest(request models.LLMRequest) (*models.LLMResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:4000/v1/chat/completions", "application/json", io.NopCloser(bytes.NewBuffer(reqBody)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var llmResponse models.LLMResponse
	err = json.NewDecoder(resp.Body).Decode(&llmResponse)
	if err != nil {
		return nil, err
	}

	return &llmResponse, nil
}
