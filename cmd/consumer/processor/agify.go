package processor

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type AgifyProcessor struct {
	client *http.Client
}

func NewAgifyProcessor() *AgifyProcessor {
	return &AgifyProcessor{
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

var cache sync.Map

func (a *AgifyProcessor) Process(input string) (string, error) {
	// Expecting input like "name:abigail"
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return "", fmt.Errorf("invalid input format, expected 'name:<value>', got: '%s'", input)
	}
	name := strings.TrimSpace(parts[1])

	// Check cache
	if result, ok := cache.Load(name); ok {
		return result.(string), nil
	}

	url := fmt.Sprintf("https://api.agify.io?name=%s", name)

	resp, err := a.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("error calling agify: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	result := string(body)
	cache.Store(name, result)

	return result, nil
}
