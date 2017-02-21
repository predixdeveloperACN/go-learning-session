package vcap

import (
	"os"
	"encoding/json"
)

type Services map[string][]Service

const ServiceKey = "VCAP_SERVICES"

type Service struct {
	Credentials map[string]interface{} `json:"credentials"`
	Label       string                 `json:"label"`
	Name        string                 `json:"name"`
	Plan        string                 `json:"plan"`
	Tags        []string               `json:"tags"`
}

func LoadServices() (Services, error) {

	var svc Services

	j := os.Getenv(ServiceKey)
	if err := json.Unmarshal([]byte(j), &svc); err != nil {
		return nil, err
	}

	return svc, nil
}