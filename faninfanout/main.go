package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func main() {
	startTime := time.Now()
	defer func() {
		slog.Info("All operations completed.", slog.String("uptime", time.Since(startTime).String()))
	}()

	wg := sync.WaitGroup{}
	wg.Add(3)
	orders := generateOrders(20)

	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrderStatuses(orders)
	}()

	go func() {
		defer wg.Done()
		reportOrderStatuses(orders)
	}()

	wg.Wait()
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

func reportOrderStatuses(orders []*Order) {
	for i := range 5 {
		time.Sleep(1 * time.Second) // Simulate periodic reporting

		fmt.Println("--Order Status Report --")
		slog.Info("Order Status Report", slog.Int("iteration", i+1))
		for _, order := range orders {
			fmt.Println(fmt.Sprintf("Order %d - Status: %s", order.ID, order.Status))
		}
		fmt.Println("--End of Order Status Report --")
	}
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
