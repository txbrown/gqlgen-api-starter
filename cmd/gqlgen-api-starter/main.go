package main

import (
	"strings"

	"github.com/txbrown/gqlgen-api-starter/internal/logger"
	"github.com/txbrown/gqlgen-api-starter/internal/orm"
	"github.com/txbrown/gqlgen-api-starter/pkg/server"
	"github.com/txbrown/gqlgen-api-starter/pkg/utils"
)

func main() {
	var serverconf = &utils.ServerConfig{
		Host:          utils.MustGet("SERVER_HOST"),
		Port:          utils.MustGet("SERVER_PORT"),
		URISchema:     utils.MustGet("SERVER_URI_SCHEMA"),
		Version:       utils.MustGet("SERVER_PATH_VERSION"),
		SessionSecret: utils.MustGet("SESSION_SECRET"),
		JWT: utils.JWTConfig{
			Secret:    utils.MustGet("AUTH_JWT_SECRET"),
			Algorithm: utils.MustGet("AUTH_JWT_SIGNING_ALGORITHM"),
		},
		GraphQL: utils.GQLConfig{
			Path:                utils.MustGet("GQL_SERVER_GRAPHQL_PATH"),
			PlaygroundPath:      utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH"),
			IsPlaygroundEnabled: utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED"),
		},
		Database: utils.DBConfig{
			Dialect:     utils.MustGet("GORM_DIALECT"),
			DSN:         utils.MustGet("GORM_CONNECTION_DSN"),
			SeedDB:      utils.MustGetBool("GORM_SEED_DB"),
			LogMode:     utils.MustGetBool("GORM_LOGMODE"),
			AutoMigrate: utils.MustGetBool("GORM_AUTOMIGRATE"),
		},
		AuthProviders: []utils.AuthProvider{
			utils.AuthProvider{
				Provider:  "google",
				ClientKey: utils.MustGet("PROVIDER_GOOGLE_KEY"),
				Secret:    utils.MustGet("PROVIDER_GOOGLE_SECRET"),
			},
			utils.AuthProvider{
				Provider:  "auth0",
				ClientKey: utils.MustGet("PROVIDER_AUTH0_KEY"),
				Secret:    utils.MustGet("PROVIDER_AUTH0_SECRET"),
				Domain:    utils.MustGet("PROVIDER_AUTH0_DOMAIN"),
				Scopes:    strings.Split(utils.MustGet("PROVIDER_AUTH0_SCOPES"), ","),
			},
			utils.AuthProvider{
				Provider:  "facebook",
				ClientKey: utils.MustGet("PROVIDER_FACEBOOK_KEY"),
				Secret:    utils.MustGet("PROVIDER_FACEBOOK_SECRET"),
			},
			utils.AuthProvider{
				Provider:  "twitter",
				ClientKey: utils.MustGet("PROVIDER_TWITTER_KEY"),
				Secret:    utils.MustGet("PROVIDER_TWITTER_SECRET"),
			},
		},
	}
	orm, err := orm.New(serverconf)

	if err != nil {
		logger.Panic(err)
	}
	server.Run(serverconf, orm)
}
