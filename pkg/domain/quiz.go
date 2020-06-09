package quiz

//Quiz is a simple example in Golang
type Quiz struct {
	ID             int    `json:"id"`
	TitleESP       string `json:"titleSpanish"`
	TitleENG       string `json:"titleEnglish"`
	URL            string `json:"url"`
	Code           string `json:"code"`
	ExplanationESP string `json:"explanationSpanish"`
	ExplanationENG string `json:"explanationEnglish"`
	Tags           string `json:"tags"`
}

//NewQuiz returns a new quiz
func NewQuiz(id int, titleSpanish, titleEnglish, url, code, explanationSpanish, explanationEnglish, tags string) (q Quiz) {
	q = Quiz{
		ID:             id,
		TitleESP:       titleSpanish,
		TitleENG:       titleEnglish,
		URL:            url,
		Code:           code,
		ExplanationESP: explanationSpanish,
		ExplanationENG: explanationEnglish,
		Tags:           tags,
	}
	return
}

//Catalog is a list of quizes
type Catalog struct {
	Quizes []Quiz `json:"quizes"`
}

//Repository is a interface which defines methods to create and list quizes
type Repository interface {
	CreateQuiz(quiz Quiz) (*Quiz, error)
	GetCatalog() (*Catalog, error)
	GetQuiz(QuizID int) (*Quiz, error)
}
