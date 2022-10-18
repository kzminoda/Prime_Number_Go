package main

import (
	"fmt"
	// 	"math"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var pList = []int64{}
var pList_file string = "./pList.txt"

func main() {
	init_pLIst()
	pList = append(pList, 7, 11, 13, 17, 19, 23)
	fmt.Println(pList)
	sync_pList()
}

func init_pLIst() {
	_, error := os.Stat(pList_file)

	if error != nil {
		pList = append(pList, 2)
	} else {
		fp, error := os.Open(pList_file)

		if error != nil {
			panic(error)
		}

		defer fp.Close()

		s := bufio.NewScanner(fp)
		for s.Scan() {
			tmp := s.Text()
			// fmt.Printf("[[%s]]", tmp)
			tmp = strings.TrimRight(tmp, "\n")
			if strings.HasSuffix(tmp, "\r") {
				tmp = strings.TrimRight(tmp, "\r")
			}
			i, _ := strconv.ParseInt(tmp, 10, 64)
			pList = append(pList, i)
		}
	}
}

func sync_pList() {
	file, error := os.Create(pList_file)

	if error != nil {
		panic(error)
	}

	defer file.Close()

	for _, v := range pList {
		line := fmt.Sprintf("%v", v)
		_, error := file.WriteString(fmt.Sprintf("%s\n", line))

		if error != nil {
			panic(error)
		}
	}
}

func is_prime_number(n int64) bool {
	var c int64
	
	for _, x := range pList {
		
		if 
	}
}
