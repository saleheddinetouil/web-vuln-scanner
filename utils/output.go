package utils

import (
    "encoding/json"
    "fmt"
    "go-web-vuln-scanner/scanner"
)



func PrintVulnerabilities(vulnerabilities []scanner.Vulnerability) {
    if len(vulnerabilities) == 0 {
        fmt.Println("No vulnerabilities found.")
        return
    }

    // JSON output
    jsonData, err := json.MarshalIndent(vulnerabilities, "", "  ")
    if err != nil {
        fmt.Println("Error marshaling JSON:", err) // Handle error
        return
    }

    fmt.Println(string(jsonData))

}
