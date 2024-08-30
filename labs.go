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
	GoogleLabsBulkKeywordDifficultyPath         = "dataforseo_labs/google/bulk_keyword_difficulty/live"
	GoogleLabsSearchIntentPath                  = "dataforseo_labs/google/search_intent/live"
	GoogleLabsRankedKeywordsPath                = "dataforseo_labs/google/ranked_keywords/live"
	GoogleLabsSERPCompetitorsPath               = "dataforseo_labs/google/serp_competitors/live"
	GoogleLabsCompetitorsDomainPath             = "dataforseo_labs/google/competitors_domain/live"
	GoogleLabsDomainIntersectionPath            = "dataforseo_labs/google/domain_intersection/live"
	GoogleLabsSubdomainsPath                    = "dataforseo_labs/google/subdomains/live"
	GoogleLabsRelevantPagesPath                 = "dataforseo_labs/google/relevant_pages/live"
	GoogleLabsDomainRankPath                    = "dataforseo_labs/google/domain_rank_overview/live"
	GoogleLabsPageIntersectionPath              = "dataforseo_labs/google/page_intersection/live"
	GoogleLabsBulkTrafficEstimationPath         = "dataforseo_labs/google/bulk_traffic_estimation/live"
	GoogleLabsHistoricalBulkTrafficPath         = "dataforseo_labs/google/historical_bulk_traffic_estimation/live"
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
					Keyword           string            `json:"keyword"`
					LocationCode      int64             `json:"location_code"`
					LanguageCode      string            `json:"language_code"`
					KeywordInfo       KeywordInfo       `json:"keyword_info"`
					KeywordProperties KeywordProperties `json:"keyword_properties"`
					SERPInfo          SERPInfo          `json:"serp_info"`
					AvgBacklinkInfo   AvgBacklinkInfo   `json:"avg_backlinks_info"`
					SearchIntentInfo  SearchIntentInfo  `json:"search_intent_info"`
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

type GoogleLabsBulkKeywordDifficultyRequest struct {
	Keywords     []string `json:"keywords"`
	LocationName string   `json:"location_name,omitempty"`
	LocationCode int64    `json:"location_code,omitempty"`
	LanguageName string   `json:"language_name,omitempty"`
	LanguageCode string   `json:"language_code,omitempty"`
	Tag          string   `json:"tag,omitempty"`
}

type GoogleLabsBulkKeywordDifficultyResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType       string `json:"se_type"`
			LocationCode int64  `json:"location_code"`
			LanguageCode string `json:"language_code"`
			TotalCount   int64  `json:"total_count"`
			ItemsCount   int64  `json:"items_count"`
			Items        []struct {
				SEType            string `json:"se_type"`
				Keyword           string `json:"keyword"`
				KeywordDifficulty int64  `json:"keyword_difficulty"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) BulkKeywordDifficulty(ctx context.Context, data *GoogleLabsBulkKeywordDifficultyRequest) (*GoogleLabsBulkKeywordDifficultyResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsBulkKeywordDifficultyPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsBulkKeywordDifficultyResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsSearchIntentRequest struct {
	Keywords     []string `json:"keywords"`
	LanguageName string   `json:"language_name,omitempty"`
	LanguageCode string   `json:"language_code,omitempty"`
	Tag          string   `json:"tag,omitempty"`
}

type GoogleLabsSearchIntentResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			LanguageCode string `json:"language_code"`
			ItemsCount   int64  `json:"items_count"`
			Items        []struct {
				Keyword                string          `json:"keyword"`
				KeywordIntent          KeywordIntent   `json:"keyword_intent"`
				SecondaryKeywordIntent []KeywordIntent `json:"secondary_keyword_intents"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) SearchIntent(ctx context.Context, data *GoogleLabsSearchIntentRequest) (*GoogleLabsSearchIntentResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsSearchIntentPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsSearchIntentResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsRankedKeywordsRequest struct {
	Target                 string   `json:"target"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	IgnoreSynonyms         string   `json:"ignore_synonyms,omitempty"`
	ItemTypes              []string `json:"item_types,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	Limit                  int64    `json:"limit,omitempty"`
	Offset                 int64    `json:"offset,omitempty"`
	LoadRankAbsolute       bool     `json:"load_rank_absolute,omitempty"`
	HistoricalSERPMode     bool     `json:"historical_serp_mode,omitempty"`
	Filters                []string `json:"filters,omitempty"`
	OrderBy                []string `json:"order_by,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

// TODO: Incomplete Response Data
type GoogleLabsRankedKeywordsResponse struct {
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
			Metrics      struct {
				Organic         ItemMetric `json:"organic"`
				Paid            ItemMetric `json:"paid"`
				FeaturedSnippet ItemMetric `json:"featured_snippet"`
				LocalPack       ItemMetric `json:"local_pack"`
			} `json:"metrics"`
			MetricsAbsolute struct {
				Organic         ItemMetric `json:"organic"`
				Paid            ItemMetric `json:"paid"`
				FeaturedSnippet ItemMetric `json:"featured_snippet"`
				LocalPack       ItemMetric `json:"local_pack"`
			} `json:"metrics_absolute"`
			Items []struct {
				SEType      string `json:"se_type"`
				KeywordData struct {
					SEType            string            `json:"se_type"`
					Keyword           string            `json:"keyword"`
					LocationCode      int64             `json:"location_code"`
					LanguageCode      string            `json:"language_code"`
					KeywordInfo       KeywordInfo       `json:"keyword_info"`
					KeywordProperties KeywordProperties `json:"keyword_properties"`
					SERPInfo          SERPInfo          `json:"serp_info"`
					AvgBacklinkInfo   AvgBacklinkInfo   `json:"avg_backlinks_info"`
					SearchIntentInfo  SearchIntentInfo  `json:"search_intent_info"`
				} `json:"keyword_data"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

// RankeKeywords endpoint will provide you with the list of keywords that any domain or webpage is ranking for. You will also get SERP elements related to the keyword position, as well as impressions, monthly searches and other data relevant to the returned keywords.
func (s *GoogleLabsService) RankedKeywords(ctx context.Context, data *GoogleLabsRankedKeywordsRequest) (*GoogleLabsRankedKeywordsResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsRankedKeywordsPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsRankedKeywordsResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsSERPCompetitorsRequest struct {
	Keywords          []string `json:"keywords"`
	LocationName      string   `json:"location_name,omitempty"`
	LocationCode      int64    `json:"location_code,omitempty"`
	LanguageName      string   `json:"language_name,omitempty"`
	LanguageCode      string   `json:"language_code,omitempty"`
	IncludeSubdomains bool     `json:"include_subdomains,omitempty"`
	ItemTypes         []string `json:"item_types,omitempty"`
	Limit             int64    `json:"limit,omitempty"`
	Offset            int64    `json:"offset,omitempty"`
	Filters           []string `json:"filters,omitempty"`
	OrderBy           []string `json:"order_by,omitempty"`
	Tag               string   `json:"tag,omitempty"`
}

type GoogleLabsSERPCompetitorsResponse struct {
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
			Items        []struct {
				SEType            string                 `json:"se_type"`
				Domain            string                 `json:"domain"`
				AvgPosition       int64                  `json:"avg_position"`
				MedianPosition    int64                  `json:"median_position"`
				Rating            int64                  `json:"rating"`
				ETV               float64                `json:"etv"`
				KeywordsCount     int64                  `json:"keywords_count"`
				Visibility        float64                `json:"visibility"`
				RelevantSERPItems int64                  `json:"relevant_serp_items"`
				KeywordsPositions map[string]interface{} `json:"keywords_positions"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

// SERPCompetitors endpoint will provide you with a list of domains ranking for the keywords you specify. You will also get SERP rankings, rating, estimated traffic volume, and visibility values the provided domains gain from the specified keywords.
func (s *GoogleLabsService) SERPCompetitors(ctx context.Context, data *GoogleLabsSERPCompetitorsRequest) (*GoogleLabsSERPCompetitorsResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsSERPCompetitorsPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsSERPCompetitorsResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsCompetitorsDomainRequest struct {
	Target                 string   `json:"target"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	ItemTypes              []string `json:"item_types,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	Filters                []string `json:"filters,omitempty"`
	OrderBy                []string `json:"order_by,omitempty"`
	Limit                  int64    `json:"limit,omitempty"`
	Offset                 int64    `json:"offset,omitempty"`
	MaxRankGroup           int64    `json:"max_rank_group,omitempty"`
	ExcludeTopDomains      bool     `json:"exclude_top_domains,omitempty"`
	IntersectingDomain     []string `json:"intersecting_domains,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

type GoogleLabsCompetitorsDomainResponse struct {
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
			Items        []struct {
				SEType            string  `json:"se_type"`
				Domain            string  `json:"domain"`
				AvgPosition       float64 `json:"avg_position"`
				SumPosition       int64   `json:"sum_position"`
				Intersections     int64   `json:"intersections"`
				FullDomainMetrics struct {
					Organic         ItemMetric `json:"organic"`
					Paid            ItemMetric `json:"paid"`
					FeaturedSnippet ItemMetric `json:"featured_snippet"`
					LocalPack       ItemMetric `json:"local_pack"`
				} `json:"full_domain_metrics"`
				Metrics struct {
					Organic         ItemMetric `json:"organic"`
					Paid            ItemMetric `json:"paid"`
					FeaturedSnippet ItemMetric `json:"featured_snippet"`
					LocalPack       ItemMetric `json:"local_pack"`
				} `json:"metrics"`
				CompetitorMetrics struct {
					Organic         ItemMetric `json:"organic"`
					Paid            ItemMetric `json:"paid"`
					FeaturedSnippet ItemMetric `json:"featured_snippet"`
					LocalPack       ItemMetric `json:"local_pack"`
				} `json:"competitor_metrics"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

// CompetitorsDomain endpoint will provide you with a full overview of ranking and traffic data of the competitor domains from organic and paid search. In addition to that, you will get the metrics specific to the keywords both competitor domains and your domain rank for within the same SERP.
func (s *GoogleLabsService) CompetitorsDomain(ctx context.Context, data *GoogleLabsCompetitorsDomainRequest) (*GoogleLabsCompetitorsDomainResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsCompetitorsDomainPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsCompetitorsDomainResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsDomainIntersectionRequest struct {
	Target1                string   `json:"target1"`
	Target2                string   `json:"target2"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	Intersections          bool     `json:"intersections,omitempty"`
	ItemTypes              []string `json:"item_types,omitempty"`
	IncludeSERPInfo        bool     `json:"include_serp_info,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	Limit                  int64    `json:"limit,omitempty"`
	Offset                 int64    `json:"offset,omitempty"`
	Filters                []string `json:"filters,omitempty"`
	OrderBy                []string `json:"order_by,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

type GoogleLabsDomainIntersectionResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType       string `json:"se_type"`
			Target1      string `json:"target1"`
			Target2      string `json:"target2"`
			LocationCode int64  `json:"location_code"`
			LanguageCode string `json:"language_code"`
			TotalCount   int64  `json:"total_count"`
			ItemsCount   int64  `json:"items_count"`
			Items        []struct {
				SEType      string `json:"se_type"`
				KeywordData struct {
					SEType            string            `json:"se_type"`
					Keyword           string            `json:"keyword"`
					LocationCode      int64             `json:"location_code"`
					LanguageCode      string            `json:"language_code"`
					KeywordInfo       KeywordInfo       `json:"keyword_info"`
					KeywordProperties KeywordProperties `json:"keyword_properties"`
					SERPInfo          SERPInfo          `json:"serp_info"`
					AvgBacklinkInfo   AvgBacklinkInfo   `json:"avg_backlinks_info"`
					SearchIntentInfo  SearchIntentInfo  `json:"search_intent_info"`
				} `json:"keyword_data"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

// DomainIntersection endpoint will provide you with the keywords for which both specified domains rank within the same SERP. You will get search volume, competition, cost-per-click and impressions data on each intersecting keyword. Along with that, you will get data on the first and second domain’s SERP element discovered for this keyword, as well as the estimated traffic volume and cost of ad traffic. Domain Intersection endpoint supports organic, paid, local pack, and featured snippet results.
func (s *GoogleLabsService) DomainIntersection(ctx context.Context, data *GoogleLabsDomainIntersectionRequest) (*GoogleLabsDomainIntersectionResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsDomainIntersectionPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsDomainIntersectionResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsSubdomainsRequest struct {
	Target                 string   `json:"target"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	ItemTypes              []string `json:"item_types,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	HistoricalSERPMode     bool     `json:"historical_serp_mode,omitempty"`
	Filters                []string `json:"filters,omitempty"`
	OrderBy                []string `json:"order_by,omitempty"`
	Limit                  int64    `json:"limit,omitempty"`
	Offset                 int64    `json:"offset,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

type GoogleLabsSubdomainsResponse struct {
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
			Items        []struct {
				SEType    string `json:"se_type"`
				Subdomain string `json:"subdomain"`
				Metrics   struct {
					Organic         ItemMetric `json:"organic"`
					Paid            ItemMetric `json:"paid"`
					FeaturedSnippet ItemMetric `json:"featured_snippet"`
					LocalPack       ItemMetric `json:"local_pack"`
				} `json:"metrics"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

// Subdomains endpoint will provide you with a list of subdomains of the specified domain, along with the ranking distribution across organic and paid search. In addition to that, you will also get the estimated traffic volume of subdomains based on search volume and impressions.
func (s *GoogleLabsService) Subdomains(ctx context.Context, data *GoogleLabsSubdomainsRequest) (*GoogleLabsSubdomainsResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsSubdomainsPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsSubdomainsResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsRelevantPagesRequest struct {
	Target                 string   `json:"target"`
	LocationName           string   `json:"location_name,omitempty"`
	LocationCode           int64    `json:"location_code,omitempty"`
	LanguageName           string   `json:"language_name,omitempty"`
	LanguageCode           string   `json:"language_code,omitempty"`
	ItemTypes              []string `json:"item_types,omitempty"`
	Limit                  int64    `json:"limit,omitempty"`
	IncludeClickStreamData bool     `json:"include_clickstream_data,omitempty"`
	Offset                 int64    `json:"offset,omitempty"`
	HistoricalSERPMode     bool     `json:"historical_serp_mode,omitempty"`
	Filters                []string `json:"filters,omitempty"`
	OrderBy                []string `json:"order_by,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
}

type GoogleLabsRelevantPagesResponse struct {
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
			Items        []struct {
				SEType      string `json:"se_type"`
				PageAddress string `json:"page_address"`
				Metrics     struct {
					Organic         ItemMetric `json:"organic"`
					Paid            ItemMetric `json:"paid"`
					FeaturedSnippet ItemMetric `json:"featured_snippet"`
					LocalPack       ItemMetric `json:"local_pack"`
				} `json:"metrics"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) RelevantPages(ctx context.Context, data *GoogleLabsRelevantPagesRequest) (*GoogleLabsRelevantPagesResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsRelevantPagesPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsRelevantPagesResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsDomainRankRequest struct {
	Target         string `json:"target"`
	LocationName   string `json:"location_name,omitempty"`
	LocationCode   int64  `json:"location_code,omitempty"`
	LanguageName   string `json:"language_name,omitempty"`
	LanguageCode   string `json:"language_code,omitempty"`
	IgnoreSynonyms string `json:"ignore_synonyms,omitempty"`
	Limit          int64  `json:"limit,omitempty"`
	Offset         int64  `json:"offset,omitempty"`
	Tag            string `json:"tag,omitempty"`
}

type GoogleLabsDomainRankResponse struct {
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
			Items        []struct {
				SEType       string `json:"se_type"`
				LocationCode int64  `json:"location_code"`
				LanguageCode string `json:"language_code"`
				Metrics      struct {
					Organic ItemMetric `json:"organic"`
					Paid    ItemMetric `json:"paid"`
				} `json:"metrics"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) DomainRank(ctx context.Context, data *GoogleLabsDomainRankRequest) (*GoogleLabsDomainRankResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsDomainRankPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsDomainRankResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsPageIntersectionRequest struct {
	Pages                  map[string]string `json:"pages"`
	ExcludePages           []string          `json:"exclude_pages"`
	LocationName           string            `json:"location_name,omitempty"`
	LocationCode           int64             `json:"location_code,omitempty"`
	LanguageName           string            `json:"language_name,omitempty"`
	LanguageCode           string            `json:"language_code,omitempty"`
	ItemTypes              []string          `json:"item_types,omitempty"`
	Limit                  int64             `json:"limit,omitempty"`
	Offset                 int64             `json:"offset,omitempty"`
	IncludeSubdomains      bool              `json:"include_subdomains,omitempty"`
	IntersectionMode       string            `json:"intersection_mode,omitempty"`
	IncludeSERPInfo        bool              `json:"include_serp_info,omitempty"`
	IncludeClickStreamData bool              `json:"include_clickstream_data,omitempty"`
	IgnoreSynonyms         string            `json:"ignore_synonyms,omitempty"`
	Filters                []string          `json:"filters,omitempty"`
	OrderBy                []string          `json:"order_by,omitempty"`
	Tag                    string            `json:"tag,omitempty"`
}

type GoogleLabsPageIntersectionResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType       string            `json:"se_type"`
			Pages        map[string]string `json:"pages"`
			ExcludePages []string          `json:"exclude_pages"`
			LocationCode int64             `json:"location_code"`
			LanguageCode string            `json:"language_code"`
			TotalCount   int64             `json:"total_count"`
			ItemsCount   int64             `json:"items_count"`
			Items        []struct {
				SEType       string `json:"se_type"`
				LocationCode int64  `json:"location_code"`
				LanguageCode string `json:"language_code"`
				KeywordData  struct {
					SEType            string            `json:"se_type"`
					Keyword           string            `json:"keyword"`
					LocationCode      int64             `json:"location_code"`
					LanguageCode      string            `json:"language_code"`
					KeywordInfo       KeywordInfo       `json:"keyword_info"`
					KeywordProperties KeywordProperties `json:"keyword_properties"`
					SERPInfo          SERPInfo          `json:"serp_info"`
					AvgBacklinkInfo   AvgBacklinkInfo   `json:"avg_backlinks_info"`
					SearchIntentInfo  SearchIntentInfo  `json:"search_intent_info"`
				} `json:"keyword_data"`
				IntersectionResult map[string]interface{} `json:"intersection_result"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) PageIntersection(ctx context.Context, data *GoogleLabsPageIntersectionRequest) (*GoogleLabsPageIntersectionResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsPageIntersectionPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsPageIntersectionResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsBulkTrafficEstimationRequest struct {
	Targets      []string `json:"targets"`
	LocationName string   `json:"location_name,omitempty"`
	LocationCode int64    `json:"location_code,omitempty"`
	LanguageName string   `json:"language_name,omitempty"`
	LanguageCode string   `json:"language_code,omitempty"`
	ItemTypes    []string `json:"item_types,omitempty"`
	Tag          string   `json:"tag,omitempty"`
}

type GoogleLabsBulkTrafficEstimationResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType       string `json:"se_type"`
			LocationCode int64  `json:"location_code"`
			LanguageCode string `json:"language_code"`
			TotalCount   int64  `json:"total_count"`
			ItemsCount   int64  `json:"items_count"`
			Items        []struct {
				SEType  string `json:"se_type"`
				Target  string `json:"target"`
				Metrics struct {
					Organic struct {
						ETV   float64 `json:"etv"`
						Count int64   `json:"count"`
					} `json:"organic"`
					Paid struct {
						ETV   float64 `json:"etv"`
						Count int64   `json:"count"`
					} `json:"paid"`
					FeaturedSnippet struct {
						ETV   float64 `json:"etv"`
						Count int64   `json:"count"`
					} `json:"featured_snippet"`
					LocalPack struct {
						ETV   float64 `json:"etv"`
						Count int64   `json:"count"`
					} `json:"local_pack"`
				} `json:"metrics"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

func (s *GoogleLabsService) BulkTrafficEstimation(ctx context.Context, data *GoogleLabsBulkTrafficEstimationRequest) (*GoogleLabsBulkTrafficEstimationResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsBulkTrafficEstimationPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsBulkTrafficEstimationResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GoogleLabsHistoricalBulkTrafficRequest struct {
	Targets      []string `json:"targets"`
	LocationName string   `json:"location_name,omitempty"`
	LocationCode int64    `json:"location_code,omitempty"`
	LanguageName string   `json:"language_name,omitempty"`
	LanguageCode string   `json:"language_code,omitempty"`
	ItemTypes    []string `json:"item_types,omitempty"`
	Tag          string   `json:"tag,omitempty"`
}

type GoogleLabsHistoricalBulkTrafficResponse struct {
	*BaseResponse
	Tasks []struct {
		*BaseResponseTaskList
		Result []struct {
			SEType       string `json:"se_type"`
			LocationCode int64  `json:"location_code"`
			LanguageCode string `json:"language_code"`
			TotalCount   int64  `json:"total_count"`
			ItemsCount   int64  `json:"items_count"`
			Items        []struct {
				SEType  string `json:"se_type"`
				Target  string `json:"target"`
				Metrics struct {
					Organic []struct {
						Year  int     `json:"year"`
						Month int     `json:"month"`
						ETV   float64 `json:"etv"`
						Count int64   `json:"count"`
					} `json:"organic"`
					Paid []struct {
						Year  int     `json:"year"`
						Month int     `json:"month"`
						ETV   float64 `json:"etv"`
						Count int64   `json:"count"`
					} `json:"paid"`
					FeaturedSnippet []struct {
						Year  int     `json:"year"`
						Month int     `json:"month"`
						ETV   float64 `json:"etv"`
						Count int64   `json:"count"`
					} `json:"featured_snippet"`
					LocalPack []struct {
						Year  int     `json:"year"`
						Month int     `json:"month"`
						ETV   float64 `json:"etv"`
						Count int64   `json:"count"`
					} `json:"local_pack"`
				} `json:"metrics"`
			} `json:"items"`
		} `json:"result"`
	} `json:"tasks"`
}

// HistoricalBulkTrafficEstimation endpoint will provide you with historical monthly traffic volumes for up to 1,000 domains collected within the specified time range through October 2020. If you do not specify the range, data will be returned for the previous 12 months. Along with organic search traffic estimations, you will also get separate values for paid search, featured snippet, and local pack results.
func (s *GoogleLabsService) HistoricalBulkTrafficEstimation(ctx context.Context, data *GoogleLabsHistoricalBulkTrafficRequest) (*GoogleLabsHistoricalBulkTrafficResponse, error) {
	req, err := s.client.NewRequest("POST", GoogleLabsHistoricalBulkTrafficPath, []interface{}{data})
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	res := &GoogleLabsHistoricalBulkTrafficResponse{}
	_, err = s.client.Do(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
