# Cервис gw-currency-wallet
### сервис обеспечивает хранение данных о пользователях и их счетах (postgres)
### сервер реализован таким образом что в любой момент можно добавить другие СУБД
### поддерживает функции внесения средств и вычета средств с валидацией
### поддерживает перевод средвст из одной валюты в другую в пределах кошелька(для курсов валют обращается к другому микросервису по gRPC)
### по адресу http://localhost:8080/swagger/index.html можно получить swagger-документацию
### взаимодейтсвие происходит по REST API
### имеет Dockerfile


# Сервис gw-exchanger
### доступ к серверу осущетсвляется только по gRPC 
### возвращает курсы валют относительно рубля
### возвращает конкретный курс между двумя валютами
### поддерживает кэширование курса валют