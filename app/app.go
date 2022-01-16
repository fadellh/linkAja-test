package app

import (
	"context"
	"fmt"
	"link-test/api"
	accController "link-test/api/account"
	accService "link-test/business/account"
	"link-test/modules"
	accRepo "link-test/modules/account"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDatabaseConnection() *gorm.DB {

	// configDB := map[string]string{
	// 	"DB_Username": os.Getenv("DB_USERNAME"),
	// 	"DB_Password": os.Getenv("DB_PASSWORD"),
	// 	"DB_Port":     os.Getenv("DB_PORT"),
	// 	"DB_Host":     os.Getenv("DB_ADDRESS"),
	// 	"DB_Name":     os.Getenv("DB_NAME"),
	// }
	configDB := map[string]string{
		"DB_Username": "postgres",
		"DB_Password": "password",
		"DB_Port":     "5432",
		"DB_Host":     "db",
		"DB_Name":     "link",
	}

	//connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//connectionString := fmt.Sprintf("host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		configDB["DB_Host"],
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Name"],
		configDB["DB_Port"])

	db, e := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	modules.InitMigrate(db)

	return db
}

func Start() {
	dbConnection := newDatabaseConnection()

	ar := accRepo.NewGormDBRepository(dbConnection)
	as := accService.NewService(ar)
	ac := accController.NewController(as)

	e := echo.New()
	api.RegisterPath(e, ac)

	go func() {
		// address := fmt.Sprintf("localhost:%d", config.AppPort)
		address := fmt.Sprintf("0.0.0.0:2801")
		// address := fmt.Sprintf("localhost:2801")
		fmt.Println(address)
		if err := e.Start(address); err != nil {
			fmt.Println(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
