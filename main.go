package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

const (
	portFlag      = "port"
	jsonInputFlag = "jsonInput"
)

func main() {
	app := cli.NewApp()
	app.Name = "Kyber Reserve Exporter"
	app.Usage = "Export metrics for Prometheus of Kyber Reserve"
	app.Version = "0.0.1"
	app.Action = runtHTTPServer

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   portFlag,
			Usage:  "port number",
			EnvVar: "PORT",
			Value:  fmt.Sprintf("8080"),
		},
		cli.StringFlag{
			Name:   jsonInputFlag,
			Usage:  "JSON Input Flag",
			EnvVar: "JSON_INPUT",
			Value:  fmt.Sprintf("config.json"),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newLogger(c *cli.Context) (*zap.Logger, error) {
	return zap.NewProduction()
}

func runtHTTPServer(c *cli.Context) error {
	logger, err := newLogger(c)
	if err != nil {
		return err
	}
	defer logger.Sync()
	inputLocation := c.GlobalString(jsonInputFlag)
	jsonFile, err := os.Open(inputLocation)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var input []Input
	err = json.Unmarshal(byteValue, &input)
	if err != nil {
		return err
	}
	contractData := make([]Contract, len(input))
	proofData := make(map[string]map[string]Proof, len(input))
	for index, i := range input {
		proofData[i.Name] = make(map[string]Proof, len(i.Claims))
		contractData[index] = Contract{
			Name:     i.Name,
			Contract: i.Contract,
		}
		for k, v := range i.Claims {
			proofData[i.Name][strings.ToLower(k)] = v
		}
	}

	port := strings.Join([]string{":", c.GlobalString(portFlag)}, "")
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/contracts", func(c *gin.Context) {
		c.JSON(http.StatusOK, contractData)
	})
	r.GET("/get-proof/:name/:address", func(c *gin.Context) {
		address := strings.ToLower(c.Param("address"))
		name := c.Param("name")
		c.JSON(http.StatusOK, proofData[name][address])
	})
	return r.Run(port)
}
