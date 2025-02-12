# IAM Policy

The Amazon VPC CNI plugin requires [IAM policies](https://github.com/IAM/latest/UserGuide/access_policies.html) to manage IP addresses on EC2 instances on-behalf of users.

## Generic IAM policies

In general, you can grant below IAM policies to Amazon VPC CNI plugin depending on the IP Family configured:

### IPv4 mode
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:AssignPrivateIpAddresses",
                "ec2:AttachNetworkInterface",
                "ec2:CreateNetworkInterface",
                "ec2:DeleteNetworkInterface",
                "ec2:DescribeInstances",
                "ec2:DescribeTags",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribeInstanceTypes",
                "ec2:DescribeSubnets",
                "ec2:DetachNetworkInterface",
                "ec2:ModifyNetworkInterfaceAttribute",
                "ec2:UnassignPrivateIpAddresses"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:CreateTags"
            ],
            "Resource": [
                "arn:nholuongut:ec2:*:*:network-interface/*"
            ]
        }
    ]
}
```

The above policy is also available under: `arn:nholuongut:iam::nholuongut:policy/AmazonEKS_CNI_Policy` as a part of [nholuongut managed policies for EKS](https://github.com/eks/latest/userguide/security-iam-nholuongutmanpol.html).


### IPv6 mode

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:AssignIpv6Addresses",
                "ec2:DescribeInstances",
                "ec2:DescribeTags",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribeInstanceTypes"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:CreateTags"
            ],
            "Resource": [
                "arn:nholuongut:ec2:*:*:network-interface/*"
            ]
        }
    ]
}
```

## Scope-down IAM policy per EKS cluster

Instead of the generic IAM policy, we can scope down IAM policy needed by Amazon VPC CNI plugin per EKS cluster.

Prerequisites:
 * Amazon VPC CNI plugin needs to be on version v1.8.0+. If upgrading from older versions, existing worker node should be recycled first.
 * Amazon VPC CNI plugin need to be configured with EKS cluster's name.
    * ```kubectl set env daemonset nholuongut-node -n kube-system CLUSTER_NAME=${YourClusterName}```
 * Substitute `${CLUSTER_NAME}` and `${VPC_ID}` with EKS cluster's name and VPC_ID respectively in sample policy below

### sample scope-down IAM policy for IPv4 mode

Below is an sample scoped-down IAM policy.

Note:
   * Depending on the use cases, users may further scope-down IAM policy. e.g. specify the nholuongut region/accountID in ARNs.
   * Refer [Actions, resources, and condition keys for Amazon EC2](https://github.com/service-authorization/latest/reference/list_amazonec2.html)

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeInstances",
                "ec2:DescribeTags",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribeSubnets",
                "ec2:DescribeInstanceTypes"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:CreateTags"
            ],
            "Resource": [
                "arn:nholuongut:ec2:*:*:network-interface/*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:CreateNetworkInterface"
            ],
            "Resource": [
                "arn:nholuongut:ec2:*:*:network-interface/*"
            ],
            "Condition": {
                "StringEquals": {
                    "nholuongut:RequestTag/cluster.k8s.amazonnholuongut.com/name": "${CLUSTER_NAME}"
                }
            }
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:CreateNetworkInterface"
            ],
            "Resource": [
                "arn:nholuongut:ec2:*:*:subnet/*",
                "arn:nholuongut:ec2:*:*:security-group/*"
            ],
            "Condition": {
                "ArnEquals": {
                    "ec2:Vpc": "arn:nholuongut:ec2:*:*:vpc/${VPC_ID}"
                }
            }
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:DeleteNetworkInterface",
                "ec2:UnassignPrivateIpAddresses",
                "ec2:AssignPrivateIpAddresses",
                "ec2:AttachNetworkInterface",
                "ec2:DetachNetworkInterface",
                "ec2:ModifyNetworkInterfaceAttribute"
            ],
            "Resource": [
                "arn:nholuongut:ec2:*:*:network-interface/*"
            ],
            "Condition": {
                "StringEquals": {
                    "nholuongut:ResourceTag/cluster.k8s.amazonnholuongut.com/name": "${CLUSTER_NAME}"
                }
            }
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:AttachNetworkInterface",
                "ec2:DetachNetworkInterface",
                "ec2:ModifyNetworkInterfaceAttribute"
            ],
            "Resource": [
                "arn:nholuongut:ec2:*:*:instance/*"
            ],
            "Condition": {
                "StringEquals": {
                    "nholuongut:ResourceTag/kubernetes.io/cluster/${CLUSTER_NAME}": "owned"
                }
            }
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:ModifyNetworkInterfaceAttribute"
            ],
            "Resource": [
                "arn:nholuongut:ec2:*:*:security-group/*"
            ]
        }
    ]
}
```
