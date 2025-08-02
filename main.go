package main

import (
	"flag"
	"fmt"
	"srujanpakanati.com/go-nodeadm-updater/internal"
)

 func main() {
  // Define command-line flags
  clusterName := flag.String("cluster-name", "", "The name of the EKS cluster")
  region := flag.String("region", "", "The AWS region") // Add region flag

  // Parse command-line arguments
  flag.Parse()

  if *clusterName == "" {
  	fmt.Printf("cluster-name is required") // Exit if cluster name is missing
  }

  if *region == "" {
  	fmt.Printf("region is required") //exit if region is missing
  }
  // Call the describeCluster function from the internal package
  describeOutput := internal.DescribeCluster(*clusterName, *region)
  fmt.Printf("Completed describing cluster\n")
  error := internal.UpdateNodeadm(describeOutput)
  if error != nil {
  	fmt.Printf("Error updating nodeadm: %v", error)
  }
  fmt.Printf("nodeadm.yaml updated successfully with cluster details\n")

 }
