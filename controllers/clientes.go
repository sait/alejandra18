package controllers

import (
	"encoding/json"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"sait.mx/demosqlx/model"
)

/*
	Lista
*/
func ListClientes(c *gin.Context) {
	clientes, err := model.ListClientes()
	if err != nil {
		c.JSON(500, "Error en el servidor")
		return
	}
	c.JSON(200, clientes)
}

/*
	Muestra
*/
func GetCliente(c *gin.Context){
	id := c.Param("id")
	cliente, err := model.GetCliente(id)
	if err != nil {
		c.JSON(500, "Error en el servidor")
		return
	}
	c.JSON(200, cliente)
}

/*
	Insertar
*/
func InsertCliente(c *gin.Context){
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	var cliente model.Cliente
	err = json.Unmarshal(body, &cliente)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	err = model.InsertCliente(cliente)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Insertado Correctamente")
}

/*
	Actualización
*/
func UpdateCliente(c *gin.Context){
	id := c.Param("id")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	var cliente model.Cliente
	err = json.Unmarshal(body, &cliente)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	cliente.ID = id
	err = model.UpdateCliente(cliente)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Actualizado Correctamente")
}

/*
	Eliminación
*/
func DeleteCliente(c *gin.Context){
	id := c.Param("id")
	err := model.DeleteCliente(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Eliminado Correctamente")
}