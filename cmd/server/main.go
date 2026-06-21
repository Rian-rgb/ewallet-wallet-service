package main

import (
	"ewallet-wallet/infra"
	"ewallet-wallet/infra/grpc"
	"ewallet-wallet/internal/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	// load config
	infra.InitConfig()

	// load log
	infra.InitLogger()

	// load db
	postgresDB := infra.InitPostgresql()

	// load redis
	redisRepo := infra.InitRedis()

	// load jwt
	jwtManager := infra.InitJWT()

	// run grpc
	gRPCRegistry, cleanup := grpc.NewConnRegistry()
	defer cleanup()

	appDeps := &infra.AppDependencies{
		PostgresDB:   postgresDB,
		RedisRepo:    redisRepo,
		JWTManager:   jwtManager,
		GrpcRegistry: gRPCRegistry,
	}

	// Inject dependency
	dependencies := infra.DependencyInject(appDeps)

	// run http
	http.Serve(dependencies, appDeps)
}
