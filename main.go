package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	store       *URLStore
	port        = flag.String("port", ":4000", "port to run program")
	storageFile = flag.String("file", "urls.json", "file to store the saved urls")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	store = NewURLStore(*storageFile)

	http.HandleFunc("/add", makeHandler(Add))
	http.HandleFunc("/", makeHandler(Redirect))
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}

func makeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Runtime Panic: %v\n", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		fn(w, req)
	}
}
func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if url := r.FormValue("url"); url == "" {
		fmt.Fprint(w, AddForm)
	} else {
		val := store.Put(url)
		fmt.Fprintf(w, "Short URL: %s", val)
	}

	return

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Path[1:]

	if url := store.Get(val); url != "" {
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
	http.NotFound(w, r)
}

const AddForm = `<html><body>
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form></body></html>`
