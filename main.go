package main

import (
	"fmt"
	"github.com/bedrockmyshit/backend/app"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var dev bool

	args := os.Args
	if len(args) >= 2 && strings.EqualFold(args[1], "dev") {
		dev = true
	}

	addr := ":8080"
	if !dev {
		addr = ":443"

		go func() {
			err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("Redirecting to https")
				http.Redirect(w, r, "https://decryptmypack.com", http.StatusMovedPermanently)
			}))
			if err != nil {
				panic(err)
			}
		}()
	}
	log.Printf("Backend listening on %s\n", addr)

	a := app.App{}
	err := a.ListenAndServe(addr, dev)
	if err != nil {
		panic(err)
	}
}
