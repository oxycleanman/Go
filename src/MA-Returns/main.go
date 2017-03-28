package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type configElement struct {
	EligibilityURL    string `json:"eligibilityUrl"`
	EligibilityMethod string `json:"eligibilityMethod"`
}

//Config type used to get configuration settings for the API
type Config struct {
	Development configElement `json:"development"`
	Production  configElement `json:"production"`
}

type replyStructure struct {
	Data   string
	Status int
	Id     string
}

var CONFIG Config

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
	//Get configurations
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	} else {
		if err := json.Unmarshal(configFile, &CONFIG); err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println(CONFIG)
			fmt.Println("Server started on port 8080")
			http.HandleFunc("/", homePage)
			http.HandleFunc("/orders", orderPage)
			http.ListenAndServe(":8080", nil)
		}
	}
}
