package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//// Receive request by ID
// Read request body for ID
// GET request to database
// Put result in object
// Write result in response

// handleQuestionByID returns a specific question for given ID
func (s *server) handleQuestionByID() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		data := fmt.Sprint(r.Header.Get("Oh"))
		fmt.Println(ps.ByName("limit"))
		//results, err := dao.QueryAllQuotes(s.database)
		//if err != nil {
		//	log.Println(err)
		//}
		resp := response{}
		resp.Data = append(resp.Data, data)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = nil
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

func (s *server) HandleAllQuestions() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")


	}
}