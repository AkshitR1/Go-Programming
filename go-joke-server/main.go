package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Joke represents a joke structure
type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

// Jokes is a slice of Joke
var Jokes = []Joke{
	{Setup: "Why don't scientists trust atoms?", Punchline: "Because they make up everything!"},
	{Setup: "What do you get if you cross a cat with a dark horse?", Punchline: "Kitty Perry"},
	{Setup: "Why was the math book sad?", Punchline: "Because it had too many problems."},
	{Setup: "What do you call fake spaghetti?", Punchline: "An impasta."},
	{Setup: "Why don’t skeletons fight each other?", Punchline: "They don’t have the guts."},
	{Setup: "Why did the scarecrow win an award?", Punchline: "Because he was outstanding in his field!"},
	{Setup: "How does a penguin build its house?", Punchline: "Igloos it together!"},
	{Setup: "Why don’t skeletons fight each other?", Punchline: "They don’t have the guts."},
	{Setup: "Why couldn't the bicycle stand up by itself?", Punchline: "It was two-tired."},
	{Setup: "How do you organize a space party?", Punchline: "You planet!"},
}

// RandomJokeHandler handles the requests and responds with a random joke in HTML
func RandomJokeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Random Joke Generator</title>
			<style>
				body {
					font-family: 'Arial', sans-serif;
					text-align: center;
					background-color: #f0f0f0;
					padding: 50px;
				}
				.container {
					max-width: 600px;
					margin: 0 auto;
					background-color: #fff;
					padding: 30px;
					border-radius: 10px;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
				}
				h1 {
					font-size: 2em;
					color: #333;
				}
				p {
					font-size: 1.5em;
					color: #666;
				}
				.button {
					display: inline-block;
					padding: 10px 20px;
					font-size: 1em;
					background-color: #4DAA10;
					color: #fdf;
					text-decoration: none;
					border-radius: 8px;
					cursor: pointer;
					transition: background-color 0.3s ease;
				}
				.button:hover {
					background-color: #45a049;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Random Joke</h1>
				<p id="setup"></p>
				<p><strong id="punchline"></strong></p>
				<button class="button" onclick="getJoke()">Get Another Joke</button>
			</div>
			<script>
				async function getJoke() {
					const response = await fetch('/api/joke');
					const joke = await response.json();
					document.getElementById('setup').innerText = joke.setup;
					document.getElementById('punchline').innerText = joke.punchline;
				}
				// Load the first joke on page load
				getJoke();
			</script>
		</body>
		</html>
	`
	fmt.Fprintf(w, html)
}

// ApiJokeHandler handles the requests and responds with a random joke in JSON format
func ApiJokeHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(Jokes))
	randomJoke := Jokes[randomIndex]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(randomJoke)
}

func main() {
	http.HandleFunc("/joke", RandomJokeHandler)
	http.HandleFunc("/api/joke", ApiJokeHandler)
	fmt.Println("Joke server is running on http://localhost:8080/joke")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
