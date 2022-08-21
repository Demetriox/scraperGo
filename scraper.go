package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgURL string `json:"imageURL"`
}

func main() {
	c := colly.NewCollector()

	var items []item

	c.OnHTML("[data-hook=product-list-grid-item]", func(h *colly.HTMLElement) {

		item := item{
			Name:   h.ChildText("[data-hook=product-item-name]"),
			Price:  h.ChildText("[data-hook=product-item-price-to-pay]"),
			ImgURL: h.ChildAttr("[data-hook=wix-media-image]", "src"),
		}
		items = append(items, item)
	})
	c.Visit("https://www.ermana.mx/shop")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("products.json", content, 0644)
}
