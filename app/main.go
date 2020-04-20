package main

import (
	"fmt"
	"log"
	"net/http"
	"web-wallet/app/config"
	"web-wallet/app/routers"
	"web-wallet/app/wallet"
)

var (
	configFileNmae = "app"
)

func main() {
	config.LoadConfig(configFileNmae)
	wallet.InitWallet()

	routes := routers.InitRouters()
	addr := fmt.Sprintf(":%d", config.GetInt("app.http_port"))

	srv := http.Server{
		Addr:    addr,
		Handler: routes,
	}

	log.Println("Now the http server listen on:", addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Server listen err:", err)
	}
}
