CRUD ke-5

POST
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{
  "name": "Monitor LED 24 Inch",
  "category": "Elektronik",
  "price": 1500000
}'

curl -X POST http://localhost:8080/customers \
-H "Content-Type: application/json" \
-d '{
  "name": "Ahmad",
  "email": "ahmad@example.com",
  "phone": "081234567890"
}'

curl -X POST http://localhost:8080/orders \
-H "Content-Type: application/json" \
-d '{
  "customer_id": 1,
  "status": "pending"
}'

curl -X POST http://localhost:8080/order-items \
-H "Content-Type: application/json" \
-d '{
  "order_id": 1,
  "product_id": 1,
  "quantity": 2
}'


GET
curl http://localhost:8080/order-items

curl http://localhost:8080/order-items/1

PUT
curl -X PUT http://localhost:8080/order-items/1 \
-H "Content-Type: application/json" \
-d '{
  "order_id": 1,
  "product_id": 1,
  "quantity": 3
}'