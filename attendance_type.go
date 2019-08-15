package dingtalk

type AttendanceListRecordsParam struct {
	UserIds       []string `json:"userIds,omitempty"` // 企业内的员工id列表，最多不能超过50个
	CheckDateFrom string   `json:"checkDateFrom"`     // 查询考勤打卡记录的起始工作日。格式为“yyyy-MM-dd hh:mm:ss”。
	CheckDateTo   string   `json:"checkDateTo"`       // 查询考勤打卡记录的结束工作日。格式为“yyyy-MM-dd hh:mm:ss”。注意，起始与结束工作日最多相隔7天
	IsI18n        string   `json:"isI18n"`            // 取值为true和false，表示是否为海外企业使用，默认为false。其中，true：海外平台使用，false：国内平台使用
}

type AttendanceListRecordsRsp struct {
	Response
	RecordResult []*AttendanceRecord `json:"recordresult"`
}

type AttendanceRecord struct {
	Id             int64   `json:"id"`             // 唯一标识ID
	IsLegal        string  `json:"isLegal"`        // 是否合法，当timeResult和locationResult都为Normal时，该值为Y；否则为N
	BaseCheckTime  int64   `json:"baseCheckTime"`  // 计算迟到和早退，基准时间；也可作为排班打卡时间
	UserId         string  `json:"userId"`         // 用户ID
	UserAddress    string  `json:"userAddress"`    // 用户打卡地址
	CheckType      string  `json:"checkType"`      // 考勤类型， OnDuty：上班, OffDuty：下班
	TimeResult     string  `json:"timeResult"`     // 时间结果
	DeviceId       string  `json:"deviceId"`       // 设备id
	CorpId         string  `json:"corpId"`         // 企业ID
	SourceType     string  `json:"sourceType"`     // 数据来源
	WorkDate       int64   `json:"workDate"`       // 工作日
	PlanCheckTime  int64   `json:"planCheckTime"`  // 排班打卡时间
	LocationMethod string  `json:"locationMethod"` // 定位方法
	LocationResult string  `json:"locationResult"` // 位置结果
	UserLongitude  float64 `json:"userLongitude"`  // 用户打卡经度
	UserLatitude   float64 `json:"userLatitude"`   // 用户打卡纬度
	PlanId         int64   `json:"planId"`         // 排班ID
	GroupId        int64   `json:"groupId"`        // 考勤组ID
	UserAccuracy   int     `json:"userAccuracy"`   // 用户打卡定位精度
	UserCheckTime  int64   `json:"userCheckTime"`  // 实际打卡时间
	ProcInstId     string  `json:"procInstId"`     // 关联的审批实例id，当该字段非空时，表示打卡记录与请假、加班等审批有关。可以与获取单个审批数据配合使用
}
