package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Product struct {
	ID      int
	Product string
	Price   int
}

type Sale struct {
	Product int
	Volume  int
	Date    string
}

func (p Product) String() string {
	return fmt.Sprintf("ID: %d, Product: %s, Price: %d", p.ID, p.Product, p.Price)
}

func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s",
		s.Product, s.Volume, s.Date)
}

func selectSales(clientID int, db *sql.DB) ([]Sale, error) {
	var sales []Sale
	// код
	rows, err := db.Query(`select product, volume, date FROM sales WHERE id = :id`,
		sql.Named("id", clientID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		sale := Sale{}
		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return sales, nil
}

func main() {
	db, err := sql.Open("sqlite", "bd/demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, product, price FROM products")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}

		err := rows.Scan(&product.ID, &product.Product, &product.Price)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(product)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return
	}

	// часть для SelectSales()
	client := 208

	sales, err := selectSales(client, db)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(sales)
	for _, sale := range sales {
		fmt.Println(sale)
	}
}
