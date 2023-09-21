package cron

import (
	"testing"
)

func TestParseCronSegment_3378aa11c7(t *testing.T) {
	// Test case 1: Test with valid segment, min, and max
	segment := "1,2,3,4,5"
	min := 1
	max := 5
	expected := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
	}
	result, err := parseCronSegment(segment, min, max)
	if err != nil {
		t.Error("Expected no error, but got: ", err)
	} else if !compareMaps(result, expected) {
		t.Error("Expected: ", expected, " but got: ", result)
	} else {
		t.Log("Test case 1 passed")
	}

	// Test case 2: Test with invalid segment, min, and max
	segment = "1,2,3,4,6"
	min = 1
	max = 5
	_, err = parseCronSegment(segment, min, max)
	if err == nil {
		t.Error("Expected an error, but got none")
	} else {
		t.Log("Test case 2 passed")
	}

	// Test case 3: Test with valid segment, min, and max but with range
	segment = "1-3"
	min = 1
	max = 5
	expected = map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}
	result, err = parseCronSegment(segment, min, max)
	if err != nil {
		t.Error("Expected no error, but got: ", err)
	} else if !compareMaps(result, expected) {
		t.Error("Expected: ", expected, " but got: ", result)
	} else {
		t.Log("Test case 3 passed")
	}
}

// compareMaps is a helper function to compare two maps
func compareMaps(map1, map2 map[int]struct{}) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key := range map1 {
		if _, ok := map2[key]; !ok {
			return false
		}
	}
	return true
}
