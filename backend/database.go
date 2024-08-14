package backend

import (
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
  "golang.org/x/crypto/bcrypt"
)

func InitDB() (*sql.DB, error) {
  // Note this does not establish any connections to the database stored in the file "./webapp.db".
  // The sql.Open function simply prepares the database handle.
  db, err := sql.Open("sqlite3", "./webapp.db")
  if err != nil {
    return nil, err
  }

  // Use defer to close the database connection when the function exits
  // statement is placed once right after opening the database, ensuring the connection is closed when the function exits, either after a successful operation or an error.
  defer db.Close()

  // Begin a transaction
  tx, err := db.Begin()
  if err != nil {
    return nil, err
  }

  // Execute SQL statements in the transaction
  _, err = tx.Exec(`
    CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      password TEXT NOT NULL,
      frostdate TEXT
    )
  `)
  if err != nil {
    tx.Rollback()
    return nil, err
  }

  _, err = tx.Exec(`
    CREATE TABLE IF NOT EXISTS plants (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      plantName TEXT NOT NULL,
      numWeeksIn INT,
      weeksRelOut INT,
      totalGrowth INT
    )
  `)
  if err != nil {
    tx.Rollback()
    return nil, err
  }

  _, err = tx.Exec(`
    CREATE TABLE IF NOT EXISTS gardens (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      userId INTEGER NOT NULL,
      plantId INTEGER NOT NULL,
      FOREIGN KEY (userId) REFERENCES users(id),
      FOREIGN KEY (plantId) REFERENCES plants(id)
    )
  `)
  if err != nil {
    tx.Rollback()
    return nil, err
  }

  _, err = tx.Exec(`
    INSERT INTO users (name, password) VALUES ('testUser01', 'afljas;fkasf');
    INSERT INTO users (name, password) VALUES ('testUser02', 'afljas;fkasf');
    INSERT INTO users (name, password) VALUES ('testUser03', 'afljas;fkasf');
  `)
  if err != nil {
    tx.Rollback()
    return nil, err
  }

  _, err = tx.Exec(`
    INSERT INTO plants (plantName, numWeeksIn, weeksRelOut, totalGrowth) VALUES ('Basil', 5, 1, 4);
    INSERT INTO plants (plantName, numWeeksIn, weeksRelOut, totalGrowth) VALUES ('Apple', 10, 4, 2);
    INSERT INTO plants (plantName, numWeeksIn, weeksRelOut, totalGrowth) VALUES ('Celery', 6, 3, 3);
  `)
  if err != nil {
    tx.Rollback()
    return nil, err
  }

  _, err = tx.Exec(`
    INSERT INTO gardens (userId, plantId) VALUES (1, 2);
    INSERT INTO gardens (userId, plantId) VALUES (1, 3);
    INSERT INTO gardens (userId, plantId) VALUES (2, 1);
    INSERT INTO gardens (userId, plantId) VALUES (3, 2);
  `)
  if err != nil {
    tx.Rollback()
    return nil, err
  }

  // Commit the transaction
  err = tx.Commit()
  if err != nil {
    return nil, err
  }

  // Close the db at the end of function execution
  err = db.Close()
  if err != nil {
    return nil, err
  }

  // Reopen the database connection to return the handle
  db, err = sql.Open("sqlite3", "./webapp.db")
  if err != nil {
    return nil, err
  }

  // Else return the database handle and a nil error to indicate success.
  return db, nil
}

type Plant struct {
  ID          int    `json:"id"`
  PlantName   string `json:"plantName"`
  NumWeeksIn  int    `json:"numWeeksIn"`
  WeeksRelOut int    `json:"weeksRelOut"`
  TotalGrowth int    `json:"totalGrowth"`
}

type Garden struct {
  ID      int `json:"id"`
  UserID  int `json:"userId"`
  PlantID int `json:"plantId"`
}

// VerifyLogin checks if the provided username and password are correct
func VerifyLogin(db *sql.DB, name, password string) bool {
  var hashedPassword string
  err := db.QueryRow("SELECT password FROM users WHERE name = ?", name).Scan(&hashedPassword)
  if err != nil {
    // If there's an error querying the database, the user doesn't exist or another issue occurred
    return false
  }

  // Compare the provided password with the hashed password
  err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
  return err == nil
}

// GetUserByName retrieves a user from the database by their name
func GetUserByName(db *sql.DB, name string) (int, string, error) {
  var id int
  var password string
  err := db.QueryRow("SELECT id, password FROM users WHERE name = ?", name).Scan(&id, &password)
  if err != nil {
    return 0, "", err
  }
  return id, password, nil
}

// CreateUser adds a new user to the database
func CreateUser(db *sql.DB, name, hashedPassword string) error {
  _, err := db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", name, hashedPassword)
  return err
}

// GetAllPlants retrieves all plants from the database
func GetAllPlants(db *sql.DB) ([]Plant, error) {
  rows, err := db.Query("SELECT id, plantName, numWeeksIn, weeksRelOut, totalGrowth FROM plants")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var plants []Plant
  for rows.Next() {
    var plant Plant
    if err := rows.Scan(&plant.ID, &plant.PlantName, &plant.NumWeeksIn, &plant.WeeksRelOut, &plant.TotalGrowth); err != nil {
      return nil, err
    }
    plants = append(plants, plant)
  }
  return plants, nil
}

// GetUserGardens retrieves gardens associated with a user from the database
func GetUserGardens(db *sql.DB, userID int) ([]Garden, error) {
  rows, err := db.Query("SELECT id, userId, plantId FROM gardens WHERE userId = ?", userID)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var gardens []Garden
  for rows.Next() {
    var garden Garden
    if err := rows.Scan(&garden.ID, &garden.UserID, &garden.PlantID); err != nil {
      return nil, err
    }
    gardens = append(gardens, garden)
  }
  return gardens, nil
}

