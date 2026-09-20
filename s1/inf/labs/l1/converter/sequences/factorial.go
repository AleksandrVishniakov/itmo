package sequences

type Factorial struct {
	nums []int
}

func NewFactorial() *Factorial {
	return &Factorial{
		nums: []int{1, 1},
	}
}

func (f *Factorial) Get(i int) int {
	for j := len(f.nums); j <= i; j++ {
		f.nums = append(f.nums, j*f.nums[j-1])
	}

	return f.nums[i]
}
