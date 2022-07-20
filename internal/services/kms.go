package services

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func EncryptString(data string) []byte {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create KMS service client
	svc := kms.New(sess)

	// Encrypt data key
	//
	// Replace the fictitious key ARN with a valid key ID

	keyId := "arn:aws:kms:us-west-2:740864632595:key/e6f70a12-3a6c-4271-94aa-0eab95c401f2"

	// Encrypt the data
	result, err := svc.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(keyId),
		Plaintext: []byte(data),
	})

	if err != nil {
		fmt.Println("Got error encrypting data: ", err)
		os.Exit(1)
	}

	return result.CiphertextBlob
}
