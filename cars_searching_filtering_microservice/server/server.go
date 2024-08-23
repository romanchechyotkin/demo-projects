package server

import (
	"bytes"
	"cars_searching_filtering/proto/pb"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"time"
)

type server struct {
	elasticsearch *elasticsearch.Client
}

func NewServer(client *elasticsearch.Client) (*server, error) {
	srv := &server{
		elasticsearch: client,
	}

	return srv, nil
}

func (s *server) CreateCarDoc(ctx context.Context, req *pb.CreateCarDocReq) (*pb.CreateCarDocRes, error) {
	_doc := struct {
		CarId       string    `json:"car_id"`
		Brand       string    `json:"brand"`
		Model       string    `json:"model"`
		Year        uint32    `json:"year"`
		PricePerDay float32   `json:"price_per_day"`
		IsAvailable bool      `json:"is_available"`
		Rating      float32   `json:"rating"`
		Username    string    `json:"username"`
		City        string    `json:"city"`
		PublishedAt time.Time `json:"published_at"`
	}{
		CarId:       req.CarId,
		Brand:       req.Brand,
		Model:       req.Model,
		Year:        req.Year,
		PricePerDay: req.PricePerDay,
		IsAvailable: req.IsAvailable,
		Rating:      req.Rating,
		Username:    req.Username,
		City:        req.City,
		PublishedAt: time.Now(),
	}

	data, err := json.Marshal(_doc)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	query := esapi.IndexRequest{
		Index:   "cars",
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	res, err := query.Do(ctx, s.elasticsearch)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer res.Body.Close()

	var m map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return nil, err
	}

	id := m["_id"]
	log.Println("INSERTED", id)

	grpcResponse := &pb.CreateCarDocRes{Id: id.(string)}

	return grpcResponse, nil
}
