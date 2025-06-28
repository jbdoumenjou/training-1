package main

import (
	"log/slog"
	"math/rand"
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

	orders := generateOrders(20)

	processOrders(orders)

	updateOrderStatuses(orders)

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

func reportOrderStatuses(orders []*Order) {
	for i := range 5 {
		time.Sleep(1 * time.Second) // Simulate periodic reporting

		slog.Info("Order Status Report", slog.Int("iteration", i+1))
		for _, order := range orders {
			slog.Info("Order Status",
				slog.Int("orderID", order.ID),
				slog.String("status", order.Status),
			)
		}
		slog.Info("End of Report", slog.Int("iteration", i+1))
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
