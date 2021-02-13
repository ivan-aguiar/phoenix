package main
 
import (
    "log"
    "github.com/ivan-aguiar/phoenix/handlers"
    "github.com/ivan-aguiar/phoenix/bd"
)
func main(){
    if bd.checkConnection() == 0 {
        log.Fatal("No connection with database")
        return
    }
    handlers.Handlers()
 
}
