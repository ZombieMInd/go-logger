package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ZombieMInd/go-logger/internal/logger"
)

func (s *server) handleLog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &logger.LogRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		err := req.Validate()
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		req.IP = r.RemoteAddr

		err = s.services.Log.Save(req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, nil)
	}
}

func (s *server) handleLogRaw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		err = s.services.Log.SaveRaw(body, r.RemoteAddr)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, nil)
	}
}
