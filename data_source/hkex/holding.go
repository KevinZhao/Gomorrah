//https://sc.hkexnews.hk/TuniS/www.hkexnews.hk/sdw/search/mutualmarket_c.aspx?t=sh&t=sz
//https://sc.hkexnews.hk/TuniS/www.hkexnews.hk/sdw/search/mutualmarket_c.aspx?t=sh&t=sh

package hkex

import (
	"fmt"
	"Gomorrah/http_operation"
	"net/http"
    "github.com/PuerkitoBio/goquery"
)

func Tbd()  {
	//processing url
	fmt.Println("processing url")

	resp := http_operation.HttpRequest(
		"GET", 
		"https://sc.hkexnews.hk/TuniS/www.hkexnews.hk/sdw/search/mutualmarket_c.aspx?t=sh&t=sh",
		"")

	//holdingPaser(resp)

}

func holdingPaser(response *http.Response)  {

    var headings, row []string
    var rows [][]string

    doc, err := goquery.NewDocumentFromResponse(response)
    if err != nil {
        fmt.Printf("httpPaser wrong")
    }

    doc.Find("table").Each(func(index int, tablehtml *goquery.Selection){

        //Todo: should judge id="mutualmarket-result"
        //if tablehtml.HasClass()

        tablehtml.Find("thead").Each(func(indexthead int, headhtml *goquery.Selection){
            headhtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection){
                rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection){

                    headings = append(headings, tableheading.Text())
                })
            })
        })

        tablehtml.Find("tbody").Each(func(indextbody int, bodyhtml *goquery.Selection){
            bodyhtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection){
                rowhtml.Find("td").Each(func(indextd int, tablecell *goquery.Selection) {
                    tablecell.Find("div").Each(func (indexcell int, value *goquery.Selection) {

                        if value.HasClass("mobile-list-body"){
                            row = append(row, value.Text())
                        }

                        fmt.Println("row = ", len(row), row)
                    })
                })
                rows = append(rows, row)
                row = nil
            })
        })
    })

    fmt.Println("####### headings = ", len(headings), headings)
    fmt.Println("####### rows = ", len(rows), rows)

}
