CAN TEST THIS REST API USING POSTMAN OR CURL

For cURL, 

Create an Item

curl -X POST -H "Content-Type: application/json" -d '{"name": "Laptop", "price": 1000.50}' http://localhost:8080/items

Get all items

curl http://localhost:8080/items

Get an item by id

curl http://localhost:8080/items/1

