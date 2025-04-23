package survey

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
