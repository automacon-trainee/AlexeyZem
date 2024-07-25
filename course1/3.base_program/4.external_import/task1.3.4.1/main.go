package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func DecimalSum(a, b string) (string, error) {
	numA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	numB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	res := numA.Add(numB)

	return res.String(), nil
}

func DecimalSubtract(a, b string) (string, error) {
	numA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	numB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	res := numA.Sub(numB)

	return res.String(), nil
}

func DecimalMultiply(a, b string) (string, error) {
	numA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	numB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	res := numA.Mul(numB)

	return res.String(), nil
}

func DecimalDivide(a, b string) (string, error) {
	numA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	numB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	res := numA.Div(numB)

	return res.String(), nil
}

func DecimalRound(a string, precision int32) (string, error) {
	numA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	res := numA.Round(precision)

	return res.String(), nil
}

func DecimalGreaterThan(a, b string) (bool, error) {
	numA, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	numB, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}

	res := numA.GreaterThan(numB)

	return res, nil
}

func DecimalLessThan(a, b string) (bool, error) {
	numA, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	numB, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}

	res := numA.LessThan(numB)

	return res, nil
}

func DecimalEqual(a, b string) (bool, error) {
	numA, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	numB, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}

	res := numA.Equal(numB)

	return res, nil
}

func main() {
	res, err := DecimalSum("123", "456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = DecimalSubtract("123", "456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = DecimalMultiply("123", "456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = DecimalDivide("123", "456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = DecimalRound("45.898868", 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	greater, err := DecimalGreaterThan("123", "456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greater)
	}

	less, err := DecimalLessThan("123", "456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(less)
	}

	equal, err := DecimalEqual(`123`, `456`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(equal)
	}

}
