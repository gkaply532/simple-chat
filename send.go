package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/send", send)
}

func send(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msg := r.PostForm.Get("msg")
	if msg != "" {
		pollers := getPollers()

		fmt.Printf("got message: %q, current pollers: %d\n", msg, len(pollers))

		for _, poller := range pollers {
			poller <- []byte(msg)
		}
	}

	w.Header().Set("Content-Type", "text/html")
	templates.ExecuteTemplate(w, "send.html", struct {
		HasParam bool
		HasMsg   bool
	}{
		HasParam: r.PostForm.Has("msg"),
		HasMsg:   msg != "",
	})
}
