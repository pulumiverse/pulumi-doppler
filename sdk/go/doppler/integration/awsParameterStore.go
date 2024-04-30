// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package integration

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-doppler/sdk/go/doppler/internal"
)

// Manage an AWS Parameter Store Doppler integration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-doppler/sdk/go/doppler/integration"
//	"github.com/pulumiverse/pulumi-doppler/sdk/go/doppler/secretsSync"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"Version": "2012-10-17",
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Effect": "Allow",
//						"Action": "sts:AssumeRole",
//						"Principal": map[string]interface{}{
//							"AWS": "arn:aws:iam::299900769157:user/doppler-integration-operator",
//						},
//						"Condition": map[string]interface{}{
//							"StringEquals": map[string]interface{}{
//								"sts:ExternalId": "<YOUR_WORKPLACE_SLUG>",
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"Version": "2012-10-17",
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Action": []string{
//							"ssm:PutParameter",
//							"ssm:LabelParameterVersion",
//							"ssm:DeleteParameter",
//							"ssm:RemoveTagsFromResource",
//							"ssm:GetParameterHistory",
//							"ssm:AddTagsToResource",
//							"ssm:GetParametersByPath",
//							"ssm:GetParameters",
//							"ssm:GetParameter",
//							"ssm:DeleteParameters",
//						},
//						"Effect":   "Allow",
//						"Resource": "*",
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			dopplerParameterStore, err := iam.NewRole(ctx, "dopplerParameterStore", &iam.RoleArgs{
//				AssumeRolePolicy: pulumi.String(json0),
//				InlinePolicies: iam.RoleInlinePolicyArray{
//					&iam.RoleInlinePolicyArgs{
//						Policy: pulumi.String(json1),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			prod, err := integration.NewAwsParameterStore(ctx, "prod", &integration.AwsParameterStoreArgs{
//				Name:          pulumi.String("Production"),
//				AssumeRoleArn: dopplerParameterStore.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = secretsSync.NewAwsParameterStore(ctx, "backendProd", &secretsSync.AwsParameterStoreArgs{
//				Integration:  prod.ID(),
//				Project:      pulumi.String("backend"),
//				Config:       pulumi.String("prd"),
//				Region:       pulumi.String("us-east-1"),
//				Path:         pulumi.String("/backend/"),
//				SecureString: pulumi.Bool(true),
//				Tags: pulumi.StringMap{
//					"myTag": pulumi.String("enabled"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AwsParameterStore struct {
	pulumi.CustomResourceState

	// The ARN of the AWS role for Doppler to assume
	AssumeRoleArn pulumi.StringOutput `pulumi:"assumeRoleArn"`
	// The name of the integration
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAwsParameterStore registers a new resource with the given unique name, arguments, and options.
func NewAwsParameterStore(ctx *pulumi.Context,
	name string, args *AwsParameterStoreArgs, opts ...pulumi.ResourceOption) (*AwsParameterStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssumeRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'AssumeRoleArn'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AwsParameterStore
	err := ctx.RegisterResource("doppler:integration/awsParameterStore:AwsParameterStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsParameterStore gets an existing AwsParameterStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsParameterStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsParameterStoreState, opts ...pulumi.ResourceOption) (*AwsParameterStore, error) {
	var resource AwsParameterStore
	err := ctx.ReadResource("doppler:integration/awsParameterStore:AwsParameterStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsParameterStore resources.
type awsParameterStoreState struct {
	// The ARN of the AWS role for Doppler to assume
	AssumeRoleArn *string `pulumi:"assumeRoleArn"`
	// The name of the integration
	Name *string `pulumi:"name"`
}

type AwsParameterStoreState struct {
	// The ARN of the AWS role for Doppler to assume
	AssumeRoleArn pulumi.StringPtrInput
	// The name of the integration
	Name pulumi.StringPtrInput
}

func (AwsParameterStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsParameterStoreState)(nil)).Elem()
}

type awsParameterStoreArgs struct {
	// The ARN of the AWS role for Doppler to assume
	AssumeRoleArn string `pulumi:"assumeRoleArn"`
	// The name of the integration
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a AwsParameterStore resource.
type AwsParameterStoreArgs struct {
	// The ARN of the AWS role for Doppler to assume
	AssumeRoleArn pulumi.StringInput
	// The name of the integration
	Name pulumi.StringInput
}

func (AwsParameterStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsParameterStoreArgs)(nil)).Elem()
}

type AwsParameterStoreInput interface {
	pulumi.Input

	ToAwsParameterStoreOutput() AwsParameterStoreOutput
	ToAwsParameterStoreOutputWithContext(ctx context.Context) AwsParameterStoreOutput
}

func (*AwsParameterStore) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsParameterStore)(nil)).Elem()
}

func (i *AwsParameterStore) ToAwsParameterStoreOutput() AwsParameterStoreOutput {
	return i.ToAwsParameterStoreOutputWithContext(context.Background())
}

func (i *AwsParameterStore) ToAwsParameterStoreOutputWithContext(ctx context.Context) AwsParameterStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsParameterStoreOutput)
}

// AwsParameterStoreArrayInput is an input type that accepts AwsParameterStoreArray and AwsParameterStoreArrayOutput values.
// You can construct a concrete instance of `AwsParameterStoreArrayInput` via:
//
//	AwsParameterStoreArray{ AwsParameterStoreArgs{...} }
type AwsParameterStoreArrayInput interface {
	pulumi.Input

	ToAwsParameterStoreArrayOutput() AwsParameterStoreArrayOutput
	ToAwsParameterStoreArrayOutputWithContext(context.Context) AwsParameterStoreArrayOutput
}

type AwsParameterStoreArray []AwsParameterStoreInput

func (AwsParameterStoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsParameterStore)(nil)).Elem()
}

func (i AwsParameterStoreArray) ToAwsParameterStoreArrayOutput() AwsParameterStoreArrayOutput {
	return i.ToAwsParameterStoreArrayOutputWithContext(context.Background())
}

func (i AwsParameterStoreArray) ToAwsParameterStoreArrayOutputWithContext(ctx context.Context) AwsParameterStoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsParameterStoreArrayOutput)
}

// AwsParameterStoreMapInput is an input type that accepts AwsParameterStoreMap and AwsParameterStoreMapOutput values.
// You can construct a concrete instance of `AwsParameterStoreMapInput` via:
//
//	AwsParameterStoreMap{ "key": AwsParameterStoreArgs{...} }
type AwsParameterStoreMapInput interface {
	pulumi.Input

	ToAwsParameterStoreMapOutput() AwsParameterStoreMapOutput
	ToAwsParameterStoreMapOutputWithContext(context.Context) AwsParameterStoreMapOutput
}

type AwsParameterStoreMap map[string]AwsParameterStoreInput

func (AwsParameterStoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsParameterStore)(nil)).Elem()
}

func (i AwsParameterStoreMap) ToAwsParameterStoreMapOutput() AwsParameterStoreMapOutput {
	return i.ToAwsParameterStoreMapOutputWithContext(context.Background())
}

func (i AwsParameterStoreMap) ToAwsParameterStoreMapOutputWithContext(ctx context.Context) AwsParameterStoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsParameterStoreMapOutput)
}

type AwsParameterStoreOutput struct{ *pulumi.OutputState }

func (AwsParameterStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsParameterStore)(nil)).Elem()
}

func (o AwsParameterStoreOutput) ToAwsParameterStoreOutput() AwsParameterStoreOutput {
	return o
}

func (o AwsParameterStoreOutput) ToAwsParameterStoreOutputWithContext(ctx context.Context) AwsParameterStoreOutput {
	return o
}

// The ARN of the AWS role for Doppler to assume
func (o AwsParameterStoreOutput) AssumeRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsParameterStore) pulumi.StringOutput { return v.AssumeRoleArn }).(pulumi.StringOutput)
}

// The name of the integration
func (o AwsParameterStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsParameterStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type AwsParameterStoreArrayOutput struct{ *pulumi.OutputState }

func (AwsParameterStoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsParameterStore)(nil)).Elem()
}

func (o AwsParameterStoreArrayOutput) ToAwsParameterStoreArrayOutput() AwsParameterStoreArrayOutput {
	return o
}

func (o AwsParameterStoreArrayOutput) ToAwsParameterStoreArrayOutputWithContext(ctx context.Context) AwsParameterStoreArrayOutput {
	return o
}

func (o AwsParameterStoreArrayOutput) Index(i pulumi.IntInput) AwsParameterStoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AwsParameterStore {
		return vs[0].([]*AwsParameterStore)[vs[1].(int)]
	}).(AwsParameterStoreOutput)
}

type AwsParameterStoreMapOutput struct{ *pulumi.OutputState }

func (AwsParameterStoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsParameterStore)(nil)).Elem()
}

func (o AwsParameterStoreMapOutput) ToAwsParameterStoreMapOutput() AwsParameterStoreMapOutput {
	return o
}

func (o AwsParameterStoreMapOutput) ToAwsParameterStoreMapOutputWithContext(ctx context.Context) AwsParameterStoreMapOutput {
	return o
}

func (o AwsParameterStoreMapOutput) MapIndex(k pulumi.StringInput) AwsParameterStoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AwsParameterStore {
		return vs[0].(map[string]*AwsParameterStore)[vs[1].(string)]
	}).(AwsParameterStoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsParameterStoreInput)(nil)).Elem(), &AwsParameterStore{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsParameterStoreArrayInput)(nil)).Elem(), AwsParameterStoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsParameterStoreMapInput)(nil)).Elem(), AwsParameterStoreMap{})
	pulumi.RegisterOutputType(AwsParameterStoreOutput{})
	pulumi.RegisterOutputType(AwsParameterStoreArrayOutput{})
	pulumi.RegisterOutputType(AwsParameterStoreMapOutput{})
}
