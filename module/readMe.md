1. Создаем папку module
2. Инициализируем go mod init github.com/VladislavCheremisin/repo-go/module
go: creating new go.mod: module github.com/VladislavCheremisin/repo-go/module
3. Создаем файл main с наполнением 
package geekbrains_go_module

import _ "github.com/valyala/fasthttp"

func main() {

}
4. Выполняем go mod tidy
go: finding module for package github.com/valyala/fasthttp
go: found github.com/valyala/fasthttp in github.com/valyala/fasthttp v1.37.0
5. Добовляем _ "github.com/gorilla/websocket"
6. Выполняем go mod tidy чтобы проинициализировать  "github.com/gorilla/websocket"
go: finding module for package github.com/gorilla/websocket
go: found github.com/gorilla/websocket in github.com/gorilla/websocket v1.5.0
7. Удаляем  github.com/valyala/fasthttp v1.37.0
8. Делаем go mod tidy, чтобы почистить после удаления
9. Делаем go mod download 
10. Делаем go mod vendor , появился каталог vendor
11. Делаем go build -mod=vendor
12. Добавляем в репо git commit -a -m 'initial commit.'
13. Делаем git push не прошел пушь через консоль (так как не настраивал его для работы через консоль) сделал через ГитКракен
14. Добавил таг через Гит Кракен v0.0.1
15. Добавил папку v2 в проект
16. Проинициализировал директорию go mod init github.com/VladislavCheremisin/repo-go/module/v2
Создался файл go.mod
17. Добавил main.go в v2
18. Добавил в файл строку _ "github.com/valyala/fasthttp"
19. Сделал go mod tidy
go: finding module for package github.com/gorilla/websocket
go: finding module for package github.com/valyala/fasthttp
go: found github.com/gorilla/websocket in github.com/gorilla/websocket v1.5.0
go: found github.com/valyala/fasthttp in github.com/valyala/fasthttp v1.37.0
20. Сделал комит и запушил
21. Добавил таг v2.0.0 и запушил таг
22. Попробовал добавить этот модуль в другой проект
23. Проделал много танцев с бубном, ну очень много, в результате удалил теги и переделал модули много раз так как у меня не получалось импортировать 
созданные модули в новый проет, но в результате сделал отдельно на main ветке и всё получилось, видимо в следствии того что модуль находится 
в состоянии пул реквест были проблемы. Не буду расписывать все свои шаги, но поверьте их было много и они были крайне разнообразными)

