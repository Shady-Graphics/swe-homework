package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type response struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func main() {
	r := chi.NewRouter()

	r.Get("/min", handleMathRequest(handleMin))
	r.Get("/max", handleMathRequest(handleMax))
	r.Get("/avg", handleMathRequest(handleAverage))
	r.Get("/median", handleMathRequest(handleMedian))
	r.Get("/percentile", handleMathRequest(handlePercentile))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}

type mathHandler func(numbers []float64, quantifier int) (interface{}, error)

func handleMathRequest(handler mathHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		numbers, err := parseNumbers(r.URL.Query().Get("numbers"))
		if err != nil {
			http.Error(w, "Invalid numbers", http.StatusBadRequest)
			return
		}

		quantifier, _ := strconv.Atoi(r.URL.Query().Get("quantifier"))

		result, err := handler(numbers, quantifier)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJSONResponse(w, result)
	}
}

func handleMin(numbers []float64, quantifier int) (interface{}, error) {
	count := int(math.Min(float64(quantifier), float64(len(numbers))))
	sortedNumbers := sortNumbers(numbers)
	minimums := sortedNumbers[:count]
	return minimums, nil
}

func handleMax(numbers []float64, quantifier int) (interface{}, error) {
	count := int(math.Min(float64(quantifier), float64(len(numbers))))
	sortedNumbers := sortNumbers(numbers)
	maximums := sortedNumbers[len(sortedNumbers)-count:]
	return maximums, nil
}

func handleAverage(numbers []float64, _ int) (interface{}, error) {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	average := sum / float64(len(numbers))
	return average, nil
}

func handleMedian(numbers []float64, _ int) (interface{}, error) {
	sortedNumbers := sortNumbers(numbers)
	median := calculateMedian(sortedNumbers)
	return median, nil
}

func handlePercentile(numbers []float64, quantifier int) (interface{}, error) {
	if quantifier < 0 || quantifier > 100 {
		return nil, fmt.Errorf("Invalid percentile")
	}

	sortedNumbers := sortNumbers(numbers)
	index := int(math.Ceil((float64(quantifier)/100)*float64(len(sortedNumbers)))) - 1
	percentile := sortedNumbers[index]
	return percentile, nil
}

func sortNumbers(numbers []float64) []float64 {
	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Float64s(sortedNumbers)
	return sortedNumbers
}

func calculateMedian(numbers []float64) float64 {
	length := len(numbers)
	if length%2 == 0 {
		return (numbers[length/2-1] + numbers[length/2]) / 2.0
	}
	return numbers[length/2]
}

func parseNumbers(numbersQuery string) ([]float64, error) {
	numberStrings := strings.Split(numbersQuery, ",")
	numbers := make([]float64, len(numberStrings))

	for i, numStr := range numberStrings {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid numbers")
		}
		numbers[i] = num
	}

	return numbers, nil
}

func writeJSONResponse(w http.ResponseWriter, result interface{}) {
	response := response{
		Message: "Success",
		Result:  result,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
