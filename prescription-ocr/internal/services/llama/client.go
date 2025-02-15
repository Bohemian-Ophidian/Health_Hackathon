package llama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/common" // Importing the common interface
)

// Client implements the LLaMAClient interface
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
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

// AnalyzeText sends a POST request to the LLaMA API to analyze the provided text
func (c *Client) AnalyzeText(ctx context.Context, text string) (*common.AnalysisResponse, error) {
	if text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	requestBody, err := json.Marshal(common.AnalysisRequest{Text: text})
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

	var result common.AnalysisResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if result.ProcessedAt.IsZero() {
		result.ProcessedAt = time.Now()
	}

	return &result, nil
}

// AnalyzeMedicationNames extracts medication names from OCR text
func (c *Client) AnalyzeMedicationNames(ctx context.Context, text string) ([]string, error) {
	// Use a prompt to extract medication names from the OCR text
	prompt := fmt.Sprintf(`Extract all medication names from the following text:
	%s`, text)

	// Call the LLaMA analyze method
	response, err := c.AnalyzeText(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze text: %w", err)
	}

	// Extract medication names from LLaMA response (assuming names are strings)
	var medications []string
	for _, medication := range response.Analysis {
		medications = append(medications, medication.(string))
	}

	return medications, nil
}

// AnalyzeMedication analyzes a specific medication name for its details
func (c *Client) AnalyzeMedication(ctx context.Context, medicationName string) (map[string]interface{}, error) {
	prompt := fmt.Sprintf(`Analyze the following medication:
	Name: %s
	
	Provide:
	1. Common uses
	2. Typical dosage
	3. Side effects
	4. Interactions
	5. Precautions`, medicationName)

	// Send the prompt to the LLaMA API
	response, err := c.AnalyzeText(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze medication: %w", err)
	}

	return map[string]interface{}{
		"medication_name": medicationName,
		"analysis":        response.Analysis,
		"confidence":      response.Confidence,
		"analyzed_at":     response.ProcessedAt,
	}, nil
}
