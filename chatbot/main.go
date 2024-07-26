package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jdkato/prose/v2"
	"github.com/tealeg/xlsx"
)

// Chatbot struct to store chatbot data
type Chatbot struct {
	Name        string
	Responses   map[string]string
	Questions   map[string]string
	ExcelFile   *xlsx.File
	Concurrency int
	mutex       sync.Mutex
}

// NewChatbot creates a new chatbot instance
func NewChatbot() *Chatbot {
	return &Chatbot{
		Name:        "Chatbot",
		Responses:   make(map[string]string),
		Questions:   make(map[string]string),
		Concurrency: 5,
	}
}

// StartChatbot starts the chatbot server
func (c *Chatbot) StartChatbot() {
	router := mux.NewRouter()
	router.HandleFunc("/chat", c.handleChat).Methods("POST")
	go func() {
		log.Fatal(http.ListenAndServe(":8080", router))
	}()
	log.Println("Chatbot started on port 8080")
}

// handleChat handles incoming chat requests
func (c *Chatbot) handleChat(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Message string `json:"message"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Handle concurrency
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Process the input message using NLP
	doc, err := prose.NewDocument(input.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	intent := classifyIntent(doc)

	// Handle the intent
	switch intent {
	case "greeting":
		w.Write([]byte(c.Responses["greeting"]))
	case "question":
		// Use Excel file to find matching question and response
		for question, response := range c.Questions {
			if strings.Contains(input.Message, question) {
				w.Write([]byte(response))
				return
			}
		}
		w.Write([]byte("I didn't understand that. Can you try again?"))
	default:
		w.Write([]byte("I didn't understand that. Can you try again?"))
	}
}

// classifyIntent classifies the intent of the message using tokenization
func classifyIntent(doc *prose.Document) string {
	for _, tok := range doc.Tokens() {
		if strings.ToLower(tok.Text) == "hello" || strings.ToLower(tok.Text) == "hi" || strings.ToLower(tok.Text) == "hey" {
			return "greeting"
		} else if strings.Contains(strings.ToLower(tok.Text), "name") || strings.Contains(strings.ToLower(tok.Text), "weather") {
			return "question"
		}
	}
	return "default"
}

// LoadResponses loads chatbot responses from an Excel file
func (c *Chatbot) LoadResponses(filePath string) error {
	var err error
	c.ExcelFile, err = xlsx.OpenFile(filePath)
	if err != nil {
		return err
	}
	sheet := c.ExcelFile.Sheets[0]
	for _, row := range sheet.Rows {
		if len(row.Cells) < 2 {
			continue
		}
		c.Responses[row.Cells[0].String()] = row.Cells[1].String()
	}
	return nil
}

// LoadQuestions loads chatbot questions from an Excel file
func (c *Chatbot) LoadQuestions(filePath string) error {
	var err error
	c.ExcelFile, err = xlsx.OpenFile(filePath)
	if err != nil {
		return err
	}
	sheet := c.ExcelFile.Sheets[0]
	for _, row := range sheet.Rows {
		if len(row.Cells) < 2 {
			continue
		}
		c.Questions[row.Cells[0].String()] = row.Cells[1].String()
	}
	return nil
}

func main() {
	c := NewChatbot()

	// Load responses and questions from Excel files
	err := c.LoadResponses("responses.xlsx")
	if err != nil {
		log.Fatalf("Failed to load responses: %v", err)
	}

	err = c.LoadQuestions("questions.xlsx")
	if err != nil {
		log.Fatalf("Failed to load questions: %v", err)
	}

	// Set default responses
	c.Responses["greeting"] = "Hello! How can I help you today?"
	c.Responses["default"] = "I didn't understand that. Can you try again?"

	// Start the chatbot server
	c.StartChatbot()

	// Keep the server running
	select {}
}
