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
	orders := generateOrders(20)

	/*	go func() {
		defer wg.Done()
		processOrders(orders)
	}()*/

	for range 3 {
		go func() {
			defer wg.Done()
			for _, order := range orders {
				updateOrderStatus(order)
			}

		}()
	}

	wg.Wait()

	reportOrderStatuses(orders)
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

func reportOrderStatuses(orders []*Order) {
	fmt.Println("--Order Status Report --")
	for _, order := range orders {
		fmt.Println(fmt.Sprintf("Order %d - Status: %s", order.ID, order.Status))
	}
	fmt.Println("--End of Order Status Report --")
}

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate processing time
		slog.Info("Processing Order", slog.Int("orderID", order.ID), slog.String("status", order.Status))
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
