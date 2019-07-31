package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(mainLogicHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8000", nil)
}
