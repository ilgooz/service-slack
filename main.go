package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/bluele/slack"
	api "github.com/mesg-foundation/core/api/service"
	services "github.com/mesg-foundation/core/service"
	"google.golang.org/grpc"
)

type NotifyInputs struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Title   string `json:"title"`
}

type ErrorOutput struct {
	Err error `json:"err"`
}

type SuccessOutput struct {
	Err error `json:"err"`
}

type Output struct {
	Key  string      `json:"key"`
	Data interface{} `json:"data"`
}

var mesg api.ServiceClient
var service *services.Service

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func notify(inputs *NotifyInputs) (output Output) {
	apiSlack := slack.New(inputs.Token)
	err := apiSlack.ChatPostMessage(inputs.Channel, inputs.Title, &slack.ChatPostMessageOpt{})
	if err != nil {
		output = Output{
			Key: "error",
			Data: ErrorOutput{
				Err: err,
			},
		}
	} else {
		output = Output{
			Key:  "success",
			Data: SuccessOutput{},
		}
	}
	return
}

func main() {
	var err error
	service, err = services.ImportFromPath("./")
	handleError(err)

	endpoint := os.Getenv("MESG_ENDPOINT")

	connection, err := grpc.Dial(endpoint, grpc.WithInsecure())
	handleError(err)
	mesg = api.NewServiceClient(connection)

	stream, err := mesg.ListenTask(context.Background(), &api.ListenTaskRequest{
		Service: service,
	})
	handleError(err)
	for {
		taskData, err := stream.Recv()
		handleError(err)
		var inputs NotifyInputs
		err = json.Unmarshal([]byte(taskData.InputData), &inputs)
		handleError(err)
		var output Output
		switch taskData.TaskKey {
		case "notify":
			output = notify(&inputs)
		}
		data, err := json.Marshal(output.Data)
		handleError(err)
		_, err = mesg.SubmitResult(context.Background(), &api.SubmitResultRequest{
			ExecutionID: taskData.ExecutionID,
			OutputKey:   output.Key,
			OutputData:  string(data),
		})
		handleError(err)
	}
}
