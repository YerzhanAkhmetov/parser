package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/xuri/excelize/v2"
)

type Product struct {
	Href string
	Name string
}

func GetProductsKeramic(products []*Product) []*Product {
	//products := []*Product{}
	for j := 1; j < 31; j++ {

		// Request the HTML page.
		res, err := http.Get("https://mirceramiki.kz/catalog/keramogranit?sort=stock_count&page=" + strconv.Itoa(j))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find(".wrap .content .card-inner .card-price .add-favorites a").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			//title := s.Find("a").Text()
			node := s.Get(0)
			href := ""
			name := ""
			for _, a := range node.Attr {
				if a.Key == "href" {
					href = a.Val
				}

				if a.Key == "data-product-name" {
					name = a.Val
				}
			}

			product := new(Product)
			product.Href = href
			product.Name = name
			products = append(products, product)
			//fmt.Printf("Review %d: href %s name %s\n", i, href, name)
		})

	}

	return products
}

func GetLaminam(products []*Product) []*Product {
	for j := 1; j < 6; j++ {

		// Request the HTML page.
		res, err := http.Get("https://mirceramiki.kz/catalog/keramicheskaya-plitka/napolnaya?sort=stock_count&page=" + strconv.Itoa(j))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find(".wrap .content .card-inner .card-price .add-favorites a").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			//title := s.Find("a").Text()
			node := s.Get(0)
			href := ""
			name := ""
			for _, a := range node.Attr {
				if a.Key == "href" {
					href = a.Val
				}

				if a.Key == "data-product-name" {
					name = a.Val
				}
			}

			product := new(Product)
			product.Href = href
			product.Name = name
			products = append(products, product)
			//fmt.Printf("Review %d: href %s name %s\n", i, href, name)
		})

	}
	return products
}

func GetGrigio(products []*Product) []*Product {

	// Request the HTML page.
	res, err := http.Get("https://mirceramiki.kz/collection/luxor")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".wrap .content-full .collection-list .card-inner .card-price .add-favorites a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		//title := s.Find("a").Text()
		node := s.Get(0)
		href := ""
		name := ""
		for _, a := range node.Attr {
			if a.Key == "href" {
				href = a.Val
			}

			if a.Key == "data-product-name" {
				name = a.Val
			}
		}

		product := new(Product)
		product.Href = href
		product.Name = name
		products = append(products, product)
		//fmt.Printf("Review %d: href %s name %s\n", i, href, name)
	})
	return products
}
func GetGrigio1(products []*Product) []*Product {

	// Request the HTML page.
	res, err := http.Get("https://mirceramiki.kz/collection/magic")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".wrap .content-full .collection-list .card-inner .card-price .add-favorites a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		//title := s.Find("a").Text()
		node := s.Get(0)
		href := ""
		name := ""
		for _, a := range node.Attr {
			if a.Key == "href" {
				href = a.Val
			}

			if a.Key == "data-product-name" {
				name = a.Val
			}
		}

		product := new(Product)
		product.Href = href
		product.Name = name
		products = append(products, product)
		//fmt.Printf("Review %d: href %s name %s\n", i, href, name)
	})
	return products
}
func GetGrigio2(products []*Product) []*Product {

	// Request the HTML page.
	res, err := http.Get("https://mirceramiki.kz/collection/umberto")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".wrap .content-full .collection-list .card-inner .card-price .add-favorites a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		//title := s.Find("a").Text()
		node := s.Get(0)
		href := ""
		name := ""
		for _, a := range node.Attr {
			if a.Key == "href" {
				href = a.Val
			}

			if a.Key == "data-product-name" {
				name = a.Val
			}
		}

		product := new(Product)
		product.Href = href
		product.Name = name
		products = append(products, product)
		//fmt.Printf("Review %d: href %s name %s\n", i, href, name)
	})
	return products
}
func GetGrigio3(products []*Product) []*Product {

	// Request the HTML page.
	res, err := http.Get("https://mirceramiki.kz/collection/wizard")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".wrap .content-full .collection-list .card-inner .card-price .add-favorites a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		//title := s.Find("a").Text()
		node := s.Get(0)
		href := ""
		name := ""
		for _, a := range node.Attr {
			if a.Key == "href" {
				href = a.Val
			}

			if a.Key == "data-product-name" {
				name = a.Val
			}
		}

		product := new(Product)
		product.Href = href
		product.Name = name
		products = append(products, product)
		//fmt.Printf("Review %d: href %s name %s\n", i, href, name)
	})
	return products
}
func GetGrigio4(products []*Product) []*Product {

	// Request the HTML page.
	res, err := http.Get("https://mirceramiki.kz/collection/laredo")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".wrap .content-full .collection-list .card-inner .card-price .add-favorites a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		//title := s.Find("a").Text()
		node := s.Get(0)
		href := ""
		name := ""
		for _, a := range node.Attr {
			if a.Key == "href" {
				href = a.Val
			}

			if a.Key == "data-product-name" {
				name = a.Val
			}
		}

		product := new(Product)
		product.Href = href
		product.Name = name
		products = append(products, product)
		//fmt.Printf("Review %d: href %s name %s\n", i, href, name)
	})
	return products
}
func SaveExcel(products []*Product) {
	//запись в EXCEL
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1") //Sheet2 до 2 страниц
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set value of a cell.
	counter := 2
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(1), "Name")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(1), "Href")
	for _, n := range products {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(counter), n.Name)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(counter), n.Href)
		counter++
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Products.xlsx"); err != nil {
		fmt.Println(err)
	}

}
func main() {
	products := []*Product{}

	products = GetProductsKeramic(products)
	//products = GetGrigio(products)
	//products = GetGrigio1(products)
	//products = GetGrigio2(products)
	//products = GetGrigio3(products)
	//products = GetLaminam(products)
	SaveExcel(products)

}
