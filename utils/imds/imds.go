package imds

import (
	"fmt"

	"github.com/nholuongut/nholuongut-sdk-go/nholuongut"
	ec2metadatasvc "github.com/nholuongut/nholuongut-sdk-go/nholuongut/ec2metadata"
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut/session"
)

// EC2Metadata wraps the methods from the amazon-sdk-go's ec2metadata package
type EC2Metadata interface {
	GetMetadata(path string) (string, error)
	Region() (string, error)
}

func GetMetaData(key string) (string, error) {
	nholuongutSession := session.Must(session.NewSession(nholuongut.NewConfig().
		WithMaxRetries(10),
	))
	var ec2Metadata EC2Metadata
	ec2Metadata = ec2metadatasvc.New(nholuongutSession)
	requestedData, err := ec2Metadata.GetMetadata(key)
	if err != nil {
		return "", fmt.Errorf("get instance metadata: failed to retrieve %s - %s", key, err)
	}
	return requestedData, nil
}
