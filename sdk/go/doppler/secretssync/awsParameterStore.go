// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secretssync

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-doppler/sdk/go/doppler/internal"
)

// Manage an AWS Parameter Store Doppler sync.
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

	// The name of the Doppler config
	Config pulumi.StringOutput `pulumi:"config"`
	// The slug of the integration to use for this sync
	Integration pulumi.StringOutput `pulumi:"integration"`
	// The path to the parameters in AWS
	Path pulumi.StringOutput `pulumi:"path"`
	// The name of the Doppler project
	Project pulumi.StringOutput `pulumi:"project"`
	// The AWS region
	Region pulumi.StringOutput `pulumi:"region"`
	// Whether or not the parameters are stored as a secure string
	SecureString pulumi.BoolPtrOutput `pulumi:"secureString"`
	// AWS tags to attach to the parameters
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAwsParameterStore registers a new resource with the given unique name, arguments, and options.
func NewAwsParameterStore(ctx *pulumi.Context,
	name string, args *AwsParameterStoreArgs, opts ...pulumi.ResourceOption) (*AwsParameterStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Integration == nil {
		return nil, errors.New("invalid value for required argument 'Integration'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AwsParameterStore
	err := ctx.RegisterResource("doppler:secretsSync/awsParameterStore:AwsParameterStore", name, args, &resource, opts...)
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
	err := ctx.ReadResource("doppler:secretsSync/awsParameterStore:AwsParameterStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsParameterStore resources.
type awsParameterStoreState struct {
	// The name of the Doppler config
	Config *string `pulumi:"config"`
	// The slug of the integration to use for this sync
	Integration *string `pulumi:"integration"`
	// The path to the parameters in AWS
	Path *string `pulumi:"path"`
	// The name of the Doppler project
	Project *string `pulumi:"project"`
	// The AWS region
	Region *string `pulumi:"region"`
	// Whether or not the parameters are stored as a secure string
	SecureString *bool `pulumi:"secureString"`
	// AWS tags to attach to the parameters
	Tags map[string]string `pulumi:"tags"`
}

type AwsParameterStoreState struct {
	// The name of the Doppler config
	Config pulumi.StringPtrInput
	// The slug of the integration to use for this sync
	Integration pulumi.StringPtrInput
	// The path to the parameters in AWS
	Path pulumi.StringPtrInput
	// The name of the Doppler project
	Project pulumi.StringPtrInput
	// The AWS region
	Region pulumi.StringPtrInput
	// Whether or not the parameters are stored as a secure string
	SecureString pulumi.BoolPtrInput
	// AWS tags to attach to the parameters
	Tags pulumi.StringMapInput
}

func (AwsParameterStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsParameterStoreState)(nil)).Elem()
}

type awsParameterStoreArgs struct {
	// The name of the Doppler config
	Config string `pulumi:"config"`
	// The slug of the integration to use for this sync
	Integration string `pulumi:"integration"`
	// The path to the parameters in AWS
	Path string `pulumi:"path"`
	// The name of the Doppler project
	Project string `pulumi:"project"`
	// The AWS region
	Region string `pulumi:"region"`
	// Whether or not the parameters are stored as a secure string
	SecureString *bool `pulumi:"secureString"`
	// AWS tags to attach to the parameters
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AwsParameterStore resource.
type AwsParameterStoreArgs struct {
	// The name of the Doppler config
	Config pulumi.StringInput
	// The slug of the integration to use for this sync
	Integration pulumi.StringInput
	// The path to the parameters in AWS
	Path pulumi.StringInput
	// The name of the Doppler project
	Project pulumi.StringInput
	// The AWS region
	Region pulumi.StringInput
	// Whether or not the parameters are stored as a secure string
	SecureString pulumi.BoolPtrInput
	// AWS tags to attach to the parameters
	Tags pulumi.StringMapInput
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

// The name of the Doppler config
func (o AwsParameterStoreOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsParameterStore) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

// The slug of the integration to use for this sync
func (o AwsParameterStoreOutput) Integration() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsParameterStore) pulumi.StringOutput { return v.Integration }).(pulumi.StringOutput)
}

// The path to the parameters in AWS
func (o AwsParameterStoreOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsParameterStore) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The name of the Doppler project
func (o AwsParameterStoreOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsParameterStore) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The AWS region
func (o AwsParameterStoreOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsParameterStore) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Whether or not the parameters are stored as a secure string
func (o AwsParameterStoreOutput) SecureString() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AwsParameterStore) pulumi.BoolPtrOutput { return v.SecureString }).(pulumi.BoolPtrOutput)
}

// AWS tags to attach to the parameters
func (o AwsParameterStoreOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AwsParameterStore) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
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