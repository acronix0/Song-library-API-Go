package songdetailapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/acronix0/song-libary-api/internal/external_api"
	"net/http"
	"time"
)

type APIClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *APIClient) FetchSongDetails(ctx context.Context, group, song string) (*externalapi.SongDetail, error) {
	url := fmt.Sprintf("%s/info?group=%s&song=%s", c.baseURL, group, song)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch song details: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned error: %s", resp.Status)
	}

	var details externalapi.SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	return &details, nil
}
