package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	totalUpdate atomic.Int64
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func main() {
	startTime := time.Now()
	defer func() {
		slog.Info("All operations completed.",
			slog.String("uptime", time.Since(startTime).String()),
			slog.Int64("total_updates", totalUpdate.Load()))
	}()

	wg := sync.WaitGroup{}
	wg.Add(3)

	orderChan := make(chan *Order, 20)
	processedChan := make(chan *Order, 20)

	go func() {
		defer wg.Done()
		for _, order := range generateOrders(20) {
			orderChan <- order
		}
		close(orderChan)

		slog.Info("Done with generating orders")
	}()

	orders := generateOrders(20)

	go processOrders(orderChan, processedChan, &wg)

	go func() {
		defer wg.Done()

		for {
			select {
			case processedOrder, ok := <-processedChan:
				if !ok {
					slog.Info("Processed channel closed, stopping processing")
					return
				}
				slog.Info("Order processed", slog.Int("orderID", processedOrder.ID), slog.String("status", processedOrder.Status))
			case <-time.After(10 * time.Second):
				slog.Info("No orders processed in the last 10 seconds, checking for updates")
				return
			}
		}
	}()

	wg.Wait()

	reportOrderStatuses(orders)
}

func reportOrderStatuses(orders []*Order) {
	fmt.Println("--Order Status Report --")
	for _, order := range orders {
		fmt.Println(fmt.Sprintf("Order %d - Status: %s", order.ID, order.Status))
	}

	fmt.Println("--End of Order Status Report --")
}

func processOrders(inChan <-chan *Order, outChan chan<- *Order, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(outChan)
	}()

	for order := range inChan {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate processing time

		order.mu.Lock()
		order.Status = "Processed"
		order.mu.Unlock()
		outChan <- order
		// slog.Info("Processing Order", slog.Int("orderID", order.ID), slog.String("status", order.Status))
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)

	for i := range count {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pending",
		}
	}

	return orders
}

func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate processing time

		status := []string{"Pending", "Processing", "Shipped", "Delivered"}[rand.Intn(4)]
		order.Status = status

		slog.Info("Updated Order Status",
			slog.Int("orderID", order.ID),
			slog.String("status", order.Status),
		)
	}
}

func updateOrderStatus(order *Order) {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond) // Simulate processing time

	order.mu.Lock()
	defer order.mu.Unlock()
	status := []string{"Pending", "Processing", "Shipped", "Delivered"}[rand.Intn(4)]
	order.Status = status

	slog.Info("Updated Order Status",
		slog.Int("orderID", order.ID),
		slog.String("status", order.Status),
	)

	totalUpdate.Add(1)
}
