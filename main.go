package main

import (
	"fmt"
	"log"
	"net/http"

	flags "github.com/jessevdk/go-flags"
)

type options struct {
	StartPort     int `long:"start-port" default:"20000" description:"The first port which the server listens to."`
	NumberOfPorts int `long:"number-of-ports" short:"n" default:"200" description:"The total number of ports that the server listens to."`
}

func main() {
	var opts options
	_, err := flags.Parse(&opts)
	if err != nil {
		if flags.WroteHelp(err) {
			return
		}
		log.Fatal(err)
	}

	fmt.Printf("Listen from :%d to :%d\n", opts.StartPort, opts.StartPort+opts.NumberOfPorts-1)
	for i := 1; i < opts.NumberOfPorts; i++ {
		port := opts.StartPort + i
		go func(port int) {
			http.ListenAndServe(fmt.Sprintf(":%d", port), &httpPortHandler{port})
		}(port)
	}
	http.ListenAndServe(fmt.Sprintf(":%d", opts.StartPort), &httpPortHandler{opts.StartPort})
}

type httpPortHandler struct {
	Port int
}

func (h *httpPortHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("HTTP Server on port %d", h.Port)))
}
