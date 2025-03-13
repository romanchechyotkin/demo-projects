package main

import (
	"fmt"

	extractor "github.com/huantt/plaintext-extractor"
)

func main() {
	a := "<html><body> <p>вот и html показываем</p> <p>Вам пришло письмо c внутреннего сервиса <a href=\"https://www.ozon.ru/product/vlazhnyy-korm-dlya-sobak-cesar-adult-s-govyadinoy-i-ovoshchami-kusochki-v-souse-28-h-85-g-181790517/?avtc=1&avte=1&avts=1675334014&sh=sSqPgMOn3Q&comm=41qfLA0GQzIckIlnlY9002WSgM40000004000000000bT1OrShRoTgLtCNxuCxKunAJqSZOriRAr7BxbndLoC5HbmdBsS5Obm5AtmNQbncJpSZSum5AqmVLuiRFbmZSrTdEoSxxrmAJqTlPrSdEqSAJtyRPrTlPpiQOe2REbjwRbmsJcjwNdPAMdj4TbPZxtDhzfj4ConpQpjQN9C5St7cZcjoTdjcPd30Nd2pPq3RPkT5gpQRfrzdh\">КАКАЯ ТО ССЫЛКА</a></p> <p>С уважением,<br> Команда COMMS.</p> <img src=\"https://api.ozon.ru/web-api.comms/mo/12tZfU0GQzIckIlnlY9002WSgM40000004000000000.gif\" width=\"1\" height=\"1\"></body></html>"

	extractor := extractor.NewHtmlExtractor()
	output, err := extractor.PlainText(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(*output)

	a = "text"
	output, err = extractor.PlainText(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(*output)

}
