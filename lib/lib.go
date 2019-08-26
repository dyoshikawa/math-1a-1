package lib

import (
	"math"
)

// DEPRECATED
func CalcLcm(a int, b int) int {
	x := a * b

	if a < b {
		tmp := a
		a = b
		b = tmp
	}

	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return x / b
}

func Lcm(a int, b int) int {
	x := a * b

	if a < b {
		tmp := a
		a = b
		b = tmp
	}

	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return x / b
}

func Soe(n int) []int {
	nums := []int{}
	for i := 2; i <= n; i++ {
		nums = append(nums, i)
	}
	pnums := []int{}
	sqrtVal := int(math.Sqrt(float64(n)))
	for true {
		if sqrtVal <= nums[0] {
			for _, v := range nums {
				pnums = append(pnums, v)
			}
			break
		}
		pnums = append(pnums, nums[0])
		newNums := []int{}
		for _, v := range nums {
			if v%nums[0] > 0 {
				newNums = append(newNums, v)
			}
		}
		nums = newNums
	}
	return pnums
}

func Pf(n int) []int {
	pNums := Soe(n)
	pNumCnt := len(pNums)
	flg := true
	divs := []int{}
	for n > 1 && flg {
		for k, v := range pNums {
			if n%v == 0 {
				divs = append(divs, v)
				n = n / v
				break
			}
			if k == pNumCnt-1 {
				flg = false
			}
		}
	}
	return divs
}

func PowInt(n int, x int) int {
	return int(math.Pow(float64(n), float64(x)))
}

func Include(list []interface{}, el interface{}) bool {
	for _, v := range list {
		if el == v {
			return true
		}
	}
	return false
}

func CombCnt(all int, selected int) int {
	a := 1
	for i := all; i > all-selected; i-- {
		a *= i
	}
	b := 1
	for i := selected; i > 0; i-- {
		b *= i
	}
	return a / b
}
