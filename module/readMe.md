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
5. Удаляем  github.com/valyala/fasthttp v1.37.0
6. Делаем go mod tidy, чтобы почистить после удаления
7. Делаем go mod download 
8. Делаем go mod vendor , появился каталог vendor
9. Делаем go build -mod=vendor
10. Добавляем в репо git commit -a -m 'initial commit.'
11. Делаем git push не прошел пушь через консоль (так как не настраивал его для работы через консоль) сделал через ГитКракен
12. Добавил таг через Гит Кракен v0.0.1
13. Добавил папку v2 в проект
14. Проинициализировал директорию go mod init github.com/VladislavCheremisin/repo-go/module/v2
Создался файл go.mod
15. Добавил main.go в v2
16. Добавил в файл строку _ "github.com/valyala/fasthttp"
17. Сделал go mod tidy
go: finding module for package github.com/gorilla/websocket
go: finding module for package github.com/valyala/fasthttp
go: found github.com/gorilla/websocket in github.com/gorilla/websocket v1.5.0
go: found github.com/valyala/fasthttp in github.com/valyala/fasthttp v1.37.0
18. Сделал комит и запушил
19. Добавил таг v2.0.0 и запушил таг
20. Попробовал добавить этот модуль в другой проект


