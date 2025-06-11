// 5.	Создайте приложение для отслеживания состояния акций. Когда цена акций изменяется, все подписанные пользователи должны получать уведомления
//Важные моменты:
//	Убедитесь, что вы правильно обрабатываете соединения и отключения клиентов.
//	Следите за производительностью, чтобы избегать излишней нагрузки при обновлении состояния.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Stock struct {
	Symbol    string
	Name      string
	Price     float64
	LastPrice float64
	mutex     sync.RWMutex
}

func NewStock(symbol, name string, initialPrice float64) *Stock {
	return &Stock{
		Symbol:    symbol,
		Name:      name,
		Price:     initialPrice,
		LastPrice: initialPrice,
	}
}

func (s *Stock) UpdatePrice(newPrice float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.LastPrice = s.Price
	s.Price = newPrice
}

func (s *Stock) GetPrice() float64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.Price
}

func (s *Stock) GetPriceChange() (float64, float64) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	change := s.Price - s.LastPrice
	changePercent := 0.0
	if s.LastPrice != 0 {
		changePercent = (change / s.LastPrice) * 100
	}

	return change, changePercent
}

func (s *Stock) String() string {
	change, changePercent := s.GetPriceChange()
	changeSymbol := ""
	if change > 0 {
		changeSymbol = "↗"
	} else if change < 0 {
		changeSymbol = "↘"
	} else {
		changeSymbol = "→"
	}

	return fmt.Sprintf("%s (%s): $%.2f %s (%.2f, %.2f%%)",
		s.Symbol, s.Name, s.Price, changeSymbol, change, changePercent)
}

type Observer interface {
	Update(stock *Stock, change float64, changePercent float64)
}

type Subscriber struct {
	Name           string
	MinChangeAlert float64
}

func (sub *Subscriber) Update(stock *Stock, change float64, changePercent float64) {
	if abs(changePercent) >= sub.MinChangeAlert {
		fmt.Printf("🔔 [%s] Уведомление для %s: %s изменилась на %.2f%% (%.2f)\n",
			time.Now().Format("15:04:05"), sub.Name, stock.Symbol, changePercent, change)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

type StockTracker struct {
	stocks    map[string]*Stock
	observers []Observer
	mutex     sync.RWMutex
}

func NewStockTracker() *StockTracker {
	return &StockTracker{
		stocks:    make(map[string]*Stock),
		observers: make([]Observer, 0),
	}
}

func (st *StockTracker) AddStock(stock *Stock) {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	st.stocks[stock.Symbol] = stock
	fmt.Printf("✅ Добавлена акция для отслеживания: %s\n", stock)
}

func (st *StockTracker) Subscribe(observer Observer) {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	st.observers = append(st.observers, observer)
	if sub, ok := observer.(*Subscriber); ok {
		fmt.Printf("👤 Новый подписчик: %s (мин. изменение: %.1f%%)\n",
			sub.Name, sub.MinChangeAlert)
	}
}

func (st *StockTracker) UpdateStockPrice(symbol string, newPrice float64) error {
	st.mutex.RLock()
	stock, exists := st.stocks[symbol]
	st.mutex.RUnlock()

	if !exists {
		return fmt.Errorf("акция %s не найдена", symbol)
	}

	stock.UpdatePrice(newPrice)

	change, changePercent := stock.GetPriceChange()

	st.mutex.RLock()
	observers := make([]Observer, len(st.observers))
	copy(observers, st.observers)
	st.mutex.RUnlock()

	for _, observer := range observers {
		observer.Update(stock, change, changePercent)
	}

	return nil
}

func (st *StockTracker) GetAllStocks() []*Stock {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	stocks := make([]*Stock, 0, len(st.stocks))
	for _, stock := range st.stocks {
		stocks = append(stocks, stock)
	}
	return stocks
}

func (st *StockTracker) PrintPortfolio() {
	fmt.Println("\n📊 === ТЕКУЩЕЕ СОСТОЯНИЕ ПОРТФЕЛЯ ===")
	stocks := st.GetAllStocks()

	if len(stocks) == 0 {
		fmt.Println("Портфель пуст")
		return
	}

	for _, stock := range stocks {
		fmt.Println(stock)
	}
	fmt.Println("=====================================")
}

func (st *StockTracker) SimulateMarket(duration time.Duration, interval time.Duration) {
	fmt.Printf("🎲 Запуск симуляции рынка на %v (обновления каждые %v)\n", duration, interval)

	ticker := time.NewTicker(interval)
	timeout := time.After(duration)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			stocks := st.GetAllStocks()
			if len(stocks) > 0 {
				stock := stocks[rand.Intn(len(stocks))]
				currentPrice := stock.GetPrice()

				changePercent := (rand.Float64() - 0.5) * 20
				newPrice := currentPrice * (1 + changePercent/100)

				if newPrice < 0.01 {
					newPrice = 0.01
				}

				st.UpdateStockPrice(stock.Symbol, newPrice)
			}

		case <-timeout:
			fmt.Println("⏰ Симуляция рынка завершена")
			return
		}
	}
}

func main() {
	tracker := NewStockTracker()

	fmt.Println("🚀 === СИСТЕМА ОТСЛЕЖИВАНИЯ АКЦИЙ ===")

	stocks := []*Stock{
		NewStock("AAPL", "Apple Inc.", 150.00),
		NewStock("GOOGL", "Alphabet Inc.", 2500.00),
		NewStock("TSLA", "Tesla Inc.", 800.00),
		NewStock("MSFT", "Microsoft Corp.", 300.00),
		NewStock("AMZN", "Amazon.com Inc.", 3200.00),
	}

	for _, stock := range stocks {
		tracker.AddStock(stock)
	}

	subscribers := []*Subscriber{
		{"Alice (Консервативный инвестор)", 2.0},
		{"Bob (Активный трейдер)", 1.0},
		{"Charlie (Дневной трейдер)", 0.5},
	}

	for _, sub := range subscribers {
		tracker.Subscribe(sub)
	}

	tracker.PrintPortfolio()

	fmt.Println("\n🔄 === РУЧНЫЕ ОБНОВЛЕНИЯ ===")
	tracker.UpdateStockPrice("AAPL", 155.50)
	tracker.UpdateStockPrice("TSLA", 792.00)
	tracker.UpdateStockPrice("GOOGL", 2525.00)

	time.Sleep(1 * time.Second)
	tracker.PrintPortfolio()

	fmt.Println("\n🎯 === АВТОМАТИЧЕСКАЯ СИМУЛЯЦИЯ ===")
	rand.Seed(time.Now().UnixNano())

	go tracker.SimulateMarket(10*time.Second, 500*time.Millisecond)
	go tracker.SimulateMarket(10*time.Second, 500*time.Millisecond)

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		tracker.PrintPortfolio()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("\n🏁 === ФИНАЛЬНОЕ СОСТОЯНИЕ ===")
	tracker.PrintPortfolio()
	fmt.Println("\nНажмите Enter для выхода...")
	fmt.Scanln()
}
