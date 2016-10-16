package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float32 `json:"price"`
	Category     string  `json:"category"`
	ThumbnailURL string  `json:"thumbnail_url"`
	Description  string  `json:"description"`
	Stars        int     `json:"stars"`
	Reviews      int     `json:"reviews"`
}

// INSERT INTO PRODUCT (ID, NAME, PRICE, CATEGORY, THUMBNAIL_URL, DESCRIPTION, STARS, REVIEWS) VALUES (1, 'Breaking Bad S2', 19.99, 'DVD',  'https://fanart.tv/fanart/tv/81189/tvthumb/B_81189%20(6).jpg', 'Walter White, a struggling high school chemistry teacher, is diagnosed with advanced lung cancer. ', 4, 10);

func products(w http.ResponseWriter, r *http.Request) {
	p := []Product{
		{1, "Breaking Bad S2", 19.99, "DVD", "https://fanart.tv/fanart/tv/81189/tvthumb/B_81189%20(6).jpg", "Walter White, a struggling high school chemistry teacher, is diagnosed with advanced lung cancer. ", 4, 10},
		{2, "Mr. Robot S1", 21.99, "DVD", "https://fanart.tv/fanart/tv/289590/tvthumb/mr-robot-556beff62203d.jpg", "The show follows Elliot, who is a cyber-security tech by day and vigilante hacker by night.", 4, 5},
		{3, "Fight Club", 16.99, "DVD", "https://fanart.tv/detailpreview/fanart/movies/550/moviethumb/fight-club-51b0f879f12e2.jpg", "A ticking-time-bomb insomniac and a slippery soap salesman channel primal male aggression into a shocking new form of therapy.", 5, 2},
		{4, "Dexter S4", 15.99, "DVD", "https://fanart.tv/fanart/tv/79349/tvthumb/D_79349%20(10).jpg", "He's Dexter Morgan, everyone's favorite serial killer. As a Miami forensics expert, he spends his days solving crimes, and nights committing them.", 3, 0},
		{5, "True Detective S1", 16.99, "DVD", "https://fanart.tv/fanart/tv/270633/tvthumb/true-detective-52d09d848dbe5.jpg", "An anthology series in which police investigations unearth the personal and professional secrets of those involved, both within and outside the law.", 5, 18},
	}
	json.NewEncoder(w).Encode(p)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Fatalf("Error opening template: %sn", err.Error())
	}
	t.Execute(w, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	mux.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/api/products", products)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe(), nil)
}
