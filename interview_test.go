package interview

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestName1(t *testing.T) {
	a := make([]int, 3, 4)
	a[0], a[1], a[2] = 0, 1, 2

	b := append(a, 11)
	b[0] = 7
	// 1

	c := append(a, 22)
	c[0] = 8
	// 2

	d := append(a, 33, 44)
	d[0] = 9
	// 3
	// predict outputs
}

func TestName2(t *testing.T) {
	a := []int{1, 2, 3}

	func(b []int) {
		b = append(b, 4)
	}(a)

	fmt.Println(a)
	// predict output
}

func TestName3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, i := range a {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second)
	// predict output
}

func TestName4(t *testing.T) {
	defer fmt.Println(1)

	defer fmt.Println(2)

	defer fmt.Println(3)

	fmt.Println(4)
	// predict output
}

func TestName5(t *testing.T) {
	fn := func(x int) (y int) {
		defer func() {
			y += 10
		}()

		y = x * 2
		return
	}

	out := fn(3)

	fmt.Println(out)
	// predict output
}

func TestName6(t *testing.T) {
	fn := func() func() int {
		var i int

		return func() int {
			i++
			return i
		}
	}()

	fmt.Println(fn())
	// predict output
}

func TestName7(t *testing.T) {
	var m map[int]bool

	fmt.Println(m[1])

	m[1] = true
	fmt.Println(m[1])
	// predict behavior
}

func TestName8(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	copy(a, a[1:])

	fmt.Println(a)
	// predict output
}

func TestName9(t *testing.T) {
	fn := func() error {
		err := errors.New("one")

		defer func() {
			err = errors.New("two")
		}()

		return err
	}

	fmt.Println(fn())
	// predict output
}
