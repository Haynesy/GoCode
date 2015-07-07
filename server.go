package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	// Routing
	http.HandleFunc("/foo/", handleFoo)
	http.HandleFunc("/hello/", handleHello)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	// Listen and Serve
	http.ListenAndServe(":8080", nil)
}

func editHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(writer, page)
}

func saveHandler(writer http.ResponseWriter, request *http.Request) {

}

func handleHello(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	writer.Header().Set("Content-type", "application/json")

	jsonMsg, err := getResponse()
	if err != nil {
		http.Error(writer, "Oops", http.StatusInternalServerError)
	}

	fmt.Fprintf(writer, jsonMsg)
}

func handleFoo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-type", "application/json")

	msg := Message{"Foo response", "Foo body", 12345}
	bytes, err := json.Marshal(msg)

	if err != nil {
		fmt.Fprintf(writer, "Error occured!")
	}

	message := string(bytes)
	fmt.Fprintf(writer, message)

}

func getResponse() (string, error) {
	unixTime := int32(time.Now().Unix())
	msg := Message{"Hi", "Hello All!", unixTime}
	jbMsg, err := json.Marshal(msg)

	if err != nil {
		return "", err
	}

	jsonMsg := string(jbMsg[:]) // Convert byte array to string
	return jsonMsg, nil
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

type Message struct {
	Title string
	Body  string
	Time  int32
}

type Page struct {
	Title string
	Body  []byte
}
