package proxy_bonanza

import (
	"encoding/json"
	"strings"
	"time"
)

type ProxyBonanzaTime time.Time

func (pbt *ProxyBonanzaTime) UnmarshalJSON(b []byte) error {
	s := strings.Replace(string(b), "\"", "", -1)
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	*pbt = ProxyBonanzaTime(t)

	return nil
}

func (pbt ProxyBonanzaTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(pbt)
}
