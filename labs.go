package dataforseo

import (
	"context"
	"errors"
	"os"
)

type GoogleLabsService Service

var (
	GoogleLabsKeywordForSitePath                = "dataforseo_labs/google/keywords_for_site/live"
	GoogleLabsRelatedKeywordPath                = "dataforseo_labs/google/related_keywords/live"
	GoogleLabsKeywordSuggestionsPath            = "dataforseo_labs/google/keyword_suggestions/live"
	GoogleLabsKeywordIdeasPath                  = "dataforseo_labs/google/keyword_ideas/live"
	GoogleLabsKeywordHistoricalSearchVolumePath = "dataforseo_labs/google/historical_search_volume/live"
)

type GoogleLabsKeywordForSiteRequest struct {
	Target                 string   `json:"target"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	IncludeSERPInfo        bool     `json:"include_serp_info,omitempty"`
	IncludeSubdomains      bool     `json:"include_subdomains,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	IgnoreSynonyms         string   `json:"ignore_synonyms,omitempty"`
	Limit                  int64    `json:"limit,omitempty"`
	Offset                 int64    `json:"offset,omitempty"`
	OffsetToken            string   `json:"offset_token,omitempty"`
	Filters                []string `json:"filters,omitempty"`
	OrderBy                []string `json:"order_by,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

type GoogleLabsKeywordForSiteResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType       string `json:"se_type"`
			Target       string `json:"target"`
			LocationCode int64  `json:"location_code"`
			LanguageCode string `json:"language_code"`
			TotalCount   int64  `json:"total_count"`
			ItemsCount   int64  `json:"items_count"`
			Offset       int64  `json:"offset"`
			OffsetToken  string `json:"offset_token"`
			Items        []struct {
				SEType       string `json:"se_type"`
				Keyword      string `json:"keyword"`
				LocationCode int64  `json:"location_code"`
				LanguageCode string `json:"language_code"`
				KeywordInfo  struct {
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
				} `json:"keyword_info"`
				KeywordProperties struct {
					SEType                     string `json:"se_type"`
					CoreKeyword                string `json:"core_keyword"`
					SynonymClusteringAlgorithm string `json:"synonym_clustering_algorithm"`
					KeywordDifficulty          int64  `json:"keyword_difficulty"`
					DetectedLanguage           string `json:"detected_language"`
					IsAnotherLanguage          bool   `json:"is_another_language"`
				} `json:"keyword_properties"`
				SERPInfo struct {
					SEType              string   `json:"se_type"`
					CheckURL            string   `json:"check_url"`
					SERPItemTypes       []string `json:"serp_item_types"`
					SEResultsCount      int64    `json:"se_results_count"`
					LastUpdatedTime     string   `json:"last_updated_time"`
					PreviousUpdatedTime string   `json:"previous_updated_time"`
				} `json:"serp_info"`
				AvgBacklinkInfo struct {
					SEType               string  `json:"se_type"`
					Backlinks            float64 `json:"backlinks"`
					Dofollow             float64 `json:"dofollow"`
					ReferringPages       float64 `json:"referring_pages"`
					ReferringDomains     float64 `json:"referring_domains"`
					ReferringMainDomains float64 `json:"referring_main_domains"`
					Rank                 float64 `json:"rank"`
					MainDomainRank       float64 `json:"main_domain_rank"`
					LastUpdatedTime      string  `json:"last_updated_time"`
				} `json:"avg_backlinks_info"`
				SearchIntentInfo struct {
					SEType          string   `json:"se_type"`
					MainIntent      string   `json:"main_intent"`
					ForeignIntent   []string `json:"foreign_intent"`
					LastUpdatedTime string   `json:"last_updated_time"`
				} `json:"search_intent_info"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) KeywordForSite(ctx context.Context, data *GoogleLabsKeywordForSiteRequest) (*GoogleLabsKeywordForSiteResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsKeywordForSitePath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsKeywordForSiteResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsRelatedKeywordsRequest struct {
	Keyword                string   `json:"keyword"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	Depth                  int64    `json:"depth,omitempty"`
	IncludeSeedKeywordInfo bool     `json:"include_seed_keyword,omitempty"`
	IncludeSERPInfo        bool     `json:"include_serp_info,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	IgnoreSynonyms         string   `json:"ignore_synonyms,omitempty"`
	ReplaceWithCoreKeyword bool     `json:"replace_with_core_keyword,omitempty"`
	Filters                []string `json:"filters,omitempty"`
	OrderBy                []string `json:"order_by,omitempty"`
	Limit                  int64    `json:"limit,omitempty"`
	Offset                 int64    `json:"offset,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

type GoogleLabsRelatedKeywordsResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType       string `json:"se_type"`
			SeedKeyword  string `json:"seed_keyword"`
			LocationCode int64  `json:"location_code"`
			LanguageCode string `json:"language_code"`
			TotalCount   int64  `json:"total_count"`
			ItemsCount   int64  `json:"items_count"`
			Items        []struct {
				SEType      string `json:"se_type"`
				KeywordData struct {
					Keyword      string `json:"keyword"`
					LocationCode int64  `json:"location_code"`
					LanguageCode string `json:"language_code"`
					KeywordInfo  struct {
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
					} `json:"keyword_info"`
					KeywordProperties struct {
						SEType                     string `json:"se_type"`
						CoreKeyword                string `json:"core_keyword"`
						SynonymClusteringAlgorithm string `json:"synonym_clustering_algorithm"`
						KeywordDifficulty          int64  `json:"keyword_difficulty"`
						DetectedLanguage           string `json:"detected_language"`
						IsAnotherLanguage          bool   `json:"is_another_language"`
					} `json:"keyword_properties"`
					SERPInfo struct {
						SEType              string   `json:"se_type"`
						CheckURL            string   `json:"check_url"`
						SERPItemTypes       []string `json:"serp_item_types"`
						SEResultsCount      int64    `json:"se_results_count"`
						LastUpdatedTime     string   `json:"last_updated_time"`
						PreviousUpdatedTime string   `json:"previous_updated_time"`
					} `json:"serp_info"`
					AvgBacklinkInfo struct {
						SEType               string  `json:"se_type"`
						Backlinks            float64 `json:"backlinks"`
						Dofollow             float64 `json:"dofollow"`
						ReferringPages       float64 `json:"referring_pages"`
						ReferringDomains     float64 `json:"referring_domains"`
						ReferringMainDomains float64 `json:"referring_main_domains"`
						Rank                 float64 `json:"rank"`
						MainDomainRank       float64 `json:"main_domain_rank"`
						LastUpdatedTime      string  `json:"last_updated_time"`
					} `json:"avg_backlinks_info"`
					SearchIntentInfo struct {
						SEType          string   `json:"se_type"`
						MainIntent      string   `json:"main_intent"`
						ForeignIntent   []string `json:"foreign_intent"`
						LastUpdatedTime string   `json:"last_updated_time"`
					} `json:"search_intent_info"`
				} `json:"keyword_data"`
				Depth           int64    `json:"depth"`
				RelatedKeywords []string `json:"related_keywords"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) RelatedKeywords(ctx context.Context, data *GoogleLabsRelatedKeywordsRequest) (*GoogleLabsRelatedKeywordsResponse, error) {
	// TODO: Set hard defaults for now
	data.IncludeSeedKeywordInfo = false
	data.IncludeSERPInfo = true

	req, err := s.client.NewRequest("POST", GoogleLabsRelatedKeywordPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsRelatedKeywordsResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsKeywordSuggestionsRequest struct {
	Keyword                string   `json:"keyword"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	Depth                  int64    `json:"depth,omitempty"`
	IncludeSeedKeywordInfo bool     `json:"include_seed_keyword,omitempty"`
	IncludeSERPInfo        bool     `json:"include_serp_info,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	ExactMatch             bool     `json:"exact_match,omitempty"`
	IgnoreSynonyms         string   `json:"ignore_synonyms,omitempty"`
	Filters                []string `json:"filters,omitempty"`
	OrderBy                []string `json:"order_by,omitempty"`
	Limit                  int64    `json:"limit,omitempty"`
	Offset                 int64    `json:"offset,omitempty"`
	OffsetToken            string   `json:"offset_token,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

type GoogleLabsKeywordSuggestionsResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType          string `json:"se_type"`
			SeedKeyword     string `json:"seed_keyword"`
			SeedKeywordData struct {
				SEType            string            `json:"se_type"`
				Keyword           string            `json:"keyword"`
				LocationCode      int64             `json:"location_code"`
				LanguageCode      string            `json:"language_code"`
				KeywordInfo       KeywordInfo       `json:"keyword_info"`
				SERPInfo          SERPInfo          `json:"serp_info"`
				KeywordProperties KeywordProperties `json:"keyword_properties"`
				SearchIntentInfo  SearchIntentInfo  `json:"search_intent_info"`
				AvgBacklinkInfo   AvgBacklinkInfo   `json:"avg_backlinks_info"`
			} `json:"seed_keyword_data"`
			LocationCode int64  `json:"location_code"`
			LanguageCode string `json:"language_code"`
			TotalCount   int64  `json:"total_count"`
			ItemsCount   int64  `json:"items_count"`
			Offset       int64  `json:"offset"`
			OffsetToken  string `json:"offset_token"`
			Items        []struct {
				SEType            string            `json:"se_type"`
				Keyword           string            `json:"keyword"`
				LocationCode      int64             `json:"location_code"`
				LanguageCode      string            `json:"language_code"`
				KeywordInfo       KeywordInfo       `json:"keyword_info"`
				KeywordProperties KeywordProperties `json:"keyword_properties"`
				SERPInfo          SERPInfo          `json:"serp_info"`
				AvgBacklinkInfo   AvgBacklinkInfo   `json:"avg_backlinks_info"`
				SearchIntentInfo  SearchIntentInfo  `json:"search_intent_info"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) KeywordSuggestions(ctx context.Context, data *GoogleLabsKeywordSuggestionsRequest) (*GoogleLabsKeywordSuggestionsResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsKeywordSuggestionsPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsKeywordSuggestionsResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsKeywordIdeasRequest struct {
	Keywords               []string `json:"keywords"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	CloselyVariants        bool     `json:"closely_variants,omitempty"`
	IgnoreSynonyms         string   `json:"ignore_synonyms,omitempty"`
	IncludeSERPInfo        bool     `json:"include_serp_info,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	Limit                  int64    `json:"limit,omitempty"`
	Offset                 int64    `json:"offset,omitempty"`
	OffsetToken            string   `json:"offset_token,omitempty"`
	Filters                []string `json:"filters,omitempty"`
	OrderBy                []string `json:"order_by,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

type GoogleLabsKeywordIdeasResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType       string   `json:"se_type"`
			SeedKeywords []string `json:"seed_keywords"`
			LocationCode int64    `json:"location_code"`
			LanguageCode string   `json:"language_code"`
			TotalCount   int64    `json:"total_count"`
			ItemsCount   int64    `json:"items_count"`
			Offset       int64    `json:"offset"`
			OffsetToken  string   `json:"offset_token"`
			Items        []struct {
				SEType            string            `json:"se_type"`
				Keyword           string            `json:"keyword"`
				LocationCode      int64             `json:"location_code"`
				LanguageCode      string            `json:"language_code"`
				KeywordInfo       KeywordInfo       `json:"keyword_info"`
				KeywordProperties KeywordProperties `json:"keyword_properties"`
				SERPInfo          SERPInfo          `json:"serp_info"`
				AvgBacklinkInfo   AvgBacklinkInfo   `json:"avg_backlinks_info"`
				SearchIntentInfo  SearchIntentInfo  `json:"search_intent_info"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) KeywordIdeas(ctx context.Context, data *GoogleLabsKeywordIdeasRequest) (*GoogleLabsKeywordIdeasResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsKeywordIdeasPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsKeywordIdeasResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsHistoricalSearchVolumeRequest struct {
	Keywords               []string `json:"keywords"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	IncludeSERPInfo        bool     `json:"include_serp_info,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

type GoogleLabsHistoricalSearchVolumeResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType       string `json:"se_type"`
			LocationCode int64  `json:"location_code"`
			LanguageCode string `json:"language_code"`
			ItemsCount   int64  `json:"items_count"`
			Items        []struct {
				SEType            string            `json:"se_type"`
				Keyword           string            `json:"keyword"`
				LocationCode      int64             `json:"location_code"`
				LanguageCode      string            `json:"language_code"`
				SearchPartners    bool              `json:"search_partners"`
				KeywordInfo       KeywordInfo       `json:"keyword_info"`
				KeywordProperties KeywordProperties `json:"keyword_properties"`
				SERPInfo          SERPInfo          `json:"serp_info"`
				AvgBacklinkInfo   AvgBacklinkInfo   `json:"avg_backlinks_info"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) HistoricalSearchVolume(ctx context.Context, data *GoogleLabsHistoricalSearchVolumeRequest) (*GoogleLabsHistoricalSearchVolumeResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsKeywordHistoricalSearchVolumePath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsHistoricalSearchVolumeResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
