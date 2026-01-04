package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "bd/demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	// product := "Облачное хранилище"
	// price := 300
	// // название продукта и цена передаются через параметры
	// res, err := db.Exec("INSERT INTO products (product, price) VALUES (:product, :price)",
	// 	sql.Named("product", product),
	// 	sql.Named("price", price))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(res.LastInsertId())
	// fmt.Println(res.RowsAffected())

	productID := 1
	price := 700
	// обновление цены у продукта с заданным идентификатором
	// цена и идентификатор передаются через параметры запроса
	_, err = db.Exec(`UPDATE products SET price = :price WHERE id = :id`,
		sql.Named("price", price),
		sql.Named("id", productID))
	if err != nil {
		log.Println(err)
		return
	}

	// DELETE запросы
	_, err = db.Exec("DELETE FROM clients WHERE id = :id", sql.Named("id", 3))
	if err != nil {
		fmt.Println(err)
		return
	}
}
