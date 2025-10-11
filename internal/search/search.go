package search

import (
	"fmt"
	"strings"
	"time"

	"govosint/internal/models"
	"govosint/internal/storage"
)

type Service struct {
	store *storage.SQLite
}

func NewService(store *storage.SQLite) *Service {
	return &Service{store: store}
}

// RunQuery is a safe, demo-ready search that generates a few authoritative seeds
// and persists them. Replace/extend with real ingestors later (USAspending, SAM, SEC, LDA).
func (s *Service) RunQuery(query string) ([]models.Result, error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return nil, nil
	}

	seed := []models.Result{
		{
			Query:   q,
			Source:  "USAspending.gov",
			Summary: fmt.Sprintf("Starter link for awards & recipients matching \"%s\"", q),
			URL:     "https://api.usaspending.gov/",
		},
		{
			Query:   q,
			Source:  "SAM.gov",
			Summary: "Entity information (UEI, CAGE) and registration status.",
			URL:     "https://sam.gov/",
		},
		{
			Query:   q,
			Source:  "SEC EDGAR",
			Summary: "Public company filings (if applicable): find CIK and filings.",
			URL:     "https://www.sec.gov/edgar/search/",
		},
		{
			Query:   q,
			Source:  "Senate LDA",
			Summary: "Lobbying disclosures search.",
			URL:     "https://lda.senate.gov/system/public/",
		},
	}

	// Save seeds, then fetch recent results for this query (history).
	for i := range seed {
		_ = s.store.SaveResult(&seed[i])
	}

	// Merge new seeds with the latest from DB for this query
	recent, _ := s.store.GetResults(q, 50)
	all := append(recent, seed...) // recent first is fine for demo
	// Ensure CreatedAt is set for seeds
	now := time.Now()
	for i := range all {
		if all[i].CreatedAt.IsZero() {
			all[i].CreatedAt = now
		}
	}
	return all, nil
}

// Recent returns the latest N results for display on the dashboard.
func (s *Service) Recent(limit int) ([]models.Result, error) {
	return s.store.GetRecent(limit)
}
