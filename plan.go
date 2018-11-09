package proxy_bonanza

type Plan struct {
	ID           int              `json:"id"`
	Login        string           `json:"login"`
	Password     string           `json:"password"`
	Expires      ProxyBonanzaTime `json:"expires"`
	Bandwith     int64            `json:"bandwidth"`
	LastIPChange ProxyBonanzaTime `json:"last_ip_change"`
	AuthIPS      []*AuthIP        `json:"authips"`
	Package      *Package         `json:"package"`
	IPPacks      []*IPPack        `json:"ippacks"`
}
