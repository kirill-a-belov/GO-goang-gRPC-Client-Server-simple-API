/*
Приложение №1
Представляет из себя CRD сервер работающий по протоколу gRPC
*/

package main

import (
	pb "./kranxapi"
	"google.golang.org/grpc"
	"fmt"
	"log"
	"golang.org/x/net/context"
	"net/http"
	"time"
	"regexp"
	"strings"
	"html/template"
)
// RPC сервер
const (
	SERER_ADDRESS     = "localhost:54321"
)

var client pb.KranxApiClient
// WEB сервер
var server = &http.Server{
	Addr:           ":80",
	Handler:        nil,
	ReadTimeout:    1000 * time.Second,
	WriteTimeout:   1000 * time.Second,
	MaxHeaderBytes: 1 << 20,
}


//Функционал клиента
func AddData(key string, value string)string{

	r, err := client.AddData(context.Background(), &pb.AddRequest{Key: key, Value: value})
	if err != nil {
		log.Fatalf("Could not add data: %v", err)
	}
	return r.Message

}

func GetData(key string)string{

	r, err := client.GetData(context.Background(), &pb.GetRequest{Key: key})
	if err != nil {
		log.Fatalf("Could not get data: %v", err)
	}
	return r.Value

}
func DelData(key string)string{

	r, err := client.DelData(context.Background(), &pb.DelRequest{Key: key})
	if err != nil {
		log.Fatalf("Could not del data: %v", err)
	}
	return r.Message

}

//Функционал вэб сервера

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var re = regexp.MustCompile(`[a-zA-Z0-9]+`)
	type statVars struct {
		Rtype string
		Rkey      string
		Response   string
	}

	fmt.Println(r.URL.Path)
	match := re.FindAllString(r.URL.Path, -1)
	fmt.Println("Request is: "+ strings.ToUpper(match[0]))


	switch strings.ToUpper(match[0]) {
	case "GET":
		p := statVars{Rtype: "GET", Rkey: match[1], Response: GetData(match[1]) }
		t, _ := template.ParseFiles("./statistic.html")
		t.Execute(w, p)
	case "PUT":
		p := statVars{Rtype: "GET", Rkey: match[1], Response: AddData(match[1],match[2]) }
		t, _ := template.ParseFiles("./statistic.html")
		t.Execute(w, p)
	case "DELETE":
		p := statVars{Rtype: "GET", Rkey: match[1], Response: DelData(match[1]) }
		t, _ := template.ParseFiles("./statistic.html")
		t.Execute(w, p)

	}
}










func main() {

	// Настраиваем соединение с RPC сервером
	conn, err := grpc.Dial(SERER_ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can not connect to gRPC server: %v", err)
	} else {
		fmt.Println("Connected t gRPC server!")
	}
	defer conn.Close()
	client = pb.NewKranxApiClient(conn)

	// Настраиваем WEB сервер
	http.HandleFunc("/", requestHandler)
	fmt.Println("WEB server starts")
	fmt.Println()
	fmt.Println("Requests:")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("WEB server wont started!!!")
		fmt.Println(err)
	}


	/*
	AddData("1", "first")
	AddData("2", "A")
	AddData("3", "@")
	AddData("4", "@")
	AddData("5", "@")
	AddData("6", "@")
	AddData("7", "@")
	AddData("8", "@")
	AddData("9", "@")
	AddData("10", "@")
	AddData("11", "@")
	GetData("5")
	DelData("3")
	*/

}
