package main

import (
	"net/http"
)

var pollers = make(chan chan<- []byte)

func init() {
	http.HandleFunc("/poll", poll)
}

func getPollers() []chan<- []byte {
	result := []chan<- []byte{}
	for {
		select {
		case poller := <-pollers:
			result = append(result, poller)
		default:
			return result
		}
	}
}

func poll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var msg []byte

	ch := make(chan []byte)
	select {
	case <-ctx.Done():
	case pollers <- ch:
		msg = <-ch
	}

	w.Header().Set("Content-Type", "text/html")
	templates.ExecuteTemplate(w, "poll.html", string(msg))
}
