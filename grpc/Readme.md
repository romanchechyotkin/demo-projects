# grpc demo

## run server 
```shell
cd server
go run main.go
```

## postman testing
для начала импортировать proto файл для доступа к методамммм

GetOrder, получение заказа по id
![image](https://github.com/user-attachments/assets/7f03dd10-c887-40ea-8fca-b3abb0863aec)

UpdateOrders, обновляется вся сущность, так как стриминг, можно отправлять сразу несколько запросов в рамках одного коннекта (см кнопки Send, End Streaming в Postman)
![image](https://github.com/user-attachments/assets/7e47fd3e-d4b0-4cf7-9120-fa446803d0e5)

SearchOrders, даем название айтема в запросе, получаем стримом все подходящие заказы, указан таймаут в 3 секунды, чтобы имитировать стриминг
![image](https://github.com/user-attachments/assets/9c56b9b0-f4f4-403e-ab59-bc99d5fee4f5)
