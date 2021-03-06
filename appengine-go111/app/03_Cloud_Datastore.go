package main

import (
	"cloud.google.com/go/datastore"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

const PROJ_NAME = "chida-test-012"

// 書き込みテスト
func handleCloudDatastore(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	dsClient, err := datastore.NewClient(c, PROJ_NAME)
	if err != nil {
		log.Errorf(c, "datastore client create error. %s", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	b := Benchmarker{}

	// 30回実行
	for i := 0; i <= 30 ; i++ {
		count := fmt.Sprintf("%d", i + 1)
		k := datastore.NameKey("Benkyo", "BK" + count, nil)
		e := Benkyo{
			Name: "Benkyo-name",
		}
		// benchmark実行
		b.Do(c, func() {
			_, err = dsClient.Put(c, k, &e)
		})
		if err != nil {
			log.Errorf(c, "datastore put error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "datastore put success.")
	log.Infof(c, "Result : %s", b.Result())

	// response
	g.String(http.StatusOK, "03_Cloud_Datastore")
}

// 読み出しテスト
func handleCloudDatastoreRead(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	dsClient, err := datastore.NewClient(c, PROJ_NAME)
	if err != nil {
		log.Errorf(c, "datastore client create error. %s", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	b := Benchmarker{}

	// 30回実行
	for i := 0; i <= 30 ; i++ {
		k := datastore.NameKey("Benkyo", "BK1", nil)
		e := Benkyo{}
		// benchmark実行
		b.Do(c, func() {
			err = dsClient.Get(c, k, &e)
		})
		if err != nil {
			log.Errorf(c, "datastore get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "datastore get success.")
	log.Infof(c, "Result : %s", b.Result())

	// response
	g.String(http.StatusOK, "02_Cloud_Datastore/get")
}