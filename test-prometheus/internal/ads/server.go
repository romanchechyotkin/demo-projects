package ads

import (
	realip "github.com/ferluci/fast-realip"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"github.com/valyala/fasthttp"
	"log"
	"net"
	"net/http"
)

type Server struct {
	geoip *geoip2.Reader
}

func NewServer(geoip *geoip2.Reader) *Server {
	return &Server{
		geoip: geoip,
	}
}

func (srv *Server) Listen() error {
	return fasthttp.ListenAndServe(":8080", srv.handler)
}

func (srv *Server) handler(ctx *fasthttp.RequestCtx) {
	ip := realip.FromRequest(ctx)
	country, err := srv.geoip.Country(net.ParseIP(ip))
	if err != nil {
		log.Println("cant parse country", err)
		return
	}

	userAgent := ctx.Request.Header.UserAgent()
	parsedUserAgent := user_agent.New(string(userAgent))
	browserName, _ := parsedUserAgent.Browser()

	u := &User{
		Country: country.Country.IsoCode,
		Browser: browserName,
	}
	campaigns := GetStaticCampaigns()
	winner := MakeAuction(campaigns, u)
	if winner == nil {
		ctx.Redirect("https://example.com", http.StatusSeeOther)
		return
	}

	ctx.Redirect(winner.ClickUrl, http.StatusSeeOther)
}
