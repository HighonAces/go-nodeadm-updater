package internal

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type DescribeClusterOutput struct {
	clusterName string
	apiSeerverEndpoint string
	certificateAuthorityData string
}

func DescribeCluster(clusterName, region string) (DescribeClusterOutput) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Printf("failed to load configuration: %v", err)
	}

	eksClient := eks.NewFromConfig(cfg)

	input := &eks.DescribeClusterInput{
		Name: &clusterName,
	}

	output, err := eksClient.DescribeCluster(context.TODO(), input)
	if err != nil {
		fmt.Printf("failed to describe cluster: %v, check cluster name and region", err)
	}
	describeOutput := DescribeClusterOutput{
		clusterName:             *output.Cluster.Name,
		apiSeerverEndpoint:      *output.Cluster.Endpoint,
		certificateAuthorityData: *output.Cluster.CertificateAuthority.Data,
	}
	return describeOutput
}