// Knapsack problem:
// You have a knapsack to carry products for selling.
// It holds up to a certain weight, not enough for carrying all your products-
// you must choose which ones to carry.
// Knowing the weight and sales value of each product, which choice of products gives the highest revenue?

package algorithms

type product struct {
	Name   string
	Weight float64
	Value  float64
}

type products []product

func (ps products) TotalWeight() float64 {
	var total float64
	for _, p := range ps {
		total += p.Weight
	}
	return total
}

func (ps products) SalesValue() float64 {
	var total float64
	for _, p := range ps {
		total += p.Value
	}
	return total
}

func Add(ps []products, e product) {
	for i := range ps {
		ps[i] = append(ps[i], e)
	}
}

func BruteForceKnapsack(items []product, maxWeight float64) ([]product, float64) {
	var result []product
	bestValue := 0.0
	for _, candidate := range PowerSet(items) {
		got := candidate.SalesValue()
		if candidate.TotalWeight() <= maxWeight && got > bestValue {
			bestValue = got
			result = candidate
		}
	}

	return result, bestValue
}

// PowerSet returns power set of ps
func PowerSet(items []product) []products {
	var result = make([]products, 0, 2^len(items))

	// add a empty set
	result = append(result, products{})

	// keeps track of the next product to consider
	for _, e := range items {
		// duplicate the result we already have
		var copyResult = make([]products, len(result))
		copy(copyResult, result)

		// adding the current product to the duplicates
		// Bug: for _, r := range result
		for i := range result {
			result[i] = append(result[i], e)
		}
		result = append(result, copyResult...)
	}

	return result
}

// RecursivePowerSet returns power set of ps
func RecursivePowerSet(items []product) []products {
	var result []products
	// TODO

	return result
}
