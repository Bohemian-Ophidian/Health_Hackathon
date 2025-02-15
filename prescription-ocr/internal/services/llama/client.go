package llama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

// NewClient initializes the LLaMA client with the provided base URL
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// AnalyzeText sends the OCR text to the LLaMA API for analysis
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

// AnalyzeMedication analyzes medication names from extracted text
func (c *Client) AnalyzeMedicationNames(text string) ([]string, error) {
	// Use a prompt to extract medication names from the OCR text
	prompt := fmt.Sprintf(`Extract all medication names from the following text:
	%s`, text)

	// Call the LLaMA analyze method
	response, err := c.AnalyzeText(context.Background(), prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze text: %w", err)
	}

	// Extract medication names from LLaMA response (assuming it's part of the analysis)
	// You may need to adjust based on LLaMA's response structure
	var medications []string
	for _, medication := range response.Analysis {
		medications = append(medications, medication.(string)) // Assuming names are strings
	}

	return medications, nil
}

// AnalyzeMedication fetches medication details and analysis using LLaMA
func (c *Client) AnalyzeMedication(ctx context.Context, medicationName string) (map[string]interface{}, error) {
	// Structured prompt for medication analysis
	prompt := fmt.Sprintf(`Analyze the following medication:
	Name: %s
	
	Provide:
	1. Common uses
	2. Typical dosage
	3. Side effects
	4. Interactions
	5. Precautions`, medicationName)

	// Call the LLaMA analyze method
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
