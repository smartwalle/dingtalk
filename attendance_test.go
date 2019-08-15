package dingtalk_test

import (
	"fmt"
	"github.com/smartwalle/dingtalk"
	"testing"
	"time"
)

func TestClient_AttendanceLists(t *testing.T) {
	var p = dingtalk.AttendanceListRecordsParam{}
	p.IsI18n = "false"
	p.UserIds = []string{"28582062691668864798"}
	p.CheckDateFrom = "2019-08-01 00:00:00"
	p.CheckDateTo = "2019-08-07 00:00:00"
	rsp, err := client.AttendanceLists(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp.ErrCode, rsp.ErrMsg)

	for _, record := range rsp.RecordResult {
		fmt.Println(record.UserId, record.UserCheckTime, time.Now().Unix())
	}
}
