package main

import (
	"context"
	"fmt"

	sqlc "github.com/vietddude/tss-impl/db/sqlc"
	"github.com/vietddude/tss-impl/utils"

	"github.com/vietddude/tss-impl/config"
	"github.com/vietddude/tss-impl/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	pool, err := db.InitDB(&cfg.DB)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	shareData, err := sqlc.New(pool).GetShareKey1(ctx, utils.StringToPgUUID("4eb3d697-09b3-404c-8b58-a49475cb0dfa"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", shareData)
}
