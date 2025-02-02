package core

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// OptimizeJSONData reads the JSON from inputFile, consolidates records by host,
// and writes the optimized JSON to outputFile.
func OptimizeJSONData(inputFile, outputFile string) error {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	trimmed := strings.TrimSpace(string(content))
	var entries []Entry

	// If the file is a JSON array (starts with "[" and ends with "]"), decode normally.
	if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
		if err := json.Unmarshal(content, &entries); err != nil {
			return fmt.Errorf("failed to unmarshal JSON array: %w", err)
		}
	} else {
		// Otherwise, assume multiple JSON objects.
		// Wrap them in square brackets by replacing newlines between objects.
		wrapped := "[" + strings.ReplaceAll(trimmed, "}\n{", "},{") + "]"
		if err := json.Unmarshal([]byte(wrapped), &entries); err != nil {
			return fmt.Errorf("failed to unmarshal wrapped JSON objects: %w", err)
		}
	}

	// Use a map to consolidate entries by host.
	optimizedMap := make(map[string]struct {
		Input   string
		Sources map[string]struct{}
	})
	for _, e := range entries {
		if e.Host == "" {
			continue
		}
		item, exists := optimizedMap[e.Host]
		if !exists {
			item = struct {
				Input   string
				Sources map[string]struct{}
			}{
				Input:   e.Input,
				Sources: make(map[string]struct{}),
			}
		}
		// Overwrite the input field (if needed) and add new sources.
		item.Input = e.Input
		for _, src := range e.Sources {
			item.Sources[src] = struct{}{}
		}
		optimizedMap[e.Host] = item
	}

	// Build the slice of optimized entries.
	var optimizedEntries []Entry
	for host, item := range optimizedMap {
		var sources []string
		for src := range item.Sources {
			sources = append(sources, src)
		}
		optimizedEntries = append(optimizedEntries, Entry{
			Host:    host,
			Input:   item.Input,
			Sources: sources,
		})
	}

	// Marshal the optimized data.
	outBytes, err := json.MarshalIndent(optimizedEntries, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal optimized entries: %w", err)
	}

	// Write to output file.
	if err := os.WriteFile(outputFile, outBytes, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("Optimized data written to %s\n", outputFile)
	return nil
}
