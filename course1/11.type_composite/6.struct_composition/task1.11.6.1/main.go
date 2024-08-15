package main

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
}

func (order *Order) RemoveDish(dish Dish) {
	for i, curDish := range order.Dishes {
		if curDish == dish {
			order.Dishes[i] = order.Dishes[len(order.Dishes)-1]
			order.Dishes = order.Dishes[:len(order.Dishes)-1]
		}
	}
}

func (order *Order) CalculateTotal() {
	sum := 0.0
	for _, dish := range order.Dishes {
		sum += dish.Price
	}
	order.Total = sum
}

func main() {
}
