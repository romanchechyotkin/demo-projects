package server

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	base             *http.Server
	router           *gin.Engine
	pgClient         *pgxpool.Pool
	redisClient      *redis.Client
	clickhouseClient *sql.DB
}

var counter int32

func New(pg *pgxpool.Pool, rc *redis.Client, cc *sql.DB) *Server {
	srv := &Server{
		router:           gin.Default(),
		pgClient:         pg,
		redisClient:      rc,
		clickhouseClient: cc,
	}

	srv.router.GET("/status", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "status")
	})
	srv.router.POST("/projects", func(ctx *gin.Context) {
		atomic.AddInt32(&counter, 1)
		name := "project" + strconv.Itoa(int(counter))
		exec, err := srv.pgClient.Exec(ctx, "insert into projects (name) values ($1)", name)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		log.Println(exec.RowsAffected())

		ctx.String(http.StatusOK, name)
	})

	srv.base = &http.Server{
		Addr:    ":8080",
		Handler: srv.router,
	}

	return srv
}

func (s *Server) Run() {
	log.Fatal(s.base.ListenAndServe())
}
