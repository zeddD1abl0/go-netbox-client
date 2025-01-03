package client

import "encoding/json"

func convertMapToStruct(m map[string]any, s any) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
