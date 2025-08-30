# Go EKS Nodeadm Updater

A command-line tool to update a `nodeadm.yaml` file with cluster details from an existing Amazon EKS cluster.

This tool simplifies the process of configuring `nodeadm` for a new EKS cluster by automatically fetching the required cluster information from AWS and updating the local `nodeadm.yaml` configuration file.

## Prerequisites

Before using this tool, ensure you have the following installed and configured:

*   **Go**: Version 1.22.0 or later.
*   **AWS Credentials**: Your AWS credentials must be configured in your environment. The tool uses the default credential provider chain, so you can configure them in a `~/.aws/credentials` file or via environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`). The credentials must have the necessary IAM permissions to describe an EKS cluster (e.g., `eks:DescribeCluster`).

## Installation

To build the tool from the source, clone the repository and run the following command in the project's root directory:

```sh
go build
```

This will create an executable file named `go-nodeadm-updater` (or `go-nodeadm-updater.exe` on Windows) in the current directory.

## Usage

To run the tool, use the following command:

```sh
./go-nodeadm-updater -cluster-name <your-cluster-name> -region <your-aws-region>
```

### Flags

*   `-cluster-name` (required): The name of your Amazon EKS cluster.
*   `-region` (required): The AWS region where your EKS cluster is located.

### Example

```sh
./go-nodeadm-updater -cluster-name my-eks-cluster -region us-west-2
```

Upon successful execution, the tool will update the `nodeadm.yaml` file in the same directory with the details of the `my-eks-cluster`.

## Configuration

The tool is designed to update a `nodeadm.yaml` file, which is a configuration file for `nodeadm`, a tool for bootstrapping EKS nodes.

The following fields in `nodeadm.yaml` are automatically updated by this tool:

*   `spec.cluster.name`
*   `spec.cluster.apiServerEndpoint`
*   `spec.cluster.certificateAuthority`

All other values in the file will be preserved.

### Example `nodeadm.yaml`

Below is an example of a `nodeadm.yaml` file. The `name`, `apiServerEndpoint`, and `certificateAuthority` fields under `spec.cluster` will be populated by the tool.

```yaml
apiVersion: node.eks.aws/v1alpha1
kind: NodeConfig
spec:
    cluster:
        apiServerEndpoint: # <-- This will be updated
        certificateAuthority: # <-- This will be updated
        cidr: 10.100.0.0/16
        name: # <-- This will be updated
    kubelet:
        config:
            kubeReserved:
                cpu: 100m
                ephemeral-storage: 1Gi
                memory: 200Mi
            systemReserved:
                cpu: 100m
                memory: 200Mi
```
