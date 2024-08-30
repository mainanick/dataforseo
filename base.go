package dataforseo

type BaseResponseTaskList struct {
	ID            string   `json:"id"`
	StatusCode    int      `json:"status_code"`
	StatusMessage string   `json:"status_message"`
	Time          string   `json:"time"`
	Cost          float64  `json:"cost"`
	ResultCount   int64    `json:"result_count"`
	Path          []string `json:"path"`
}

type BaseResponse struct {
	Version       string                 `json:"version"`
	StatusCode    int                    `json:"status_code"`
	StatusMessage string                 `json:"status_message"`
	Time          string                 `json:"time"`
	Cost          float64                `json:"cost"`
	TasksCount    int64                  `json:"tasks_count"`
	TasksError    int64                  `json:"tasks_error"`
	Tasks         []BaseResponseTaskList `json:"tasks"`
}

type AvgBacklinkInfo struct {
	SEType               string  `json:"se_type"`
	Backlinks            float64 `json:"backlinks"`
	Dofollow             float64 `json:"dofollow"`
	ReferringPages       float64 `json:"referring_pages"`
	ReferringDomains     float64 `json:"referring_domains"`
	ReferringMainDomains float64 `json:"referring_main_domains"`
	Rank                 float64 `json:"rank"`
	MainDomainRank       float64 `json:"main_domain_rank"`
	LastUpdatedTime      string  `json:"last_updated_time"`
}

type SERPInfo struct {
	SEType              string   `json:"se_type"`
	CheckURL            string   `json:"check_url"`
	SERPItemTypes       []string `json:"serp_item_types"`
	SEResultsCount      int64    `json:"se_results_count"`
	LastUpdatedTime     string   `json:"last_updated_time"`
	PreviousUpdatedTime string   `json:"previous_updated_time"`
}

type KeywordProperties struct {
	SEType                     string `json:"se_type"`
	CoreKeyword                string `json:"core_keyword"`
	SynonymClusteringAlgorithm string `json:"synonym_clustering_algorithm"`
	KeywordDifficulty          int64  `json:"keyword_difficulty"`
	DetectedLanguage           string `json:"detected_language"`
	IsAnotherLanguage          bool   `json:"is_another_language"`
}

type KeywordInfo struct {
	SEType           string  `json:"se_type"`
	LastUpdatedTime  string  `json:"last_updated_time"`
	Competition      float64 `json:"competition"`
	CompetitionLevel string  `json:"competition_level"`
	CPC              float64 `json:"cpc"`
	SearchVolume     int64   `json:"search_volume"`
	LowTopOfPageBid  float64 `json:"low_top_of_page_bid"`
	HighTopOfPageBid float64 `json:"high_top_of_page_bid"`
	Categories       []int64 `json:"categories"`
	MonthlySearches  []struct {
		Year         int64 `json:"year"`
		Month        int64 `json:"month"`
		SearchVolume int64 `json:"search_volume"`
	} `json:"monthly_searches"`
}

type SearchIntentInfo struct {
	SEType          string   `json:"se_type"`
	MainIntent      string   `json:"main_intent"`
	ForeignIntent   []string `json:"foreign_intent"`
	LastUpdatedTime string   `json:"last_updated_time"`
}
