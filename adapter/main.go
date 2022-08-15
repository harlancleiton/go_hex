package main

import (
	"log"

	rpcserver "github.com/harlancleiton/gopersonalities/adapter/grpc"
	rpcuser "github.com/harlancleiton/gopersonalities/adapter/grpc/user"
)

func init() {
	log.Println("Initializing Application...")
}

func main() {
	grpcAdapter := rpcserver.NewAdapter(rpcuser.GRPCUserAdapter{})

	err := grpcAdapter.StartServer()

	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
