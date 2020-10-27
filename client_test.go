package stravaclient

import (
	"reflect"
	"testing"
)

func TestShowAthlete(t *testing.T) {
	t.Run("Gets athlete profile correctly", func(t *testing.T) {
		api := mockApi{}
		client := Client{"TOKEN", &api}
		got, _ := client.ShowAthlete()
		want := Athlete{
			11011101,
			"mrossi",
			2,
			"Mario",
			"Rossi",
			"Berlin",
			"Germany ",
			"M",
			true,
			true,
			"2018-09-11T21:23:01Z",
			"2020-10-27T14:38:25Z",
			1,
			"https://dgalywyr863hv.cloudfront.net/pictures/athletes/11011101/100100/1/medium.jpg",
			"https://dgalywyr863hv.cloudfront.net/pictures/athletes/11011101/100100/1/large.jpg",
		}
		assertAthlete(t, got, &want)
	})
}

func assertAthlete(t *testing.T, got *Athlete, want *Athlete) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
