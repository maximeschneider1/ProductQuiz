package handler

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"productManagement/dao"
)

func (s *server) HandleAllScores() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		arg := ps.ByName("arg")
		var data interface{}
		var err error

		if govalidator.IsInt(arg)  {
			data, err = dao.QueryAllScores(s.database, arg)
			if err != nil {
				log.Println("Error querying questions", err)
			}
		} else {
			data, err = dao.QueryUserScore(s.database, arg)
			if err != nil {
				log.Println("Error querying user score for user", arg)
			}
		}

		resp := response{}
		resp.Data = data
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = nil
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

func (s *server) HandleUserScore() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		username := ps.ByName("limit")

		data, err := dao.QueryUserScore(s.database, username)
		if err != nil {
			log.Println("Error querying questions", err)
		}

		resp := response{}
		resp.Data = data
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = nil
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}