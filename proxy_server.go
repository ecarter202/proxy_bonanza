package proxy_bonanza

type ProxyServer struct {
	GeoRegionID int        `json:"georegion_id"`
	GeoRegion   *GeoRegion `json:"georegion"`
}
