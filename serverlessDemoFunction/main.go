package main

import (
    "time"
    "github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
    Message   string    `json:"message"`
    Timestamp time.Time `json:"timestamp"`
}

func hello() (Response, error) {
    return Response{
        Message:   "Hello, Habr!",
        Timestamp: time.Now(),
    }, nil
}

func main() {
    lambda.Start(hello)
}