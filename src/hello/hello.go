package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://ma-returns.apps-np.homedepot.com/?id=" + r.URL.Path[1:], 301)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}