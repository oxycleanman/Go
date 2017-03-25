package main

import "net/http"

const myVar string = "My Test String"
const myVar2 string = "My Test String 2"

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://ma-returns.apps-np.homedepot.com/?id="+r.URL.Path[1:], 301)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
