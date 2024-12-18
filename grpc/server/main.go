package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "server/service"
)

const PORT = ":5000"
const NETWORK = "tcp"

type server struct {
	orderMap map[string]*pb.OrderMessage
	pb.UnimplementedOrderManagementServer
}

func (s *server) GetOrder(ctx context.Context, in *pb.ValueMessage) (*pb.OrderMessage, error) {
	order, exists := s.orderMap[in.Value]
	if exists {
		return order, nil
	}
	return nil, errors.New("no such order")
}

func (s *server) SearchOrders(searchQuery *pb.ValueMessage, stream pb.OrderManagement_SearchOrdersServer) error {
	log.Println("searching for", searchQuery.Value)
	for k, order := range s.orderMap {
		log.Println(k, order)

		for _, item := range order.Items {
			log.Println(item)

			if strings.Contains(item, searchQuery.Value) {
				time.Sleep(3 * time.Second)
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " + k)
				break
			}
		}
	}
	return nil
}

func (s *server) UpdateOrders(stream grpc.ClientStreamingServer[pb.OrderMessage, emptypb.Empty]) error {
	ordersStr := "Updated Order IDs : "

	for {
		var order pb.OrderMessage
		err := stream.RecvMsg(&order)
		if err == io.EOF {
			return stream.SendMsg(&emptypb.Empty{})
		}
		s.orderMap[fmt.Sprint("%d", order.Id)] = &order
		log.Println("Order ID", order.Id, "Updated")
		ordersStr += fmt.Sprintf("%d, ", order.Id)
	}
}

// func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
// 	const orderBatchSize = 3
// 	combinedShipmentMap := make(map[string]pb.CombinedShipment)
//
// 	for {
// 		orderId, err := stream.Recv()
// 		fmt.Println(orderId)
// 		if err == io.EOF {
// 			for _, comb := range combinedShipmentMap {
// 				stream.Send(&comb)
// 			}
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
//
// 		order, err := s.GetOrder(context.Background(), orderId)
// 		log.Println("got", order)
// 		if err != nil {
// 			log.Println(err)
// 			return err
// 		}
//
// 		val, ok := combinedShipmentMap[order.Destination]
// 		log.Println(val)
// 		if !ok {
// 			log.Println("created ")
// 			combinedShipmentMap[order.Destination] = pb.CombinedShipment{}
// 		}
//
// 		val.OrdersList = append(val.OrdersList, order)
// 		log.Println("added", val.OrdersList, len(val.OrdersList), cap(val.OrdersList))
//
// 		if len(val.OrdersList) == orderBatchSize {
// 			// передаем клиенту поток заказов, объединенных в партии
// 			for _, comb := range combinedShipmentMap {
// 				log.Println(comb)
// 				// передаем клиенту партию объединенных заказов
// 				stream.Send(&comb)
// 			}
// 		}
//
// 	}
//
// }

func main() {
	var srv = &server{
		orderMap: map[string]*pb.OrderMessage{
			"1": {
				Id:          1,
				Items:       []string{"Apple", "Orange"},
				Description: "qwerty",
				Price:       1230,
				Destination: "Minsk",
			},
			"2": {
				Id:          2,
				Items:       []string{"Apple"},
				Description: "qwerty",
				Price:       99901.1,
				Destination: "Minsk",
			},
			"3": {
				Id:          3,
				Items:       nil,
				Description: "qwerty",
				Price:       0,
				Destination: "Minsk",
			},
			"4": {
				Id:          4,
				Items:       []string{"Orange"},
				Description: "qwerty",
				Price:       12.321,
				Destination: "Minsk",
			},
			"5": {
				Id:          5,
				Items:       []string{"Coca-cola"},
				Description: "qwerty",
				Price:       1234,
				Destination: "Minsk",
			},
		},
	}

	listen, err := net.Listen(NETWORK, PORT)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(orderUnaryServerInterceptor),
		grpc.StreamInterceptor(orderStreamServerInterceptor),
	)
	pb.RegisterOrderManagementServer(s, srv)
	log.Printf("Starting gRPC listener on port " + PORT)
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func orderUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("------------------ [Server Interceptor]", info.FullMethod)

	m, err := handler(ctx, req)

	log.Printf(" Post Proc Message : %s", m)
	return m, err
}

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("====== [Server Stream Interceptor Wrapper] "+
		"Receive a message (Type: %T) at %s",
		m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("====== [Server Stream Interceptor Wrapper] "+
		"Send a message (Type: %T) at %v",
		m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func orderStreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("====== [Server Stream Interceptor] ", info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}
