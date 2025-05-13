package survey

import (
	"encoding/json"
	"social_bot_backend/internal/user"
)

type SurveyService struct {
	SurveyRepository *SurveyRepository
}

func NewSurveyService(repository *SurveyRepository) *SurveyService {
	return &SurveyService{SurveyRepository: repository}
}

func (service *SurveyService) CreateSurvey(survey *SurveyCreateRequest) error {
	var questions []Question
	for _, v := range survey.Questions {
		questions = append(questions, Question{Question: v})
	}

	err := service.SurveyRepository.CreateSurvey(&Survey{
		Name:           survey.Name,
		Description:    survey.Description,
		CountQuestions: len(questions),
		IsPublic:       survey.IsPublic,
		Questions:      questions,
	})

	return err
}
func (service *SurveyService) GetAllSurvey() (*[]AllSurveyResponse, error) {
	survey := service.SurveyRepository.GetAllSurvey()
	var response []AllSurveyResponse

	for _, v := range *survey {
		response = append(response, AllSurveyResponse{
			Id:   v.ID,
			Name: v.Name,
		})
	}

	return &response, nil
}

func (service *SurveyService) GetQuestionWithAnswers(id uint) (*[]SurveyResponse, error) {
	survey, err := service.SurveyRepository.GetQuestionWithAnswers(id)
	if err != nil {
		return nil, err
	}
	var resp []SurveyResponse
	for _, question := range survey.Questions {
		response := SurveyResponse{}
		questionData := make(map[string]interface{})
		err := json.Unmarshal(question.Question, &questionData)
		if err != nil {
			return nil, err
		}

		response.Question = questionData["question"].(string)
		response.Type = questionData["type"].(string)

		answerCounts := make(map[string]int)
		countAnswers(answerCounts, question.UserAnswer)

		for answerText, count := range answerCounts {
			response.Answers = append(response.Answers, Answer{
				Answer: answerText,
				Count:  count,
			})
		}

		resp = append(resp, response)
	}
	return &resp, nil
}

func countAnswers(answerCounts map[string]int, answers []user.UserAnswer) {
	for _, userAnswer := range answers {
		for _, answer := range []string(userAnswer.Answer) {
			answerCounts[answer]++
		}
	}
}
