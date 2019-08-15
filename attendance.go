package dingtalk

import "net/http"

func (this *Client) AttendanceLists(param AttendanceListRecordsParam) (result *AttendanceListRecordsRsp, err error) {
	err = this.doRequestWithAccessToken(http.MethodPost, "attendance/listRecord", nil, param, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
