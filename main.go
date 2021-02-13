package main
 
import (
    "log"
    "github.com/ivan-aguiar/phoenix/handlers"
    "github.com/ivan-aguiar/phoenix/db"
)
func main(){
    if db.CheckConnection() == 0 {
        log.Fatal("No connection with database")
        return
    }
    handlers.Handlers()
 
}
