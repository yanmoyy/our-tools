package auth

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/yanmoyy/our-tools/internal/sns/browser"
	"github.com/yanmoyy/our-tools/internal/sns/color"
)

const successHTML = "internal/sns/templates/success.html"

// get auth code from Kakao Talk.
// See more info:
// https://developers.kakao.com/docs/latest/en/kakaologin/rest-api#request-code
func getAuthCode(apiKey, redirectURI string) (string, error) {
	srv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}
	// channel for signal when callback is received
	codeCh := make(chan string)
	http.HandleFunc("/oauth", handleCallback(codeCh))
	// Context for server shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %s\n", err)
		}
	}()

	if err := requestGetAuthCode(apiKey, redirectURI); err != nil {
		cancel()
		return "", fmt.Errorf("failed to request auth code: %s", err)
	}

	var code string
	// wait for callback
	select {
	case code = <-codeCh:
		// callback received
	case <-ctx.Done():
		// parent context canceled
		return "", ctx.Err()
	}
	// shutdown server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		return "", fmt.Errorf("failed to shutdown server: %s", err)
	}
	// wait for server fully shutdown
	wg.Wait()

	if code == "" {
		return "", fmt.Errorf("no auth code received")
	}
	return code, nil
}

// save auth code to config
func handleCallback(ch chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "no code", http.StatusBadRequest)
			return
		}

		// Send HTML response with JavaScript to attempt auto-close
		w.Header().Set("Content-Type", "text/html")
		tmpl, err := template.ParseFiles(successHTML)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ch <- code
		close(ch)
	}
}

// send request to get auth code, the code will be handled by callback server
func requestGetAuthCode(apiKey, redirectURI string) error {
	req, err := http.NewRequest("GET", getAuthCodeURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err)
	}
	q := req.URL.Query()
	q.Add("client_id", apiKey)
	q.Add("redirect_uri", redirectURI)
	q.Add("response_type", "code")
	q.Add("scope", "talk_message,friends,profile_nickname")
	req.URL.RawQuery = q.Encode()

	// print the URL for user to open
	fmt.Println("Please open the following URL in your browser:")
	fmt.Println(color.Yellow.ColorString(req.URL.String()))

	// Optionally, attempt to open the URL in the default browser
	if err := browser.Open(req.URL.String()); err != nil {
		fmt.Printf("Failed to open browser automatically: %v\n", err)
		fmt.Println("Please manually open the URL above.")
	}

	return nil
}
