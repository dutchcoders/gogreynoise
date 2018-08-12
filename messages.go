package greynoise

// ContextResult is the result of the /context api
type ContextResult struct {
	Actor     string `json:"actor"`
	FirstSeen string `json:"first_seen"`
	IP        string `json:"ip"`
	LastSeen  string `json:"last_seen"`
	Metadata  struct {
		ASN          string `json:"asn"`
		Category     string `json:"category"`
		City         string `json:"city"`
		Country      string `json:"country"`
		CountryCode  string `json:"country_code"`
		Organization string `json:"organization"`
		Os           string `json:"os"`
		Rdns         string `json:"rdns"`
		Tor          bool   `json:"tor"`
	} `json:"metadata"`
	RawData struct {
		Scan []struct {
			Port     int64  `json:"port"`
			Protocol string `json:"protocol"`
		} `json:"scan"`
		Web struct {
		} `json:"web"`
	} `json:"raw_data"`
	Seen bool     `json:"seen"`
	Tags []string `json:"tags"`
}

// ErrorResult is the result when an error occurs
type ErrorResult struct {
	Error string `json:"error"`
}
