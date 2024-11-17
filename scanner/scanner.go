package scanner

import (
    "fmt"
    "go-web-vuln-scanner/utils"
    "net/http"
    "net/url"
    "strings"
    "sync"
)

// Scan performs the vulnerability scan.
func Scan(targetURL string, vulnTypes []string) []Vulnerability {
    var vulnerabilities []Vulnerability
    var wg sync.WaitGroup
    client := NewHTTPClient()

    for _, vulnType := range vulnTypes {
        payloads := GetPayloads(vulnType)
        if payloads == nil {
            fmt.Printf("Warning: No payloads found for vulnerability type: %s\n", vulnType)
            continue
        }

        for _, payload := range payloads {
            wg.Add(1)
            go func(payload, vulnType string) { // Pass vulnType to the goroutine
                defer wg.Done()

                // Construct test URL based on vulnerability type (improve this logic as needed)
                testURL := buildTestURL(targetURL, vulnType, payload)


                resp, err := client.Get(testURL)
                if err != nil {
                    fmt.Printf("Error during request: %s\n", err) // Log the error
                    return
                }
                defer resp.Body.Close()


                if checkVulnerability(resp, vulnType, payload) { // Separate check function
                    vulnerabilities = append(vulnerabilities, Vulnerability{
                        Type:    vulnType,    // Use vulnType from the goroutine
                        Payload: payload,
                        URL:     testURL,
                    })
                }

            }(payload, vulnType)
        }
    }

    wg.Wait()
    return vulnerabilities
}



// buildTestURL constructs the URL for testing, handling different injection points.
func buildTestURL(targetURL, vulnType, payload string) string {
        // Implement more sophisticated logic for different vulnerability types here
        // Example for XSS in a query parameter:
        return targetURL + "?q=" + url.QueryEscape(payload) 
}




// checkVulnerability performs the actual vulnerability check based on type.
func checkVulnerability(resp *http.Response, vulnType, payload string) bool {
    switch vulnType {
    case "xss":
        return strings.Contains(resp.Status, payload) // Basic XSS check - improve this
    case "sqli":
        // Add SQLi check logic here (e.g., error messages, time-based)
        // ...
        return false // Placeholder for now
    // ... other cases

    default:
        return false
    }
}



// ... (NewHTTPClient from previous example)
