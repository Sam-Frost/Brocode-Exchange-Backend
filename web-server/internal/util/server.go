package util

import (
	db "github.com/Sam-Frost/db/generated"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	Router *gin.Engine
	DB     *pgxpool.Pool
	Query  *db.Queries
}
