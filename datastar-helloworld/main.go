package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"net/http"

	"github.com/duongphuhiep/datastar-helloworld/pkg/toolspack"
	datastar "github.com/starfederation/datastar/sdk/go"
)

var globalReqId = 0

type Store struct {
	Myinput string `json:"myinput"` // delay in milliseconds between each character of the message.
}

func setInputHandler(w http.ResponseWriter, r *http.Request) {
	store := &Store{}
	//deserialize signals state
	if err := datastar.ReadSignals(r, store); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	globalReqId++
	reqId := strconv.Itoa(globalReqId)
	slog.InfoContext(r.Context(), "Received input from client "+reqId, "myinput", store.Myinput)
	sse := datastar.NewSSE(w, r)
	for range 5 {
		slog.InfoContext(r.Context(), "Send response for "+reqId)

		//update the state with new value
		store.Myinput = fmt.Sprintf("%s - time on server is %s", reqId, time.Now().Format(time.RFC3339))

		//merge the new signals state back to the client
		newStoreBytes, err := json.Marshal(store)
		if err != nil {
			panic(err)
		}
		sse.MergeSignals(newStoreBytes)

		time.Sleep(2 * time.Second)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "./web/hello-world.html", http.StatusFound)
}

func main() {
	ctx := context.Background()

	tracerDispose := toolspack.InitTracer()
	defer tracerDispose()

	loggerDispose := toolspack.InitOtelLogger(ctx)
	defer loggerDispose()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)

	// Apply LoggingMiddleware only to /actions/setinput
	mux.Handle("POST /actions/setinput", toolspack.LoggingMiddleware(http.HandlerFunc(setInputHandler)))

	mux.Handle("GET /web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	slog.Info("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
