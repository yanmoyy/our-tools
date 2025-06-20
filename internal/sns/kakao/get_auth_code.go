package kakao

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/yanmoyy/our-tools/internal/sns/color"
)

const successHTML = "internal/sns/templates/success.html"

func (cfg *Config) getAuthCode() (string, error) {
	srv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}
	// channel for signal when callback is received
	codeCh := make(chan string)
	http.HandleFunc("/oauth", cfg.handleCallback(codeCh))
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

	if err := requestAuthCode(cfg.apiKey, cfg.redirectURI); err != nil {
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
func (cfg *Config) handleCallback(ch chan string) http.HandlerFunc {
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
func requestAuthCode(clientID, redirectURI string) error {
	req, err := http.NewRequest("GET", getAuthCodeURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err)
	}
	q := req.URL.Query()
	q.Add("client_id", clientID)
	q.Add("redirect_uri", redirectURI)
	q.Add("response_type", "code")
	req.URL.RawQuery = q.Encode()

	// print the URL for user to open
	fmt.Println("Please open the following URL in your browser:")
	fmt.Println(color.Yellow.ColorString(req.URL.String()))

	// Optionally, attempt to open the URL in the default browser
	if err := openBrowser(req.URL.String()); err != nil {
		fmt.Printf("Failed to open browser automatically: %v\n", err)
		fmt.Println("Please manually open the URL above.")
	}

	return nil
}
