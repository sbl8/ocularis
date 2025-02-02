package core

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode"
)

// Entry represents one JSON record.
type Entry struct {
	Host    string   `json:"host"`
	Input   string   `json:"input"`
	Sources []string `json:"sources"`
}

// DomainDisplay represents the data to be displayed in the report table.
type DomainDisplay struct {
	Host       string   `json:"host"`
	RootDomain string   `json:"root_domain"`
	Sources    []string `json:"sources"`
}

// TermFreq represents a term and its frequency count.
type TermFreq struct {
	Term      string `json:"term"`
	Frequency int    `json:"frequency"`
}

// ReportData holds all data needed by the report.
type ReportData struct {
	Domains      []DomainDisplay `json:"domains"`
	TermFreqData []TermFreq      `json:"term_freq_data"`
	SummaryStats map[string]int  `json:"summary_stats"`
	RawData      []Entry         `json:"raw_data"`
}

// pad applies PKCS7 padding.
func pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// encryptData encrypts the given data (as string) using AES-CBC with PKCS7 padding.
func EcryptData(data string, key []byte) (string, string, error) {
	plaintext := []byte(data)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", fmt.Errorf("failed to create cipher: %w", err)
	}
	blockSize := block.BlockSize()

	// Apply PKCS7 padding.
	plaintext = pad(plaintext, blockSize)

	// Generate a random IV.
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", "", fmt.Errorf("failed to generate IV: %w", err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)

	encryptedB64 := base64.StdEncoding.EncodeToString(ciphertext)
	ivB64 := base64.StdEncoding.EncodeToString(iv)
	return encryptedB64, ivB64, nil
}

// GenerateReportData processes the input JSON data and generates report data.
func GenerateReportData(data []Entry) (ReportData, error) {
	if len(data) == 0 {
		return ReportData{}, errors.New("input data cannot be empty")
	}

	var domains []DomainDisplay
	uniqueRoot := make(map[string]struct{})
	uniqueHosts := make(map[string]struct{})
	var allTerms []string

	for _, e := range data {
		domains = append(domains, DomainDisplay{
			Host:       e.Host,
			RootDomain: e.Input,
			Sources:    e.Sources,
		})
		uniqueHosts[e.Host] = struct{}{}
		uniqueRoot[e.Input] = struct{}{}

		// Break the host into terms by splitting on non-letter characters.
		for _, term := range strings.FieldsFunc(e.Host, func(r rune) bool {
			return !unicode.IsLetter(r)
		}) {
			if len(term) > 3 {
				allTerms = append(allTerms, strings.ToLower(term))
			}
		}
	}

	// Count term frequencies.
	termCount := make(map[string]int)
	for _, term := range allTerms {
		termCount[term]++
	}

	var termFreqs []TermFreq
	for term, count := range termCount {
		termFreqs = append(termFreqs, TermFreq{
			Term:      term,
			Frequency: count,
		})
	}
	sort.Slice(termFreqs, func(i, j int) bool {
		return termFreqs[i].Term < termFreqs[j].Term
	})

	summaryStats := map[string]int{
		"total_subdomains":    len(data),
		"unique_hosts":        len(uniqueHosts),
		"unique_root_domains": len(uniqueRoot),
	}

	return ReportData{
		Domains:      domains,
		TermFreqData: termFreqs,
		SummaryStats: summaryStats,
		RawData:      data,
	}, nil
}
