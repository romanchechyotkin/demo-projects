package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
)

// Student struct
type Student struct {
	Name   string  `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"name"`
	Age    int32   `parquet:"name=age, type=INT32, encoding=PLAIN" json:"age"`
	Id     int64   `parquet:"name=id, type=INT64" json:"id"`
	Weight float32 `parquet:"name=weight, type=FLOAT" json:"weight"`
	Sex    bool    `parquet:"name=sex, type=BOOLEAN" json:"sex"`
	Day    int32   `parquet:"name=day, type=INT32, convertedtype=DATE" json:"day"`
}

const totalRecords = 1_000_000

func main() {
	start := time.Now()

	
	parquetFile := "output/data.parquet" // parquet файл
	jsonFile := "output/data.json"	     // json файл

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeParquet(parquetFile, totalRecords)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeJSON(jsonFile, totalRecords)
	}()

	wg.Wait()
	log.Printf("Finished writing %d records to Parquet & JSON in %v\n", totalRecords, time.Since(start))
}

// функция для записи в parquet файл
func writeParquet(filename string, num int) {
	fw, err := local.NewLocalFileWriter(filename)
	if err != nil {
		log.Fatal("Can't create local file:", err)
	}
	defer fw.Close()

	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
	if err != nil {
		log.Fatal("Can't create parquet writer:", err)
	}
	defer pw.WriteStop()

	pw.RowGroupSize = 128 * 1024 * 1024 // 128M
	pw.PageSize = 8 * 1024              // 8K
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	// Writing records
	for i := 0; i < num; i++ {
		record := Student{
			Name:   "StudentName",
			Age:    20 + int32(i%5),
			Id:     int64(i),
			Weight: 50.0 + float32(i)*0.1,
			Sex:    i%2 == 0,
			Day:    int32(time.Now().Unix() / 3600 / 24),
		}
		if err := pw.Write(record); err != nil {
			log.Println("Parquet Write error:", err)
		}
	}

	log.Println("Parquet write completed")
}

// функция для записи в json файл
func writeJSON(filename string, num int) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Can't create JSON file:", err)
	}
	defer file.Close()

	res := make([]Student, 0, num)
	for i := 0; i < num; i++ {
		res = append(res, Student{
			Name:   "StudentName",
			Age:    20 + int32(i%5),
			Id:     int64(i),
			Weight: 50.0 + float32(i)*0.1,
			Sex:    i%2 == 0,
			Day:    int32(time.Now().Unix() / 3600 / 24),
		})
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(res)

	log.Println("JSON write completed")
}
