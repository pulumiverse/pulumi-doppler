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

// Manage a CircleCI Doppler sync.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-doppler/sdk/go/doppler/integration"
//	"github.com/pulumiverse/pulumi-doppler/sdk/go/doppler/secretsSync"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			prod, err := integration.NewCircleci(ctx, "prod", &integration.CircleciArgs{
//				Name:     pulumi.String("Production"),
//				ApiToken: pulumi.String("my_api_token"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = secretsSync.NewCircleci(ctx, "backend_prod", &secretsSync.CircleciArgs{
//				Integration:      prod.ID(),
//				Project:          pulumi.String("backend"),
//				Config:           pulumi.String("prd"),
//				ResourceType:     pulumi.String("project"),
//				ResourceId:       pulumi.String("github/myorg/myproject"),
//				OrganizationSlug: pulumi.String("myorg"),
//				DeleteBehavior:   pulumi.String("leave_in_target"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Circleci struct {
	pulumi.CustomResourceState

	// The name of the Doppler config
	Config pulumi.StringOutput `pulumi:"config"`
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior pulumi.StringPtrOutput `pulumi:"deleteBehavior"`
	// The slug of the integration to use for this sync
	Integration pulumi.StringOutput `pulumi:"integration"`
	// The organization slug where the resource is located
	OrganizationSlug pulumi.StringOutput `pulumi:"organizationSlug"`
	// The name of the Doppler project
	Project pulumi.StringOutput `pulumi:"project"`
	// The resource ID (either project or context) to sync to
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Either "project" or "context", based on the resource type to sync to
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
}

// NewCircleci registers a new resource with the given unique name, arguments, and options.
func NewCircleci(ctx *pulumi.Context,
	name string, args *CircleciArgs, opts ...pulumi.ResourceOption) (*Circleci, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Integration == nil {
		return nil, errors.New("invalid value for required argument 'Integration'")
	}
	if args.OrganizationSlug == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationSlug'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Circleci
	err := ctx.RegisterResource("doppler:secretsSync/circleci:Circleci", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCircleci gets an existing Circleci resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCircleci(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CircleciState, opts ...pulumi.ResourceOption) (*Circleci, error) {
	var resource Circleci
	err := ctx.ReadResource("doppler:secretsSync/circleci:Circleci", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Circleci resources.
type circleciState struct {
	// The name of the Doppler config
	Config *string `pulumi:"config"`
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior *string `pulumi:"deleteBehavior"`
	// The slug of the integration to use for this sync
	Integration *string `pulumi:"integration"`
	// The organization slug where the resource is located
	OrganizationSlug *string `pulumi:"organizationSlug"`
	// The name of the Doppler project
	Project *string `pulumi:"project"`
	// The resource ID (either project or context) to sync to
	ResourceId *string `pulumi:"resourceId"`
	// Either "project" or "context", based on the resource type to sync to
	ResourceType *string `pulumi:"resourceType"`
}

type CircleciState struct {
	// The name of the Doppler config
	Config pulumi.StringPtrInput
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior pulumi.StringPtrInput
	// The slug of the integration to use for this sync
	Integration pulumi.StringPtrInput
	// The organization slug where the resource is located
	OrganizationSlug pulumi.StringPtrInput
	// The name of the Doppler project
	Project pulumi.StringPtrInput
	// The resource ID (either project or context) to sync to
	ResourceId pulumi.StringPtrInput
	// Either "project" or "context", based on the resource type to sync to
	ResourceType pulumi.StringPtrInput
}

func (CircleciState) ElementType() reflect.Type {
	return reflect.TypeOf((*circleciState)(nil)).Elem()
}

type circleciArgs struct {
	// The name of the Doppler config
	Config string `pulumi:"config"`
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior *string `pulumi:"deleteBehavior"`
	// The slug of the integration to use for this sync
	Integration string `pulumi:"integration"`
	// The organization slug where the resource is located
	OrganizationSlug string `pulumi:"organizationSlug"`
	// The name of the Doppler project
	Project string `pulumi:"project"`
	// The resource ID (either project or context) to sync to
	ResourceId string `pulumi:"resourceId"`
	// Either "project" or "context", based on the resource type to sync to
	ResourceType string `pulumi:"resourceType"`
}

// The set of arguments for constructing a Circleci resource.
type CircleciArgs struct {
	// The name of the Doppler config
	Config pulumi.StringInput
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior pulumi.StringPtrInput
	// The slug of the integration to use for this sync
	Integration pulumi.StringInput
	// The organization slug where the resource is located
	OrganizationSlug pulumi.StringInput
	// The name of the Doppler project
	Project pulumi.StringInput
	// The resource ID (either project or context) to sync to
	ResourceId pulumi.StringInput
	// Either "project" or "context", based on the resource type to sync to
	ResourceType pulumi.StringInput
}

func (CircleciArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*circleciArgs)(nil)).Elem()
}

type CircleciInput interface {
	pulumi.Input

	ToCircleciOutput() CircleciOutput
	ToCircleciOutputWithContext(ctx context.Context) CircleciOutput
}

func (*Circleci) ElementType() reflect.Type {
	return reflect.TypeOf((**Circleci)(nil)).Elem()
}

func (i *Circleci) ToCircleciOutput() CircleciOutput {
	return i.ToCircleciOutputWithContext(context.Background())
}

func (i *Circleci) ToCircleciOutputWithContext(ctx context.Context) CircleciOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CircleciOutput)
}

// CircleciArrayInput is an input type that accepts CircleciArray and CircleciArrayOutput values.
// You can construct a concrete instance of `CircleciArrayInput` via:
//
//	CircleciArray{ CircleciArgs{...} }
type CircleciArrayInput interface {
	pulumi.Input

	ToCircleciArrayOutput() CircleciArrayOutput
	ToCircleciArrayOutputWithContext(context.Context) CircleciArrayOutput
}

type CircleciArray []CircleciInput

func (CircleciArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Circleci)(nil)).Elem()
}

func (i CircleciArray) ToCircleciArrayOutput() CircleciArrayOutput {
	return i.ToCircleciArrayOutputWithContext(context.Background())
}

func (i CircleciArray) ToCircleciArrayOutputWithContext(ctx context.Context) CircleciArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CircleciArrayOutput)
}

// CircleciMapInput is an input type that accepts CircleciMap and CircleciMapOutput values.
// You can construct a concrete instance of `CircleciMapInput` via:
//
//	CircleciMap{ "key": CircleciArgs{...} }
type CircleciMapInput interface {
	pulumi.Input

	ToCircleciMapOutput() CircleciMapOutput
	ToCircleciMapOutputWithContext(context.Context) CircleciMapOutput
}

type CircleciMap map[string]CircleciInput

func (CircleciMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Circleci)(nil)).Elem()
}

func (i CircleciMap) ToCircleciMapOutput() CircleciMapOutput {
	return i.ToCircleciMapOutputWithContext(context.Background())
}

func (i CircleciMap) ToCircleciMapOutputWithContext(ctx context.Context) CircleciMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CircleciMapOutput)
}

type CircleciOutput struct{ *pulumi.OutputState }

func (CircleciOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Circleci)(nil)).Elem()
}

func (o CircleciOutput) ToCircleciOutput() CircleciOutput {
	return o
}

func (o CircleciOutput) ToCircleciOutputWithContext(ctx context.Context) CircleciOutput {
	return o
}

// The name of the Doppler config
func (o CircleciOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *Circleci) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
func (o CircleciOutput) DeleteBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Circleci) pulumi.StringPtrOutput { return v.DeleteBehavior }).(pulumi.StringPtrOutput)
}

// The slug of the integration to use for this sync
func (o CircleciOutput) Integration() pulumi.StringOutput {
	return o.ApplyT(func(v *Circleci) pulumi.StringOutput { return v.Integration }).(pulumi.StringOutput)
}

// The organization slug where the resource is located
func (o CircleciOutput) OrganizationSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *Circleci) pulumi.StringOutput { return v.OrganizationSlug }).(pulumi.StringOutput)
}

// The name of the Doppler project
func (o CircleciOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Circleci) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The resource ID (either project or context) to sync to
func (o CircleciOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Circleci) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Either "project" or "context", based on the resource type to sync to
func (o CircleciOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Circleci) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

type CircleciArrayOutput struct{ *pulumi.OutputState }

func (CircleciArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Circleci)(nil)).Elem()
}

func (o CircleciArrayOutput) ToCircleciArrayOutput() CircleciArrayOutput {
	return o
}

func (o CircleciArrayOutput) ToCircleciArrayOutputWithContext(ctx context.Context) CircleciArrayOutput {
	return o
}

func (o CircleciArrayOutput) Index(i pulumi.IntInput) CircleciOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Circleci {
		return vs[0].([]*Circleci)[vs[1].(int)]
	}).(CircleciOutput)
}

type CircleciMapOutput struct{ *pulumi.OutputState }

func (CircleciMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Circleci)(nil)).Elem()
}

func (o CircleciMapOutput) ToCircleciMapOutput() CircleciMapOutput {
	return o
}

func (o CircleciMapOutput) ToCircleciMapOutputWithContext(ctx context.Context) CircleciMapOutput {
	return o
}

func (o CircleciMapOutput) MapIndex(k pulumi.StringInput) CircleciOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Circleci {
		return vs[0].(map[string]*Circleci)[vs[1].(string)]
	}).(CircleciOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CircleciInput)(nil)).Elem(), &Circleci{})
	pulumi.RegisterInputType(reflect.TypeOf((*CircleciArrayInput)(nil)).Elem(), CircleciArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CircleciMapInput)(nil)).Elem(), CircleciMap{})
	pulumi.RegisterOutputType(CircleciOutput{})
	pulumi.RegisterOutputType(CircleciArrayOutput{})
	pulumi.RegisterOutputType(CircleciMapOutput{})
}