package main

import (
	"github.com/carrot-systems/cs-session/src/adapters/gateway"
	"github.com/carrot-systems/cs-session/src/adapters/persistence/postgres"
	"github.com/carrot-systems/cs-session/src/adapters/rest"
	"github.com/carrot-systems/cs-session/src/config"
	"github.com/carrot-systems/cs-session/src/core/usecases"
	configurationClient "github.com/carrot-systems/csl-configuration-client"
	discoveryClient "github.com/carrot-systems/csl-discovery-client"
	env "github.com/carrot-systems/csl-env"
	"gorm.io/gorm"
	"log"
)

func main() {
	env.LoadEnv()

	discovery := discoveryClient.NewClient()
	err := discovery.Register("session")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	configuration := configurationClient.NewClient(discovery)
	err = configuration.LoadConfiguration()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	ginConfiguration := config.LoadGinConfiguration()

	dbConfig := config.LoadGormConfiguration()

	var sessionRepo usecases.SessionRepo
	var userClientGateway usecases.UserClientGateway

	var db *gorm.DB
	if dbConfig.Engine == "POSTGRES" {
		db = postgres.StartGormDatabase(dbConfig)
		err = postgres.Migrate(db, "./migrations", "sessions_migration")
		if err != nil {
			log.Fatalln(err.Error())
		}
		sessionRepo = postgres.NewSessionRepo(db)
	}

	userClientGateway = gateway.NewUserClientGateway(discovery)

	usecasesHandler := usecases.NewInteractor(sessionRepo, userClientGateway)

	restServer := rest.NewServer(ginConfiguration)
	routesHandler := rest.NewRouter(usecasesHandler)

	rest.SetRoutes(restServer.Router, routesHandler)
	restServer.Start()
}
