package stravaclient

import (
	"io"
	"strings"
)

type mockApi struct {
}

func (a *mockApi) GetAthlete(accessToken string) (io.Reader, error) {
	r := strings.NewReader("{\"id\":11011101,\"username\":\"mrossi\",\"resource_state\":2,\"firstname\":\"Mario\",\"lastname\":\"Rossi\",\"city\":\"Berlin\",\"state\":\"Germany \",\"country\":null,\"sex\":\"M\",\"premium\":true,\"summit\":true,\"created_at\":\"2018-09-11T21:23:01Z\",\"updated_at\":\"2020-10-27T14:38:25Z\",\"badge_type_id\":1,\"profile_medium\":\"https://dgalywyr863hv.cloudfront.net/pictures/athletes/11011101/100100/1/medium.jpg\",\"profile\":\"https://dgalywyr863hv.cloudfront.net/pictures/athletes/11011101/100100/1/large.jpg\",\"friend\":null,\"follower\":null}")
	return r, nil
}
