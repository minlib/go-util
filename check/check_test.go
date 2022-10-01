package check

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestCheckMobile(t *testing.T) {
	fmt.Println(CheckMobile("13422221990")) // ture
	fmt.Println(CheckMobile("19422221990")) // ture

	fmt.Println(CheckMobile("1342222199"))   // false
	fmt.Println(CheckMobile("03422221990"))  // false
	fmt.Println(CheckMobile("12222219901"))  // false
	fmt.Println(CheckMobile("134222219912")) // false
}

func TestCheckIdCard(t *testing.T) {
	fmt.Println(CheckIdCard("440421200001015333")) // true
	fmt.Println(CheckIdCard("44042120000101533X")) // true
	fmt.Println(CheckIdCard("44042120000101533x")) // true

	fmt.Println(CheckIdCard("44042120000101533"))   // false
	fmt.Println(CheckIdCard("4404212X0001015333"))  // false
	fmt.Println(CheckIdCard("4404212000010153331")) // false
	fmt.Println(CheckIdCard("x40421200001015333"))  // false
	fmt.Println(CheckIdCard("040441200009215"))     // false
	fmt.Println(CheckIdCard("040421200001015333"))  // false
}

func TestCheckUserName(t *testing.T) {
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int31())
	fmt.Println(rand.Intn(100))
}
