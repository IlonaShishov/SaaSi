package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/RHEcosystemAppEng/SaaSi/deployer/pkg/config"
	"gopkg.in/yaml.v2"
)

type Numbers struct {
	Num1 float64 `yaml:"num1"`
	Num2 float64 `yaml:"num2"`
}

type Result struct {
	Sum float64 `yaml:"sum"`
}

func addNumbers(w http.ResponseWriter, r *http.Request) {
	// Read the request body into a byte slice
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the YAML data into a Numbers struct
	var numbers Numbers
	err = yaml.Unmarshal(body, &numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate the sum of the two numbers
	sum := numbers.Num1 + numbers.Num2

	// Return the sum as a YAML response
	result := Result{Sum: sum}
	w.Header().Set("Content-Type", "application/x-yaml")
	w.WriteHeader(http.StatusOK)
	err = yaml.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleRequests() {
	http.HandleFunc("/add", config.ReadRestRequest)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
