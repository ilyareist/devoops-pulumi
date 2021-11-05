package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		sg, err := ec2.NewSecurityGroup(ctx, "allow http", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow all incoming http traffic"),
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Description: pulumi.String("Allow all incoming http traffic"),
					FromPort:    pulumi.Int(80),
					ToPort:      pulumi.Int(80),
					Protocol:    pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					Description: pulumi.String("Allow all outgoing traffic"),
					FromPort:    pulumi.Int(-1),
					ToPort:      pulumi.Int(-1),
					Protocol:    pulumi.String("All"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
		})
		if err != nil {
			return err
		}

		initScriptFile, err := os.ReadFile("./provision.sh")
		if err != nil {
			return err
		}

		// Create an AWS resource (EC2 instance)
		instance, err := ec2.NewInstance(ctx,
			"my-instance",
			&ec2.InstanceArgs{
				Ami:                 pulumi.String("ami-0ed961fa828560210"),
				InstanceType:        pulumi.String("t3.medium"),
				UserData:            pulumi.String(initScriptFile),
				VpcSecurityGroupIds: pulumi.StringArray{sg.ID()},
			})
		if err != nil {
			return err
		}

		ctx.Export("instance IP", instance.PublicIp)
		return nil
	})
}
