package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if _, err := w.Write([]byte(`
			<!DOCTYPE html>
			<html>
			<head>
				<meta charset="utf-8">
				<title>Random GIF Activity</title>
				<style>
					body {
						background-color: #dff7ff;
						font-family: sans-serif;
						text-align: center;
						padding: 2rem;
					}
					h1 {
						color: #1f3b5b;
					}
				</style>
			</head>
			<body>
				<h1>Hello, world!</h1>
				<img src="https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CatPawPsychadelia.gif" alt="Vibe Cat">
			</body>
			</html>
		`)); err != nil {
			log.Printf("write response: %v", err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Open http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
