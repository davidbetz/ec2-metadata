package main

import (
        "fmt"
        "log"
        "strings"
		"bytes"
		"os"

        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/aws/ec2metadata"
)

func main() {
		bucket := os.Getenv("BUCKET")
		if len(bucket) == 0 {
			log.Fatal("BUCKET is required.")
		}

        sess := session.Must(session.NewSessionWithOptions(session.Options {
                SharedConfigState: session.SharedConfigEnable,
        }))

        ec2svc := ec2metadata.New(sess)
        id, err := ec2svc.GetInstanceIdentityDocument()
        if err != nil {
                log.Fatal(err)
        }

		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("DevpayProductCodes: %q\n", id.DevpayProductCodes))
		sb.WriteString(fmt.Sprintf("MarketplaceProductCodes: %q\n", id.MarketplaceProductCodes))
		sb.WriteString(fmt.Sprintf("AvailabilityZone: %q\n", id.AvailabilityZone))
		sb.WriteString(fmt.Sprintf("PrivateIP: %q\n", id.PrivateIP))
		sb.WriteString(fmt.Sprintf("Version: %q\n", id.Version))
		sb.WriteString(fmt.Sprintf("Region: %q\n", id.Region))
		sb.WriteString(fmt.Sprintf("InstanceID: %q\n", id.InstanceID))
		sb.WriteString(fmt.Sprintf("BillingProducts: %q\n", id.BillingProducts))
		sb.WriteString(fmt.Sprintf("InstanceType: %q\n", id.InstanceType))
		sb.WriteString(fmt.Sprintf("AccountID: %q\n", id.AccountID))
		sb.WriteString(fmt.Sprintf("PendingTime: %q\n", id.PendingTime))
		sb.WriteString(fmt.Sprintf("ImageID: %q\n", id.ImageID))
		sb.WriteString(fmt.Sprintf("KernelID: %q\n", id.KernelID))
		sb.WriteString(fmt.Sprintf("RamdiskID: %q\n", id.RamdiskID))
		sb.WriteString(fmt.Sprintf("Architecture: %q\n", id.Architecture))

		fmt.Println(sb.String())
        s3svc := s3.New(sess)
        _, erra := s3svc.PutObject(&s3.PutObjectInput {
                Bucket: aws.String(bucket),
                Key: aws.String(id.InstanceID + ".txt"),
                Body: bytes.NewReader([]byte(sb.String())),
        })
        if erra != nil {
                log.Fatal(erra)
        }
}
