package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/samuel/go-zookeeper/zk"
)

type Controller struct {
	zkConn *zk.Conn
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := cleanPath(req.URL.Path)
	switch req.Method {
	case "GET":
		log.Printf("retrieving node %q", path)
		contents, _, err := c.zkConn.Get(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(contents)
		return
	case "PUT", "POST":
		log.Printf("updating node %q", path)
		exists, stats, err := c.zkConn.Exists(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("invalid body: %q", reqBody)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !exists {
			_, err := c.zkConn.Create(path, reqBody, 0, zk.WorldACL(zk.PermAll))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		} else {
			if _, err := c.zkConn.Set(path, reqBody, stats.Version); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
		return
	case "DELETE":
		log.Printf("deleting node %q", path)
		exists, stats, err := c.zkConn.Exists(path)
		if !exists || err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := c.zkConn.Delete(path, stats.Version); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	case "OPTIONS":
		// TODO: Implement CORs correctly
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
}

func cleanPath(str string) string {
	if str == "" || str == "/" {
		return str
	}
	return path.Clean(str)
}
