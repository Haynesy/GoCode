package main

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
)

func main() {
    http.HandleFunc("/foo", handleFoo)
    http.HandleFunc("/hello", handleHello)
    http.ListenAndServe(":8080", nil)
}

func handleHello(writer http.ResponseWriter, request *http.Request){
    writer.Header().Set("Access-Control-Allow-Origin", "*")

    writer.Header().Set("Content-type", "application/json")

    jsonMsg, err := getResponse()
    if err != nil {
        http.Error(writer, "Oops", http.StatusInternalServerError)
    }

    fmt.Fprintf(writer, jsonMsg)
}

func handleFoo(writer http.ResponseWriter, request *http.Request){
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

func getResponse() (string, error){
    unixTime := int32(time.Now().Unix())
    msg := Message{"Hi", "Hello All!", unixTime}
    jbMsg, err := json.Marshal(msg)

    if err != nil {
        return "", err
    }

    jsonMsg := string(jbMsg[:]) // Convert byte array to string
    return jsonMsg, nil
}

type Message struct {
    Title string
    Body string
    Time int32
}
