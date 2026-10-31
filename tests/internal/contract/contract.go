package contract

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func ReadArchitecture() (string, error) {
	data, err := os.ReadFile(filepath.Join("..", "..", "arquitetura-final-trading.md"))
	if err != nil {
		return "", fmt.Errorf("read architecture contract: %w", err)
	}
	return string(data), nil
}

type Scenario struct {
	Name    string
	Context string
	Rule    string
}

func RunScenarios(t *testing.T, document string, scenarios []Scenario) {
	t.Helper()
	if len(scenarios) < 50 {
		t.Fatalf("need at least 50 named scenarios, got %d", len(scenarios))
	}
	seen := make(map[string]bool, len(scenarios))
	for _, scenario := range scenarios {
		scenario := scenario
		if scenario.Name == "" || scenario.Context == "" || scenario.Rule == "" {
			t.Fatal("every scenario needs a name, market context, and contract rule")
		}
		if seen[scenario.Name] {
			t.Fatalf("duplicate scenario name %q", scenario.Name)
		}
		seen[scenario.Name] = true
		t.Run(scenario.Name, func(t *testing.T) {
			if !strings.Contains(document, scenario.Rule) {
				t.Fatalf("context %q cites a rule missing from architecture contract: %s", scenario.Context, scenario.Rule)
			}
		})
	}
}

func RequireInvariants(document string, terms ...string) error {
	var missing []string
	for _, term := range terms {
		if !strings.Contains(document, term) {
			missing = append(missing, term)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("architecture contract is missing invariants: %s", strings.Join(missing, "; "))
	}
	return nil
}
