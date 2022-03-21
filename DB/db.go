package DB
import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Get_db()(*gorm.DB, error){
	dsn := "host=localhost user=postgres password=36network dbname=newuser port=5433"
	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println("Error Occured",err)
		return db,err
	}
	return db,nil
}