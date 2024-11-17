package scanner

import (
    "net/http"
    "time"
)

// Create a custom HTTP client with timeout and other settings
func NewHTTPClient() *http.Client {
    return &http.Client{
        Timeout: 10 * time.Second, // Set a timeout
        // ... other settings (e.g., custom headers, cookies)
    }
}
