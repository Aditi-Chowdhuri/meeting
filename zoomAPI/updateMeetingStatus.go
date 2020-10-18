package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client Client) UpdateMeetingStatus(meetingId int, status string) (err error) {
	
	updateMeetingStatusRequest := UpdateMeetingStatusRequest{
		Action: status,
	}

	var reqBody []byte
	reqBody, err = json.Marshal(updateMeetingStatusRequest)
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("/meetings/%d/status", meetingId)
	httpMethod := http.MethodPut

	_, err = client.executeRequestWithBody(endpoint, httpMethod, reqBody)
	if err != nil {
		return
	}

	return

}
