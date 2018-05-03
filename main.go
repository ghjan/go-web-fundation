package main

import (
	"log"
	"net/http"
	"time"
	"os"
	"os/signal"
	"fmt"
)

const VERSION = "4"

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, v" + VERSION + ", the request url is :" + r.URL.String()))
}

func main() {
	version4()
	//version3()
	//version2()
	//version1()
}

func version4() {
	server := http.Server{Addr: ":4000", WriteTimeout: 4 * time.Second}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayByte)
	server.Handler = mux
	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()
	log.Println("Starting server... v4")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Print("Server closed unexpected")
		}
	}
	fmt.Println("Server exit")

}
func version3() {
	server := http.Server{Addr: ":4000", WriteTimeout: 4 * time.Second}
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayByte)
	server.Handler = mux
	log.Println("Starting server... v3")
	log.Fatal(server.ListenAndServe())

}
func version2() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayByte)
	log.Println("Starting server... v2")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func sayByte(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("Bye byte, this is version " + VERSION + "!"))
}

func version1() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, this is version 1!"))
	})
	http.HandleFunc("/bye", sayByte)
	log.Println("Starting server... v1")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
