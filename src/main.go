package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"apiresp"
)

type configElement struct {
	EligibilityURL    string `json:"EligibilityUrl"`
	EligibilityMethod string
	OrderAggURL       string `json:"OrderAggUrl"`
	OrderAggMethod    string
}

//Config type used to get configuration settings for the API
type config struct {
	Environment string
	Development configElement
	Production  configElement
}

type CurrentConfig struct {
	Settings configElement
}

type replyStructure struct {
	Data   string
	Status int
	Id     string
}

var client = &http.Client{}
var Config CurrentConfig

func handleRequestError(errorMessage error, errorCode int, w http.ResponseWriter) {
	w.WriteHeader(errorCode)
	log.Fatalln(errorMessage)
}

func handleOrderAggResponse(resp []byte) {
	fmt.Println(string(resp))
}

func homePage(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		//Detect SSO token and return appropriate response
		if ssoCookie, err := req.Cookie("THDSSO"); err != nil {
			fmt.Println("Unauthorized Request From: ", req.Referer())
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			//Make Order Agg callout and respond with data
			orderAggBody := apiresp.GetOrderAggRequest("W111111111")
			if aggRequest, err := http.NewRequest(Config.Settings.OrderAggMethod, Config.Settings.OrderAggURL, bytes.NewBufferString(orderAggBody)); err != nil {
				fmt.Println("Error getting http request object")
				handleRequestError(err, http.StatusInternalServerError, w)
			} else {
				aggRequest.Header.Set("Authorization", "Bearer "+ssoCookie.Value)
				aggRequest.Header.Set("Content-Type", "application/json")
				if resp, err := client.Do(aggRequest); err != nil {
					fmt.Println("Error making service callout")
					handleRequestError(err, http.StatusInternalServerError, w)
				} else {
					respBody, _ := ioutil.ReadAll(resp.Body)
					defer resp.Body.Close()
					w.Write(respBody)
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
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
	var c config
	if err != nil {
		log.Fatalln(err)
	} else {
		if err := json.Unmarshal(configFile, &c); err != nil {
			log.Fatalln(err)
		} else {
			//Set global configuration variable
			if c.Environment == "dev" {
				Config.Settings = c.Development
			} else {
				Config.Settings = c.Production
			}

			fmt.Println("Config Settings: ", Config)
			fmt.Println("Server started on port 8080")
			http.HandleFunc("/", homePage)
			http.HandleFunc("/orders", orderPage)
			http.ListenAndServe("localhost.homedepot.com:8080", nil)
		}
	}
}
