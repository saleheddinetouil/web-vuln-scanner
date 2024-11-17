# Web Vulnerability Scanner

This project is a web vulnerability scanner written in Go. It's designed to help security professionals and developers identify common web application vulnerabilities.  This is a work in progress and is intended for educational and ethical hacking purposes only.


## Features

* **Scans for multiple vulnerabilities:**  Currently supports Cross-Site Scripting (XSS) and SQL Injection (SQLi) with plans to expand to other vulnerability types (CSRF, directory traversal, etc.).
* **Modular design:**  Organized into packages for better maintainability and extensibility (`scanner`, `utils`).
* **Concurrent scanning:**  Uses goroutines for faster scanning of multiple payloads.
* **Customizable HTTP client:**  Allows control over timeouts and other request parameters.
* **Structured output:**  Provides results in JSON format for easy parsing and integration with other tools.
* **Command-line flags:**  Flexible usage with options to specify the target URL and vulnerability types to scan.


## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/go-web-vuln-scanner.git  // Replace with your repo URL
   ```

2. **Build the scanner:**
   ```bash
   go build
   ```

3. **Run the scanner:**
   ```bash
   ./go-web-vuln-scanner -url <target_url> -types <vulnerability_types>
   ```
   * `-url`: The URL of the target website (required).
   * `-types`: A comma-separated list of vulnerability types to scan (e.g., `xss,sqli`).  Defaults to `xss,sqli`.

   Example:
   ```bash
   ./go-web-vuln-scanner -url http://example.com -types xss,sqli
   ```

## Usage Examples


**Scan for XSS and SQLi:**
```bash
./go-web-vuln-scanner -url https://www.example.com -types xss,sqli
```

**Scan only for XSS:**
```bash
./go-web-vuln-scanner -url https://www.example.com -types xss 
```


## Disclaimer

This scanner is provided for educational and ethical hacking purposes only.  **It is crucial to obtain explicit permission before scanning any website.**  Unauthorized scanning is illegal and unethical.  The developers of this project are not responsible for any misuse of this tool.  Use it responsibly.


## Contributing

Contributions are welcome!  Feel free to open issues or submit pull requests to improve the scanner's functionality, add new vulnerability checks, or enhance the codebase.


## Roadmap

* Implement more vulnerability checks (CSRF, directory traversal, command injection, etc.).
* Improve payload management and context-aware scanning.
* Enhance output formatting and reporting.
* Add rate limiting and other safety features.
* Integrate with other security tools.


## License

This project is licensed under the [MIT License](LICENSE). 


## Contact

Saleh Eddine Touil - saleh.touil@icloud.com

