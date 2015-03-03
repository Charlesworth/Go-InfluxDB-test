package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"strconv"
	"time"
	"encoding/json"
)

func main() {

	//setDescription("test6", "test description", time.Now(), "Pizza", "Pasta", "Chilli", "Soup")
	//setVote("test6", "test comment", "1", "0", "0", "0", "0")
	//setVote("test6", "test comment", "1", "0", "0", "0", "0")
	//setVote("test6", "test comment", "0", "0", "0", "1", "0")

	//getDescripDB()
	getPooDB()

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

func getDescripDB() {

	query := "select+LAST(description),+LAST(v1),+LAST(v2),+LAST(v3),+LAST(v4),+LAST(endTime)+from+test+limit+1;"
	url := "http://178.62.74.225:8086/db/test/series?q=" + query + "&u=root&p=root"

	client := &http.Client{}
	resp, err := client.Get(url)
	handleErr(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	
}

func getVoteDB() {

	query := "select+SUM(v1),+SUM(v2),+SUM(v3),+SUM(v4),+SUM(NotA)+from+test6;"
	url := "http://178.62.74.225:8086/db/test/series?q=" + query + "&u=root&p=root"

	client := &http.Client{}
	resp, err := client.Get(url)
	handleErr(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func getPooDB() {

	query := "select+SUM(v1),+SUM(v2),+SUM(v3),+SUM(v4),+SUM(NotA)+from+test6;"
	url := "http://178.62.74.225:8086/db/test/series?q=" + query + "&u=root&p=root"

	client := &http.Client{}
	resp, err := client.Get(url)
	handleErr(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	slcB, _ := json.Marshal(string(body))
    	fmt.Println(string(slcB))
	fmt.Println(slcB.points[1])
	fmt.Println(slcB.points[2])
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
