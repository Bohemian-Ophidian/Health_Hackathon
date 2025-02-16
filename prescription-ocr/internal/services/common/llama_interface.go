package common

import (
	"context"
)

// LLaMAClient is an interface for analyzing medication names
type LLaMAClient interface {
	AnalyzeMedicationNames(text string) ([]string, error)
}
