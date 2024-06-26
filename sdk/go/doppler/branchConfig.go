// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package doppler

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-doppler/sdk/go/doppler/internal"
)

// Manage a Doppler config.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-doppler/sdk/go/doppler"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := doppler.NewBranchConfig(ctx, "backend_ci_github", &doppler.BranchConfigArgs{
//				Project:     pulumi.String("backend"),
//				Environment: pulumi.String("ci"),
//				Name:        pulumi.String("ci_github"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import doppler:index/branchConfig:BranchConfig default <project-name>.<environment-slug>.<config-name>
// ```
type BranchConfig struct {
	pulumi.CustomResourceState

	// The name of the Doppler environment where the config is located
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The name of the Doppler config
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Doppler project where the config is located
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewBranchConfig registers a new resource with the given unique name, arguments, and options.
func NewBranchConfig(ctx *pulumi.Context,
	name string, args *BranchConfigArgs, opts ...pulumi.ResourceOption) (*BranchConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("doppler:index/config:Config"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BranchConfig
	err := ctx.RegisterResource("doppler:index/branchConfig:BranchConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchConfig gets an existing BranchConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchConfigState, opts ...pulumi.ResourceOption) (*BranchConfig, error) {
	var resource BranchConfig
	err := ctx.ReadResource("doppler:index/branchConfig:BranchConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchConfig resources.
type branchConfigState struct {
	// The name of the Doppler environment where the config is located
	Environment *string `pulumi:"environment"`
	// The name of the Doppler config
	Name *string `pulumi:"name"`
	// The name of the Doppler project where the config is located
	Project *string `pulumi:"project"`
}

type BranchConfigState struct {
	// The name of the Doppler environment where the config is located
	Environment pulumi.StringPtrInput
	// The name of the Doppler config
	Name pulumi.StringPtrInput
	// The name of the Doppler project where the config is located
	Project pulumi.StringPtrInput
}

func (BranchConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchConfigState)(nil)).Elem()
}

type branchConfigArgs struct {
	// The name of the Doppler environment where the config is located
	Environment string `pulumi:"environment"`
	// The name of the Doppler config
	Name *string `pulumi:"name"`
	// The name of the Doppler project where the config is located
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a BranchConfig resource.
type BranchConfigArgs struct {
	// The name of the Doppler environment where the config is located
	Environment pulumi.StringInput
	// The name of the Doppler config
	Name pulumi.StringPtrInput
	// The name of the Doppler project where the config is located
	Project pulumi.StringInput
}

func (BranchConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchConfigArgs)(nil)).Elem()
}

type BranchConfigInput interface {
	pulumi.Input

	ToBranchConfigOutput() BranchConfigOutput
	ToBranchConfigOutputWithContext(ctx context.Context) BranchConfigOutput
}

func (*BranchConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchConfig)(nil)).Elem()
}

func (i *BranchConfig) ToBranchConfigOutput() BranchConfigOutput {
	return i.ToBranchConfigOutputWithContext(context.Background())
}

func (i *BranchConfig) ToBranchConfigOutputWithContext(ctx context.Context) BranchConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchConfigOutput)
}

// BranchConfigArrayInput is an input type that accepts BranchConfigArray and BranchConfigArrayOutput values.
// You can construct a concrete instance of `BranchConfigArrayInput` via:
//
//	BranchConfigArray{ BranchConfigArgs{...} }
type BranchConfigArrayInput interface {
	pulumi.Input

	ToBranchConfigArrayOutput() BranchConfigArrayOutput
	ToBranchConfigArrayOutputWithContext(context.Context) BranchConfigArrayOutput
}

type BranchConfigArray []BranchConfigInput

func (BranchConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchConfig)(nil)).Elem()
}

func (i BranchConfigArray) ToBranchConfigArrayOutput() BranchConfigArrayOutput {
	return i.ToBranchConfigArrayOutputWithContext(context.Background())
}

func (i BranchConfigArray) ToBranchConfigArrayOutputWithContext(ctx context.Context) BranchConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchConfigArrayOutput)
}

// BranchConfigMapInput is an input type that accepts BranchConfigMap and BranchConfigMapOutput values.
// You can construct a concrete instance of `BranchConfigMapInput` via:
//
//	BranchConfigMap{ "key": BranchConfigArgs{...} }
type BranchConfigMapInput interface {
	pulumi.Input

	ToBranchConfigMapOutput() BranchConfigMapOutput
	ToBranchConfigMapOutputWithContext(context.Context) BranchConfigMapOutput
}

type BranchConfigMap map[string]BranchConfigInput

func (BranchConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchConfig)(nil)).Elem()
}

func (i BranchConfigMap) ToBranchConfigMapOutput() BranchConfigMapOutput {
	return i.ToBranchConfigMapOutputWithContext(context.Background())
}

func (i BranchConfigMap) ToBranchConfigMapOutputWithContext(ctx context.Context) BranchConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchConfigMapOutput)
}

type BranchConfigOutput struct{ *pulumi.OutputState }

func (BranchConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchConfig)(nil)).Elem()
}

func (o BranchConfigOutput) ToBranchConfigOutput() BranchConfigOutput {
	return o
}

func (o BranchConfigOutput) ToBranchConfigOutputWithContext(ctx context.Context) BranchConfigOutput {
	return o
}

// The name of the Doppler environment where the config is located
func (o BranchConfigOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchConfig) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The name of the Doppler config
func (o BranchConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the Doppler project where the config is located
func (o BranchConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type BranchConfigArrayOutput struct{ *pulumi.OutputState }

func (BranchConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchConfig)(nil)).Elem()
}

func (o BranchConfigArrayOutput) ToBranchConfigArrayOutput() BranchConfigArrayOutput {
	return o
}

func (o BranchConfigArrayOutput) ToBranchConfigArrayOutputWithContext(ctx context.Context) BranchConfigArrayOutput {
	return o
}

func (o BranchConfigArrayOutput) Index(i pulumi.IntInput) BranchConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchConfig {
		return vs[0].([]*BranchConfig)[vs[1].(int)]
	}).(BranchConfigOutput)
}

type BranchConfigMapOutput struct{ *pulumi.OutputState }

func (BranchConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchConfig)(nil)).Elem()
}

func (o BranchConfigMapOutput) ToBranchConfigMapOutput() BranchConfigMapOutput {
	return o
}

func (o BranchConfigMapOutput) ToBranchConfigMapOutputWithContext(ctx context.Context) BranchConfigMapOutput {
	return o
}

func (o BranchConfigMapOutput) MapIndex(k pulumi.StringInput) BranchConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchConfig {
		return vs[0].(map[string]*BranchConfig)[vs[1].(string)]
	}).(BranchConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchConfigInput)(nil)).Elem(), &BranchConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchConfigArrayInput)(nil)).Elem(), BranchConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchConfigMapInput)(nil)).Elem(), BranchConfigMap{})
	pulumi.RegisterOutputType(BranchConfigOutput{})
	pulumi.RegisterOutputType(BranchConfigArrayOutput{})
	pulumi.RegisterOutputType(BranchConfigMapOutput{})
}
