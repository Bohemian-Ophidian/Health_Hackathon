package llama

import "context"

// LLaMAClient defines the interface for interacting with LLaMA
type LLaMAClient interface {
	AnalyzeMedicationNames(ctx context.Context, text string) ([]string, error)
	AnalyzeMedication(ctx context.Context, medicationName string) (map[string]interface{}, error)
}
