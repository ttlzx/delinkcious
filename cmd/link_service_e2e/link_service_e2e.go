package main

import (
	"context"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/ttlzx/delinkcious/pkg/db_util"
	"github.com/ttlzx/delinkcious/pkg/link_manager_client"
	om "github.com/ttlzx/delinkcious/pkg/object_model"
	. "github.com/ttlzx/delinkcious/pkg/test_util"
)

func initDB() {
	db, err := db_util.RunLocalDB("link_manager")
	Check(err)

	tables := []string{"tags", "links"}
	for _, table := range tables {
		err = db_util.DeleteFromTableIfExist(db, table)
		Check(err)
	}
}

func runLinkService(ctx context.Context) {
	// Set environment
	err := os.Setenv("PORT", "8080")
	Check(err)

	err = os.Setenv("MAX_LINKS_PER_USER", "10")
	Check(err)

	RunService(ctx, ".", "link_service")
}

func runSocialGraphService(ctx context.Context) {
	err := os.Setenv("PORT", "9090")
	Check(err)

	RunService(ctx, "../social_graph_service", "social_graph_service")
}

func main() {
	//// Turn on authentication
	//err := os.Setenv("DELINKCIOUS_MUTUAL_AUTH", "true")
	//Check(err)

	initDB()

	ctx := context.Background()
	defer StopService(ctx)

	if os.Getenv("RUN_SOCIAL_GRAPH_SERVICE") == "true" {
		runSocialGraphService(ctx)
	}

	if os.Getenv("RUN_LINK_SERVICE") == "true" {
		runLinkService(ctx)
	}

	// Run some tests with the client
	cli, err := link_manager_client.NewClient("localhost:8080")
	Check(err)

	links, err := cli.GetLinks(om.GetLinksRequest{Username: "ttlzx"})
	Check(err)
	log.Print("ttlzx's links:", links)

	err = cli.AddLink(om.AddLinkRequest{Username: "ttlzx",
		Url:   "https://github.com/ttlzx",
		Title: "Gigi on Github",
		Tags:  map[string]bool{"programming": true}})
	Check(err)
	links, err = cli.GetLinks(om.GetLinksRequest{Username: "ttlzx"})
	Check(err)
	log.Print("ttlzx's links:", links)

	err = cli.UpdateLink(om.UpdateLinkRequest{Username: "ttlzx",
		Url:         "https://github.com/ttlzx",
		Description: "Most of my open source code is here"},
	)

	Check(err)
	links, err = cli.GetLinks(om.GetLinksRequest{Username: "ttlzx"})
	Check(err)
	log.Print("ttlzx's links:", links)

	err = cli.DeleteLink("ttlzx", "https://github.com/ttlzx")
	Check(err)
	Check(err)
	links, err = cli.GetLinks(om.GetLinksRequest{Username: "ttlzx"})
	Check(err)
	log.Print("ttlzx's links:", links)
}
