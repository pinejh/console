package console

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Log(print ...interface{}) {
	for _, n := range print {
		fmt.Print(n)
	}
	fmt.Print("\n")
}

func RandInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min+1) + min
}

func newSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func SquareInt(num int) int {
	return int(math.Pow(float64(num), 2.0))
}

func IsPrime(num int) bool {
	is := true
	for i := 0; float64(i) <= math.Sqrt(float64(num)); i++ {
		if i > 1 && num%i == 0 {
			is = false
		}
	}
	return is
}

func IsDivBy(num, div int) bool {
	if div != 0 && num != 0 && num%div == 0 {
		return true
	}
	return false
}

type vector struct {
	number int
	val    int
}

func DivCount(num int) int {
	n := num
	d := 2
	factors := []vector{}
	for d <= n {
		if IsDivBy(n, d) {
			n /= d
			if len(factors) > 0 && factors[len(factors)-1].number == d {
				factors[len(factors)-1].val++
			} else {
				factors = append(factors, vector{d, 1})
			}
		} else {
			if d == 2 {
				d++
			} else {
				d += 2
			}
		}
	}
	return calcFactors(factors)
}

func calcFactors(v []vector) int {
	factors := 1
	for _, p := range v {
		factors *= p.val + 1
	}
	return factors
}

func Repeat(f func(), delay time.Duration) {
	go func() {
		for {
			f()
			time.Sleep(delay)
		}
	}()
}
