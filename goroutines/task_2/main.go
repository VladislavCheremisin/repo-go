package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Создаем канал слушать сигнал.
var signalChan chan (os.Signal) = make(chan os.Signal, 1)

func main() {
	// Определяем порт для службы HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: http.HandlerFunc(handler),
	}

	// SIGTERM обрабатывает сигнал завершения.
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Старт HTTP server.
	go func() {
		log.Printf("listening on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Получаем из signalChan.
	sig := <-signalChan
	log.Printf("%s signal caught", sig)

	// Таймаут завершения.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Выключаем сервер, ожидая существующих запросов (кроме веб-сокетов).
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %+v", err)
	}
	log.Print("server exited")
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("terminate") != "" {
		fmt.Fprint(w, "Goodbye World!\n")
		signalChan <- syscall.SIGTERM
		return
	}
	fmt.Fprint(w, "Hello World!\n")
}
