First version without go routines:
```bash
go run main.go 
2025/06/28 10:58:08 INFO Processing Order orderID=1 status=Pending
2025/06/28 10:58:08 INFO Processing Order orderID=2 status=Pending
2025/06/28 10:58:09 INFO Processing Order orderID=3 status=Pending
2025/06/28 10:58:09 INFO Processing Order orderID=4 status=Pending
2025/06/28 10:58:09 INFO Processing Order orderID=5 status=Pending
2025/06/28 10:58:09 INFO Processing Order orderID=6 status=Pending
2025/06/28 10:58:10 INFO Processing Order orderID=7 status=Pending
2025/06/28 10:58:10 INFO Processing Order orderID=8 status=Pending
2025/06/28 10:58:10 INFO Processing Order orderID=9 status=Pending
2025/06/28 10:58:10 INFO Processing Order orderID=10 status=Pending
2025/06/28 10:58:10 INFO Processing Order orderID=11 status=Pending
2025/06/28 10:58:11 INFO Processing Order orderID=12 status=Pending
2025/06/28 10:58:11 INFO Processing Order orderID=13 status=Pending
2025/06/28 10:58:11 INFO Processing Order orderID=14 status=Pending
2025/06/28 10:58:11 INFO Processing Order orderID=15 status=Pending
2025/06/28 10:58:11 INFO Processing Order orderID=16 status=Pending
2025/06/28 10:58:12 INFO Processing Order orderID=17 status=Pending
2025/06/28 10:58:12 INFO Processing Order orderID=18 status=Pending
2025/06/28 10:58:12 INFO Processing Order orderID=19 status=Pending
2025/06/28 10:58:12 INFO Processing Order orderID=20 status=Pending
2025/06/28 10:58:13 INFO Updated Order Status orderID=1 status=Pending
2025/06/28 10:58:13 INFO Updated Order Status orderID=2 status=Shipped
2025/06/28 10:58:13 INFO Updated Order Status orderID=3 status=Pending
2025/06/28 10:58:13 INFO Updated Order Status orderID=4 status=Delivered
2025/06/28 10:58:14 INFO Updated Order Status orderID=5 status=Delivered
2025/06/28 10:58:14 INFO Updated Order Status orderID=6 status=Shipped
2025/06/28 10:58:14 INFO Updated Order Status orderID=7 status=Shipped
2025/06/28 10:58:14 INFO Updated Order Status orderID=8 status=Shipped
2025/06/28 10:58:15 INFO Updated Order Status orderID=9 status=Pending
2025/06/28 10:58:15 INFO Updated Order Status orderID=10 status=Shipped
2025/06/28 10:58:15 INFO Updated Order Status orderID=11 status=Shipped
2025/06/28 10:58:15 INFO Updated Order Status orderID=12 status=Pending
2025/06/28 10:58:15 INFO Updated Order Status orderID=13 status=Pending
2025/06/28 10:58:15 INFO Updated Order Status orderID=14 status=Pending
2025/06/28 10:58:16 INFO Updated Order Status orderID=15 status=Processing
2025/06/28 10:58:16 INFO Updated Order Status orderID=16 status=Processing
2025/06/28 10:58:17 INFO Updated Order Status orderID=17 status=Delivered
2025/06/28 10:58:17 INFO Updated Order Status orderID=18 status=Processing
2025/06/28 10:58:17 INFO Updated Order Status orderID=19 status=Pending
2025/06/28 10:58:18 INFO Updated Order Status orderID=20 status=Pending
2025/06/28 10:58:19 INFO Order Status Report iteration=1
2025/06/28 10:58:19 INFO Order Status orderID=1 status=Pending
2025/06/28 10:58:19 INFO Order Status orderID=2 status=Shipped
2025/06/28 10:58:19 INFO Order Status orderID=3 status=Pending
2025/06/28 10:58:19 INFO Order Status orderID=4 status=Delivered
2025/06/28 10:58:19 INFO Order Status orderID=5 status=Delivered
2025/06/28 10:58:19 INFO Order Status orderID=6 status=Shipped
2025/06/28 10:58:19 INFO Order Status orderID=7 status=Shipped
2025/06/28 10:58:19 INFO Order Status orderID=8 status=Shipped
2025/06/28 10:58:19 INFO Order Status orderID=9 status=Pending
2025/06/28 10:58:19 INFO Order Status orderID=10 status=Shipped
2025/06/28 10:58:19 INFO Order Status orderID=11 status=Shipped
2025/06/28 10:58:19 INFO Order Status orderID=12 status=Pending
2025/06/28 10:58:19 INFO Order Status orderID=13 status=Pending
2025/06/28 10:58:19 INFO Order Status orderID=14 status=Pending
2025/06/28 10:58:19 INFO Order Status orderID=15 status=Processing
2025/06/28 10:58:19 INFO Order Status orderID=16 status=Processing
2025/06/28 10:58:19 INFO Order Status orderID=17 status=Delivered
2025/06/28 10:58:19 INFO Order Status orderID=18 status=Processing
2025/06/28 10:58:19 INFO Order Status orderID=19 status=Pending
2025/06/28 10:58:19 INFO Order Status orderID=20 status=Pending
2025/06/28 10:58:19 INFO End of Report iteration=1
2025/06/28 10:58:20 INFO Order Status Report iteration=2
2025/06/28 10:58:20 INFO Order Status orderID=1 status=Pending
2025/06/28 10:58:20 INFO Order Status orderID=2 status=Shipped
2025/06/28 10:58:20 INFO Order Status orderID=3 status=Pending
2025/06/28 10:58:20 INFO Order Status orderID=4 status=Delivered
2025/06/28 10:58:20 INFO Order Status orderID=5 status=Delivered
2025/06/28 10:58:20 INFO Order Status orderID=6 status=Shipped
2025/06/28 10:58:20 INFO Order Status orderID=7 status=Shipped
2025/06/28 10:58:20 INFO Order Status orderID=8 status=Shipped
2025/06/28 10:58:20 INFO Order Status orderID=9 status=Pending
2025/06/28 10:58:20 INFO Order Status orderID=10 status=Shipped
2025/06/28 10:58:20 INFO Order Status orderID=11 status=Shipped
2025/06/28 10:58:20 INFO Order Status orderID=12 status=Pending
2025/06/28 10:58:20 INFO Order Status orderID=13 status=Pending
2025/06/28 10:58:20 INFO Order Status orderID=14 status=Pending
2025/06/28 10:58:20 INFO Order Status orderID=15 status=Processing
2025/06/28 10:58:20 INFO Order Status orderID=16 status=Processing
2025/06/28 10:58:20 INFO Order Status orderID=17 status=Delivered
2025/06/28 10:58:20 INFO Order Status orderID=18 status=Processing
2025/06/28 10:58:20 INFO Order Status orderID=19 status=Pending
2025/06/28 10:58:20 INFO Order Status orderID=20 status=Pending
2025/06/28 10:58:20 INFO End of Report iteration=2
2025/06/28 10:58:21 INFO Order Status Report iteration=3
2025/06/28 10:58:21 INFO Order Status orderID=1 status=Pending
2025/06/28 10:58:21 INFO Order Status orderID=2 status=Shipped
2025/06/28 10:58:21 INFO Order Status orderID=3 status=Pending
2025/06/28 10:58:21 INFO Order Status orderID=4 status=Delivered
2025/06/28 10:58:21 INFO Order Status orderID=5 status=Delivered
2025/06/28 10:58:21 INFO Order Status orderID=6 status=Shipped
2025/06/28 10:58:21 INFO Order Status orderID=7 status=Shipped
2025/06/28 10:58:21 INFO Order Status orderID=8 status=Shipped
2025/06/28 10:58:21 INFO Order Status orderID=9 status=Pending
2025/06/28 10:58:21 INFO Order Status orderID=10 status=Shipped
2025/06/28 10:58:21 INFO Order Status orderID=11 status=Shipped
2025/06/28 10:58:21 INFO Order Status orderID=12 status=Pending
2025/06/28 10:58:21 INFO Order Status orderID=13 status=Pending
2025/06/28 10:58:21 INFO Order Status orderID=14 status=Pending
2025/06/28 10:58:21 INFO Order Status orderID=15 status=Processing
2025/06/28 10:58:21 INFO Order Status orderID=16 status=Processing
2025/06/28 10:58:21 INFO Order Status orderID=17 status=Delivered
2025/06/28 10:58:21 INFO Order Status orderID=18 status=Processing
2025/06/28 10:58:21 INFO Order Status orderID=19 status=Pending
2025/06/28 10:58:21 INFO Order Status orderID=20 status=Pending
2025/06/28 10:58:21 INFO End of Report iteration=3
2025/06/28 10:58:22 INFO Order Status Report iteration=4
2025/06/28 10:58:22 INFO Order Status orderID=1 status=Pending
2025/06/28 10:58:22 INFO Order Status orderID=2 status=Shipped
2025/06/28 10:58:22 INFO Order Status orderID=3 status=Pending
2025/06/28 10:58:22 INFO Order Status orderID=4 status=Delivered
2025/06/28 10:58:22 INFO Order Status orderID=5 status=Delivered
2025/06/28 10:58:22 INFO Order Status orderID=6 status=Shipped
2025/06/28 10:58:22 INFO Order Status orderID=7 status=Shipped
2025/06/28 10:58:22 INFO Order Status orderID=8 status=Shipped
2025/06/28 10:58:22 INFO Order Status orderID=9 status=Pending
2025/06/28 10:58:22 INFO Order Status orderID=10 status=Shipped
2025/06/28 10:58:22 INFO Order Status orderID=11 status=Shipped
2025/06/28 10:58:22 INFO Order Status orderID=12 status=Pending
2025/06/28 10:58:22 INFO Order Status orderID=13 status=Pending
2025/06/28 10:58:22 INFO Order Status orderID=14 status=Pending
2025/06/28 10:58:22 INFO Order Status orderID=15 status=Processing
2025/06/28 10:58:22 INFO Order Status orderID=16 status=Processing
2025/06/28 10:58:22 INFO Order Status orderID=17 status=Delivered
2025/06/28 10:58:22 INFO Order Status orderID=18 status=Processing
2025/06/28 10:58:22 INFO Order Status orderID=19 status=Pending
2025/06/28 10:58:22 INFO Order Status orderID=20 status=Pending
2025/06/28 10:58:22 INFO End of Report iteration=4
2025/06/28 10:58:23 INFO Order Status Report iteration=5
2025/06/28 10:58:23 INFO Order Status orderID=1 status=Pending
2025/06/28 10:58:23 INFO Order Status orderID=2 status=Shipped
2025/06/28 10:58:23 INFO Order Status orderID=3 status=Pending
2025/06/28 10:58:23 INFO Order Status orderID=4 status=Delivered
2025/06/28 10:58:23 INFO Order Status orderID=5 status=Delivered
2025/06/28 10:58:23 INFO Order Status orderID=6 status=Shipped
2025/06/28 10:58:23 INFO Order Status orderID=7 status=Shipped
2025/06/28 10:58:23 INFO Order Status orderID=8 status=Shipped
2025/06/28 10:58:23 INFO Order Status orderID=9 status=Pending
2025/06/28 10:58:23 INFO Order Status orderID=10 status=Shipped
2025/06/28 10:58:23 INFO Order Status orderID=11 status=Shipped
2025/06/28 10:58:23 INFO Order Status orderID=12 status=Pending
2025/06/28 10:58:23 INFO Order Status orderID=13 status=Pending
2025/06/28 10:58:23 INFO Order Status orderID=14 status=Pending
2025/06/28 10:58:23 INFO Order Status orderID=15 status=Processing
2025/06/28 10:58:23 INFO Order Status orderID=16 status=Processing
2025/06/28 10:58:23 INFO Order Status orderID=17 status=Delivered
2025/06/28 10:58:23 INFO Order Status orderID=18 status=Processing
2025/06/28 10:58:23 INFO Order Status orderID=19 status=Pending
2025/06/28 10:58:23 INFO Order Status orderID=20 status=Pending
2025/06/28 10:58:23 INFO End of Report iteration=5
2025/06/28 10:58:23 INFO All operations completed. uptime=14.852610605s
```