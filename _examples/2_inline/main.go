package main

import (
	"fmt"
	"os"

	"github.com/kataras/tableprinter"
)

type (
	book struct {
		Title       string    `header:"title"`
		Description string    `header:"desc"`
		Sales       int       `header:"sales"`
		Publisher   publisher `header:"inline"`
	}

	publisher struct {
		Name    string  `header:"publisher name"`
		Country country `header:"publisher country"`
	}

	country struct {
		Name string
		Code string
	}
)

func (c country) String() string {
	return c.Name
}

func main() {
	n := 42
	books := make([]book, n, n)
	var b book

	for i := 1; i <= n; i++ {
		b.Title = fmt.Sprintf("Title for Book %d", i)
		b.Description = fmt.Sprintf("Description for Book %d", i)
		b.Sales = i * 12000
		b.Publisher = publisher{
			fmt.Sprintf("Publisher Name %d", i),
			country{fmt.Sprintf("Country Name for Publisher %d", i), "Code doesn't matter"},
		}

		books[i-1] = b
	}

	tableprinter.Print(os.Stdout, books)
}