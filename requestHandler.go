package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"strconv"
	"time"
)

func main() {

	//	setDescription("fart", "a big smelly fart", time.Now(), "Pizza", "Pasta", "Chilli", "Soup")
	//	setVote("fart", "a big smelly comment", "1", "0", "0", "0", "0")

	getFromDB()

	//	http.HandleFunc("/", testOut)

	//	fmt.Println("Listening on port 3000...")
	//	http.ListenAndServe(":3000", nil)
}

func setVote(n string, d string, v1 string, v2 string, v3 string, v4 string, NotA string) {
	var jsonStr = []byte(`[
	  {
	    "name" : "` + n + `",
	    "columns" : ["description", "endTime", "v1", "v2", "v3", "v4", "NotA"],
	    "points" : [
	    ["` + d + `",
		0,
		` + v1 + `,
		` + v2 + `,
		` + v3 + `,
		` + v4 + `,
		` + NotA + `]
	    ]
	  }
	]`)

	sendToDB(jsonStr)
}

func setDescription(n string, d string, e time.Time, v1 string, v2 string, v3 string, v4 string) {
	endTime := strconv.Itoa(e.Nanosecond())

	var jsonStr = []byte(`[
	  {
	    "name" : "` + n + `",
	    "columns" : ["description", "endTime", "v1", "v2", "v3", "v4", "NotA"],
	    "points" : [
	    ["` + d + `",
		` + endTime + `,
		"` + v1 + `",
		"` + v2 + `",
		"` + v3 + `",
		"` + v4 + `",
		""]
	    ]
	  }
	]`)

	sendToDB(jsonStr)
}

func sendToDB(jsonStr []byte) {
	url := "http://178.62.74.225:8086/db/test/series?u=root&p=root"

	fmt.Println(string(jsonStr))

	client := &http.Client{}
	resp, err := client.Post(url, "text/plain", bytes.NewBuffer(jsonStr))
	handleErr(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func getFromDB() {

	//	var Url *url.URL
	//	Url, err := url.Parse("http://178.62.74.225:8086")
	//	if err != nil {
	//		panic("boom")
	//	}

	//	Url.Path += "/db/test/series"
	//	parameters := url.Values{}
	//	parameters.Add("u", "root")
	//	parameters.Add("p", "root")
	//	parameters.Add("q", "list series")
	//	Url.RawQuery = parameters.Encode()

	//	fmt.Printf("Encoded URL is %q\n", Url.String())

	url := "http://178.62.74.225:8086/db/test/series?q=list+series&u=root&p=root"

	client := &http.Client{}
	resp, err := client.Get(url)
	handleErr(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
