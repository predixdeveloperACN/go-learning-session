package vcap

import (
	"os"
	"encoding/json"
)

const AppKey = "VCAP_APPLICATION"

func LoadApplication() (Application, error) {

	var app Application

	j := os.Getenv(AppKey)
	if err := json.Unmarshal([]byte(j), &app); err != nil {
		return Application{}, err
	}

	return app, nil
}
