package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	articleDelivery "github.com/sesha04/test_kumparan/article/delivery"
	articleRepo "github.com/sesha04/test_kumparan/article/repository"
	articleUsecase "github.com/sesha04/test_kumparan/article/usecase"
)

type mysqlConfig struct {
	Username string
	Password string
	Host     string
	DbName   string
	Charset  string
	Pool     int
}

func newMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True",
		"root",
		"root",
		"127.0.0.1",
		"test_kumparan",
		"utf8"),
	)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(300 * time.Second)
	err = db.Ping()
	return db, err
}

func main() {
	e := echo.New()

	db, err := newMysql()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	ar := articleRepo.NewArticleRepository(db)
	au := articleUsecase.NewArticleUsecase(ar)
	articleDelivery.NewArticleHandler(e, au)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func(*echo.Echo) {
		if serr := e.Start(":8080"); serr != http.ErrServerClosed {
			log.Fatal(serr)
		}
	}(e)

	<-sigChan

	log.Println("\nShutting down...")
}
