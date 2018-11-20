package main

import "io/ioutil"
import "log"
import "net/http"
import "os"

type viewHandler struct{}

func (vh *viewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	println(r.URL.Path)

	data, err := ioutil.ReadFile(string(path))

	if err != nil {
		log.Printf("Error with path %s: %v", path, err)
		w.WriteHeader(404)
		w.Write([]byte("404"))
	}

	// if strings.HasSuffix(path, ".html") {
	// 	w.Header().Add("Content-Type", "text/html")
	// } else if strings.HasSuffix(path, ".mp4") {
	// 	w.Header().Add("Content-Type", "video/mp4")
	// }

	contentType := http.DetectContentType(data)
	w.Header().Set("Content-type", contentType)

	w.Write(data)
}

func main() {
	// println(os.Getenv("FSDirectory"))
	// http.Handle("/", new(viewHandler))
	http.Handle("/", http.FileServer(http.Dir(os.Getenv("FSDirectory"))))
	http.ListenAndServe(":8080", nil)
}
