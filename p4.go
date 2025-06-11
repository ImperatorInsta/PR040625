// 4. Напишите класс, представляющий деление с обработкой исключений на случай деления на ноль.

package main

import (
	"fmt"
	"math"
)

type DivisionError struct {
	Message  string
	Dividend float64
	Divisor  float64
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("ошибка деления: %s (%.2f / %.2f)", e.Message, e.Dividend, e.Divisor)
}

type Calculator struct {
	Name string
}

func NewCalculator(name string) *Calculator {
	return &Calculator{
		Name: name,
	}
}

func (c *Calculator) Divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, &DivisionError{
			Message:  "деление на ноль невозможно",
			Dividend: dividend,
			Divisor:  divisor,
		}
	}

	result := dividend / divisor

	if math.IsInf(result, 0) {
		return 0, &DivisionError{
			Message:  "результат деления стремится к бесконечности",
			Dividend: dividend,
			Divisor:  divisor,
		}
	}

	if math.IsNaN(result) {
		return 0, &DivisionError{
			Message:  "результат деления не является числом",
			Dividend: dividend,
			Divisor:  divisor,
		}
	}

	return result, nil
}

func (c *Calculator) SafeDivide(dividend, divisor float64) float64 {
	result, err := c.Divide(dividend, divisor)
	if err != nil {
		fmt.Printf("Предупреждение: %v\n", err)
		return 0
	}
	return result
}

func (c *Calculator) BatchDivide(pairs [][2]float64) []float64 {
	results := make([]float64, 0, len(pairs))

	for i, pair := range pairs {
		dividend, divisor := pair[0], pair[1]

		if result, err := c.Divide(dividend, divisor); err != nil {
			fmt.Printf("Операция %d: %v\n", i+1, err)
			results = append(results, 0)
		} else {
			fmt.Printf("Операция %d: %.2f / %.2f = %.6f\n", i+1, dividend, divisor, result)
			results = append(results, result)
		}
	}

	return results
}

func (c *Calculator) GetInfo() string {
	return fmt.Sprintf("Калькулятор '%s' для безопасного деления", c.Name)
}

func main() {
	calc := NewCalculator("Научный калькулятор")

	fmt.Println("=== Демонстрация работы калькулятора ===")
	fmt.Println(calc.GetInfo())
	fmt.Println()

	testCases := [][2]float64{
		{10, 2},
		{15, 3},
		{7, 0},
		{0, 5},
		{-10, 2},
		{10, -2},
		{0, 0},
		{3.14159, 2},
	}

	fmt.Println("--- Обычное деление с обработкой ошибок ---")
	for i, testCase := range testCases {
		dividend, divisor := testCase[0], testCase[1]

		if result, err := calc.Divide(dividend, divisor); err != nil {
			fmt.Printf("Тест %d: %v\n", i+1, err)
		} else {
			fmt.Printf("Тест %d: %.2f / %.2f = %.6f\n", i+1, dividend, divisor, result)
		}
	}

	fmt.Println("\n--- Безопасное деление ---")
	for i, testCase := range testCases {
		dividend, divisor := testCase[0], testCase[1]
		result := calc.SafeDivide(dividend, divisor)
		fmt.Printf("Безопасный тест %d: %.2f / %.2f = %.6f\n", i+1, dividend, divisor, result)
	}

	fmt.Println("\n--- Пакетное деление ---")
	results := calc.BatchDivide(testCases)
	fmt.Printf("Результаты всех операций: %v\n", results)
	fmt.Println("\nНажмите Enter для выхода...")
	fmt.Scanln()
}
