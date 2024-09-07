package models

type Article struct {
	Title       string
	Description string
	Link        string
	Published   string

	NewSummary   string
	NewSumamryPT string
	Score        int
}

func (a *Article) hasSummary() bool {
	return a.NewSummary != ""
}

func (a *Article) hasScore() bool {
	return a.Score != 0
}
