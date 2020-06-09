package inmemoryrepository

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	quiz "github.com/tatoalonso/apiquizgo/pkg/domain"
)

const (
	pathFile   = "../../data/"
	nameOfFile = "quizes"
	fileType   = ".csv"
)

//InMemoryRepository is a repository in memory
type InMemoryRepository struct {
	path string
}

//NewInMemoryRepository creates a new NewInMemoryRepository
func NewInMemoryRepository() quiz.Repository {
	return &InMemoryRepository{path: pathFile}
}

//CreateQuiz inserts a new Quiz in the repository
func (storage InMemoryRepository) CreateQuiz(q quiz.Quiz) (*quiz.Quiz, error) {

	f, err := os.OpenFile(fmt.Sprintf("%v%v%v", storage.path, nameOfFile, fileType), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return nil, errors.New(fmt.Sprintln(err))
	}

	writer := csv.NewWriter(f)

	id := strconv.Itoa(q.ID)

	record := []string{id, q.TitleESP, q.TitleENG, q.URL, q.Code, q.ExplanationESP, q.ExplanationENG, q.Tags}

	err = writer.Write(record)

	writer.Flush()
	defer f.Close()

	if err != nil {
		return nil, errors.New("Problem writting quiz")
	}

	return &q, nil

}

//GetCatalog returns a list of quizes
func (storage InMemoryRepository) GetCatalog() (*quiz.Catalog, error) {

	file, _ := os.Open(fmt.Sprintf("%v%v%v", storage.path, nameOfFile, fileType))
	reader := bufio.NewReader(file)

	var c quiz.Catalog

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		id, _ := strconv.Atoi(values[0])

		quiz := quiz.NewQuiz(
			id,
			values[1],
			values[2],
			values[3],
			values[4],
			values[5],
			values[6],
			values[7],
		)

		c.Quizes = append(c.Quizes, quiz)
	}

	return &c, nil
}

//GetQuiz returns the quiz asociated to ID
func (storage InMemoryRepository) GetQuiz(QuizID int) (*quiz.Quiz, error) {

	catalog, err := storage.GetCatalog()

	if err != nil {
		return nil, errors.New("CATALOG NOT FOUND")
	}

	var quizFound quiz.Quiz

	for _, quiz := range catalog.Quizes {

		if QuizID == quiz.ID {
			quizFound = quiz

		}

	}

	if (quiz.Quiz{}) == quizFound {

		return nil, errors.New("QUIZ NOT FOUND")

	}

	return &quizFound, nil

}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
