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

// Manage a Fly.io Doppler integration.
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

	// A Fly.io API key.
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// The name of the integration
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewFlyio registers a new resource with the given unique name, arguments, and options.
func NewFlyio(ctx *pulumi.Context,
	name string, args *FlyioArgs, opts ...pulumi.ResourceOption) (*Flyio, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Flyio
	err := ctx.RegisterResource("doppler:integration/flyio:Flyio", name, args, &resource, opts...)
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
	err := ctx.ReadResource("doppler:integration/flyio:Flyio", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flyio resources.
type flyioState struct {
	// A Fly.io API key.
	ApiKey *string `pulumi:"apiKey"`
	// The name of the integration
	Name *string `pulumi:"name"`
}

type FlyioState struct {
	// A Fly.io API key.
	ApiKey pulumi.StringPtrInput
	// The name of the integration
	Name pulumi.StringPtrInput
}

func (FlyioState) ElementType() reflect.Type {
	return reflect.TypeOf((*flyioState)(nil)).Elem()
}

type flyioArgs struct {
	// A Fly.io API key.
	ApiKey string `pulumi:"apiKey"`
	// The name of the integration
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a Flyio resource.
type FlyioArgs struct {
	// A Fly.io API key.
	ApiKey pulumi.StringInput
	// The name of the integration
	Name pulumi.StringInput
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

// A Fly.io API key.
func (o FlyioOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Flyio) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

// The name of the integration
func (o FlyioOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Flyio) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
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
