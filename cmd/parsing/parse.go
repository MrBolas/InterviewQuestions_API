package parsing

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/MrBolas/InterviewQuestions_API/cmd/models"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func readQuestionsFromFilePath(path string) []models.Question {

	var questions []models.Question

	log.Println("Reading file:", path)
	questionsArr := readCsvFile(path)

	// Parse into []models.Question
	for _, v := range questionsArr {

		questions = append(questions, models.Question{
			Question: v[0],
			Answer:   v[1],
			Category: v[2],
			Level:    v[3],
			Source:   v[4],
		})
	}

	return questions
}

func ReadQuestionsFromFilePaths(paths []string) []models.Question {

	var questions []models.Question

	for _, path := range paths {
		questions = append(questions, readQuestionsFromFilePath(path)...)
	}

	return questions
}
