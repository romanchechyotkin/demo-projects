package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/xitongsys/parquet-go-source/hdfs"
	"github.com/xitongsys/parquet-go/writer"
)

const apiURL = "https://api.binance.com/api/v3/avgPrice?symbol=BTCUSDT"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go process(ctx)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("application got signal", s.String())
		cancel()
		time.Sleep(3 * time.Second) // таумаут для завершения работы приложения
	}
}

type Record struct {
	ID         string  `parquet:"name=id, type=BYTE_ARRAY"`
	Currency   string  `parquet:"name=currency, type=BYTE_ARRAY"`
	PriceFloat float64 `parquet:"name=price_float, type=DOUBLE"`
	PriceInt   int64   `parquet:"name=price_int, type=INT64"`
	Price      string  `json:"price" parquet:"name=price_str, type=BYTE_ARRAY"`
	Timestamp  int64   `parquet:"name=timestamp, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

func process(ctx context.Context) {
	t := time.NewTicker(3 * time.Second)

	fw, err := hdfs.NewHdfsFileWriter([]string{"localhost:8020"}, "hadoop", "output.parquet")
	if err != nil {
		log.Println("Can't create file", err)
		return
	}

	//write
	pw, err := writer.NewParquetWriter(fw, new(Record), 4)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("process stopped")
			pw.WriteStop()
			fw.Close()
			t.Stop()
			return
		case <-t.C:
			var r = Record{
				ID:        uuid.NewString(),
				Currency:  "BTC",
				Timestamp: time.Now().Unix(),
			}

			resp, err := http.Get(apiURL)
			if err != nil {
				log.Println("failed to get reponse", err)
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				log.Println("failed to get successful response", resp.StatusCode)
				continue
			}

			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("failed to read reponse body", err)
				continue
			}

			if err = json.Unmarshal(respBody, &r); err != nil {
				log.Println("failed to unmarshal", err)
				continue
			}

			r.PriceFloat, err = strconv.ParseFloat(r.Price, 64)
			if err != nil {
				log.Println("failed to parse float", err)
				continue
			}

			r.PriceInt = int64(r.PriceFloat * math.Pow10(8))

			if err = pw.Write(r); err != nil {
				log.Println("failed to write to parquet file")
				continue
			}

			log.Println("successful write", r)
		}
	}
}
