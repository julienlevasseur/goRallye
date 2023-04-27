package waypoint

import (
	"encoding/json"
	"io/ioutil"

	"github.com/julienlevasseur/gps/coordinates"
	"github.com/julienlevasseur/haversine"
)

type Waypoint struct {
	ID               int             `json:"id"`
	VisibleID        bool            `json:"visible_id"`
	Coordinates      haversine.Coord `json:"coordinates"`
	VisibleCoords    bool            `json:"visible_coords"`
	VisibilityRadius int             `json:"visibility_radius"` // in meters
	Validated        bool            `json:"validated"`
	ValidationRadius int             `json:"validation_radius"` // in meters
}

//	WPV (Waypoint visible):
//
// A waypoint whose coordinates are communicated in the road book.
// Moving to a visible waypoint, all information is displayed on the GPS screen.
// To validate a WPV, a competitor must pass within 200 meters of it.
func NewVisibleWaypoint(id int, coordinates haversine.Coord) Waypoint {
	return Waypoint{
		ID:               id,
		VisibleID:        true,
		Coordinates:      coordinates,
		VisibleCoords:    true,
		VisibilityRadius: 0,
		Validated:        false,
		ValidationRadius: 200,
	}
}

// WPE (Waypoint Eclipse):
// A waypoint that becomes completely visible on the GPS once the preceding
// waypoint has been validated or within a radius of 1000 meters if the previous
// waypoint was missed. To validate a WPE, a competitor must pass within 90 meters of it.

func NewEclipseWaypoint(id int, coordinates haversine.Coord) Waypoint {
	return Waypoint{
		ID:               id,
		VisibleID:        false,
		Coordinates:      coordinates,
		VisibleCoords:    false,
		VisibilityRadius: 1000,
		Validated:        false,
		ValidationRadius: 90,
	}
}

// WPM (Masked Waypoint):
// A waypoint whose coordinates are not revealed to competitors. The GPS directs
// the competitor to this point only once within 800 meters of the latter.
// To validate a WPM, a competitor must pass within 90 meters of it.

func NewMaskedWaypoint(id int, coordinates haversine.Coord) Waypoint {
	return Waypoint{
		ID:               id,
		VisibleID:        false,
		Coordinates:      coordinates,
		VisibleCoords:    false,
		VisibilityRadius: 800,
		Validated:        false,
		ValidationRadius: 90,
	}
}

// WPS (Waypoint Security):
// A waypoint used to guarantee the safety of competitors, indicated in the
// road book and whose coordinates are not revealed to the competitors.
// The GPS only directs the competitor towards this point once he has arrived
// within a radius of 1000 meters of the latter. To validate a WPS,
// the competitor must pass within 30 meters of it.

func NewSecurityWaypoint(id int, coordinates haversine.Coord) Waypoint {
	return Waypoint{
		ID:               id,
		VisibleID:        false,
		Coordinates:      coordinates,
		VisibleCoords:    false,
		VisibilityRadius: 1000,
		Validated:        false,
		ValidationRadius: 30,
	}
}

// WPC (Control Waypoint):
// A Control waypoint is a waypoint which allows verifying the respect of the
// Road-Book, without any information of navigation being provided by the GPS
// other than the order of passage compared to the other waypoints or boxes of
// the Road Book and its name. A WPC should never be placed off track.
// In addition, the organizer will use as many WPCs as needed to avoid any
// possibility of shortcuts. To validate a WPC, the competitor must pass within
// 300 meters of it.

func NewControlWaypoint(id int, coordinates haversine.Coord) Waypoint {
	return Waypoint{
		ID:               id,
		VisibleID:        true,
		Coordinates:      coordinates,
		VisibleCoords:    false,
		VisibilityRadius: 0,
		Validated:        false,
		ValidationRadius: 300,
	}
}

// WPP (Precise Waypoint):
// A WPP is a waypoint that allows to check precisely the respect of the
// Roadbook follow-up on the tracks, without navigation information provided by
// the NAV-GPS. Its number and its order of passage in relation to other
// waypoints are only shown in the waypoint list of the road book.

func NewPreciseWaypoint(id int, coordinates haversine.Coord) Waypoint {
	return Waypoint{
		ID:               id,
		VisibleID:        true,
		Coordinates:      coordinates,
		VisibleCoords:    false,
		VisibilityRadius: 0,
		Validated:        false,
		ValidationRadius: 300,
	}
}

// WPN (Navigation Waypoint):
// The argument for this validation radius of 200 m is to allow the competitors
// more freedom to validate a WPN especially in off piste or dunes.
// The organizer will define the exact position during his reconnaissance and
// doing so he will consider the ground (gravel, sand, etc) for location of the waypoint.
// Even when doing so, in the dunes, with many vehicles passing the waypoint the
// situation may change (e.g. vehicles get stuck, the dune may change, etc.)
// during the rally. The organizer may use this waypoint to prevent competitors
// from avoiding challenging routes (e.g. dunes) or navigation difficulties.
// The GPS directs the competitors to this point once they have entered a radius
// of 800 meters from it.

func NewNavigationWaypoint(id int, coordinates haversine.Coord) Waypoint {
	return Waypoint{
		ID:               id,
		VisibleID:        false,
		Coordinates:      coordinates,
		VisibleCoords:    false,
		VisibilityRadius: 800,
		Validated:        false,
		ValidationRadius: 200,
	}
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

func HeadingToWaypoint(bearing, heading float64) float64 {
	return bearing - heading
}

func DistanceToWaypoint(position, waypoint haversine.Coord) (mi, km float64) {
	return haversine.Distance(position, waypoint)
}
