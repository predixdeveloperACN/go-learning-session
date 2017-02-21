package vcap

type Application struct {
	ID         string   `json:"application_id"`
	Name       string   `json:"application_name"`
	URIs       []string `json:"application_uris"`
	Version    string   `json:"application_version"`
	InstanceID string   `json:"instance_id"`
	Limits     struct {
			   Mem  int `json:"mem"`
			   Disk int `json:"disk"`
			   FDs  int `json:"fds"`
		   } `json:"limits"`
	SpaceID            string `json:"space_id"`
	Start              string `json:"start"`
	StartedAtTimestamp int64  `json:"started_at_timestamp"`
}
