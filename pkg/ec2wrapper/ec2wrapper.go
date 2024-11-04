// Package ec2wrapper is used to wrap around the ec2 service APIs
package ec2wrapper

import (
	"github.com/nholuongut/amazon-vpc-cni-k8s/pkg/nholuongututils/nholuongutsession"
	"github.com/nholuongut/amazon-vpc-cni-k8s/pkg/ec2metadatawrapper"
	"github.com/nholuongut/amazon-vpc-cni-k8s/pkg/utils/logger"
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut"
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut/ec2metadata"
	"github.com/nholuongut/nholuongut-sdk-go/service/ec2"
	"github.com/nholuongut/nholuongut-sdk-go/service/ec2/ec2iface"
	"github.com/pkg/errors"
)

const (
	resourceID   = "resource-id"
	resourceKey  = "key"
	clusterIDTag = "CLUSTER_ID"
)

var log = logger.Get()

// EC2Wrapper is used to wrap around EC2 service APIs to obtain ClusterID from
// the ec2 instance tags
type EC2Wrapper struct {
	ec2ServiceClient         ec2iface.EC2API
	instanceIdentityDocument ec2metadata.EC2InstanceIdentityDocument
}

// NewMetricsClient returns an instance of the EC2 wrapper
func NewMetricsClient() (*EC2Wrapper, error) {
	sess := nholuongutsession.New()
	ec2MetadataClient := ec2metadatawrapper.New(sess)

	instanceIdentityDocument, err := ec2MetadataClient.GetInstanceIdentityDocument()
	if err != nil {
		return &EC2Wrapper{}, err
	}

	nholuongutCfg := nholuongut.NewConfig().WithRegion(instanceIdentityDocument.Region)
	sess = sess.Copy(nholuongutCfg)
	ec2ServiceClient := ec2.New(sess)

	return &EC2Wrapper{
		ec2ServiceClient:         ec2ServiceClient,
		instanceIdentityDocument: instanceIdentityDocument,
	}, nil
}

// GetClusterTag is used to retrieve a tag from the ec2 instance
func (e *EC2Wrapper) GetClusterTag(tagKey string) (string, error) {
	input := ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{
			{
				Name: nholuongut.String(resourceID),
				Values: []*string{
					nholuongut.String(e.instanceIdentityDocument.InstanceID),
				},
			}, {
				Name: nholuongut.String(resourceKey),
				Values: []*string{
					nholuongut.String(tagKey),
				},
			},
		},
	}

	log.Infof("Calling DescribeTags with key %s", tagKey)
	results, err := e.ec2ServiceClient.DescribeTags(&input)
	if err != nil {
		return "", errors.Wrap(err, "GetClusterTag: Unable to obtain EC2 instance tags")
	}

	if len(results.Tags) < 1 {
		return "", errors.Errorf("GetClusterTag: No tag matching key: %s", tagKey)
	}

	return nholuongut.StringValue(results.Tags[0].Value), nil
}
