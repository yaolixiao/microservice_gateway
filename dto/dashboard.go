package dto

type PanelGroupDataOutput struct {
	ServiceNum 		int64 `json:"service_num"`
	AppNum 			int64 `json:"app_num"`
	CurrentQps 		int64 `json:"current_qps"`
	TodayRequestNum int64 `json:"today_request_num"`
}

type DashServiceStatOutput struct {
	Legend 		[]string `json:"legend"` //表示有几种分类
	Data 		[]DashServiceStatItemOutput `json:"data"` //详细数据
}

type DashServiceStatItemOutput struct {
	Name string `json:"name"`
	LoadType int `json:"load_type"`
	Value int64 `json:"value"`
}