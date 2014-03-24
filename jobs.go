package genkins


type JobsService struct {
	client *Client
}

func (s JobsService) GetAll() (jobView *JobView, err error)  {

	req, err := s.client.NewRequest("GET", "/api/json?tree=jobs[name,url,color]", nil)

	if err != nil {
		return nil, err
	}

	var view JobView
	err = s.client.Do(req, &view)

	if err != nil {
		return nil, err
	}

	return &view, nil
}

type JobView struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Url string   `json:"url"`
}
