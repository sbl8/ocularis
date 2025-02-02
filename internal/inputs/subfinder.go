package inputs

import (
	"encoding/json"
	"fmt"
	"os"

	"ocularis/internal/core"
)

// internal/inputs/subfinder.go
type SubfinderEntry struct {
	Host     string   `json:"host"`     // Matches Subfinder's "host" field
	Root     string   `json:"root"`     // Subfinder's "root" instead of "input"
	Sources  []string `json:"sources"`  // Already correct
	IP       string   `json:"ip"`       // Additional Subfinder field
	Provider string   `json:"provider"` // Additional Subfinder field
}

func (s *SubfinderEntry) ToCoreEntry() core.Entry {
	return core.Entry{
		Host:    s.Host,
		Input:   s.Root, // Map "root" to "Input"
		Sources: s.Sources,
		// Add other fields as needed
	}
}

// LoadSubfinderData loads Subfinder JSON data from a file.
func LoadSubfinderData(filePath string) ([]core.Entry, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var entries []core.Entry
	if err := json.Unmarshal(content, &entries); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return entries, nil
}

// GenerateSampleData generates sample data for testing.
func GenerateSampleData() []core.Entry {
	var data []core.Entry
	for i := 0; i < 50; i++ {
		entry := core.Entry{
			Host:  fmt.Sprintf("sub%d.example%d.com", i, i%7),
			Input: "example.com",
		}
		if i%2 == 0 {
			entry.Sources = []string{"waybackarchive", "hackertarget"}
		} else {
			entry.Sources = []string{"hudsonrock"}
		}
		data = append(data, entry)
	}
	return data
}
