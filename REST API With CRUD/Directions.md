CAN TEST THIS REST API USING POSTMAN OR CURL

For cURL, 

Create an Item

curl -X POST -H "Content-Type: application/json" -d '{"name": "Laptop", "price": 1000.50}' http://localhost:8080/items

Get all items

curl http://localhost:8080/items

Get an item by id

curl http://localhost:8080/items/1

Update an item

curl -X PUT -H "Content-Type: application/json" -d '{"name": "Gaming Laptop", "price": 1500.75}' http://localhost:8080/items?id=1

Delete an item

curl -X DELETE http://localhost:8080/items?id=1


