package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type replyStructure struct {
	Data   string
	Status int
	Id     string
}

func homePage(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		idValue := req.URL.Query().Get("id")
		myReply := replyStructure{Data: "This is a test of the data attribute", Status: 200, Id: idValue}
		bs, _ := json.Marshal(myReply)
		w.Write(bs)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		/*w.Write([]byte(`
			<html>
				<div>
					Error: Unsupported Method: ` + req.Method + `
				</div>
			</html>
		`))*/
	}

}

func orderPage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`
		<html>
			<title>This is the test page</title>
			<div>
				<p>
					This is the orders page
				</p>
			</div>
		</html>
		`))
}

func main() {
	fmt.Println("Server started on port 8080")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/orders", orderPage)
	http.ListenAndServe(":8080", nil)
}
