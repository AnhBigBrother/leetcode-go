package main

type ProductOfNumbers struct {
	product   []int
	countZero []int
}

func NewProductOfNumbers() ProductOfNumbers {
	return ProductOfNumbers{
		product:   []int{},
		countZero: []int{},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if len(this.product) == 0 {
		if num == 0 {
			this.countZero = append(this.countZero, 1)
			this.product = append(this.product, 1)
		} else {
			this.countZero = append(this.countZero, 0)
			this.product = append(this.product, num)
		}
	} else {
		if num == 0 {
			this.countZero = append(this.countZero, this.countZero[len(this.countZero)-1]+1)
			this.product = append(this.product, 1)
		} else {
			this.countZero = append(this.countZero, this.countZero[len(this.countZero)-1])
			this.product = append(this.product, this.product[len(this.product)-1]*num)
		}
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k == len(this.product) {
		if this.countZero[len(this.countZero)-1] > 0 {
			return 0
		}
		return this.product[len(this.product)-1]
	}
	if this.countZero[len(this.countZero)-1]-this.countZero[len(this.countZero)-1-k] > 0 {
		return 0
	}
	if this.product[len(this.product)-1-k] == 0 {
		return -1
	}
	return this.product[len(this.product)-1] / this.product[len(this.product)-1-k]
}
