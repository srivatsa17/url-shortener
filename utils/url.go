package utils

import (
	"fmt"
	"net/url"
)

// validateURL checks if the URL has a valid protocol
func ValidateURL(longURL string) error {
	parsedURL, err := url.ParseRequestURI(longURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %v", err)
	}

	// Check if protocol is present and valid
	if parsedURL.Scheme == "" {
		return fmt.Errorf("URL must include protocol (http:// or https://)")
	}

	// Only allow http and https
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("only http and https protocols are supported")
	}

	// Check if host is present
	if parsedURL.Host == "" {
		return fmt.Errorf("URL must include a valid host")
	}

	return nil
}

func IsValidURL(urlStr string) bool {
	if err := ValidateURL(urlStr); err != nil {
		return false
	}

	return true
}
