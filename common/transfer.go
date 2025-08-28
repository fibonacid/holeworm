package common

type Transfer struct {
	SessionID    string `json:"session_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	WormholeCode string `json:"wormhole_code,omitempty"`
}
