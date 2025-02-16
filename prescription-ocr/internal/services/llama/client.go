package llama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/common"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

type AnalysisRequest struct {
	Text string `json:"text"`
}

type AnalysisResponse struct {
	Analysis    map[string]interface{} `json:"analysis"`
	Confidence  float64                `json:"confidence"`
	ProcessedAt time.Time              `json:"processed_at"`
}

// NewClient initializes a new LLaMA client
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// AnalyzeMedicationNames extracts medication names from the OCR text
func (c *Client) AnalyzeMedicationNames(text string) ([]string, error) {
	// Use a prompt to extract medication names from the OCR text
	prompt := fmt.Sprintf("Extract all medication names from the following text:\n%s", text)

	response, err := c.AnalyzeText(context.Background(), prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze text: %w", err)
	}

	// Extract medication names from LLaMA response (assuming they are in a list)
	medications := []string{}
	for _, medication := range response.Analysis {
		if medName, ok := medication.(string); ok {
			medications = append(medications, medName)
		}
	}

	return medications, nil
}

// AnalyzeText sends text to LLaMA API for processing
func (c *Client) AnalyzeText(ctx context.Context, text string) (*AnalysisResponse, error) {
	if text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	requestBody, err := json.Marshal(AnalysisRequest{Text: text})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.BaseURL+"/analyze", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result AnalysisResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if result.ProcessedAt.IsZero() {
		result.ProcessedAt = time.Now()
	}

	return &result, nil
}
