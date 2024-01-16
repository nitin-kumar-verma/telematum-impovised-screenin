package startup

import (
	"screening/internal/config"
	"screening/internal/db"
)

/*Write all the code which is supposed to be invoked during app startup
  like env loading, connection initialization*/

func Init() {
	config.LoadConfig()
	db.InitSQLConn()
}
