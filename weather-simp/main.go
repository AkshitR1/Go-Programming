package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const apiKey = "e28bd1b534d3174e92afa77cada6e08a"

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather(city string) (WeatherResponse, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherResponse{}, fmt.Errorf("failed to get weather: %s", resp.Status)
	}

	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return WeatherResponse{}, err
	}

	return weatherResponse, nil
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		city := r.FormValue("city")
		if city == "" {
			http.Error(w, "City name is required", http.StatusBadRequest)
			return
		}

		weather, err := getWeather(city)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get weather: %v", err), http.StatusInternalServerError)
			log.Printf("Error fetching weather for city %s: %v", city, err)
			return
		}

		tmpl := template.Must(template.New("weather").Parse(`
			<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Weather in {{.Name}}</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" integrity="sha384-JcKb8q3iqJ61gNV9KGb8thSsNjpSL0n8PARn9HuZOnIxN0hoP+VmmDGMN5t9UJ0Z" crossorigin="anonymous">
    <style>
        body {
            background-color: #f8f9fa;
        }
        .card {
            border: none;
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        }
        .card-body {
            padding: 2rem;
        }
        .card-title {
            text-align: center;
            margin-bottom: 2rem;
        }
        .card-text {
            font-size: 1.2rem;
            margin-bottom: 1rem;
        }
    </style>
</head>
<body>
    <div class="container mt-5">
        <div class="card">
            <div class="card-body">
                <h2 class="card-title">Weather in {{.Name}}</h2>
                {{range .Weather}}
                    <p class="card-text">Description: {{.Description}}</p>
                {{end}}
                <p class="card-text">Temperature: {{.Main.Temp}}°C</p>
                <p class="card-text">Humidity: {{.Main.Humidity}}%</p>
                <p class="card-text">Wind Speed: {{.Wind.Speed}} m/s</p>
            </div>
        </div>
    </div>
</body>
</html>
		`))

		if err := tmpl.Execute(w, weather); err != nil {
			http.Error(w, fmt.Sprintf("Failed to execute template: %v", err), http.StatusInternalServerError)
			log.Printf("Error executing template for city %s: %v", city, err)
			return
		}
	} else if r.Method == http.MethodGet {
		// Display the form to enter the city
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintln(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Weather Search</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" integrity="sha384-JcKb8q3iqJ61gNV9KGb8thSsNjpSL0n8PARn9HuZOnIxN0hoP+VmmDGMN5t9UJ0Z" crossorigin="anonymous">
    <style>
        body {
            background-color: #f8f9fa;
        }
        .card {
            border: none;
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        }
        .card-body {
            padding: 2rem;
        }
        .form-group {
            margin-bottom: 1rem;
        }
        .btn-primary {
            background-color: #007bff;
            border-color: #007bff;
        }
        .btn-primary:hover {
            background-color: #0069d9;
            border-color: #0062cc;
        }
    </style>
</head>
<body>
        <div class="container mt-5">
        <h1 class="mt-5 text-center">Weather App</h1>
        <div class="row justify-content-center">
            <div class="col-md-6">
                <form method="post" class="mt-3">
                    <div class="input-group mb-3">
                        <input type="text" id="city" name="city" class="form-control" placeholder="Enter city name" required>
                        <div class="input-group-append">
                            <button style="background-color: #FFA500;" class="btn btn-primary" type="submit">Get Weather</button>
                        </div>
                    </div>
                </form>
        </div>
    </div>
</body>
</html>
		`)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/weather", weatherHandler)

	fmt.Println("Server is running on http://localhost:8080/weather")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
