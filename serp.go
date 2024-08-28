package dataforseo

import "context"

type SerpListIDRequest struct {
	DatetimeFrom    string `json:"datetime_from"`
	DatetimeTo      string `json:"datetime_to"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
	Sort            string `json:"sort"`
	IncludeMetadata bool   `json:"include_metadata"`
}

type SerpListIDResult struct {
	ID             string            `json:"id"`
	URL            []string          `json:"url"`
	DatetimePosted string            `json:"datetime_posted"`
	DatetimeDone   string            `json:"datetime_done"`
	Status         string            `json:"status"`
	Cost           float64           `json:"cost"`
	MetaData       SerpListIDRequest `json:"metadata"`
}
type SerpListIDResponse struct {
	Version       string  `json:"version"`
	StatusCode    int     `json:"status_code"`
	StatusMessage string  `json:"status_message"`
	Time          string  `json:"time"`
	Cost          float64 `json:"cost"`
	TasksCount    bool    `json:"tasks_count"`
	TasksError    bool    `json:"tasks_error"`
	Tasks         []struct {
		ID            string            `json:"id"`
		StatusCode    int               `json:"status_code"`
		StatusMessage string            `json:"status_message"`
		Time          string            `json:"time"`
		Cost          float64           `json:"cost"`
		ResultCount   int64             `json:"result_count"`
		Path          []string          `json:"path"`
		Data          SerpListIDRequest `json:"data"`
		Result        SerpListIDRequest
	} `json:"tasks"`
}

type GoogleOrganicRequest struct {
	Keyword             string `json:"keyword"`
	URL                 string `json:"url,omitempty"`
	LocationName        string `json:"location_name,omitempty"`
	LocationCode        int64  `json:"location_code,omitempty"`
	LocationCoordinate  string `json:"location_coordinate,omitempty"`
	LanguageName        string `json:"language_name,omitempty"`
	LanguageCode        string `json:"language_code,omitempty"`
	Device              string `json:"device,omitempty"`
	OS                  string `json:"os,omitempty"`
	SEDomain            string `json:"se_domain,omitempty"`
	Depth               int    `json:"depth,omitempty"`
	Target              string `json:"target,omitempty"`
	GroupOrganicResults bool   `json:"group_organic_results,omitempty"`
	MaxCrawlPages       int    `json:"max_crawl_pages,omitempty"`
	SearchParams        string `json:"search_param,omitempty"`
	Tag                 string `json:"tag,omitempty"`
}

type GoogleOrganicResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			Keyword      string `json:"keyword"`
			Type         string `json:"type"`
			SEDomain     string `json:"se_domain"`
			LocationCode int64  `json:"location_code"`
			LanguageCode string `json:"language_code"`
			CheckURL     string `json:"check_url"`
			Datetime     string `json:"datetime"`
			Spell        struct {
				Keyword string `json:"keyword"`
				Type    string `json:"type"`
			} `json:"spell"`
			ItemTypes      []string `json:"item_types"`
			SEResultsCount int64    `json:"se_results_count"`
			ItemsCount     int64    `json:"items_count"`
			Items          []struct {
				Type         string `json:"type"`
				RankGroup    int64  `json:"rank_group"`
				RankAbsolute int64  `json:"rank_absolute"`
				Domain       string `json:"domain"`
				Title        string `json:"title"`
				Description  string `json:"description"`
				URL          string `json:"url"`
				Breadcrumb   string `json:"breadcrumb"`
			}
		} `json:"result"`
	} `json:"tasks"`
}

type SerpService Service

func (s *SerpService) GoogleOrganicRegular(ctx context.Context, data GoogleOrganicRequest) (*GoogleOrganicResponse, error) {
	req, err := s.client.NewRequest("POST", "serp/google/organic/live/regular", []interface{}{data})
	if err != nil {
		return nil, err
	}
	response := &GoogleOrganicResponse{}
	_, err = s.client.Do(ctx, req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
