package serp

var (
	IDListURL = "https://api.dataforseo.com/v3/serp/id_list"
)

type ListIDRequest struct {
	DatetimeFrom    string `json:"datetime_from"`
	DatetimeTo      string `json:"datetime_to"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
	Sort            string `json:"sort"`
	IncludeMetadata bool   `json:"include_metadata"`
}

type ListIDTask struct {
	ID            string        `json:"id"`
	StatusCode    int           `json:"status_code"`
	StatusMessage string        `json:"status_message"`
	Time          string        `json:"time"`
	Cost          float64       `json:"cost"`
	ResultCount   int64         `json:"result_count"`
	Path          []string      `json:"path"`
	Data          ListIDRequest `json:"data"`
	Result        ListIDResult
}

type ListIDResult struct {
	ID             string        `json:"id"`
	URL            []string      `json:"url"`
	DatetimePosted string        `json:"datetime_posted"`
	DatetimeDone   string        `json:"datetime_done"`
	Status         string        `json:"status"`
	Cost           float64       `json:"cost"`
	MetaData       ListIDRequest `json:"metadata"`
}
type ListIDResponse struct {
	Version       string       `json:"version"`
	StatusCode    int          `json:"status_code"`
	StatusMessage string       `json:"status_message"`
	Time          string       `json:"time"`
	Cost          float64      `json:"cost"`
	TasksCount    bool         `json:"tasks_count"`
	TasksError    bool         `json:"tasks_error"`
	Tasks         []ListIDTask `json:"tasks"`
}
