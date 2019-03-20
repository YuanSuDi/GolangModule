package main

import (
	"bufio"
	"fmt"
	"fujitsu.com/inter/ExeclToVcf/vcf"
	"github.com/Luxurioust/excelize"
	"os"
	"strconv"
	"strings"
)

//将execl文件转化成标准联系人vcf文件
func main()  {
	start :
		tip := `please input execl filename : 
input 'exit' is exit --->`
		fmt.Printf(tip)
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			goto start
		}
		filename := string(line)
		if filename == "exit" {
			os.Exit(0)
		}
		//start read execl file
		xlsx, err := excelize.OpenFile(filename)
		if err != nil {
			fmt.Println(err)
			goto start
		}
		vcfFile, err := os.Create(strings.Split(filename, ".")[0]+".vcf")
		if err != nil {
			fmt.Println(err)
			goto start
		}

		vcfWrite := bufio.NewWriter(vcfFile)
		contact := vcf.Contact{Header: "BEGIN:VCARD\nVERSION:3.0",Footer:"END:VCARD"}
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
			fmt.Fprintln(vcfWrite,contact.Header)
			contact.ReadRow(row)
			fmt.Fprintf(vcfWrite,"%s\n", contact.GetName())
			fmt.Fprintf(vcfWrite,"%s\n", contact.GetCellPhone())
			fmt.Fprintln(vcfWrite,contact.Footer)
		}
		fmt.Println("successful!")
		vcfWrite.Flush()
		vcfFile.Close()
	goto start
}


