package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DeviceCodeResponse is the response from a device code request.
type DeviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

// TokenResponse is the response from a token request.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

// DeviceCodeFlow runs the OAuth device code flow.
func DeviceCodeFlow(ctx context.Context, deviceAuthURL, tokenURL, clientID string) (*TokenResponse, error) {
	// Request device code
	resp, err := http.PostForm(deviceAuthURL, url.Values{
		"client_id": {clientID},
		"scope":     {"read:user"},
	})
	if err != nil {
		return nil, fmt.Errorf("device code request failed: %w", err)
	}
	defer resp.Body.Close()

	var dcResp DeviceCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&dcResp); err != nil {
		return nil, fmt.Errorf("failed to decode device code response: %w", err)
	}

	fmt.Printf("Please visit %s and enter code: %s\n", dcResp.VerificationURI, dcResp.UserCode)

	// Poll for token
	interval := time.Duration(dcResp.Interval) * time.Second
	if interval == 0 {
		interval = 5 * time.Second
	}
	deadline := time.Now().Add(time.Duration(dcResp.ExpiresIn) * time.Second)

	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(interval):
		}

		tokenResp, err := pollForToken(tokenURL, clientID, dcResp.DeviceCode)
		if err != nil {
			continue
		}
		return tokenResp, nil
	}
	return nil, fmt.Errorf("device code flow timed out")
}

func pollForToken(tokenURL, clientID, deviceCode string) (*TokenResponse, error) {
	resp, err := http.PostForm(tokenURL, url.Values{
		"client_id":   {clientID},
		"device_code": {deviceCode},
		"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}
	if tokenResp.AccessToken == "" {
		return nil, fmt.Errorf("no access token in response")
	}
	return &tokenResp, nil
}

// RefreshToken refreshes an OAuth token.
func RefreshToken(tokenURL, clientID, refreshToken string) (*TokenResponse, error) {
	resp, err := http.PostForm(tokenURL, url.Values{
		"client_id":     {clientID},
		"refresh_token": {refreshToken},
		"grant_type":    {"refresh_token"},
	})
	if err != nil {
		return nil, fmt.Errorf("token refresh failed: %w", err)
	}
	defer resp.Body.Close()

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}
	return &tokenResp, nil
}

// ParseBearerToken extracts the token from an Authorization header.
func ParseBearerToken(authHeader string) string {
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(authHeader, "Bearer ")
}
