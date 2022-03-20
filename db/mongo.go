package db

import (
	config "kienmatu/go-todos/config"

	mgo "gopkg.in/mgo.v2"
)

var instance *mgo.Session

var err error

// GetInstance return copy of db session
func GetMongoInstance(c *config.Configuration) *mgo.Session {

	if instance == nil {
		instance, err = mgo.Dial(c.DatabaseConnectionURL)
		if err != nil {
			panic(err)
		}
	}
	return instance.Copy()

}
