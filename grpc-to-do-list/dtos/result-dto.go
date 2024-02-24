package dtos

type ResultDto struct {
	Error  []ErrorDto  `json:"error"`  // 錯誤物件陣列
	State  int         `json:"state"`  // 狀態碼
	Result interface{} `json:"result"` // 請求成功時回傳的物件
}

type ErrorDto struct {
	Message string `json:"message"` // 錯誤訊息
	Field   string `json:"field"`
}

type StatusDto struct {
	State   int    `json:"state"`
	Message string `json:"message"`
}

type StatusWtishResultDto struct {
	State   int        `json:"state"`
	Message string     `json:"message"` // 錯誤訊息
	Result  []StatusID `json:"result"`
}

type StatusID struct {
	ID int64 `json:"id"`
}

type PanicDto struct {
	HTTPCode int
	State    int
	Result   interface{}
	Message  interface{}
}
