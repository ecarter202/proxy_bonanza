package proxy_bonanza

type StatDaily struct {
	UserPackageID int              `json:"userpackage_id"`
	Date          ProxyBonanzaDate `json:"date"`
	HttpBnd       int64            `json:"bnd_http"`
	HttpConn      int64            `json:"conn_http"`
	SocksBnd      int64            `json:"bnd_socks"`
	SocksConn     int64            `json:"conn_socks"`
	TotalBnd      int64            `json:"bnd_total"`
	TotalConn     int64            `json:"conn_total"`
}
