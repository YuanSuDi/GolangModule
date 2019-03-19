package main

import (
	"bufio"
	"fmt"
	"github.com/Luxurioust/excelize"
	"os"
	"strconv"
)

func main()  {
	xlsx, err := excelize.OpenFile("quite.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vcf, err := os.Create("quite.vcf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer vcf.Close()
	vcfWrite := bufio.NewWriter(vcf)
	defer vcfWrite.Flush()
	//Get sheet index.
	index := xlsx.GetSheetIndex("Sheet1")
	// Get all the rows in a sheet.
	rows := xlsx.GetRows("Sheet" + strconv.Itoa(index))
	count := 0
	for _, row := range rows {
		count++
		if count==1{
			continue
		}
		fmt.Fprintln(vcfWrite,"BEGIN:VCARD")
		fmt.Fprintln(vcfWrite,"VERSION:3.0")

		for j:=0 ;j<2 ; j++ {
			if j==0{
				fmt.Fprintf(vcfWrite,"FN:%s\n", row[j])
			}else if j==1{
				fmt.Fprintf(vcfWrite,"TEL;TYPE=CELL:%s\n", row[j])
			}

		}
		fmt.Fprintln(vcfWrite,"END:VCARD")

	}
}
