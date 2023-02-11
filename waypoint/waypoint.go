package waypoint

import (
	"encoding/json"
	"io/ioutil"

	"github.com/julienlevasseur/gps/coordinates"
	"github.com/julienlevasseur/haversine"
)

type Waypoint struct {
	ID          int             `json:"id"`
	Coordinates haversine.Coord `json:"coordinates"`
	Validated   bool            `json:"validated"`
}

// Validate takes reference coordinates, if they match with the waypoint's
// coordinates, this function returns a waypoint with Validated: true
func (w *Waypoint) Validate(coords haversine.Coord) *Waypoint {
	roundedRefCoords := coordinates.RoundTo3Decimals(w.Coordinates)
	roundedCoords := coordinates.RoundTo3Decimals(coords)
	if roundedCoords.Lat == roundedRefCoords.Lat && roundedCoords.Lon == roundedRefCoords.Lon {
		return &Waypoint{
			ID:          w.ID,
			Coordinates: w.Coordinates,
			Validated:   true,
		}
	}
	return &Waypoint{
		ID:          w.ID,
		Coordinates: w.Coordinates,
		Validated:   false,
	}
}

// ParseWaypoints read a json file from the given path and return a slice of
// Waypoints.
func ParseWaypoints(path string) ([]Waypoint, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []Waypoint{}, err
	}

	var waypoints []Waypoint
	err = json.Unmarshal(content, &waypoints)
	if err != nil {
		return []Waypoint{}, err
	}

	return waypoints, nil
}
