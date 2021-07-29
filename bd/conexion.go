package bd

import (
	"database/sql"
)

func Conexion() (*sql.DB, error) {

	return sql.Open("mysql", "Server=Localhost:3306;Database=time_gest;Uid=root;Pwd=Milbert1981;")

}
