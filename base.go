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
