// 3. Реализуйте класс "Стек" с методами push, pop и проверки пустоты. Используйте динамическое выделение памяти.

package main

import (
	"fmt"
)

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
	fmt.Printf("Добавлен элемент: %v\n", item)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("стек пуст")
	}

	index := len(s.items) - 1
	item := s.items[index]

	s.items = s.items[:index]

	fmt.Printf("Удален элемент: %v\n", item)
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("стек пуст")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) PrintStack() {
	if s.IsEmpty() {
		fmt.Println("Стек пуст")
		return
	}

	fmt.Print("Стек (снизу вверх): ")
	for _, item := range s.items {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}

func main() {
	stack := NewStack()

	fmt.Println("=== Демонстрация работы стека ===")

	fmt.Printf("Стек пуст: %t\n", stack.IsEmpty())

	stack.Push(10)
	stack.Push("Hello")
	stack.Push(3.14)
	stack.Push(true)

	stack.PrintStack()
	fmt.Printf("Размер стека: %d\n", stack.Size())

	if top, err := stack.Peek(); err == nil {
		fmt.Printf("Верхний элемент: %v\n", top)
	}

	for !stack.IsEmpty() {
		if _, err := stack.Pop(); err == nil {
			fmt.Printf("Размер стека после удаления: %d\n", stack.Size())
		} else {
			fmt.Printf("Ошибка: %v\n", err)
		}
	}

	if _, err := stack.Pop(); err != nil {
		fmt.Printf("Ошибка при попытке удаления из пустого стека: %v\n", err)
	}
	fmt.Println("\nНажмите Enter для выхода...")
	fmt.Scanln()
}
