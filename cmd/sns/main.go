package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yanmoyy/our-tools/internal/sns"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg, err := sns.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	sns.StartRepl(cfg)
}

//
// func runCallbackServer() {
// 	fmt.Println("Listening on port 8080")
// 	callbackServer := &http.Server{
// 		Addr:              ":8080",
// 		ReadHeaderTimeout: 3 * time.Second,
// 	}
// 	log.Fatal(callbackServer.ListenAndServe())
// }
