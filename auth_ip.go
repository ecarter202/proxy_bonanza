package proxy_bonanza

type AuthIP struct {
	ID            int    `json:"id"`
	IP            string `json:"ip"`
	UserPackageID int    `json:"userpackage_id"`
}
