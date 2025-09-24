package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"tgsms/config"
	"tgsms/external/rest_api"
	"tgsms/internal/transaction"
	"tgsms/rimport"
	"tgsms/tools/logger"
	"tgsms/uimport"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-telegram/bot"
	"github.com/jmoiron/sqlx"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	conf := config.NewConfig()
	fmt.Println("DbConnectionString: ", conf.DbConnectionString())
	db, err := sqlx.Connect("mysql", conf.DbConnectionString())
	if err != nil {
		log.Panic("db connection error: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Panic("ошибка при пинге базы данных")
	}

	b, err := bot.New(conf.BotToken)
	if err != nil {
		log.Panic("ошибка при чтении токена бота: ", err)
	}
	defer b.Close(ctx)

	ginConfig := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}

	router := gin.Default()
	router.Use(cors.New(ginConfig))
	srv := &http.Server{
		Addr:    conf.HttpServerPort,
		Handler: router,
	}

	logger, err := logger.InitLogger()
	if err != nil {
		log.Panic("не удалось создать логер: ", err)
	}

	sessionManager := transaction.NewSQLSessionManager(db)
	ri := rimport.NewRepositoryImports()
	ui := uimport.NewUsecaseImport(ri, logger, conf, b)

	rest_api.NewIncomingMessagesHandler(ui, router, conf, logger, sessionManager)

	wg.Add(2)
	// запуск http сервера
	go func() {
		defer wg.Done()

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic("Ошибка при запуске HTTP-сервера: ", err)
		}
	}()

	// запуск бота
	go func() {
		defer wg.Done()

		log.Println("bot started")
		b.Start(ctx)
	}()

	// Ожидание сигнала завершения
	<-ctx.Done()
	log.Println("Останавливаем приложение...")

	// Даём время на завершение запросов
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// Остановка HTTP-сервера
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Println("Ошибка при остановке HTTP-сервера:", err)
	}

	// Ожидание завершения всех горутин
	wg.Wait()
	log.Println("Приложение остановлено")
}
