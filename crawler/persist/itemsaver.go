package persist

import (
	"github.com/olivere/elastic"
	"log"
	"os"
)

func ItemSaver() chan interface{}{
	out := make(chan interface{})
	go func() {
		itemCount :=0
		for{
			item := <-out
			log.Panicf("Item Saver: got item "+ "#%d: %v",itemCount,item)
			itemCount++

			id, err :=save(item)
			if err != nil{
				log.Printf("Item  Saver: error"+"saving item %v: %v",item,err)
			}
		}

	}()
	return out
}

func ESClient() (client *elastic.Client,err error){
	file := "./log.log"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766) // 应该判断error，此处简略
	cfg := []elastic.ClientOptionFunc{
		//elastic.SetURL("http://10.1.1.1:9200"),
		elastic.SetURL("localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(logFile, "ES-INFO: ", 0)),
		elastic.SetTraceLog(log.New(logFile, "ES-TRACE: ", 0)),
		elastic.SetErrorLog(log.New(logFile, "ES-ERROR: ", 0)),
	}
	client,err =elastic.NewClient(cfg ...)
	return
}

func save(item interface{}) (id string,err error) {
	//http.Post()
	client, err := ESClient()

	defer client.Stop()
	if err != nil {
		log.Fatalf("Error creating the client: %s\n", err)
		return "",err

	}


	res,err :=client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do()

	if err !=nil{
		return "",err
	}
	log.Println()
	return res.Id,nil
}



