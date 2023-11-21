package response

// JsonResponseData 表示统一响应的JSON格式
type JsonResponseData struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}
