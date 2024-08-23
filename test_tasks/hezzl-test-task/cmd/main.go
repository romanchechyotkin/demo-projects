package main

import (
	"github.com/romanchechyotkin/hezzl-test-task/internal/server"
	"github.com/romanchechyotkin/hezzl-test-task/pkg/clickhouse"
	"github.com/romanchechyotkin/hezzl-test-task/pkg/postgresql"
	"github.com/romanchechyotkin/hezzl-test-task/pkg/redis"
)

func main() {
	pgCfg := postgresql.Config{
		User:     "postgres",
		Password: "5432",
		Host:     "localhost",
		Port:     "5432",
		Database: "hezzl",
	}

	pgClient := postgresql.New(&pgCfg)

	redisCfg := redis.Config{
		Host:     "localhost",
		Port:     "6379",
		Password: "6379",
		Database: 0,
	}

	redisClient := redis.NewClient(&redisCfg)

	clickhouseCfg := clickhouse.Config{
		Host:     "localhost",
		Port:     "8123",
		User:     "clickhouse",
		Password: "8123",
	}

	clickhouseClient := clickhouse.NewClient(&clickhouseCfg)

	srv := server.New(pgClient, redisClient, clickhouseClient)
	srv.Run()
}
