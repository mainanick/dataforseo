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

type KeywordIntent struct {
	Label       string  `json:"label"`
	Probability float64 `json:"probability"`
}

type ItemMetric struct {
	Pos_1                         int64   `json:"pos_1"`
	Pos_2_3                       int64   `json:"pos_2_3"`
	Pos_4_10                      int64   `json:"pos_4_10"`
	Pos_11_20                     int64   `json:"pos_11_20"`
	Pos_21_30                     int64   `json:"pos_21_30"`
	Pos_31_40                     int64   `json:"pos_31_40"`
	Pos_41_50                     int64   `json:"pos_41_50"`
	Pos_51_60                     int64   `json:"pos_51_60"`
	Pos_61_70                     int64   `json:"pos_61_70"`
	Pos_71_80                     int64   `json:"pos_71_80"`
	Pos_81_90                     int64   `json:"pos_81_90"`
	Pos_91_100                    int64   `json:"pos_91_100"`
	ETV                           float64 `json:"etv"`
	ImpressionsETV                float64 `json:"impressions_etv"`
	Count                         int64   `json:"count"`
	EstimatedPaidTrafficCost      float64 `json:"estimated_paid_traffic_cost"`
	IsNew                         int64   `json:"is_new"`
	IsUp                          int64   `json:"is_up"`
	IsDown                        int64   `json:"is_down"`
	IsLost                        int64   `json:"is_lost"`
	ClickstreamETV                int64   `json:"clickstream_etv"`
	ClickstreamGenderDistribution struct {
		Female int64 `json:"female"`
		Male   int64 `json:"male"`
	} `json:"clickstream_gender_distribution"`
	ClickstreamAgeDistribution struct {
		Age_18_24 int64 `json:"18-24"`
		Age_25_34 int64 `json:"25-34"`
		Age_35_44 int64 `json:"35-44"`
		Age_45_54 int64 `json:"45-54"`
		Age_55_64 int64 `json:"55-64"`
	} `json:"clickstream_age_distribution"`
}
