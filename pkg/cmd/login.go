package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/urfave/cli/v3"
)

var loginCommand = cli.Command{
	Name:     "login",
	Category: "AUTH",
	Usage:    "Authenticate with Spotify via browser",
	UsageText: `spotted login [options]

Examples:
  spotted login
  spotted login --server http://localhost:3000
  spotted login --scopes "user-read-private user-read-email"
  spotted login --scopes all`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "server",
			Usage: "OAuth server URL",
			Value: "https://cjav.dev",
		},
		&cli.IntFlag{
			Name:  "port",
			Usage: "Local callback server port",
			Value: 8765,
		},
		&cli.StringFlag{
			Name:  "scopes",
			Usage: "Space-separated scopes to request, or 'all' for all scopes",
			Value: "all",
		},
		&cli.BoolFlag{
			Name:  "no-browser",
			Usage: "Print the URL instead of opening a browser",
			Value: false,
		},
	},
	Action: loginAction,
}

func loginAction(ctx context.Context, cmd *cli.Command) error {
	serverURL := cmd.String("server")
	port := cmd.Int("port")
	scopes := cmd.String("scopes")
	noBrowser := cmd.Bool("no-browser")

	// Build the authorization URL
	authURL, err := url.Parse(serverURL)
	if err != nil {
		return fmt.Errorf("invalid server URL: %w", err)
	}
	authURL.Path = "/spotify_connections/cli"
	q := authURL.Query()
	q.Set("port", strconv.Itoa(int(port)))
	q.Set("scopes", scopes)
	authURL.RawQuery = q.Encode()

	// Channel to receive credentials from callback
	credsChan := make(chan *Credentials, 1)
	errChan := make(chan error, 1)

	// Start local callback server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to start callback server on port %d: %w", port, err)
	}

	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/callback" {
				http.NotFound(w, r)
				return
			}

			query := r.URL.Query()
			accessToken := query.Get("access_token")
			if accessToken == "" {
				errChan <- fmt.Errorf("no access token received")
				http.Error(w, "Authentication failed: no access token", http.StatusBadRequest)
				return
			}

			expiresIn, _ := strconv.Atoi(query.Get("expires_in"))
			var expiresAt time.Time
			if expiresIn > 0 {
				expiresAt = time.Now().Add(time.Duration(expiresIn) * time.Second)
			}

			creds := &Credentials{
				AccessToken:  accessToken,
				RefreshToken: query.Get("refresh_token"),
				ExpiresAt:    expiresAt,
				TokenType:    query.Get("token_type"),
				Scope:        query.Get("scope"),
			}

			credsChan <- creds

			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head><title>Spotted CLI - Login Success</title></head>
<body style="font-family: system-ui; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; background: #1DB954;">
<div style="text-align: center; color: white;">
<h1>Login Successful!</h1>
<p>You can close this window and return to the terminal.</p>
</div>
</body>
</html>`)
		}),
	}

	go func() {
		if err := server.Serve(listener); err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	fmt.Printf("Starting authentication flow...\n")
	fmt.Printf("Callback server listening on http://localhost:%d\n\n", port)

	if noBrowser {
		fmt.Printf("Open this URL in your browser:\n%s\n\n", authURL.String())
	} else {
		fmt.Printf("Opening browser to authenticate...\n")
		if err := openBrowser(authURL.String()); err != nil {
			fmt.Printf("Could not open browser automatically.\n")
			fmt.Printf("Please open this URL manually:\n%s\n\n", authURL.String())
		}
	}

	fmt.Printf("Waiting for authentication...\n")

	// Wait for callback or timeout
	select {
	case creds := <-credsChan:
		server.Shutdown(ctx)

		if err := SaveCredentials(creds); err != nil {
			return fmt.Errorf("failed to save credentials: %w", err)
		}

		path, _ := CredentialsPath()
		fmt.Printf("\nLogin successful!\n")
		fmt.Printf("Credentials saved to: %s\n", path)
		return nil

	case err := <-errChan:
		server.Shutdown(ctx)
		return fmt.Errorf("authentication failed: %w", err)

	case <-time.After(5 * time.Minute):
		server.Shutdown(ctx)
		return fmt.Errorf("authentication timed out after 5 minutes")

	case <-ctx.Done():
		server.Shutdown(ctx)
		return ctx.Err()
	}
}

func openBrowser(url string) error {
	return openBrowserOS(url)
}
