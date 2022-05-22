package problem6056

import (
	"fmt"
	"testing"
)

func Test_largestGoodInteger(t *testing.T) {
	fmt.Println(largestGoodInteger("6777133339999999999"))
	fmt.Println(largestGoodInteger("677713333000000"))
	fmt.Println(largestGoodInteger("2300019111"))
	fmt.Println(largestGoodInteger("42352338413857329479813"))
	fmt.Println(largestGoodInteger(""))
	fmt.Println(largestGoodInteger("000"))
}
