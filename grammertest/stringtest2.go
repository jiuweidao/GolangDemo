package main

import (
	"os"
	"strings"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"log"
)

type Song struct {
	Title string
	Seconds int
}

func main()  {
	if len(os.Args)==1||!strings.HasSuffix(os.Args[1],".txt"){
		fmt.Printf("useage: %s  <file.txt>\n",filepath.Base(os.Args[0]))
		fmt.Println(os.Args)
	}

	if rawBytes, err:= ioutil.ReadFile(os.Args[1]);err!=nil{
		log.Fatal(err)
	}else {
		fmt.Printf("rawBytes:%v\n",rawBytes)
		fmt.Printf("string(rawBytes):%s\n",string(rawBytes))
		songs := readSongs(string(rawBytes))
		writeSongs(songs)
	}
}

func readSongs(songMassge string)  (songs []Song){
	var song Song//{"大白天 ",23}
	for i,line :=range strings.Split(songMassge,"\r\n")   {
		fmt.Printf("%d 读取到%v\n",i,line)
		line = strings.TrimSpace(line)
		if strings.Contains(line,"name"){
			song.Title,song.Seconds = "we",12
		}
		if song.Title!="" && song.Seconds!=0{
			songs = append(songs, song)
		}

	}
	return  songs
}
func writeSongs(songs []Song)  {

	fmt.Println("打印歌曲")
	for i,song:= range songs {
		i++
		fmt.Printf("Title=%s,second =%d\n",song.Title,song.Seconds)
	}
	fmt.Printf("一共%d首歌曲",len(songs))

}

