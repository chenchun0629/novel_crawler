package data

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/novel_crawler/internal/data/ent"
	"github.com/novel_crawler/internal/data/ent/migrate"
	"github.com/novel_crawler/pkg/conf"
	"log"
)

type Config struct {
	Driver string `json:"driver" yaml:"driver"`
	Source string `json:"source" yaml:"source"`
}

func (c *Config) Load() error {
	return conf.Load("data.database", c)
}

var EntClient *ent.Client

func Init() {
	var cfg = new(Config)
	if err := cfg.Load(); err != nil {
		log.Fatalf("data load config fail: %v", err)
	}

	var err error
	EntClient, err = ent.Open(cfg.Driver, cfg.Source)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := EntClient.Schema.Create(context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		migrate.WithFixture(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
