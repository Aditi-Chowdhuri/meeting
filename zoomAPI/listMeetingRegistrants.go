package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client Client) ListMeetingRegistrants(meetingId int) (apiResponse ListMeetingRegistrantsResponse, err error) {

	endpoint := fmt.Sprintf("/meetings/%d/registrants", meetingId)
	httpMethod := http.MethodGet

	var b []byte
	b, err = client.executeRequest(endpoint, httpMethod)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &apiResponse)
	if err != nil {
		return
	}

	return
}

