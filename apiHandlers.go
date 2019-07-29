package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Add a new document
func addPC(c *gin.Context) {
	/* session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	} */
	session := getSession()
	defer session.Close()
	var s sysinfo
	err := c.ShouldBindJSON(&s)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	pc := session.DB("inventory").C("pc")
	err = pc.Insert(s)
	if err != nil {
		if mgo.IsDup(err) {
			c.JSON(http.StatusAlreadyReported, gin.H{"error": err.Error()})
		} else {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"error": ""})
	}
}

// Get list of all documents
func getAllPC(c *gin.Context) {
	/* session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	} */
	session := getSession()
	defer session.Close()
	var list []sysinfo
	pc := session.DB("inventory").C("pc")
	err := pc.Find(bson.M{}).All(&list)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"error": "", "result": list})
	}
}

// Get one document identified by ID or Serial Number
func getOnePC(c *gin.Context) {
	/* session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	} */
	session := getSession()
	defer session.Close()
	var s sysinfo
	pc := session.DB("inventory").C("pc")
	id, haveID := c.GetQuery("id")
	sn, haveSN := c.GetQuery("sn")
	name, haveName := c.GetQuery("name")
	if haveID {
		err := pc.FindId(bson.ObjectIdHex(id)).One(&s)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"error": "", "result": s})
	} else if haveSN {
		err := pc.Find(bson.M{"baseboard.serialnumber": sn}).One(&s)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "", "result": s})
		}
	} else if haveName {
		log.Println("Got name:", name)
		err := pc.Find(bson.M{"computersystem.name": name}).One(&s)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "", "result": s})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have to specify item ID or baseboard serial number"})
	}
}

// Update an existing document
func updatePC(c *gin.Context) {
	/* session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	} */
	session := getSession()
	defer session.Close()
	var s sysinfo
	err := c.ShouldBindJSON(&s)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	pc := session.DB("inventory").C("pc")
	sn, haveSN := c.GetQuery("sn")
	if haveSN {
		err = pc.Update(bson.M{"baseboard.serialnumber": sn}, &s)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": ""})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have to specify baseboard serial number"})
	}
}

// Add new eletronic device
func addElectronics(c *gin.Context) {
	/* session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	} */
	session := getSession()
	defer session.Close()
	var e electronics
	err := c.ShouldBindJSON(&e)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	el := session.DB("inventory").C("electronics")
	err = el.Insert(e)
	if err != nil {
		if mgo.IsDup(err) {
			c.JSON(http.StatusAlreadyReported, gin.H{"error": err.Error()})
		} else {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"error": ""})
	}
}

// Update an existing document
func updateElectronics(c *gin.Context) {
	/* session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	} */
	session := getSession()
	defer session.Close()
	var e electronics
	err := c.ShouldBindJSON(&e)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	el := session.DB("inventory").C("electronics")
	sn, haveSN := c.GetQuery("sn")
	if haveSN {
		err = el.Update(bson.M{"serialnumber": sn}, &e)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": ""})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have to specify device serial number"})
	}
}

// Add new network device
func addNetworkDevice(c *gin.Context) {
	session := getSession()
	defer session.Close()
	var n networkDevice
	err := c.ShouldBindJSON(&n)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	el := session.DB("inventory").C("network")
	err = el.Insert(n)
	if err != nil {
		if mgo.IsDup(err) {
			c.JSON(http.StatusAlreadyReported, gin.H{"error": err.Error()})
		} else {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"error": ""})
	}
}

// Update an existing document
func updateNetworkDevice(c *gin.Context) {
	session := getSession()
	defer session.Close()
	var n networkDevice
	err := c.ShouldBindJSON(&n)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	el := session.DB("inventory").C("network")
	sn, haveSN := c.GetQuery("sn")
	if haveSN {
		err = el.Update(bson.M{"serialnumber": sn}, &n)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": ""})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have to specify device serial number"})
	}
}
