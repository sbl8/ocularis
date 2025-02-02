
# Ocularis

Ocularis is a Go-based tool designed to process and generate reports from structured data, such as subdomain lists exported through tools like [Subfinder](https://github.com/projectdiscovery/subfinder). It is built to be modular, efficient, and easy to integrate into your security or data analysis workflows.

## Features

- **Report Generation**: Generate HTML reports from structured data using customizable Go templates.
- **Optimization**: Optimize generated reports for size, readability, or performance.
- **Extensibility**: Easily extend the tool to support additional data formats or output types.
- **Lightweight**: Built in Go, Ocularis is fast and resource-efficient.

## Project Structure

```
ocularis/
├── cmd
│   └── ocularis
│       └── main.go
├── go.mod
├── internal
│   └── report
│       ├── generator.go
│       └── optimizer.go
└── templates
    ├── report.gohtml
```

### Key Components

1. **`cmd/ocularis/main.go`**: The entry point of the application. Initializes and runs the report generation process.
2. **`internal/report/generator.go`**: Handles the logic for generating reports from templates.
3. **`internal/report/optimizer.go`**: Contains logic to optimize the generated reports (e.g., compressing, formatting, etc.).
4. **`templates/report.gohtml`**: An HTML template used for generating the report.

## How It Works

1. **Data Ingestion**:
   - Ocularis ingests structured data, such as subdomain lists, and prepares it for processing.
2. **Report Generation**:
   - The `generator.go` file reads the `report.gohtml` template, processes it with the ingested data, and generates a report.
3. **Report Optimization**:
   - The `optimizer.go` file applies optimizations to the report, such as minification or compression.
4. **Output**:
   - The final report is saved to the specified output directory.

## Use Cases

- **Security Teams**: Generate and analyze reports from subdomain enumeration tools like Subfinder.
- **Data Analysts**: Process and visualize structured data in a clean, readable format.
- **Developers**: Extend the tool to support custom data formats or templates.

## Installation

### Prerequisites

- Go 1.20 or higher
- Git

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/sbl8/ocularis.git
   cd ocularis
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the project:
   ```bash
   go build -o ocularis ./cmd/ocularis
   ```

4. Run the application:
   ```bash
   ./ocularis
   ```

## Usage

To generate a report, run the following command:

```bash
./ocularis -input path/to/subdomains.txt -output path/to/report.html
```

- `-input`: Path to the input file (e.g., subdomain list).
- `-output`: Path to save the generated report.

## Contributing

We welcome contributions! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeatureName`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeatureName`).
5. Open a pull request.

For more details, see our [Contributing Guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Roadmap

- [ ] Add support for JSON and CSV input formats.
- [ ] Integrate with additional security tools (e.g., Amass, Assetfinder).
- [ ] Add a CLI flag for custom template paths.
- [ ] Improve error handling and logging.

## Acknowledgments

- Inspired by tools like [Subfinder](https://github.com/projectdiscovery/subfinder) and [Amass](https://github.com/OWASP/Amass).
- Built with ❤️ using Go.

---

For questions or feedback, feel free to reach out to [sbl8](https://github.com/sbl8).
