package main

import (
	"flutter-chat/config"
	"flutter-chat/database"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	if err := database.Connect(cfg); err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer database.DB.Close()

	if err := database.RunMigrations(cfg); err != nil {
		log.Fatal("Migration failed:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Flutter Chat backend is running!"))
	})

	log.Printf("Server listening on :%s\n", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		log.Fatal("Server Failed:", err)
	}
}
