package dynamic

import (
	"fmt"
	"testing"
)

func TestGenerateShortId(t *testing.T) {
	id := GenerateShortId()
	fmt.Println(id)
}

func TestGenerateCustomerCode(t *testing.T) {
	id := GenerateCustomerCode()
	fmt.Println(id)
}
