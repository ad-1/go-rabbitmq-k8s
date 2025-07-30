package processor

import (
	"fmt"
	"strings"
)

type MockAgifyProcessor struct{}

func NewMockAgifyProcessor() *MockAgifyProcessor {
	return &MockAgifyProcessor{}
}

func (m *MockAgifyProcessor) Process(input string) (string, error) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return "", fmt.Errorf("invalid input format, expected 'name:<value>', got: '%s'", input)
	}
	name := strings.TrimSpace(parts[1])

	// Just return a fake age response for demonstration
	result := fmt.Sprintf(`{"name":"%s", "age":%d, "mock":true}`, name, len(name)+20)
	return result, nil
}
