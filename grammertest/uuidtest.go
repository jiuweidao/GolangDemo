package main

import (
	_"github.com/pborman/uuid"
	"github.com/anker-dev/infra/uuidUtil"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main()  {
	fmt.Println("uuidUtil.GetUniqueNumber():",uuidUtil.GetUniqueNumber())
	fmt.Println("uuidUtil.GetUUID():",uuidUtil.GetUUID())

	sess := session.New(&aws.Config{
		Region:      aws.String(getS3Region()),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, ""),
	})
}
