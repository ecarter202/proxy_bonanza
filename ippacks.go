package proxy_bonanza

type IPPack struct {
	IP          string       `json:"ip"`
	PortHTTP    int          `json:"port_http"`
	PortSocks   int          `json:"port_socks"`
	ProxyServer *ProxyServer `json:"proxyserver"`
}
