package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	testFile()
}

func testFile() {
	// 使用ioutil包
	file, _ := ioutil.ReadFile("./ch193/file.txt")
	fmt.Println(string(file))

	file1, _ := ioutil.ReadDir("./ch193")
	for _, v := range file1 {
		fmt.Println(v)
	}

	fmt.Println("=================================")
	// 使用os包
	f, err := os.Open("./ch193/file.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer f.Close()
		var str string

		for {
			buf := make([]byte, 1024)
			n, err := f.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			str += string(buf)
		}
		fmt.Println(str)
	}
	fmt.Println("=================================")
	f, err = os.Open("./ch193/file.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer f.Close()
		str, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(str))
		}
	}
	fmt.Println("=================================")
	f, err = os.Open("./ch193/file.txt")
	str_arr := make([]string, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		reg := regexp.MustCompile(`^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$`)
		for scanner.Scan() {
			str := scanner.Text()
			if ok := reg.MatchString(str); ok {
				if err != nil {
					fmt.Println(err.Error())
				} else {
					str_arr = append(str_arr, str)
				}
			}
		}
		fmt.Println(str_arr)
	}
}
