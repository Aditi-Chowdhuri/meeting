package zoomAPI

import (
	"fmt"
	"net/http"
)

func (client Client) DeleteMeeting(meetingId int) (err error) {

	endpoint := fmt.Sprintf("/meetings/%d", meetingId)
	httpMethod := http.MethodDelete

	_, err = client.executeRequest(endpoint, httpMethod)
	if err != nil {
		return
	}

	return
}

