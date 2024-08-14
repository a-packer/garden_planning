package main

import (
  "fmt"
  // "net/http"
  // "react-go-app/backend"
  // "github.com/gorilla/mux"
)

func main() {
  // db, err := backend.InitDB()
  // if err != nil {
  //  fmt.Println("Error initializing database:", err)
  //  return
  // }
  // defer db.Close()

  // router := mux.NewRouter()

  // API endpoints
  // router.HandleFunc("/login", backend.LoginHandler(db)).Methods(http.MethodPost)
  // router.HandleFunc("/register", backend.RegisterHandler(db)).Methods(http.MethodPost)
  // router.HandleFunc("/plants", backend.GetAllPlantsHandler(db)).Methods(http.MethodGet)
  // router.HandleFunc("/gardens", backend.GetUserPlantsHandler(db)).Methods(http.MethodGet)

  // Serve the React app
  // router.PathPrefix("/").Handler(http.FileServer(http.Dir("./build/")))

  fmt.Println("Server is listening on :8080...")
  // http.ListenAndServe(":8080", router)
}