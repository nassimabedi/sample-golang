package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type TestObject struct {
	ObjectIds []int   `json:"object_ids"`
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

	fmt.Println(objId)
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

