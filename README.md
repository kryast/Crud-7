CRUD ke-7

Product

POST
curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d '{"name":"Kabel HDMI","category":"Aksesoris","price":75000}'

GET
curl http://localhost:8080/products
curl http://localhost:8080/products/1

PUT
curl -X PUT http://localhost:8080/products/1 -H "Content-Type: application/json" -d '{"name":"Kabel HDMI Premium","category":"Aksesoris","price":95000}'

DELETE
curl -X DELETE http://localhost:8080/products/1


Supplier

POST
curl -X POST http://localhost:8080/suppliers -H "Content-Type: application/json" -d '{"name":"PT Sumber Jaya","email":"sales@sumberjaya.com","phone":"021777","address":"Jakarta"}'

GET
curl http://localhost:8080/suppliers
curl http://localhost:8080/suppliers/1

PUT
curl -X PUT http://localhost:8080/suppliers/1 -H "Content-Type: application/json" -d '{"name":"PT Sumber Jaya Updated","email":"contact@sumberjaya.com","phone":"021999","address":"Bekasi"}'

DELETE
curl -X DELETE http://localhost:8080/suppliers/1