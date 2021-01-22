package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/pingguodeli573365/delinkcious/pkg/db_util"
	om "github.com/pingguodeli573365/delinkcious/pkg/object_model"
	. "github.com/pingguodeli573365/delinkcious/pkg/test_util"
	"github.com/pingguodeli573365/delinkcious/pkg/user_client"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() {
	db, err := db_util.RunLocalDB("user_manager")
	if err != nil {
		return
	}

	tables := []string{"sessions", "users"}
	for _, table := range tables {
		err = db_util.DeleteFromTableIfExist(db, table)
		check(err)
	}
}

func main() {
	initDB()

	ctx := context.Background()
	defer StopService(ctx)
	RunService(ctx, ".", "user_service")

	// Run some tests with the client
	cli, err := user_client.NewClient("localhost:7070")
	check(err)

	err = cli.Register(om.User{"gg@gg.com", "gigi"})
	check(err)
	log.Print("gigi has registered successfully")

	session, err := cli.Login("gigi", "secret")
	check(err)
	log.Print("gigi has logged in successfully. the session is: ", session)

	err = cli.Logout("gigi", session)
	check(err)
	log.Print("gigi has logged out successfully.")

}
