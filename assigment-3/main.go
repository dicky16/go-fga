package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path"
)

type Sensor struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// render html file
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		statusWind := "AMAN"
		statusWater := "AMAN"
		water := rand.Intn(100)
		wind := rand.Intn(100)
		var data = map[string]interface{}{
			"water": water,
			"wind":  wind,
		}
		if water >= 6 && water <= 8 {
			statusWater = "Siaga"
		} else if water > 8 {
			statusWater = "Bahaya"
		}
		if wind >= 6 && wind <= 8 {
			statusWind = "Siaga"
		} else if wind > 8 {
			statusWind = "Bahaya"
		}
		data["statusWater"] = statusWater
		data["statusWind"] = statusWind
		file, _ := json.MarshalIndent(data, "", " ")

		_ = ioutil.WriteFile("data/sensor.json", file, 0644)
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
