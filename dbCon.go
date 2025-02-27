package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

/* Creates a new session if mgoSession is nil i.e there is no active mongo session.
If there is an active mongo session it will return a Clone*/
func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("172.16.2.204")
		if err != nil {
			log.Fatal("Failed to start the Mongo session")
		}
	}
	return mgoSession.Clone()
}
