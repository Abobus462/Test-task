package main

import (
	"fmt"
	"math/rand"
	"proj/functions"
	"time"
)

func main() {

	A := [][]float64{
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1},
		{1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1},
		{1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1},
		{1, 1, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 1},
		{1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1},
		{1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1},
		{1, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1},
		{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1},
		{1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1},
		{1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
	}

	for i := range 101 {
		fmt.Printf("##### TEST %d #####\n", i)
		len_data := rand.Intn(11) + 1
		fmt.Printf("Len_data: %d\n", len_data)

		data := functions.Generate_data(len_data)
		fmt.Printf("Data: %v\n", data)

		G := functions.Create_generator_matrix(A, len_data)

		code_message := functions.Coder(G, data)

		fmt.Printf("Code data: %.2f\n", code_message)

		sigma := float64(i) / 10

		fmt.Printf("Level of noise: %v\n", sigma)

		noisy_code := functions.AddNoise(code_message, sigma)
		fmt.Printf("Noisy code: %.2f\n", noisy_code)

		start := time.Now()

		decode_message := functions.Decoder(G, noisy_code)

		end := time.Since(start)

		fmt.Printf("TIME: %s\n", end)

		fmt.Printf("Decoding message: %v\n", decode_message)

		coincidence := 0

		for index := range data {
			if data[index] == decode_message[index] {
				coincidence++
			}
		}

		accuracy := float64(coincidence) / float64(len(data)) * 100

		fmt.Printf("Accuracy: %.2f%%\n\n\n\n", accuracy)
	}

}
