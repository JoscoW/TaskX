package tool

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func EnsureEnv() []string {
	env := os.Environ()

	if _, ok := os.LookupEnv("DB_ROOT_PASSWORD"); !ok {
		env = append(env, "DB_ROOT_PASSWORD=pass")
	}

	return env
}

func GetClientConnection() (*grpc.ClientConn, error) {

	port := 9000
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("fail to dial: %v", err)
	}

	return conn, nil

}
