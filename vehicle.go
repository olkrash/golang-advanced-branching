package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type vehicle interface {
}

type car struct {
	model       string
	make        string
	typeVehicle string
}

type bike struct {
	model string
	make  string
}

type truck struct {
	model       string
	make        string
	typeVehicle string
}

// Values array for the feedback.json file

type Values struct {
	Models []Model `json:"values"`
}

// Model array for the feedback.json file
type Model struct {
	Name     string
	Feedback []string `json:"feedback"`
}

type feedbackResult struct {
	feedbackTotal    int
	feedbackNegative int
	feedbackPositive int
	feedbackNeutral  int
}

type rating float32

var vehicleResult map[string]feedbackResult

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
)

var inventory []vehicle

func init() {

	inventory = []vehicle{
		bike{"FTR 1200", "Indian"},
		bike{"Iron 1200", "Harley"},
		car{"Sonata", "Hyundai", "Sedan"},
		car{"SantaFe", "Hyundai", "SUV"},
		car{"Civic", "Honda", "Hatchback"},
		car{"A5", "Audi", "Coupe"},
		car{"Mazda6", "Mazda", "Sedan"},
		car{"CRV", "Honda", "SUV"},
		car{"Camry", "Toyota", "Sedan"},
		truck{"F-150", "Ford", "Truck"},
		truck{"RAM1500", "Dodge", "Truck"}}

	vehicleResult = make(map[string]feedbackResult)

}

func main() {

	// Generate ratings for the different vehicles
	generateRating()

	// Print ratings for the different vehicles
}

func readJSONFile() Values {
	jsonFile, err := os.Open("feedback.json")

	if err != nil {
		log.Fatal("File not found")
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var content Values
	err = json.Unmarshal(byteValue, &content)
	if err != nil {
		return Values{}
	}

	return content
}

func generateRating() {
	f := readJSONFile()

	for _, v := range f.Models {
		var vehResult feedbackResult
		var vehRating rating
		vehResult = vehicleResult[v.Name]

		for _, msg := range v.Feedback {
			text := strings.Split(msg, " ")
			if len(text) >= 5 {
				return
			}
			vehRating = 5.0
			vehResult.feedbackTotal++
			for _, word := range text {
				s := strings.Trim(strings.ToLower(word), " ,.,!,?,\t,\n,\r")
				switch s {
				case "pleasure", "impressed", "wonderful", "fantastic", "splendid":
					vehRating += extraPositive
				case "help", "helpful", "thanks", "thank you", "happy":
					vehRating += positive
				case "not helpful", "sad", "angry", "improve", "annoy":
					vehRating += negative
				case "pathetic", "bad", "worse", "unfortunately", "agitated", "frustrated":
					vehRating += extraNegative
				}

			}
		}
	}

}
