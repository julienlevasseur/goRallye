package waypoint

import (
	"testing"

	"github.com/julienlevasseur/haversine"
	"github.com/stretchr/testify/assert"
)

func TestParseWaypoints(t *testing.T) {
	waypoints := []Waypoint{
		Waypoint{
			ID: 0,
			Coordinates: haversine.Coord{
				Lat: 47.123456,
				Lon: 73.123456,
			},
			Validated: false,
		},
		Waypoint{
			ID: 1,
			Coordinates: haversine.Coord{
				Lat: 78.901234,
				Lon: 56.789012,
			},
			Validated: true,
		},
	}

	wps, err := ParseWaypoints("test.json")
	assert.NoError(t, err)
	assert.Equal(t, len(waypoints), len(wps))
	for id, wp := range wps {
		assert.Equal(t, waypoints[id].ID, wp.ID)
		assert.Equal(t, waypoints[id].Coordinates.Lat, wp.Coordinates.Lat)
		assert.Equal(t, waypoints[id].Coordinates.Lon, wp.Coordinates.Lon)
		assert.Equal(t, waypoints[id].Validated, wp.Validated)
	}
}

func TestValidate(t *testing.T) {
	waypoint := Waypoint{
		ID: 0,
		Coordinates: haversine.Coord{
			Lat: 01.123456,
			Lon: 78.901234,
		},
		Validated: false,
	}

	wp := waypoint.Validate(
		haversine.Coord{
			Lat: 1.123,
			Lon: 78.901,
		},
	)

	assert.Equal(t, true, wp.Validated)
}
