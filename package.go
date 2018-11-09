package proxy_bonanza

type Package struct {
	Name        string  `json:"name"`
	Bandwith    int64   `json:"bandwidth"`
	Price       float64 `json:"price"`
	IPCount     int     `json:"howmany_ips"`
	PricePerGB  float64 `json:"price_per_gig"`
	PackageType string  `json:"package_type"`
}
