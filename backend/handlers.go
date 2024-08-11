// package backend

// import (
//   "encoding/json"
//   "fmt"
//   "net/http"

//   "database/sql"

//   "golang.org/x/crypto/bcrypt"
// )

// // LoginHandler handles user login requests
// func LoginHandler(db *sql.DB) http.HandlerFunc {
//   return func(w http.ResponseWriter, r *http.Request) {
//     var user struct {
//       Name     string `json:"name"`
//       Password string `json:"password"`
//     }
//     if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//       http.Error(w, "Invalid request payload", http.StatusBadRequest)
//       return
//     }

//     if VerifyLogin(db, user.Name, user.Password) {
//       w.WriteHeader(http.StatusOK)
//       fmt.Fprintln(w, "Login successful")
//     } else {
//       w.WriteHeader(http.StatusUnauthorized)
//       fmt.Fprintln(w, "Invalid username or password")
//     }
//   }
// }

// // RegisterHandler handles user registration requests
// func RegisterHandler(db *sql.DB) http.HandlerFunc {
//   return func(w http.ResponseWriter, r *http.Request) {
//     var user struct {
//       Name     string `json:"name"`
//       Password string `json:"password"`
//     }
//     if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//       http.Error(w, "Invalid request payload", http.StatusBadRequest)
//       return
//     }

//     hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
//     if err != nil {
//       http.Error(w, "Error creating user", http.StatusInternalServerError)
//       return
//     }

//     _, err = db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", user.Name, hashedPassword)
//     if err != nil {
//       http.Error(w, "Error creating user", http.StatusInternalServerError)
//       return
//     }

//     w.WriteHeader(http.StatusCreated)
//     fmt.Fprintln(w, "User registered successfully")
//   }
// }

// // GetAllPlantsHandler handles requests for retrieving all plants
// func GetAllPlantsHandler(db *sql.DB) http.HandlerFunc {
//   return func(w http.ResponseWriter, r *http.Request) {
//     rows, err := db.Query("SELECT id, name, type FROM plants")
//     if err != nil {
//       http.Error(w, "Error fetching plants", http.StatusInternalServerError)
//       return
//     }
//     defer rows.Close()

//     var plants []struct {
//       ID   int    `json:"id"`
//       Name string `json:"name"`
//       Type string `json:"type"`
//     }

//     for rows.Next() {
//       var plant struct {
//         ID   int    `json:"id"`
//         Name string `json:"name"`
//         Type string `json:"type"`
//       }
//       if err := rows.Scan(&plant.ID, &plant.Name, &plant.Type); err != nil {
//         http.Error(w, "Error scanning plants", http.StatusInternalServerError)
//         return
//       }
//       plants = append(plants, plant)
//     }

//     if err := rows.Err(); err != nil {
//       http.Error(w, "Error processing plants", http.StatusInternalServerError)
//       return
//     }

//     w.Header().Set("Content-Type", "application/json")
//     if err := json.NewEncoder(w).Encode(plants); err != nil {
//       http.Error(w, "Error encoding plants", http.StatusInternalServerError)
//       return
//     }
//   }
// }

// // GetUserPlantsHandler handles requests for retrieving plants specific to a user
// func GetUserPlantsHandler(db *sql.DB) http.HandlerFunc {
//   return func(w http.ResponseWriter, r *http.Request) {
//     userID := r.URL.Query().Get("userID")
//     if userID == "" {
//       http.Error(w, "Missing userID parameter", http.StatusBadRequest)
//       return
//     }

//     rows, err := db.Query("SELECT id, name, type FROM plants WHERE user_id = ?", userID)
//     if err != nil {
//       http.Error(w, "Error fetching user plants", http.StatusInternalServerError)
//       return
//     }
//     defer rows.Close()

//     var plants []struct {
//       ID   int    `json:"id"`
//       Name string `json:"name"`
//       Type string `json:"type"`
//     }

//     for rows.Next() {
//       var plant struct {
//         ID   int    `json:"id"`
//         Name string `json:"name"`
//         Type string `json:"type"`
//       }
//       if err := rows.Scan(&plant.ID, &plant.Name, &plant.Type); err != nil {
//         http.Error(w, "Error scanning user plants", http.StatusInternalServerError)
//         return
//       }
//       plants = append(plants, plant)
//     }

//     if err := rows.Err(); err != nil {
//       http.Error(w, "Error processing user plants", http.StatusInternalServerError)
//       return
//     }

//     w.Header().Set("Content-Type", "application/json")
//     if err := json.NewEncoder(w).Encode(plants); err != nil {
//       http.Error(w, "Error encoding user plants", http.StatusInternalServerError)
//       return
//     }
//   }
// }
