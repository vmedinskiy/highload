package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/vmedinskiy/highload/internal/cmd/server"
)

func main() {
	server.Execute()
}
