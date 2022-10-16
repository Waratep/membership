package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/Waratep/membership/src/interface/gin_server"
	"github.com/Waratep/membership/src/repository/member_repository"
	"github.com/Waratep/membership/src/use_case"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	AppName                string `env:"APP_NAME" envDefault:"membership"`
	AppVersion             string `env:"APP_VERSION"`
	Environment            string `env:"ENVIRONMENT" envDefault:"development"`
	Port                   int64  `env:"PORT" envDefault:"80"`
	Debug                  bool   `env:"DEBUG" envDefault:"true"`
	PostgresDatasourceName string `env:"POSTGRES_DATASOURCE_NAME"`
}

func main() {
	cfg := initEnvironment()

	memberRepo := initRepositories(cfg)
	useCase := use_case.New(memberRepo)

	initInterface(cfg, &useCase)
}

func initEnvironment() config {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	err := godotenv.Load()
	if err == nil {
		var cfg config
		err = env.Parse(&cfg)
		if err != nil {
			log.Fatalf("Error parse env: %s\n", err)
		}

		return cfg
	}

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Error parse env `PORT`: %s\n", err)
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalf("Error parse env `DEBUG`: %s\n", err)
	}

	return config{
		AppName:                os.Getenv("APP_NAME"),
		AppVersion:             os.Getenv("APP_VERSION"),
		Environment:            os.Getenv("ENVIRONMENT"),
		Port:                   port,
		Debug:                  debug,
		PostgresDatasourceName: os.Getenv("POSTGRES_DATASOURCE_NAME"),
	}
}

func initInterface(cfg config, useCase *use_case.UseCase) {
	wg := new(sync.WaitGroup)

	ginServer := gin_server.New(useCase, &gin_server.ServerConfig{
		AppVersion:    cfg.AppVersion,
		RequestLog:    false,
		ListenAddress: fmt.Sprintf(":%d", cfg.Port),
		Debug:         cfg.Debug,
	})

	ginServer.Start(wg)
	wg.Wait()
}

func initRepositories(cfg config) use_case.MemberRepository {
	postgresDB, err := sql.Open("postgres", cfg.PostgresDatasourceName)
	if err != nil {
		log.Fatal(err)
	}
	err = postgresDB.Ping()
	if err != nil {
		log.Fatalln("Error ping postgres database", err)
	}
	log.Println("Ping postgres database success")

	conn, err := postgresDB.Conn(context.Background())
	if err != nil {
		log.Fatalln("Error connect postgres database", err)
	}

	defer postgresDB.Close()

	memberRepo := member_repository.NewPostgres(conn)

	return memberRepo
}
