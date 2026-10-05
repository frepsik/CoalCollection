package storage

type minerTypeJSON struct {
	TypeName   string `json:"typeName"`
	Name       string `json:"name"`
	Cost       int    `json:"cost"`
	Energy     int    `json:"energy"`
	Extraction int    `json:"extraction"`
	Interval   string `json:"interval"`
	Growth     int    `json:"growth"`
}
