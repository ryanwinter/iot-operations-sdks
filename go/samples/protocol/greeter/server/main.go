// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Azure/iot-operations-sdks/go/mqtt"
	"github.com/Azure/iot-operations-sdks/go/protocol"
	"github.com/Azure/iot-operations-sdks/go/samples/protocol/greeter/envoy"
	"github.com/lmittmann/tint"
)

type Handlers struct{}

func main() {
	ctx := context.Background()
	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, nil)))
	app := must(protocol.NewApplication())

	mqttClient := must(mqtt.NewSessionClientFromEnv(
		mqtt.WithLogger(slog.Default()),
	))
	server := must(envoy.NewGreeterServer(app, mqttClient, &Handlers{}))
	defer server.Close()

	check(mqttClient.Start())
	check(server.Start(ctx))

	fmt.Println("Press enter to quit.")
	must(fmt.Scanln())
}

func (Handlers) SayHello(
	ctx context.Context,
	req *protocol.CommandRequest[envoy.HelloRequest],
) (*protocol.CommandResponse[envoy.HelloResponse], error) {
	slog.Info(
		"--> Greeter.SayHello",
		slog.String("id", req.CorrelationData),
		slog.String("client", req.ClientID),
	)
	defer slog.Info(
		"<-- Greeter.SayHello",
		slog.String("id", req.CorrelationData),
		slog.String("client", req.ClientID),
	)

	return protocol.Respond(envoy.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Payload.Name),
	})
}

func (Handlers) SayHelloWithDelay(
	ctx context.Context,
	req *protocol.CommandRequest[envoy.HelloWithDelayRequest],
) (*protocol.CommandResponse[envoy.HelloResponse], error) {
	slog.Info(
		"--> Greeter.SayHelloWithDelay",
		slog.String("id", req.CorrelationData),
		slog.String("client", req.ClientID),
	)
	defer slog.Info(
		"<-- Greeter.SayHelloWithDelay",
		slog.String("id", req.CorrelationData),
		slog.String("client", req.ClientID),
	)

	delay := time.Duration(req.Payload.Delay)
	time.Sleep(delay)
	return protocol.Respond(envoy.HelloResponse{
		Message: fmt.Sprintf("Hello %s after %s", req.Payload.Name, delay),
	})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func must[T any](t T, e error) T {
	check(e)
	return t
}
