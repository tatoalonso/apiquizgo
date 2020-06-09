package usescases

import (
	quiz "github.com/tatoalonso/apiquizgo/pkg/domain"
)

//Usecase is an interface which CRU methods
type Usecase interface {
	CreateNewQuiz(quiz quiz.Quiz) (*quiz.Quiz, error)
	GetQuizByID(ID int) (*quiz.Quiz, error)
	GetQuizLists(ID string) (*quiz.Catalog, error)
}

//DefaultUseCase is a usecase which creates a new Quiz
type DefaultUseCase struct {
	repository quiz.Repository
}

//NewDefaultUseCase creates a new usecase
func NewDefaultUseCase(repository quiz.Repository) DefaultUseCase {
	return DefaultUseCase{repository: repository}
}

//CreateNewQuiz creates a new Quiz
func (newQuizUseCase DefaultUseCase) CreateNewQuiz(quiz quiz.Quiz) (*quiz.Quiz, error) {

	createdQuiz, err := newQuizUseCase.repository.CreateQuiz(quiz)

	if err != nil {
		return nil, err
	}
	return createdQuiz, nil

}

//GetQuizByID gets a quiz by id from the storage
func (newQuizUseCase DefaultUseCase) GetQuizByID(ID int) (*quiz.Quiz, error) {

	returnedQuiz, err := newQuizUseCase.repository.GetQuiz(ID)

	if err != nil {
		return nil, err
	}
	return returnedQuiz, nil

}

//GetQuizLists gets a list of quizes from the storage
func (newQuizUseCase DefaultUseCase) GetQuizLists(ID string) (*quiz.Catalog, error) {

	returnedQuizuesList, err := newQuizUseCase.repository.GetCatalog()

	if err != nil {
		return nil, err
	}
	return returnedQuizuesList, nil

}
