# BankRestApi

Примеры запросов 


Создание аккаунта


post


http://localhost:8081/create_account
тело запроса 

{
    "Client": {
        "Id": 2,
        "Name":"Evgeny",
        "Surname":"Grishchuk"
    },
    "Balance":50.0
}

Получение аккаунта по id 

get


http://localhost:8081/?id=2

Транзакция

get

http://localhost:8081/transaction?id1=1&id2=2&sum=20.0
