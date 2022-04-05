package main

import "net/http"

func main() {
	server := http.Server{
		DB:             db,
		StagingService: &stagingService,
		TargetService: &targetService,
		TelescopeMappingService: &telescopeMappingService,
		Port: port,
	}

	log.Info("starting dsw update server http api, version: 1.0.10")
	server.Start()
}
