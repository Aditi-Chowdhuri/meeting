package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client Client) GetMeetingInvitation(meetingId int) (getMeetingInvitationResponse GetMeetingInvitationResponse, err error) {

	endpoint := fmt.Sprintf("/meetings/%d/invitation", meetingId)
	httpMethod := http.MethodGet

	var b []byte
	b, err = client.executeRequest(endpoint, httpMethod)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &getMeetingInvitationResponse)
	if err != nil {
		return
	}

	return
}
