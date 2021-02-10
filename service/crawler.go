package service

import (
	"bytes"
	"context"
	"strings"

	"github.com/thanhhm/stw/model"
	"github.com/thanhhm/stw/util"
	"golang.org/x/net/html"
)

const (
	tvOverviewPageURL  = "https://vn.tradingview.com/markets/stocks-vietnam/"
	tvOverviewDivClass = "tv-feed__item"
	tvOverviewAClass   = "tv-widget-idea__title"
)

type Crawler struct {
	Caller model.Caller
}

func NewCrawler(hc *util.HTTPC) *Crawler {
	return &Crawler{
		Caller: model.NewCaller(hc),
	}
}

type TradingViewData struct {
	Title []string
}

func (cr *Crawler) GetTradingView(ctx context.Context) (*TradingViewData, error) {
	b, err := cr.Caller.Get(ctx, tvOverviewPageURL)
	if err != nil {
		return nil, err
	}

	r := bytes.NewBuffer(b)
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	// Travel document to get div tag
	var divTags []*html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" && isContainAttr(n.Attr, "class", tvOverviewDivClass) {
			divTags = append(divTags, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// Extract title
	var title []string
	for _, v := range divTags {
		title = append(title, getBlockTitle(v))
	}

	return nil, nil
}

func getBlockTitle(n *html.Node) string {
	var result string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" && isContainAttr(n.Attr, "class", tvOverviewAClass) {
			result = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return result
}

func isContainAttr(s []html.Attribute, key, val string) bool {
	for _, v := range s {
		if v.Key == key && strings.Contains(v.Val, val) {
			return true
		}
	}
	return false
}
