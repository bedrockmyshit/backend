package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router = mux.NewRouter()
)

type App struct {
}

func (a *App) ListenAndServe(addr string, dev bool) error {
	router.HandleFunc("/download", a.download).Queries("target", "{target}")
	router.HandleFunc("/assets/{path}", serveDirFunc("./assets"))

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	if dev {
		return http.ListenAndServe(addr, router)
	}
	return http.ListenAndServeTLS(addr, "./certificate.crt", "./private.key", router)
}

// serverDirFunc serves files from the given directory.
func serveDirFunc(dir string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, dir+r.URL.Path)
	}
}

func serveFileFunc(name string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, name)
	}
}
