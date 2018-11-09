package proxy_bonanza

import (
	"encoding/json"
	"strings"
	"time"
)

type ProxyBonanzaDate time.Time

func (pbd *ProxyBonanzaDate) UnmarshalJSON(b []byte) error {
	s := strings.Replace(string(b), "\"", "", -1)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*pbd = ProxyBonanzaDate(t)

	return nil
}

func (pbd ProxyBonanzaDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(pbd)
}
