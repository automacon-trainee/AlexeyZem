package main

import (
	"reflect"
	"testing"
)

func TestPractice(t *testing.T) {
	product := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	product.ProductID = ProductCocaCola
	product.ProductID = ProductPepsi
	product.ProductID = ProductSprite

	statProfit := NewStatisticProfit(
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDifferenceProfit,
		WithAllData,
	).(*StatisticProfit)

	statProfit.SetProduct(product)
	expectedData := make([]float64, 0, 4)

	if statProfit.product.ProductID != ProductSprite {
		t.Errorf("productId is not updated")
	}

	res := statProfit.Sum(statProfit.product.Buys)
	if res != 45 {
		t.Errorf("Error in Sum: Expect 45 but got %f", res)
	}

	res = statProfit.Average(statProfit.product.Sells)
	if res != 20 {
		t.Errorf("Error in Average: Expect 20 but got %f", res)
	}

	res = statProfit.getAverageProfit()
	expected := statProfit.Average(statProfit.product.Sells) - statProfit.Average(statProfit.product.Buys)
	if res != expected {
		t.Errorf("Error in getAverageProfit: Expect %v but got %f", expected, res)
	}
	expectedData = append(expectedData, expected)

	res = statProfit.getAverageProfitPercent()
	expected = 100 * (statProfit.Average(statProfit.product.Sells) - statProfit.Average(statProfit.product.Buys)) / statProfit.Average(statProfit.product.Buys)
	if res != expected {
		t.Errorf("Error in getAverageProfitPercent: Expect %v but got %f", expected, res)
	}
	expectedData = append(expectedData, expected)

	res = statProfit.getCurrentProfit()
	expected = statProfit.product.CurrentPrice - statProfit.product.CurrentPrice*(100-statProfit.product.ProfitPercent)/100
	if res != expected {
		t.Errorf("Error in getCurrentProfit: Expect %v but got %f", expected, res)
	}
	expectedData = append(expectedData, expected)

	res = statProfit.getDifferenceProfit()
	expected = statProfit.product.CurrentPrice - statProfit.Average(statProfit.product.Sells)
	if res != expected {
		t.Errorf("Error in getDifferenceProfit: Expect %v but got %f", expected, res)
	}
	expectedData = append(expectedData, expected)

	data := statProfit.getAllData()
	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("Error in getAllData: Expect %v but got %v", expectedData, data)
	}

}
