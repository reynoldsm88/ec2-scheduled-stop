package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	//"github.com/aws/aws-lambda-go/events"
	//"github.com/aws/aws-lambda-go/lambda"
	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/ec2"
)

func handler(ctx context.Context, event events.CloudWatchEvent) {

	session, initErr := session.NewSession()

	if initErr != nil {
		fmt.Errorf("error creating ec2 client")
	} else {
		fmt.Println("client init successfully...")
		ec2Client := ec2.New(session)

		params := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				&ec2.Filter{
					Name: aws.String("tag:dart.owner"),
					Values: []*string{
						aws.String("michael"),
					},
				},
			},
		}

		fmt.Printf("params = %s", params)

		result, err := ec2Client.DescribeInstances(params)
		fmt.Printf("results = %s", result)
		if err != nil {
			fmt.Errorf("unable to describe instances")
		} else {
			for _, i := range result.Reservations[0].Instances {
				fmt.Println(i.InstanceId)
			}
		}
	}
}

func main() {
	lambda.Start(handler)
}
