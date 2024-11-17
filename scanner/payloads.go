package scanner

// Payloads for different vulnerability types
var xssPayloads = []string{
    "<script>alert('XSS');</script>",
    "<img src=x onerror=alert('XSS')>",
    // ... more payloads
}

var sqliPayloads = []string{
    "' OR '1'='1",
    "--",
    "/*",  // ... more payloads
}

// ... add payloads for other vulnerabilities (CSRF, etc.)

// Function to get payloads for a specific vulnerability type
func GetPayloads(vulnType string) []string {
    switch vulnType {
    case "xss":
        return xssPayloads
    case "sqli":
        return sqliPayloads
    // ... other cases
    default:
        return nil
    }
}
