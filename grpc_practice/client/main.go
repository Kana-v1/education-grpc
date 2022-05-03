package client

import (
	"context"
	"flag"
	"fmt"
	"io"
	"my_grpc/grpc_practice/server"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const port = ":9000"

func main() {
	option := flag.Int("o", 1, "Command to run")
	flag.Parse()
	creds, err := credentials.NewClientTLSFromFile("certificates/localhost.crt", "")
	if err != nil {
		panic(err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial("localhost"+port, opts...)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := server.NewEmployeeServiceClient(conn)

	switch *option {
	case 1:
		SendMetadata(client)
	case 2:
		GetByBadgeNumber(client)
	case 3:
		GetAll(client)
	case 4:
		AddPhoto(client)
	case 5:
		SaveAll(client)
	}
}

func SaveAll(client server.EmployeeServiceClient) {
	employees := []server.Employee{
		{
			BadgeNumber:         123,
			FirstName:           "John",
			LastName:            "Smith",
			VacationAccrualRate: 1.2,
			VacationAccured:     0,
		},
		{
			BadgeNumber:         234,
			FirstName:           "Lisa",
			LastName:            "Wu",
			VacationAccrualRate: 1.7,
			VacationAccured:     10,
		},
	}

	stream, err := client.SaveAll(context.Background())
	if err != nil {
		panic(err)
	}

	doneCH := make(chan struct{})

	// we dont know when a server would send response top our request, so its better not to block main thread
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				doneCH <- struct{}{}
				return
			}

			if err != nil {
				panic(err)
			}

			fmt.Println(res.Employee)
		}
	}()

	for _, e := range employees {
		stream.Send(&server.EmployeeRequest{Employee: &e})
	}

	stream.CloseSend()

	<-doneCH
}

func AddPhoto(client server.EmployeeServiceClient) {
	f, err := os.Open("image.jpeg")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	md := metadata.New(map[string]string{"badge_number": "2080"})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.AddPhoto(ctx)

	for {
		chunk := make([]byte, 30*100)
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		if n < len(chunk) {
			chunk = chunk[:n]
		}

		stream.Send(&server.AddPhotoRequest{Data: chunk})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	fmt.Println(res.IsOK)
}

func GetAll(client server.EmployeeServiceClient) {
	stream, err := client.GetAll(context.Background(), &server.GetAllRequest{})
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(res.Employee)
	}
}

func SendMetadata(client server.EmployeeServiceClient) {
	md := metadata.MD{}

	md["user"] = []string{"mvalsign"}
	md["password"] = []string{"password1"}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	_, err := client.GetByBadgeNumber(ctx, &server.GetByBadgeNumberRequest{})
	if err != nil {
		fmt.Println("ERRROR: ", err)
	}
}

func GetByBadgeNumber(client server.EmployeeServiceClient) {
	res, err := client.GetByBadgeNumber(context.Background(),
		&server.GetByBadgeNumberRequest{BadgeNumber: 2080})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func Start() {
	main()
}
