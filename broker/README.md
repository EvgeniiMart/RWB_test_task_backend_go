При текущем events.json ожидаются такие результаты:

0-5 секунд: [] (запросов не было)
5-10 секунд: [Strawberry, Apple] (Strawberry здесь является проверкой на expiration, можно в docker-compose поставить низкий EXPIRATION_SECONDS и убедиться, что Strawberry упадёт вниз; Apple же нужен для того, чтобы в будущем убедиться, что его обгонит Banana)
10-20 секунд: [Strawberry, Apple, Banana] (Banana единственный, кто регулярно увеличивается)
20-30 секунд: [Strawberry, Banana, Apple]
30-40 секунд: [Strawberry, Melon, Banana, Apple] (Melon должен сразу по вставке оказаться выше Banana, но ненадолго)
40-... секунд: [Strawberry, Banana, Melon, Apple]

Дальнейшее зависит от EXPIRATION_SECONDS

По аналогии можете создавать свои собственные алгоритмы поведения брокеров для ручного тестирования сервиса
