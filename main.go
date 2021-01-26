package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func handler(ctx context.Context, event events.CloudWatchEvent) {
	fmt.Println("hello lambda")

	session, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})

	if err != nil {
	}

	ec2Client := ec2.New(session)

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(""),
			},
		},
	}

	fmt.Printf("%s", params.String())

	if ec2Client == nil {

	}

}
func main() {
	lambda.Start(handler)
}
