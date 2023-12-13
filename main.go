package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/wutipong/gallery3-backend/docs"
	"github.com/wutipong/gallery3-backend/fsmeta"
)

//go:generate go run -mod=mod github.com/swaggo/swag/cmd/swag@latest init

// @title           Gallery3 API
// @version         3.0
// @description     API Server for gallery3

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	if err := godotenv.Overload(); err == nil {
		log.Info().Msg("Use .env file.")
	}

	address := ":8872"
	if v, b := os.LookupEnv("BACKEND_ADDRESS"); b {
		address = v
	}

	dataPath := "./data"
	if v, b := os.LookupEnv("DATA_PATH"); b {
		dataPath = v
	}

	fsmeta.Init(dataPath)

	if _, b := os.LookupEnv("MANGAWEB_DEVELOPMENT"); b {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
			Level(zerolog.DebugLevel)
	}

	router := httprouter.New()
	handler := cors.Default().Handler(router)

	router.GET("/doc/:any", swaggerHandler)
	router.GET("/dir/*path", fsmeta.Handler)
	log.Info().Msg("Server starts.")

	if err := http.ListenAndServe(address, handler); err != nil {
		log.Error().AnErr("error", err).Msg("Starting server fails")
		return
	}

	log.Info().Msg("shutting down the server")
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
