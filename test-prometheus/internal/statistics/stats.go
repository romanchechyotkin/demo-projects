package statistics

type Key struct {
	Timestamp  int64
	Country    string
	Os         string
	Browser    string
	CampaignId uint32
}

type Value struct {
	Requests    int64
	Impressions int64
}

type Rows map[Key]Value
