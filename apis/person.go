package apis

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/JanusHuang/gin-demo/models"
	"gopkg.in/gin-gonic/gin.v1"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	person := Person{FirstName: firstName, LastName: lastName}
	id, err := person.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("insert person with id {}", id)
	msg := fmt.Sprintf("insert successful %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	p := &Person{}
	persons, err := p.GetPersons()

	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}

func GetPersonApi(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"person": nil,
		})
		return
	}
	p := &Person{Id: id}
	person, err := p.GetPerson()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"person": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}

func UpdatePersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	person := Person{Id: id}
	err = c.Bind(&person)
	if err != nil {
		log.Fatalln(err)
	}

	count, err := person.UpdatePerson()

	msg := fmt.Sprintf("Update person %d successfully %d", person.Id, count)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DeletePersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}

	person := Person{Id: id}
	ra, err := person.DeletePerson()

	msg := fmt.Sprintf("Delete person %d successfully %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
