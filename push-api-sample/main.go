package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/SherClockHolmes/webpush-go"
)

type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     struct {
		P256dh string `json:"p256dh"`
		Auth   string `json:"auth"`
	} `json:"keys"`
}

const PORT = ":8080"

var subscriptions []Subscription

const VAPID_PUBLIC_KEY = "BIlQKyaTvb5GVgA-kw4URKBWAEWRS-OiC8jsXUg0eKgsFBGxk4MY90qzdIdtzFxWOfxLrD8LY_eTsDx3jRcsYaU"
const VAPID_PRIVATE_KEY = "fPkPZpw2LyNfeg01te54N9msnzs0TB4j6kCbmqMvwgE"

func main() {
	setupLogging()

	// Serve static files (HTML, JS, CSS, etc.) from the current directory
	fs := http.FileServer(http.Dir("."))
	slog.Info("Server started at http://localhost" + PORT)
	http.Handle("/", fs)

	slog.Info("http://localhost" + PORT + "/subscribe")
	http.HandleFunc("/subscribe", handleSubscribe)

	slog.Info("http://localhost" + PORT + "/push")
	http.HandleFunc("/push", handlePush)

	log.Fatal(http.ListenAndServe(PORT, nil))
}

// Store subscription
func handleSubscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var sub Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	subscriptions = append(subscriptions, sub)
	w.WriteHeader(http.StatusCreated)
}

// Send push notification
func handlePush(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Example: Send a notification to all subscribers
	for _, sub := range subscriptions {
		s := webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				P256dh: sub.Keys.P256dh,
				Auth:   sub.Keys.Auth,
			},
		}

		// Send notification
		resp, err := webpush.SendNotification(
			[]byte(`{"title":"Hello from Go!","body":"This is a push notification"}`),
			&s,
			&webpush.Options{
				Subscriber:      "mailto:admin@example.com", // Your contact
				VAPIDPublicKey:  VAPID_PUBLIC_KEY,
				VAPIDPrivateKey: VAPID_PRIVATE_KEY,
			},
		)
		if err != nil {
			log.Printf("Error sending notification: %v", err)
			continue
		}
		defer resp.Body.Close()
	}

	w.WriteHeader(http.StatusOK)
}

// Configure slog to write logs to stdout
func setupLogging() {
	var handler slog.Handler
	if isDev() {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug, // Set desired log level
			AddSource: false,           // Show file and line number
		})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true, // Show file and line number
		})
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

// isDev checks if the environment is development.
func isDev() bool {
	return os.Getenv("APP_ENV") == "development"
}
