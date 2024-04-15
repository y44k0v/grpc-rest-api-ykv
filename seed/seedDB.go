package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/y44k0v/grpc-rest-api-ykv/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/y44k0v/grpc-rest-api-ykv/ent/product"
	"github.com/y44k0v/grpc-rest-api-ykv/ent/user"
)

func main() {
	// Read the connection string to the database from a CLI flag.
	var dsn string
	flag.StringVar(&dsn, "dsn", "", "database DSN")
	flag.Parse()

	// Instantiate the Ent client.
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	// If we don't have any posts yet, seed the database.
	if !client.Order.Query().ExistX(ctx) {
		if err := seed(ctx, client); err != nil {
			log.Fatalf("failed seeding the database: %v", err)
		}
	}
	// ... Continue with server start.
}

func seed(ctx context.Context, client *ent.Client) error {
	// Check if the user "rotemtam" already exists.
	r, err := client.User.Query().
		Where(
			user.Name("y44k0v"),
		).
		Only(ctx)
	switch {
	// If not, create the user.
	case ent.IsNotFound(err):
		r, err = client.User.Create().
			SetName("y44k0v").
			SetEmail("y44k0v@gmail.com").
			SetAddress("123 street dr, toronto A0A 0A0").
			SetOrderID([]int{1}).
			SetLevel("admin").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating user: %v", err)
		}
	case err != nil:
		return fmt.Errorf("failed querying user: %v", err)
	}
	// create product
	client.Product.Create().
		SetSku("0001-apple").
		SetName("Apples").
		SetCategories([]string{"Fruits"}).
		SetDescription("Snowflake apple").
		SetPictures([]string{"assets/img/produtcs/apple.jpg"}).
		SetQtyAvailable(10).
		SetPrice(5.99).
		Exec(ctx)

		// Finally create order order.
	prodId, err := client.Product.Query().Where(
		product.Name("Apples")).FirstID(ctx)
	if err != nil {
		log.Fatalf("Product not found: %v", err)
	}

	return client.Order.Create().
		SetProductsRodered(map[int]int{prodId: 5}).
		SetSubTotal(5.99 * 5).
		SetUser(r).
		Exec(ctx)
}
