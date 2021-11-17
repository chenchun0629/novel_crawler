package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/novel_crawler/internal/data/ent"
	"github.com/novel_crawler/internal/data/ent/migrate"

	"context"
	"log"
)

var DBClient *ent.Client

func Init() {
	var err error
	DBClient, err = ent.Open("mysql", "root:root@tcp(localhost:3306)/novel_crawler?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := DBClient.Schema.Create(context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		migrate.WithFixture(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
