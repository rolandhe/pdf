package main

import (
	pdf "github.com/adrg/go-wkhtmltopdf"
	"log"
	"os"
	"sync"
)

func main() {
	if err := pdf.Init(); err != nil {
		log.Fatal(err)
	}
	defer pdf.Destroy()

	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan struct{}, 1)
	go func() {
		ch <- struct{}{}
		toPdf("sample.html", "out1.pdf")
		<-ch
		wg.Done()
	}()
	go func() {
		ch <- struct{}{}
		toPdf("sample2.html", "out2.pdf")
		<-ch
		wg.Done()
	}()
	wg.Wait()
}

func toPdf(htmlFile string, pdfFile string) {
	// Create object from file.
	object, err := pdf.NewObject(htmlFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create converter.
	converter, err := pdf.NewConverter()
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()

	// Add created objects to the converter.
	converter.Add(object)

	// Set converter options.
	converter.Title = "Sample document"
	converter.PaperSize = pdf.A4
	converter.Orientation = pdf.Landscape
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "10mm"
	converter.MarginRight = "10mm"

	// Convert objects and save the output PDF document.
	outFile, err := os.Create(pdfFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	if err := converter.Run(outFile); err != nil {
		log.Fatal(err)
	}
}
