package database

import (
	"encoding/json"
	"reflect"
)

// DeviceDetails - hold the details about a vehicle/vessel event
// there's really only two feilds we can count on being a string everytime here, device and generated
type DeviceDetails struct {
	DeviceId  string      `json:"device"`
	Generated string      `json:"generated"`
	Speed     interface{} `json:"speed"`
	Heading   interface{} `json:"heading"`
	Position  struct {
		Latitude  interface{} `json:"lat,omitempty"`
		Longitude interface{} `json:"long,omitempty"`
	} `json:"position,omitempty"`
}

// MarshalJSON - custom marshaler to hide generated_at and position if null on the way out
func (d DeviceDetails) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"device": d.DeviceId,
	}

	if !isNil(d.Position.Latitude) {
		position := map[string]interface{}{
			"lat":  d.Position.Latitude,
			"long": d.Position.Longitude,
		}
		m["position"] = position
	}

	if !isNil(d.Heading) {
		m["heading"] = d.Heading
	}

	if !isNil(d.Heading) {
		m["speed"] = d.Speed
	}

	return json.Marshal(m)
}

// isNil - really make sure that interface is nil
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
