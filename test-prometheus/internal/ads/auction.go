package ads

import "sort"

type User struct {
	Country string
	Browser string
}

type Campaign struct {
	ClickUrl  string
	Price     float64
	Targeting Targeting
}

type Targeting struct {
	Country string
	Browser string
}

type filterFunc func(in []*Campaign, u *User) (out []*Campaign)

var filters = []filterFunc{
	FilterByBrowser,
	FilterByCountry,
}

func MakeAuction(in []*Campaign, u *User) (winner *Campaign) {
	campaigns := make([]*Campaign, len(in))
	copy(campaigns, in)

	for _, f := range filters {
		campaigns = f(campaigns, u)
	}

	if len(campaigns) == 0 {
		return nil
	}

	sort.Slice(campaigns, func(i, j int) bool {
		return campaigns[i].Price < campaigns[j].Price
	})

	return campaigns[0]
}
func FilterByBrowser(in []*Campaign, u *User) []*Campaign {
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Browser) == 0 {
			//пустой браузер - значит нет ограничений, не удаляем кампанию
			continue
		}

		if in[i].Targeting.Browser == u.Browser {
			//браузер совпадает - всё ок, кампания проходит
			continue
		}

		in[i] = in[0]
		in = in[1:]
	}
	return in
}

func FilterByCountry(in []*Campaign, u *User) []*Campaign {
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Country) == 0 {
			//пустой браузер - значит нет ограничений, не удаляем кампанию
			continue
		}

		if in[i].Targeting.Country == u.Country {
			//браузер совпадает - всё ок, кампания проходит
			continue
		}

		in[i] = in[0]
		in = in[1:]
	}
	return in
}

func GetStaticCampaigns() []*Campaign {
	return []*Campaign{
		{
			Price: 1,
			Targeting: Targeting{
				Country: "BLR",
				Browser: "Chrome",
			},
			ClickUrl: "https://yandex.ru",
		},
		{
			Price: 1,
			Targeting: Targeting{
				Country: "DE",
				Browser: "Chrome",
			},
			ClickUrl: "https://google.com",
		},
		{
			Price: 1,
			Targeting: Targeting{
				Browser: "Firefox",
			},
			ClickUrl: "https://duckduckgo.com",
		},
	}
}
