package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_gpHandler "github.com/gas_price/gas_price/delivery/http"
	middleware "github.com/gas_price/gas_price/delivery/http/middleware"
	"github.com/gas_price/gas_price/repository/postgresql"
	"github.com/gas_price/gas_price/usecase"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	dbConn, err := sql.Open(`postgres`, psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	middL := middleware.InitMiddleware()

	e.Use(middL.CORS)

	// gasPriceRepo :=
	gasPriceRepo := postgresql.NewPostgresqlGasPriceRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	gpu := usecase.NewGasPriceUsecase(gasPriceRepo, timeoutContext)
	_gpHandler.NewGasPriceHandler(e, gpu)

	log.Fatal(e.Start(viper.GetString("server.address"))) //nolint

	// Connect to the PostgreSQL database
	// db, err := sql.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable")
	// // db, err := sql.Open("postgres", "postgres://user:postgres@localhost/gas_prices?sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// // Define the HTTP handlers for the CRUD operations
	// http.HandleFunc("/gas_price", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		getGasPrices(db, w, r)
	// 	case http.MethodPost:
	// 		createGasPrice(db, w, r)
	// 	case http.MethodPut:
	// 		updateGasPrice(db, w, r)
	// 	case http.MethodDelete:
	// 		deleteGasPrice(db, w, r)
	// 	default:
	// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 	}
	// })

	// // Start the HTTP server
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

// func getGasPrices(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("SELECT id, litre, premium, pertalite FROM gas_price")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	// Build a slice of GasPrice structs from the retrieved rows
// 	gasPrices := make([]GasPrice, 0)
// 	for rows.Next() {
// 		var gasPrice GasPrice
// 		err := rows.Scan(&gasPrice.ID, &gasPrice.Litre, &gasPrice.Premium, &gasPrice.Pertalite)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		gasPrices = append(gasPrices, gasPrice)
// 	}
// 	if err := rows.Err(); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Marshal the slice of GasPrice structs to JSON and write the response
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(gasPrices)
// }

// func createGasPrice(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(nil)
// }

// func updateGasPrice(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(nil)
// }

// func deleteGasPrice(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(nil)
// }
