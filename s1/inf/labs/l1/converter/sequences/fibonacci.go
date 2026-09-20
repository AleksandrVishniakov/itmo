package sequences

type Fibonacci struct {
	nums []int
}

func NewFibonacci() *Fibonacci {
	return &Fibonacci{
		nums: []int{1, 1},
	}
}

func (f *Fibonacci) Get(i int) int {
	for j := len(f.nums); j <= i; j++ {
		f.nums = append(f.nums, f.nums[j-1]+f.nums[j-2])
	}

	return f.nums[i]
}
