package main

import (
	"flag"
	"log"
	"context"

	m2m "github.com/justledbetter/mail2most/lib"

        "github.com/aws/aws-lambda-go/lambda"
)

var LAMBDA_BUILD string = ""

func LambdaHandler(ctx context.Context, config m2m.Config) (string, error) {

	// Lambda does not support writing files, so we need to force some reasonable defaults.
	//
	config.NoStateFile = true
	config.General.File = ""
	config.Logging.Logtype = "json"
	config.Logging.Output = "stdout"

	m, err := m2m.NewFromJson(config)
	if err != nil {
		return "done", err
	}

	err = m.Run()
	return "done", err
}


func main() {
	if LAMBDA_BUILD != "" {
		log.Println("Executing Lambda runtime")
		lambda.Start(LambdaHandler)

	} else {

		confFile := flag.String("c", "conf/mail2most.conf", "path to config file")
		flag.Parse()

		m, err := m2m.New(*confFile)
		if err != nil {
			log.Fatal(err)
		}

		err = m.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
