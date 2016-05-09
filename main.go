package main

import (
	"net/http"
	"log"
	"code.google.com/p/go.net/websocket"
	"io"
	"os"
)


func EchoServer(ws *websocket.Conn){
	log.Println("WS Connected: ",ws.RemoteAddr().String())
	defer log.Println("WS Closed: ", ws.RemoteAddr().String())

	_, err := io.Copy(ws, ws)
	if err != nil{
		log.Println("Copy error: ", err)

	}
}

var dirPath string

func IndexPage(w http.ResponseWriter, req *http.Request){
	log.Println("Client connected: " + req.RemoteAddr)
	defer log.Println("Client disconnected: " + req.RemoteAddr)

	fp, err := os.Open(dirPath + "/index.html")
	if err != nil{
		log.Println("Could not open file ", err.Error())
		w.Write([]byte("ERROR OPEN FILE: "+ dirPath + "/index.html"))
		return
	}
	defer fp.Close()
	_, err = io.Copy(w, fp)
	if err != nil {
		log.Println("Error reading file: ", err.Error())
		return
	}
}

func main(){
	dirPath = "./src"
	http.HandleFunc("/", IndexPage)
	http.Handle("/ws", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}