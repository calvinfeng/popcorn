package lowrank

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

// RandVector creates a random float64 value of K length.
func RandVector(K int) []float64 {
	vector := make([]float64, K)
	for k := 0; k < K; k++ {
		vector[k] = rand.Float64()
	}

	return vector
}

// DotProduct computes the dot product of two vector. The function will return an error if the two
// vectors are not the same length.
func DotProduct(vector1 []float64, vector2 []float64) (float64, error) {
	var sum float64
	if len(vector1) != len(vector2) {
		return sum, errors.New("dimension mismatch")
	}

	for i := 0; i < len(vector1); i++ {
		sum += vector1[i] * vector2[i]
	}

	return sum, nil
}

func randMat(row, col int) *mat.Dense {
	rand.Seed(time.Now().UTC().Unix())

	randFloats := []float64{}
	for i := 0; i < row*col; i++ {
		randFloats = append(randFloats, rand.Float64())
	}

	return mat.NewDense(row, col, randFloats)
}

func average(list []float64) float64 {
	sum := 0.0

	if len(list) == 0 {
		return sum
	}

	for _, el := range list {
		sum += el
	}

	return sum / float64(len(list))
}

func absMax(M *mat.Dense) float64 {
	I, J := M.Dims()
	max := 0.0
	for i := 0; i < I; i++ {
		for j := 0; j < J; j++ {
			if math.Abs(M.At(i, j)) > max {
				max = math.Abs(M.At(i, j))
			}
		}
	}

	return max
}

func absAverage(M *mat.Dense) float64 {
	I, J := M.Dims()
	sum := 0.0
	for i := 0; i < I; i++ {
		for j := 0; j < J; j++ {
			sum += math.Abs(M.At(i, j))
		}
	}

	return sum / float64(I*J)
}
