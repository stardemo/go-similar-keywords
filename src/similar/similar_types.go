package similar

type ReqSimilarKeywords struct {
	KeyWords string `json:"key_words" p:"key_words" v:"required#关键词keyword不可为空"`
	Num      uint   `json:"num" p:"num"`
}
