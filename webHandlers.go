package main

import (
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Add new device manually via web-interface
func addDevice(c *gin.Context) {
	c.Request.ParseMultipartForm(1000)
	for key, value := range c.Request.PostForm {
		log.Println(key, value)
	}
	log.Println("String", c.PostForm("sn"))
	/*session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	}*/
	session := getSession()
	defer session.Close()
	var e electronics
	e.SerialNumber = strings.ToLower(c.PostForm("sn"))
	e.Manufacturer = strings.ToLower(c.PostForm("manufacturer"))
	e.Model = strings.ToLower(c.PostForm("model"))
	e.Type = strings.ToLower(c.PostForm("type"))
	if e.Type == "0" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "You didn't specified device type!"})
	} else {
		/*if err != nil {
			log.Println(err.Error())
			c.HTML(http.StatusBadRequest, "error", gin.H{"error": err.Error()})
		}*/
		el := session.DB("inventory").C("electronics")
		err := el.Insert(e)
		if err != nil {
			if mgo.IsDup(err) {
				c.HTML(http.StatusAlreadyReported, "error.html", gin.H{"error": err.Error()})
			} else {
				log.Println(err.Error())
				c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
			}
		}
		if err == nil {
			c.HTML(http.StatusOK, "success.html", gin.H{"error": ""})
		}
	}
}

// Show list of all PC
func allPC(c *gin.Context) {
	/*session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	}*/
	session := getSession()
	defer session.Close()
	var list pclist
	pc := session.DB("inventory").C("pc")
	err := pc.Find(bson.M{}).All(&list)
	if err != nil {
		log.Println(err.Error())
	} else {
		sort.Slice(list, func(i, j int) bool {
			return list[i].ComputerSystem.Name < list[j].ComputerSystem.Name
		})
		c.HTML(http.StatusOK, "pclist.html", gin.H{"pc": list})
	}
}

// Show list of all devices
func allDevices(c *gin.Context) {
	/*session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	}*/
	session := getSession()
	defer session.Close()
	var list []electronics
	pc := session.DB("inventory").C("electronics")
	err := pc.Find(bson.M{}).All(&list)
	if err != nil {
		log.Println(err.Error())
	} else {
		sort.Slice(list, func(i, j int) bool {
			return list[i].Manufacturer < list[j].Manufacturer
		})
		c.HTML(http.StatusOK, "devlist.html", gin.H{"devices": list})
	}
}

// Show home page
func homePage(c *gin.Context) {
	/*session, err := mgo.Dial("172.16.2.204")
	if err != nil {
		log.Println(err.Error())
	}*/
	session := getSession()
	defer session.Close()
	var s sysinfo
	var e electronics
	var n networkDevice
	pc := session.DB("inventory").C("pc")
	el := session.DB("inventory").C("electronics")
	net := session.DB("inventory").C("network")
	pcCount, err := pc.Count()
	if err != nil {
		log.Println(err.Error())
	}
	devCount, err := el.Count()
	if err != nil {
		log.Println(err.Error())
	}
	netCount, err := net.Count()
	if err != nil {
		log.Println(err.Error())
	}
	search, haveSearch := c.GetQuery("search")
	if haveSearch && len(search) > 0 {
		err = pc.Find(bson.M{"baseboard.serialnumber": search}).One(&s)
		if err != nil {
			log.Println(err.Error())
		}
		if len(s.ComputerSystem.Name) > 0 {
			c.HTML(http.StatusOK, "pc.html", gin.H{
				"baseboard":      s.BaseBoard,
				"bios":           s.Bios,
				"computersystem": s.ComputerSystem,
				"cpu":            s.CPU,
				"diskdrive":      s.DiskDrive,
				"display":        s.Display,
				"dt":             s.Dt,
				"keyboard":       s.Keyboard,
				"mouse":          s.Mouse,
				"nic":            s.NIC,
				"os":             s.OS,
				"printer":        s.Printer,
				"tz":             s.Tz,
				"user":           s.User,
				"video":          s.Video})
		} else {
			err = el.Find(bson.M{"serialnumber": search}).One(&e)
			if err != nil {
				log.Println(err.Error())
			}
			if len(e.Model) > 0 {
				c.HTML(http.StatusOK, "device.html", gin.H{
					"Manufacturer": e.Manufacturer,
					"Model":        e.Model,
					"Purchase":     e.Purchase,
					"IP":           e.IP,
					"Name":         e.Name,
					"Dt":           e.Dt,
					"Printouts":    e.Printouts,
					"Cartridge":    e.Cartridge})
			} else {
				err = net.Find(bson.M{"serialnumber": search}).One(&n)
				if len(n.Model) > 0 {
					c.HTML(http.StatusOK, "netdevice.html", gin.H{
						"Manufacturer": n.Manufacturer,
						"Model":        n.Model,
						"MAC":          n.MAC,
						"IP":           n.IP,
						"Name":         n.Name,
						"Dt":           n.Dt,
						"SerialNumber": n.SerialNumber,
						"Uptime":       n.Uptime})
				} else {
					c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "There's no device with specified serial number"})
				}
			}
		}
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"cpu":         pcCount,
			"electronics": devCount,
			"network":     netCount})
	}
}

// Show list of all network devices
func allNetworkDevices(c *gin.Context) {
	session := getSession()
	defer session.Close()
	var list []networkDevice
	pc := session.DB("inventory").C("network")
	err := pc.Find(bson.M{}).All(&list)
	if err != nil {
		log.Println(err.Error())
	} else {
		sort.Slice(list, func(i, j int) bool {
			return list[i].Manufacturer < list[j].Manufacturer
		})
		c.HTML(http.StatusOK, "netlist.html", gin.H{"devices": list})
	}
}
