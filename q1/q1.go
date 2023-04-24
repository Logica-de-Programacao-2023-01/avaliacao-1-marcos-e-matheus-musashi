package q1

import "fmt"

func DivideWatermelon(weight int) (bool, error) {
	if weight <= 0 {
		return false, fmt.Errorf("Peso invalido.")
	}
	if weight == 2 {
		return false, nil
	} else if (weight % 2) == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
