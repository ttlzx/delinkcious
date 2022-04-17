package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/ttlzx/delinkcious/pkg/db_util"
	"github.com/ttlzx/delinkcious/pkg/social_graph_client"
	. "github.com/ttlzx/delinkcious/pkg/test_util"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() {
	db, err := db_util.RunLocalDB("social_graph_manager")
	check(err)

	// Ignore if table doesn't exist (will be created by service)
	err = db_util.DeleteFromTableIfExist(db, "social_graph")
	check(err)
}

func main() {
	initDB()

	ctx := context.Background()
	defer StopService(ctx)
	RunService(ctx, ".", "social_graph_service")

	// Run some tests with the client
	cli, err := social_graph_client.NewClient("localhost:9090")
	check(err)

	following, err := cli.GetFollowing("ttlzx")
	check(err)
	log.Print("ttlzx is following:", following)
	followers, err := cli.GetFollowers("ttlzx")
	check(err)
	log.Print("ttlzx is followed by:", followers)

	err = cli.Follow("ttlzx", "liat")
	check(err)
	err = cli.Follow("ttlzx", "guy")
	check(err)
	err = cli.Follow("guy", "ttlzx")
	check(err)
	err = cli.Follow("saar", "ttlzx")
	check(err)
	err = cli.Follow("saar", "ophir")
	check(err)

	following, err = cli.GetFollowing("ttlzx")
	check(err)
	log.Print("ttlzx is following:", following)
	followers, err = cli.GetFollowers("ttlzx")
	check(err)
	log.Print("ttlzx is followed by:", followers)
}
