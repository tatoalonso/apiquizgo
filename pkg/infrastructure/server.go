package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tatoalonso/apiquizgo/pkg/application/usescases"
	quiz "github.com/tatoalonso/apiquizgo/pkg/domain"
	inmemoryrepository "github.com/tatoalonso/apiquizgo/pkg/infrastructure/repository"
)

//Server includes a router
type Server struct {
	Router *mux.Router
}

//NewServer creates a Server
func NewServer() Server {

	r := mux.NewRouter()

	r.HandleFunc("/quizes", fetchQuizes).Methods("GET")
	r.HandleFunc("/quiz/{id:[0-9]+}", fetchQuiz).Methods("GET")
	r.HandleFunc("/newquiz", newQuiz).Methods("POST")

	s := Server{
		Router: r,
	}

	return s

}

func fetchQuizes(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("quizes!\n"))
}

func fetchQuiz(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	repository := inmemoryrepository.NewInMemoryRepository()
	useCase := usescases.NewDefaultUseCase(repository)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	returnedQuiz, err := useCase.GetQuizByID(id)
	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	handleResponse(w, returnedQuiz, nil)
}

func newQuiz(w http.ResponseWriter, r *http.Request) {

	var quiz quiz.Quiz
	err := json.NewDecoder(r.Body).Decode(&quiz)

	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	repository := inmemoryrepository.NewInMemoryRepository()
	useCase := usescases.NewDefaultUseCase(repository)

	createdQuiz, err := useCase.CreateNewQuiz(quiz)

	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	handleResponse(w, createdQuiz, nil)

}

func handleErrorResponse(writer http.ResponseWriter, err error) {
	fmt.Println(err)
	writer.WriteHeader(http.StatusBadRequest)
	_, err = writer.Write([]byte(err.Error()))
	if err != nil {
		//TODO log framework
		log.Println(err)
	}
}

func handleResponse(writer http.ResponseWriter, result interface{}, err error) {

	responseBytes, err := json.Marshal(result)
	if err != nil {
		handleErrorResponse(writer, err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(responseBytes)
	if err != nil {
		log.Println(err)
	}
}
