package internal

import (
	"fmt"
	"os"
	"log"
	yaml "gopkg.in/yaml.v3"
)

type NodeadmConfig struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Spec       NodeadmConfigSpec `yaml:"spec"`
}

type NodeadmConfigSpec struct {
	Cluster NodeadmClusterConfig `yaml:"cluster"`
}

type NodeadmClusterConfig struct {
	Name               string `yaml:"name"`
	APIServerEndpoint  string `yaml:"apiServerEndpoint"`
	CertificateAuthority string `yaml:"certificateAuthority"`
	CIDR               string `yaml:"cidr"`
}

func UpdateNodeadm(DescribeClusterOutput DescribeClusterOutput) error {
	// Read the nodeadm.yaml file
	yamlFile, err := os.ReadFile("nodeadm.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}

	// Parse into a generic map to preserve all existing configuration
	var rawConfig map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &rawConfig)
	if err != nil {
		fmt.Printf("Unmarshal to map: %v", err)
		return err
	}

	// Also parse into our struct to work with typed data
	nodeadmConfig := NodeadmConfig{}
	err = yaml.Unmarshal(yamlFile, &nodeadmConfig)
	if err != nil {
		fmt.Printf("Unmarshal to struct: %v", err)
		return err
	}

	// Update only the cluster-specific fields in the struct
	nodeadmConfig.Spec.Cluster.Name = DescribeClusterOutput.clusterName
	nodeadmConfig.Spec.Cluster.APIServerEndpoint = DescribeClusterOutput.apiSeerverEndpoint
	nodeadmConfig.Spec.Cluster.CertificateAuthority = DescribeClusterOutput.certificateAuthorityData

	// Update the cluster configuration in the raw map while preserving everything else
	if spec, ok := rawConfig["spec"].(map[string]interface{}); ok {
		// Create or update the cluster section within spec
		clusterConfig := map[string]interface{}{
			"name":                nodeadmConfig.Spec.Cluster.Name,
			"apiServerEndpoint":   nodeadmConfig.Spec.Cluster.APIServerEndpoint,
			"certificateAuthority": nodeadmConfig.Spec.Cluster.CertificateAuthority,
		}
		
		// Only add CIDR if it's not empty (preserve existing CIDR if new one is empty)
		if nodeadmConfig.Spec.Cluster.CIDR != "" {
			clusterConfig["cidr"] = nodeadmConfig.Spec.Cluster.CIDR
		} else if cluster, exists := spec["cluster"].(map[string]interface{}); exists {
			if existingCIDR, hasCIDR := cluster["cidr"]; hasCIDR {
				clusterConfig["cidr"] = existingCIDR
			}
		}
		
		spec["cluster"] = clusterConfig
	} else {
		// If spec doesn't exist, create it
		rawConfig["spec"] = map[string]interface{}{
			"cluster": map[string]interface{}{
				"name":                nodeadmConfig.Spec.Cluster.Name,
				"apiServerEndpoint":   nodeadmConfig.Spec.Cluster.APIServerEndpoint,
				"certificateAuthority": nodeadmConfig.Spec.Cluster.CertificateAuthority,
				"cidr":                nodeadmConfig.Spec.Cluster.CIDR,
			},
		}
	}

	// Marshal the updated configuration back to YAML
	updatedYaml, err := yaml.Marshal(rawConfig)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	// Write the updated YAML back to the nodeadm.yaml file
	err = os.WriteFile("nodeadm.yaml", updatedYaml, 0644)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	fmt.Println("Successfully updated nodeadm.yaml")
	return nil
}