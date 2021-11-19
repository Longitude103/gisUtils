package gisUtils

import (
	"testing"
)

func Test_InverseDW(t *testing.T) {
	testDistances := []float64{20214.0, 33821.0, 39704.0}

	weights, err := InverseDW(testDistances)
	if err != nil {
		t.Fatalf("Function Error %s", err)
	}

	// fmt.Printf("Weights: %v\n", weights)

	if weights[0] != 0.619 {
		t.Errorf("Weight 1 wrong: want 0.619, got: %v", weights[0])
	}

	if weights[1] != 0.221 {
		t.Errorf("Weight 2 wrong: want 0.221, got: %v", weights[1])
	}

	if weights[2] != 0.16 {
		t.Errorf("Weight 3 wrong: want 0.16, got: %v", weights[2])
	}

}
