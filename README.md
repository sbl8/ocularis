# Ocularis 👁️

Ocularis is a Go-based tool for processing and generating reports from structured data, with a particular focus on subdomain enumeration outputs from [Subfinder](https://github.com/projectdiscovery/subfinder).

## Features

- **Encrypted single file HTML Report Generation**: Produces structured HTML reports using Go templates decrypted with a secret key.
- **Tool Specific Ingestion**: Supports ingestion of Subfinder's JSON output with plans to support additional tools in the future.
- **Multiple Export options**: Supports multiple export options (CSV, JSON, txt, etc.).

## Project Structure

```
ocularis
├── cmd
│   └── ocularis
│       └── main.go
├── go.mod
├── internal
│   ├── core
│   │   ├── generator.go
│   │   └── optimizer.go
│   ├── inputs
│   │   └── subfinder.go
│   └── outputs
│       └── html.go
├── ocularis
├── pkg
│   └── utils
│       └── fileutils.go
├── README.md
└── templates
    └── report.gohtml
```

### Core Components

- **`cmd/ocularis/main.go`**: Primary execution point of the application.
- **`internal/core/generator.go`**: Governs structured report generation processes.
- **`internal/core/optimizer.go`**: Implements sanitization before templating.
- **`internal/inputs/subfinder.go`**: Facilitates structured ingestion of Subfinder output.
- **`internal/outputs/html.go`**: Handles structured HTML report generation and output.
- **`pkg/utils/fileutils.go`**: file-handling utilities.

## Installation

### Prerequisites

- Go 1.21+
- Git

### Setup Procedure

```bash
git clone https://github.com/sbl8/ocularis.git
cd ocularis
go mod download
go build -o ocularis ./cmd/ocularis
```

## Usage

```bash
./ocularis -input path/to/subdomains.json -output path/to/report.html -template templates/report.gohtml
```

- `-input`: Specifies the path to the structured input data (JSON).
- `-output`: Defines the destination for the generated report.
- `-template`: Points to the structured template file used for report generation.

## Contribution Guidelines

1. Fork the repository.
2. Establish a feature branch (`git checkout -b feature/YourFeatureName`).
3. Implement modifications (`git commit -m 'feature'`).
4. Push updates (`git push origin feature/YourFeatureName`).
5. Submit a pull request for review.

## Licensing

This project is licensed under the MIT License. Refer to the [LICENSE](LICENSE) file for specifics.


