package dataforseo

import "context"

type SiteKeywordRequest struct {
	Target               string `json:"target"`
	TargetType           string `json:"target_type"`
	LocationName         string `json:"location_name"`
	LocationCode         int64  `json:"location_code"`
	LocationCoordinate   string `json:"location_coordinate"`
	LanguageName         string `json:"language_name"`
	LanguageCode         string `json:"language_code"`
	SearchPartners       bool   `json:"search_partners"`
	DateFrom             string `json:"date_from"`
	DateTo               string `json:"date_to"`
	IncludeAdultKeywords bool   `json:"include_adult_keywords"`
	SortBy               string `json:"sort_by"`
	Tag                  string `json:"tag"`
}

type KeywordResult struct {
	Keyword          string  `json:"keyword"`
	LocationCode     int64   `json:"location_code"`
	LanguageCode     string  `json:"language_code"`
	SearchPartners   bool    `json:"search_partners"`
	Competition      string  `json:"competition"`
	CompetitionIndex string  `json:"competition_index"`
	SearchVolume     int64   `json:"search_volume"`
	LowTopOfPageBid  float64 `json:"low_top_of_page_bid"`
	HighTopOfPageBid float64 `json:"high_top_of_page_bid"`
	Cpc              float64 `json:"cpc"`
	MonthlySearches  []struct {
		Year         int   `json:"year"`
		Month        int   `json:"month"`
		SearchVolume int64 `json:"search_volume"`
	} `json:"monthly_searches"`
	// TODO: Refactor keyword_annotations type:object desc:the annotations for the keyword
	KeywordAnnotations string `json:"keyword_annotations"`
	Concepts           []struct {
		Name         string `json:"name"`
		ConceptGroup []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"concepts"`
	}
}
type SiteKeywordResponseTasks struct {
	*BaseReponseTaskList
	Result []KeywordResult
}
type SiteKeywordResponse struct {
	*BaseResponse
	Tasks []SiteKeywordResponseTasks
}

func GoogleAdsSearchVolume(ctx context.Context, data SiteKeywordRequest) {
}
