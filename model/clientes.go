package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Cliente struct {
	//las etiquetas de json no son iguales a las de la base de datos
	ID       string `db:"id"        json:"id"`
	Nombre   string `db:"Nombre"    json:"Nombre"`
	Apellido string `db:"Apellidos" json:"Apellidos"`
}

var DB *sqlx.DB

func OpenDB() {
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/pruebasgo")
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ListClientes() (clientes []Cliente, err error) {
	err = DB.Select(&clientes, `SELECT id, Nombre, Apellidos FROM clientes`)
	return
}

func GetCliente(clientID string) (cliente Cliente, err error) {
	err = DB.Get(&cliente, `SELECT id, Nombre, Apellidos FROM clientes
		WHERE id=?`, clientID)
	return
}

func InsertCliente(cliente Cliente) (err error) {
	_, err = DB.NamedExec(`INSERT INTO clientes (Nombre, Apellidos)
		VALUES(:Nombre, :Apellidos)`, &cliente)
	return
}

func UpdateCliente(cliente Cliente) (err error) {
	_, err = DB.NamedExec(`UPDATE clientes
		SET Nombre=:Nombre, Apellidos=:Apellidos
		WHERE id=:id`, cliente)
	return
}

func DeleteCliente(clienteID string) (err error) {
	_, err = DB.Exec(`DELETE FROM clientes
		WHERE id=?`, clienteID)
	return
}
