package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){

	seedUrl := "https://www.douban.com/"
	resp , err := http.Get(seedUrl)
	if err != nil {

		fmt.Println(err.Error())
		panic(err)
	}

	defer resp.Body.Close()

	if res, err := ioutil.ReadAll(resp.Body); err != nil {

		fmt.Println(err.Error())
		panic(err)
	}else{
		fmt.Println(string(res))
	}


}
