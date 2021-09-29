package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
	"github.com/lab259/cors"
)

type data struct {
	Status status `json:"status"`
}

type status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	fmt.Println("server run on port 8082")

	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/data":
			returnJSONFile(ctx)
		default:
			ctx.Success(fasthttp.StatusMessage(200),[]byte("Health check is okay :)"))
		}
	}

	timer := time.Tick(15 * time.Second)
	go genData(timer)

	handler := cors.Default().Handler(m)

	if err := fasthttp.ListenAndServe(":8082", handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

}

func returnJSONFile(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("./data.json")
}

func genData(timer <-chan time.Time)  {
	nData := data{}
	genData := func() {
		nWind := rand.Intn(100)
		nWater := rand.Intn(100)

		nData = data{
			Status: status{
				Water: nWater,
				Wind:  nWind,
			},
		}
	}

	for next := range timer {
		genData()
		fmt.Println("logger: a new data generated at",next.Format("2006-01-02 15:04:05 "),"data:",nData.Status)
		out, _ := json.Marshal(nData)
		_ = ioutil.WriteFile("./data.json", out, 0777)
	}

}
