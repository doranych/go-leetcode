package product_of_the_last_k_numbers

type ProductOfNumbers struct {
	prefix []int
	size   int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		prefix: []int{1},
		size:   0,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefix = []int{1}
		this.size = 0
		return
	}
	this.prefix = append(this.prefix, num*this.prefix[this.size])
	this.size++
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k > this.size {
		return 0
	}
	return this.prefix[this.size] / this.prefix[this.size-k]
}

/* naive solution
type ProductOfNumbers struct {
	slice []int
}


func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		slice: make([]int, 0),
	}
}


func (this *ProductOfNumbers) Add(num int)  {
	this.slice = append(this.slice, num)
}


func (this *ProductOfNumbers) GetProduct(k int) int {
	prod := 1
	for i:= len(this.slice)-1; i>=len(this.slice)-k; i-- {
		prod *= this.slice[i]
	}
	return prod
}*/
