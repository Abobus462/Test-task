package functions

import (
	"math"
	"math/rand"
)

func Reverse(number []float64) []float64 {
	result := []float64{}
	for i := len(number) - 1; i >= 0; i-- {
		result = append(result, number[i])
	}
	return result
}

func Bin(number int) []float64 {
	result := []float64{}
	if number == 0 {
		return []float64{0}
	}
	for number > 0 {
		result = append(result, float64(number%2))
		number /= 2
	}
	return Reverse(result)
}

func Create_generator_matrix(A [][]float64, count_cols int) [][]float64 {
	result := make([][]float64, len(A))
	for i := range A {
		result[i] = make([]float64, count_cols)
		copy(result[i], A[i][:count_cols])
	}
	return result
}

func Generate_data(length int) []float64 {
	result := make([]float64, length)
	for i := range result {
		result[i] = float64(rand.Intn(2))
	}
	return result
}

func Dot_product(vector1 []float64, vector2 []float64) float64 {
	sum := 0.0
	for i := range vector1 {
		sum += vector1[i] * vector2[i]
	}
	return sum
}

func Add_leading_zeros(bin []float64, k int) []float64 {
	if len(bin) >= k {
		return bin
	}
	result := make([]float64, k)
	copy(result[k-len(bin):], bin)
	return result
}

func Coder(g_matrix [][]float64, data []float64) []float64 {
	g_rows := len(g_matrix)
	g_cols := len(g_matrix[0])
	result := make([]float64, g_rows)
	for i := 0; i < g_rows; i++ {
		sum := 0.0
		for j := 0; j < g_cols; j++ {
			sum += g_matrix[i][j] * data[j]
		}
		result[i] = math.Mod(sum, 2)
	}
	return result
}

func Decoder(g_matrix [][]float64, code []float64) []float64 {
	k := len(g_matrix[0])
	best_d := make([]float64, k)
	max_score := -1.0
	for i := 0; i < int(math.Pow(2, float64(k))); i++ {
		d := Bin(i)
		d = Add_leading_zeros(d, k)
		c_i := Coder(g_matrix, d)
		score := Dot_product(code, c_i)
		if score > max_score {
			max_score = score
			best_d = d
		}
	}
	return best_d
}

func GenerateNoise(length int) []float64 {
	noise := make([]float64, length)
	for i := range noise {
		noise[i] = rand.NormFloat64()
	}
	return noise
}

func AddNoise(codeword []float64, sigma float64) []float64 {
	noise := GenerateNoise(len(codeword))
	noisyCodeword := make([]float64, len(codeword))
	for i := range codeword {
		noisyCodeword[i] = codeword[i] + sigma*noise[i]
	}
	return noisyCodeword
}
