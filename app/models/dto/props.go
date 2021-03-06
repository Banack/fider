package dto

import (
	"database/sql/driver"
	"encoding/json"
)

// Props is a map of key:value
type Props map[string]interface{}

// Value converts props into a database value
func (p Props) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

// Merge current props with given props
func (p Props) Merge(props Props) Props {
	new := Props{}
	if p != nil {
		for k, v := range p {
			new[k] = v
		}
	}
	if props != nil {
		for k, v := range props {
			new[k] = v
		}
	}
	return new
}
