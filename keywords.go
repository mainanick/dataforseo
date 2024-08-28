package dataforseo

import "context"

type SiteKeywordRequest struct {
	Target               string `json:"target"`
	TargetType           string `json:"target_type,omitempty"`
	LocationName         string `json:"location_name,omitempty"`
	LocationCode         int64  `json:"location_code,omitempty"`
	LocationCoordinate   string `json:"location_coordinate,omitempty"`
	LanguageName         string `json:"language_name,omitempty"`
	LanguageCode         string `json:"language_code,omitempty"`
	SearchPartners       bool   `json:"search_partners,omitempty"`
	DateFrom             string `json:"date_from,omitempty"`
	DateTo               string `json:"date_to,omitempty"`
	IncludeAdultKeywords bool   `json:"include_adult_keywords,omitempty"`
	SortBy               string `json:"sort_by,omitempty"`
	Tag                  string `json:"tag,omitempty"`
}

type KeywordForKeywordRequest struct {
	Keywords             []string `json:"keywords"`
	LocationName         string   `json:"location_name,omitempty"`
	LocationCode         int64    `json:"location_code,omitempty"`
	LocationCoordinate   string   `json:"location_coordinate,omitempty"`
	LanguageName         string   `json:"language_name,omitempty"`
	LanguageCode         string   `json:"language_code,omitempty"`
	SearchPartners       bool     `json:"search_partners,omitempty"`
	DateFrom             string   `json:"date_from,omitempty"`
	DateTo               string   `json:"date_to,omitempty"`
	IncludeAdultKeywords bool     `json:"include_adult_keywords,omitempty"`
	SortBy               string   `json:"sort_by,omitempty"`
	Tag                  string   `json:"tag,omitempty"`
}

type SiteKeywordResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			Keyword          string  `json:"keyword"`
			LocationCode     int64   `json:"location_code"`
			LanguageCode     string  `json:"language_code"`
			SearchPartners   bool    `json:"search_partners"`
			Competition      string  `json:"competition"`
			CompetitionIndex int64   `json:"competition_index"`
			SearchVolume     int64   `json:"search_volume"`
			LowTopOfPageBid  float64 `json:"low_top_of_page_bid"`
			HighTopOfPageBid float64 `json:"high_top_of_page_bid"`
			CPC              float64 `json:"cpc"`
			MonthlySearches  []struct {
				Year         int   `json:"year"`
				Month        int   `json:"month"`
				SearchVolume int64 `json:"search_volume"`
			} `json:"monthly_searches"`
			KeywordAnnotations struct {
				Concepts []struct {
					Name         string `json:"name"`
					ConceptGroup struct {
						Name string `json:"name"`
						Type string `json:"type"`
					} `json:"concept_group"`
				} `json:"concepts"`
			} `json:"keyword_annotations"`
		} `json:"result"`
	} `json:"tasks"`
}

type SiteKeywordService Service

func (s *SiteKeywordService) GoogleSiteKeywords(ctx context.Context, data SiteKeywordRequest) (*SiteKeywordResponse, error) {
	req, err := s.client.NewRequest("POST", "keywords_data/google_ads/keywords_for_site/live", []interface{}{data})
	if err != nil {
		return nil, err
	}
	keywordResponse := &SiteKeywordResponse{}
	_, err = s.client.Do(ctx, req, keywordResponse)
	if err != nil {
		return nil, err
	}
	return keywordResponse, nil
}

func (s *SiteKeywordService) KeywordsForKeywords(ctx context.Context, data KeywordForKeywordRequest) (*SiteKeywordResponse, error) {
	req, err := s.client.NewRequest("POST", "keywords_data/google_ads/keywords_for_keywords/live", []interface{}{data})
	if err != nil {
		return nil, err
	}
	keywordResponse := &SiteKeywordResponse{}
	_, err = s.client.Do(ctx, req, keywordResponse)
	if err != nil {
		return nil, err
	}
	return keywordResponse, nil
}
