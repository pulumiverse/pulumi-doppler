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

// Manage a Fly.io Doppler sync.
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
//			prod, err := integration.NewFlyio(ctx, "prod", &integration.FlyioArgs{
//				Name:   pulumi.String("TF Fly.io"),
//				ApiKey: pulumi.String("fo1_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = secretsSync.NewFlyio(ctx, "backend_prod", &secretsSync.FlyioArgs{
//				Integration:     prod.ID(),
//				Project:         pulumi.String("backend"),
//				Config:          pulumi.String("prd"),
//				AppId:           pulumi.String("my-app"),
//				RestartMachines: pulumi.Bool(true),
//				DeleteBehavior:  pulumi.String("leave_in_target"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Flyio struct {
	pulumi.CustomResourceState

	// The app ID
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The name of the Doppler config
	Config pulumi.StringOutput `pulumi:"config"`
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior pulumi.StringPtrOutput `pulumi:"deleteBehavior"`
	// The slug of the integration to use for this sync
	Integration pulumi.StringOutput `pulumi:"integration"`
	// The name of the Doppler project
	Project pulumi.StringOutput `pulumi:"project"`
	// Whether or not to restart the Fly.io machines when secrets are updated
	RestartMachines pulumi.BoolOutput `pulumi:"restartMachines"`
}

// NewFlyio registers a new resource with the given unique name, arguments, and options.
func NewFlyio(ctx *pulumi.Context,
	name string, args *FlyioArgs, opts ...pulumi.ResourceOption) (*Flyio, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Integration == nil {
		return nil, errors.New("invalid value for required argument 'Integration'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.RestartMachines == nil {
		return nil, errors.New("invalid value for required argument 'RestartMachines'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Flyio
	err := ctx.RegisterResource("doppler:secretsSync/flyio:Flyio", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlyio gets an existing Flyio resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlyio(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlyioState, opts ...pulumi.ResourceOption) (*Flyio, error) {
	var resource Flyio
	err := ctx.ReadResource("doppler:secretsSync/flyio:Flyio", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flyio resources.
type flyioState struct {
	// The app ID
	AppId *string `pulumi:"appId"`
	// The name of the Doppler config
	Config *string `pulumi:"config"`
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior *string `pulumi:"deleteBehavior"`
	// The slug of the integration to use for this sync
	Integration *string `pulumi:"integration"`
	// The name of the Doppler project
	Project *string `pulumi:"project"`
	// Whether or not to restart the Fly.io machines when secrets are updated
	RestartMachines *bool `pulumi:"restartMachines"`
}

type FlyioState struct {
	// The app ID
	AppId pulumi.StringPtrInput
	// The name of the Doppler config
	Config pulumi.StringPtrInput
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior pulumi.StringPtrInput
	// The slug of the integration to use for this sync
	Integration pulumi.StringPtrInput
	// The name of the Doppler project
	Project pulumi.StringPtrInput
	// Whether or not to restart the Fly.io machines when secrets are updated
	RestartMachines pulumi.BoolPtrInput
}

func (FlyioState) ElementType() reflect.Type {
	return reflect.TypeOf((*flyioState)(nil)).Elem()
}

type flyioArgs struct {
	// The app ID
	AppId string `pulumi:"appId"`
	// The name of the Doppler config
	Config string `pulumi:"config"`
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior *string `pulumi:"deleteBehavior"`
	// The slug of the integration to use for this sync
	Integration string `pulumi:"integration"`
	// The name of the Doppler project
	Project string `pulumi:"project"`
	// Whether or not to restart the Fly.io machines when secrets are updated
	RestartMachines bool `pulumi:"restartMachines"`
}

// The set of arguments for constructing a Flyio resource.
type FlyioArgs struct {
	// The app ID
	AppId pulumi.StringInput
	// The name of the Doppler config
	Config pulumi.StringInput
	// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
	DeleteBehavior pulumi.StringPtrInput
	// The slug of the integration to use for this sync
	Integration pulumi.StringInput
	// The name of the Doppler project
	Project pulumi.StringInput
	// Whether or not to restart the Fly.io machines when secrets are updated
	RestartMachines pulumi.BoolInput
}

func (FlyioArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flyioArgs)(nil)).Elem()
}

type FlyioInput interface {
	pulumi.Input

	ToFlyioOutput() FlyioOutput
	ToFlyioOutputWithContext(ctx context.Context) FlyioOutput
}

func (*Flyio) ElementType() reflect.Type {
	return reflect.TypeOf((**Flyio)(nil)).Elem()
}

func (i *Flyio) ToFlyioOutput() FlyioOutput {
	return i.ToFlyioOutputWithContext(context.Background())
}

func (i *Flyio) ToFlyioOutputWithContext(ctx context.Context) FlyioOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlyioOutput)
}

// FlyioArrayInput is an input type that accepts FlyioArray and FlyioArrayOutput values.
// You can construct a concrete instance of `FlyioArrayInput` via:
//
//	FlyioArray{ FlyioArgs{...} }
type FlyioArrayInput interface {
	pulumi.Input

	ToFlyioArrayOutput() FlyioArrayOutput
	ToFlyioArrayOutputWithContext(context.Context) FlyioArrayOutput
}

type FlyioArray []FlyioInput

func (FlyioArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Flyio)(nil)).Elem()
}

func (i FlyioArray) ToFlyioArrayOutput() FlyioArrayOutput {
	return i.ToFlyioArrayOutputWithContext(context.Background())
}

func (i FlyioArray) ToFlyioArrayOutputWithContext(ctx context.Context) FlyioArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlyioArrayOutput)
}

// FlyioMapInput is an input type that accepts FlyioMap and FlyioMapOutput values.
// You can construct a concrete instance of `FlyioMapInput` via:
//
//	FlyioMap{ "key": FlyioArgs{...} }
type FlyioMapInput interface {
	pulumi.Input

	ToFlyioMapOutput() FlyioMapOutput
	ToFlyioMapOutputWithContext(context.Context) FlyioMapOutput
}

type FlyioMap map[string]FlyioInput

func (FlyioMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Flyio)(nil)).Elem()
}

func (i FlyioMap) ToFlyioMapOutput() FlyioMapOutput {
	return i.ToFlyioMapOutputWithContext(context.Background())
}

func (i FlyioMap) ToFlyioMapOutputWithContext(ctx context.Context) FlyioMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlyioMapOutput)
}

type FlyioOutput struct{ *pulumi.OutputState }

func (FlyioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Flyio)(nil)).Elem()
}

func (o FlyioOutput) ToFlyioOutput() FlyioOutput {
	return o
}

func (o FlyioOutput) ToFlyioOutputWithContext(ctx context.Context) FlyioOutput {
	return o
}

// The app ID
func (o FlyioOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *Flyio) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// The name of the Doppler config
func (o FlyioOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *Flyio) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
func (o FlyioOutput) DeleteBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flyio) pulumi.StringPtrOutput { return v.DeleteBehavior }).(pulumi.StringPtrOutput)
}

// The slug of the integration to use for this sync
func (o FlyioOutput) Integration() pulumi.StringOutput {
	return o.ApplyT(func(v *Flyio) pulumi.StringOutput { return v.Integration }).(pulumi.StringOutput)
}

// The name of the Doppler project
func (o FlyioOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Flyio) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Whether or not to restart the Fly.io machines when secrets are updated
func (o FlyioOutput) RestartMachines() pulumi.BoolOutput {
	return o.ApplyT(func(v *Flyio) pulumi.BoolOutput { return v.RestartMachines }).(pulumi.BoolOutput)
}

type FlyioArrayOutput struct{ *pulumi.OutputState }

func (FlyioArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Flyio)(nil)).Elem()
}

func (o FlyioArrayOutput) ToFlyioArrayOutput() FlyioArrayOutput {
	return o
}

func (o FlyioArrayOutput) ToFlyioArrayOutputWithContext(ctx context.Context) FlyioArrayOutput {
	return o
}

func (o FlyioArrayOutput) Index(i pulumi.IntInput) FlyioOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Flyio {
		return vs[0].([]*Flyio)[vs[1].(int)]
	}).(FlyioOutput)
}

type FlyioMapOutput struct{ *pulumi.OutputState }

func (FlyioMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Flyio)(nil)).Elem()
}

func (o FlyioMapOutput) ToFlyioMapOutput() FlyioMapOutput {
	return o
}

func (o FlyioMapOutput) ToFlyioMapOutputWithContext(ctx context.Context) FlyioMapOutput {
	return o
}

func (o FlyioMapOutput) MapIndex(k pulumi.StringInput) FlyioOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Flyio {
		return vs[0].(map[string]*Flyio)[vs[1].(string)]
	}).(FlyioOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlyioInput)(nil)).Elem(), &Flyio{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlyioArrayInput)(nil)).Elem(), FlyioArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlyioMapInput)(nil)).Elem(), FlyioMap{})
	pulumi.RegisterOutputType(FlyioOutput{})
	pulumi.RegisterOutputType(FlyioArrayOutput{})
	pulumi.RegisterOutputType(FlyioMapOutput{})
}