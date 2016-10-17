package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Fatalf("Error opening template: %sn", err.Error())
	}
	t.Execute(w, "")
}

func products(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching products..")
	res, err := http.Get(productsAPI)
	if err != nil {
		log.Printf("Error fetching products: %s\n", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body for products: %s\n", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func staticResources(resources []string, mux *http.ServeMux) {
	for _, res := range resources {
		path := fmt.Sprintf("/%s/", res)
		dir := fmt.Sprintf("static/%s", res)
		mux.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
	}
}

var productsAPI string

func main() {
	proxyAddr := os.Getenv("PROXY_ADDR")
	productsAPI = fmt.Sprintf("%s/%s/api/products", proxyAddr, "products")
	log.Printf("using products api: %s\n", productsAPI)

	mux := http.NewServeMux()

	resources := []string{"css", "img", "js", "fonts"}
	staticResources(resources, mux)

	mux.HandleFunc("/", index)
	mux.HandleFunc("/products", products)
	mux.HandleFunc("/health", health)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mux,
	}

	log.Println("Starting server..")
	log.Fatal(server.ListenAndServe(), nil)
}
