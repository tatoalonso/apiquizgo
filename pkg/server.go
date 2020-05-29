package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Server includes a router
type Server struct {
	Router *mux.Router
}

//NewServer creates a Server
func NewServer() Server {

	r := mux.NewRouter()

	r.HandleFunc("/", fetchQuiz)

	s := Server{
		Router: r,
	}

	return s

}

func fetchQuiz(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("quiz!\n"))
}
