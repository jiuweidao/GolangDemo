package main

import (
	"os"
	"fmt"
	"path/filepath"
	"log"
)

func main()  {
	if len(os.Args) == 1{
		fmt.Printf("usage:%s<wholel-number>\n",filepath.Base(os.Args[0]))
		//os.Exit(1)
	}


	strOfDigits := os.Args[0]

	for column :=range strOfDigits  {
		line :=""
		digit := strOfDigits[column]
		if 0<digit &&digit<=9{
			line+= "drf"
		}else {
			log.Fatal("invalid whole number")
		}
		fmt.Println(line)
	}

}
