package main

import (
	"food_delivery_api/config"
	"food_delivery_api/pkg/adding"
	"food_delivery_api/pkg/editing"
	"food_delivery_api/pkg/http/rest"
	"food_delivery_api/pkg/listing"
	"food_delivery_api/pkg/removing"
	"food_delivery_api/pkg/storage/mysql"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	goEnv := strings.ToLower(os.Getenv("GO_ENV"))
	if goEnv == "" {
		goEnv = "local"
	}

	// Load the config
	LoadConfig(goEnv)

	log.Print(strings.ToUpper(config.AppName), " is warming up ...")

	// Run the server
	run(goEnv)
}

func LoadConfig(goEnv string) {
	var arg string

	if config := os.Getenv("CONFIG_FILE"); config != "" {
		arg = config
	} else if len(os.Args) == 2 {
		arg = "config/config." + os.Args[1] + ".yaml"
	} else {
		arg = "config/config." + goEnv + ".yaml"
	}

	err := config.Load(arg)
	if err != nil {
		log.Fatal("Error: config failed to load - ", err)
	}

	log.Println("Load config from", arg)
}

func run(goEnv string) {
	// MySQL setup
	rmy, err := mysql.NewStorage(config.My)
	if err != nil {
		log.Fatal("Error: Database failed to connect (", config.My.DSN, ") - ", err)
	}

	// Handler setup
	adder := adding.NewService(rmy)
	editer := editing.NewService(rmy)
	lister := listing.NewService(rmy)
	remover := removing.NewService(rmy)

	r := rest.Handler(adder, editer, lister, remover)

	host := config.Glb.Serv.Host
	if host == "" {
		host = GetLocalIP()
		config.Glb.Serv.Host = host
	}

	log.Println("Server Running on", goEnv, "environment, (REST APIs) listening on", host+":"+config.Serv.Port)
	log.Fatal("Error: Server failed to run - ", r.Run(config.Serv.Host+":"+config.Serv.Port))
}

func GetLocalIP() string {
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range address {
		if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				return inet.IP.String()
			}
		}
	}

	return ""
}
