package proxy_bonanza

type GeoRegion struct {
	Name   string `json:"name"`
	County *Country
}
