package util

import (
	"fmt"
	"testing"
)

func TestUtil(t *testing.T) {

	right := ToWei("1.089", 18)
	right1 := ToWei("1", 18)
	wrong := ToWei(1, 18)
	fmt.Println(right)
	fmt.Println(right1)

	fmt.Printf("right: %d     right1: %d  \n", len(right.String()), len(right1.String()))
	fmt.Println("----------------------wrong ⬇")
	fmt.Println(wrong)

}
