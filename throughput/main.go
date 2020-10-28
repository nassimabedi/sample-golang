package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type TestObject struct {
	ObjectIds []int   `json:"object_ids"`
}

type ObjectInfo struct {
	Id int   `json:"id"`
	Online bool   `json:"online"`
}






func main() {
    http.HandleFunc("/", HelloServer)
	http.HandleFunc("/callback", CallBack)
    http.ListenAndServe(":9090", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}


func getObjectInfo(objId int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	defer wg.Done()
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Println("panic occurred:", err)
	//	}
	//}()

	fmt.Println("==========", objId)
	//resp, err := http.Get("http://localhost:9010/objects/"+ string(objId))
	//resp, err := http.Get("http://localhost:9010/objects/"+objId)
	resp, err := http.Get(fmt.Sprintf("http://localhost:9010/objects/%d", objId))

	if err != nil {
		fmt.Println("errr1111111111111111111111111111")
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("errr2222222222222222222")
		panic(err)
	}

	fmt.Println(string(body))
	var t ObjectInfo
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Online)
}

func CallBack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "callBack, %s!", r.URL.Path[1:])
	fmt.Println("============================")

	//way 1 to get POST param
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(string(body))
	//var t test_struct
	//err = json.Unmarshal(body, &t)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(t.ObjectIds)
	//fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//way 2 to get POST param
	decoder := json.NewDecoder(r.Body)
	var a TestObject
	err := decoder.Decode(&a)
	if err != nil {
		panic(err)
	}
	log.Println(a.ObjectIds)
	objectList := a.ObjectIds
	var wg sync.WaitGroup
	for _, obj := range objectList {
		wg.Add(1)
		fmt.Println( obj)
		go getObjectInfo(obj, &wg)
	}
	wg.Wait()
}

