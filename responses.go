package proxy_bonanza

type plansResponse struct {
	Success bool    `json:"success"`
	Plans   []*Plan `json:"data"`
}

type planResponse struct {
	Success bool  `json:"success"`
	Plan    *Plan `json:"data"`
}

type authIPSResponse struct {
	Success bool      `json:"success"`
	AuthIPS []*AuthIP `json:"data"`
}

type addAuthIPResponse struct {
	Success  bool    `json:"success"`
	AuthIPID *AuthIP `json:"data"`
}

type monthsStatsResponse struct {
	Success bool         `json:"success"`
	Stats   []*StatDaily `json:"data"`
}

type todaysStatsResponse struct {
	Success bool          `json:"success"`
	Stats   []*StatHourly `json:"data"`
}
