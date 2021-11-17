package crawler

import (
	"github.com/novel_crawler/internal/biz/collector"
	"github.com/novel_crawler/internal/data"
	"github.com/novel_crawler/pkg/conf"
	"github.com/spf13/cobra"
)

var CrawlerCmd = &cobra.Command{
	Use: "crawler",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		conf.Init(cmd.Flag("config").Value.String(), "yaml")
		data.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		collector.NewNovelClueCollector().Visit()

		//c := colly.NewCollector()
		//
		//c.OnHTML("div.category-wrap_iQLoo", func(e *colly.HTMLElement) {
		//	var (
		//		title  = e.ChildText("a[href]>div.c-single-text-ellipsis")
		//		link   = e.ChildAttr("a[href].title_dIF3B", "href")
		//		author = e.ChildTexts("div.intro_1l0wp")
		//		intro  = e.ChildText("div.c-single-text-ellipsis.desc_3CTjT")
		//		score  = e.ChildText("div.hot-index_1Bl1a")
		//	)
		//
		//	fmt.Printf("%s, %s, %s, %s, %s \n\n\n", title, score, link, author, intro)
		//
		//	var iScore, _ = strconv.Atoi(score)
		//	data.DBClient.NovelClue.Create().
		//		SetIntro(intro).
		//		SetLink(link).
		//		SetScore(iScore).
		//		SetTitle(title).
		//		SetDate(time.Now()).
		//		SetAuthor(strings.TrimPrefix(author[0], "作者：")).
		//		SetCategoryTitle(strings.TrimPrefix(author[1], "类型：")).
		//		SaveX(context.Background())
		//
		//	//fmt.Printf("title: %s\n price: %s \n\n\n",
		//	//	e.ChildText("div.p-name-type-2 > a[target]"),
		//	//	e.ChildText("div.p-price > strong > i[data-price]"),
		//	//	//e.ChildText("strong[data-presale] > i[data-price]"),
		//	//)
		//	//link := e.Attr("href")
		//	// Print link
		//	//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		//	// Visit link found on page
		//	// Only those links are visited which are in AllowedDomains
		//	//c.Visit(e.Request.AbsoluteURL(link))
		//})
		//
		//// Before making a request print "Visiting ..."
		//c.OnRequest(func(r *colly.Request) {
		//	//r.Headers.Set("authority", "search.jd.com")
		//	//r.Headers.Set("method", "GET")
		//	//r.Headers.Set("path", "/Search?keyword=27C9C")
		//	//r.Headers.Set("scheme", "https")
		//	//r.Headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		//	//r.Headers.Set("accept-encoding", "gzip, deflate, br")
		//	//r.Headers.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
		//	//r.Headers.Set("cache-control", "max-age=0")
		//	//r.Headers.Set("cookie", "shshshfpa=84f5d758-9d7f-a7bf-7a50-e1a81f667dc9-1623894974; shshshfpb=i6zx2NctUok0GrnZDDQ0gzQ%3D%3D; __jdu=1623894973002324499844; pinId=f5BinUYwQdE; qrsc=3; pin=zkerika; unick=%E5%95%A6%E5%95%A6%E5%95%A6%E5%95%A6cc; _tp=HhwaNEc314Ny0xhuXBd9XA%3D%3D; _pst=zkerika; unpl=V2_ZzNtbRIEFhx0XENdeEtVAmIAFFRLUUdFdQ8WBisYDlJuBRIOclRCFnUURlVnGVoUZwYZXkpcQxJFCENkexhdBWMGEV5EVnMlMEsWBi8FXABlARNbRV5LFXUARVJ6EVwGZh8RXUNfQxB2CEdcexxsBmczE21CUEAXfABDXH0bXQRkARFYR19LHHANdmR7EVg1ZwITXkdRSxR1C0BUcxFs08mXx9X9jvOwo6LogfapiYzM16zbcldKFXcAR1V7Hl81ZjMTbQM5QxR3CEJSexkRBWAAEFRKUksTdwlHV3kaWQBvCxtYR2dCJXY%3d; __jdv=122270672|www.google.com.hk|-|referral|-|1635771565734; areaId=2; ipLoc-djd=2-2830-51803-0; mt_xid=V2_52007VwMVUV9RWl8XTxtdBGQBEVdYUFpTGkwpCwJgChFaXV9OCRYaHkAAYQBFTlVcUlkDTxwPVTAEG1VZWFQKL0oYXwV7AxJOXFpDWh5CHV0OZgAiUG1YYl8aThFVB24EE1JtWFReGw%3D%3D; TrackID=1Ase8TlfB-eKIbCOVnM0GNtIAyXVv6d7ZfdafAUKEITIF3ZmN0gkB2Z0P8bsN1q1eG_Qf1gloqFYKdMYLV6A1tgUb8y36GelrOVQSz4TxIxc; __jda=122270672.1623894973002324499844.1623894973.1636080870.1636534264.34; __jdc=122270672; shshshfp=bb3fad024795db62db3dc0b203f84d84; rkv=1.0; __jdb=122270672.3.1623894973002324499844|34.1636534264; shshshsID=4293c072c831a7607806efc059d4a281_3_1636534282012; 3AB9D23F7A4B3C9B=27VQG5R57OEMJZTSWVVILUOARWULHLV6KE5RYDZE4OQCVFVPRVNQVJNGXHURE6RQT2LBLNRMAZVEF63QJBSRN22EMM; RT=\"z=1&dm=jd.com&si=kiktjo7v5u8&ss=kvta94r1&sl=2&tt=0&obo=2&nu=2bbab2c065d092c20f52cc0a81e0fa25&cl=iyz&ld=31yw&r=51fcfe23bd6604b618d60b86af173b9e&ul=31yy\"")
		//	//r.Headers.Set("sec-ch-ua", "\"Chromium\";v=\"94\", \"Google Chrome\";v=\"94\", \";Not A Brand\";v=\"99\"")
		//	//r.Headers.Set("sec-ch-ua-mobile", "?0")
		//	//r.Headers.Set("sec-ch-ua-platform", "\"Windows\"")
		//	//r.Headers.Set("sec-fetch-dest", "document")
		//	//r.Headers.Set("sec-fetch-mode", "navigate")
		//	//r.Headers.Set("sec-fetch-site", "none")
		//	//r.Headers.Set("sec-fetch-user", "?1")
		//	//r.Headers.Set("upgrade-insecure-requests", "1")
		//	//r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")
		//
		//	fmt.Println("Visiting", r.URL.String())
		//})
		//
		//c.OnResponse(func(r *colly.Response) {
		//	//fmt.Println("Visit Response", r.Request.URL.String(), string(r.Body))
		//})
		//
		//err := c.Visit("https://top.baidu.com/board?tab=novel")
		//if err != nil {
		//	fmt.Println(err, "===============cv err")
		//}
	},
}

func init() {
}
