# proxy_bonanza
Proxy Bonanza API Wrapper

```go
func main() {
	myApiKey := "zyxwvutsrqponmlkjihgfedcba"
  
	fmt.Println("Getting all of your proxy IPs...")
	pb := proxy_bonanza.New(myApiKey)
	plans, err := pb.GetPlans()
	if err != nil {
		panic(err)
	}
  
	for _, p := range plans.Plans {
		plan, err := pb.GetPlan(p.ID)
		if err != nil {
			panic(err)
		}
		ips := plan.Plan.IPPacks
		for _, ip := range ips {
			fmt.Printf("IP: %s  Http: %v  Socks: %v\n", ip.IP, ip.PortHTTP, ip.PortSocks)
		}
	}
}
```

```
// Output
IP: 123.12.123.99  Http: 54321  Socks: 12345
IP: 123.12.123.98  Http: 54321  Socks: 12345
IP: 123.12.123.97  Http: 54321  Socks: 12345
IP: 123.12.123.96  Http: 54321  Socks: 12345
IP: 123.12.123.95  Http: 54321  Socks: 12345
IP: 123.12.123.94  Http: 54321  Socks: 12345
IP: 123.12.123.93  Http: 54321  Socks: 12345
IP: 123.12.123.92  Http: 54321  Socks: 12345
IP: 123.12.123.91  Http: 54321  Socks: 12345
IP: 123.12.123.90  Http: 54321  Socks: 12345
