package middleware

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	model "github.com/isfnev/postgres_integration/Models"
	"github.com/joho/godotenv"
)

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
	var POSTGRES_URL = "postgres://postgres:rage@localhost:5434/stocksdb?sslmode=disable"
	db, err := sql.Open("postgres", POSTGRES_URL)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `select * from stocks`
	row, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer row.Close()
	stocks := []model.Stock{}

	for row.Next() {
		var stock model.Stock
		if err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company); err != nil {
			panic(err)
		}
		stocks = append(stocks, stock)
	}

	json.NewEncoder(w).Encode(stocks)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	var stock model.Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		panic(err)
	}

	sqlStatement := `insert into stocks (name, price, company) values ($1, $2, $3) returning stockid`
	res, err := db.Exec(sqlStatement, stock.Name, stock.Price, stock.Company)

	if err != nil {
		panic(err)
	}

	createId, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(createId)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	params := mux.Vars(r)
	sqlStatement := `select * from stocks where stockid = $1`
	id, err := strconv.Atoi(params["id"])

	if err != nil{
		panic(err)
	}

	row := db.QueryRow(sqlStatement, id)
	stock := model.Stock{}

	if err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request){
	db := createConnection()
	defer db.Close()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err)
	}

	var stock model.Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		panic(err)
	}

	sqlStatement := `update stocks set name = $1, price = $2, company = $3 where stockid = $4`
	res, err := db.Exec(sqlStatement, stock.Name, stock.Price, stock.Company, id)

	if err != nil{
		panic(err)
	}
	rowsaffected, err := res.RowsAffected()
	if err!= nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(rowsaffected)
}

func DeleteStock(w http.ResponseWriter, r *http.Request){
	db := createConnection()
	defer db.Close()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err)
	}
	sqlStatement := `delete from stocks where stockid = $1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	rowsaffected, err := res.RowsAffected()
	if err != nil{
		panic(err)
	}
	json.NewEncoder(w).Encode(rowsaffected)
}