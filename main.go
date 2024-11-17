package main

import (
    "flag"
    "fmt"
    "go-web-vuln-scanner/scanner"
    "go-web-vuln-scanner/utils"
    "log"
)

func main() {
    // Command-line flags for target URL and vulnerability types
    targetURL := flag.String("url", "", "Target URL to scan")
    vulnTypes := flag.String("types", "xss,sqli", "Comma-separated list of vulnerability types to scan (e.g., xss,sqli,csrf)")
    flag.Parse()


    if *targetURL == "" {
        log.Fatal("Target URL (-url) is required.")
    }

    // Split the vulnerability types string into a slice
    vulnTypesSlice := strings.Split(*vulnTypes, ",")

    fmt.Println("Starting scan on:", *targetURL)
    vulnerabilities := scanner.Scan(*targetURL, vulnTypesSlice)

    // Output the results (using the utils package)
    utils.PrintVulnerabilities(vulnerabilities)

}
