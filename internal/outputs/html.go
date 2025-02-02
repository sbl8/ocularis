package outputs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"ocularis/internal/core"
)

// GenerateHTMLReport generates an HTML report from the provided data.
func GenerateHTMLReport(data []core.Entry, outputFile, templateFile string, encryptionKey []byte) error {
	reportData, err := core.GenerateReportData(data)
	if err != nil {
		return fmt.Errorf("failed to generate report data: %w", err)
	}

	reportJSONBytes, err := json.Marshal(reportData)
	if err != nil {
		return fmt.Errorf("failed to marshal report data: %w", err)
	}

	encryptedDataB64, ivB64, err := core.EcryptData(string(reportJSONBytes), encryptionKey)
	if err != nil {
		return fmt.Errorf("encryption failed: %w", err)
	}

	templateBytes, err := os.ReadFile(templateFile)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}
	htmlContent := string(templateBytes)
	htmlContent = strings.ReplaceAll(htmlContent, "{{encrypted_data}}", encryptedDataB64)
	htmlContent = strings.ReplaceAll(htmlContent, "{{iv}}", ivB64)

	if err := os.WriteFile(outputFile, []byte(htmlContent), 0644); err != nil {
		return fmt.Errorf("failed to write output report: %w", err)
	}

	encryptionKeyB64 := base64.StdEncoding.EncodeToString(encryptionKey)
	fmt.Printf("Report successfully generated: %s\n", outputFile)
	fmt.Println("IMPORTANT: To view the report, provide the following decryption key when prompted:")
	fmt.Println(encryptionKeyB64)

	return nil
}
