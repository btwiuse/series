package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/btwiuse/series/fibonacci"
	"github.com/btwiuse/series/flipper"
	"github.com/btwiuse/series/linear"
	"github.com/btwiuse/series/random"
	"github.com/btwiuse/series/square"
	// "k0s.io/pkg/middleware"
)

func metrics(name, HELP, TYPE string, value int) string {
	helpLn := fmt.Sprintf("# HELP %s %s\n", name, HELP)
	typeLn := fmt.Sprintf("# TYPE %s %s\n", name, TYPE)
	valueLn := fmt.Sprintf("%s %d\n", name, value)
	return helpLn + typeLn + valueLn
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	body := Body()
	io.WriteString(w, body)
}

func main() {
	// fmt.Print(Body())
	handler := http.HandlerFunc(handleMetrics)
	// handler = middleware.LoggingMiddleware(os.Stderr, handler)
	http.Handle("/", handler)
	log.Println("listening on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func Body() string {
	return strings.Join(
		[]string{
			metrics("flipper", "flipper series", "gauge", flipper.Value()),
			metrics("linear", "linear series", "gauge", linear.Value()),
			metrics("square", "square series", "gauge", square.Value()),
			metrics("random", "random series", "gauge", random.Value()),
			metrics("fibonacci", "fibonacci series", "gauge", fibonacci.Value()),
		},
		"",
	)
}
