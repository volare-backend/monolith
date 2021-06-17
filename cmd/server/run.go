package main

import (
	"github.com/srvc/appctx"
	"github.com/volare-backend/monolith/app/server"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/riita10069/roche/pkg/rocheserver/mysql"
)

func run() error {
	// Application context
	ctx := appctx.Global()
	sql, err := mysql.GetMySQL()
	if err != nil {
		return err
	}
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
		// TODO
		server.NewSampleServiceServer(sql),
		),
	)
	return s.Serve(ctx)
}
