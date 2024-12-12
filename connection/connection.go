/* package connection

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DSN = "root:@tcp(127.0.0.1:3306)/to-do?charset=utf8mb4&parseTime=True&loc=UTC"

func Connection() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Base de datos corriendo...")
}
*/

package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// const DSN = "root:@tcp(127.0.0.1:3306)/to-do?charset=utf8mb4&parseTime=True&loc=UTC"
// const DSN = "host=dpg-crq9rmqj1k6c738dgejg-a user=mendez password=BCso273EsUZZWP2RIfLt5YWHx411dbTc dbname=tasks_7vap port=5432"
const DSN = "postgresql://mendez:BCso273EsUZZWP2RIfLt5YWHx411dbTc@dpg-crq9rmqj1k6c738dgejg-a.oregon-postgres.render.com/tasks_7vap"

func Connection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Base de datos corriendo...")
}
