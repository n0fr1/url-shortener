URL-shortener (сервис сокращатель-ссылок)

Назначение:

Данный сервис позволяет пользователю генерировать для произвольго URL'a его короткую версию, которую удобно вставлять в различные публикации, сообщения, новости, промо-материалы и так далее. Также сервис позволяет получать статистику переходов по каждому сгенерированному URL'у.

Принцип работы:

Переходя по короткой ссылке, пользователь отправляет запрос на backend, который, в свою очередь, ищет в своем хранилище (например, в базе данных) длинную версию того короткого, URL по которому был совершен переход. В случае нахождения в БД искомой записи, сервер перенаправляет пользователя на соотетствующий длинный URL и фиксирует в базе факт перехода для того, чтобы в дальнейшем предоставить статистику владельцу URL'a.
