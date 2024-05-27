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

// ## Example Usage
type TerraformCloud struct {
	pulumi.CustomResourceState

	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// The name of the integration
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTerraformCloud registers a new resource with the given unique name, arguments, and options.
func NewTerraformCloud(ctx *pulumi.Context,
	name string, args *TerraformCloudArgs, opts ...pulumi.ResourceOption) (*TerraformCloud, error) {
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
	var resource TerraformCloud
	err := ctx.RegisterResource("doppler:integration/terraformCloud:TerraformCloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTerraformCloud gets an existing TerraformCloud resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTerraformCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TerraformCloudState, opts ...pulumi.ResourceOption) (*TerraformCloud, error) {
	var resource TerraformCloud
	err := ctx.ReadResource("doppler:integration/terraformCloud:TerraformCloud", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TerraformCloud resources.
type terraformCloudState struct {
	ApiKey *string `pulumi:"apiKey"`
	// The name of the integration
	Name *string `pulumi:"name"`
}

type TerraformCloudState struct {
	ApiKey pulumi.StringPtrInput
	// The name of the integration
	Name pulumi.StringPtrInput
}

func (TerraformCloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*terraformCloudState)(nil)).Elem()
}

type terraformCloudArgs struct {
	ApiKey string `pulumi:"apiKey"`
	// The name of the integration
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a TerraformCloud resource.
type TerraformCloudArgs struct {
	ApiKey pulumi.StringInput
	// The name of the integration
	Name pulumi.StringInput
}

func (TerraformCloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*terraformCloudArgs)(nil)).Elem()
}

type TerraformCloudInput interface {
	pulumi.Input

	ToTerraformCloudOutput() TerraformCloudOutput
	ToTerraformCloudOutputWithContext(ctx context.Context) TerraformCloudOutput
}

func (*TerraformCloud) ElementType() reflect.Type {
	return reflect.TypeOf((**TerraformCloud)(nil)).Elem()
}

func (i *TerraformCloud) ToTerraformCloudOutput() TerraformCloudOutput {
	return i.ToTerraformCloudOutputWithContext(context.Background())
}

func (i *TerraformCloud) ToTerraformCloudOutputWithContext(ctx context.Context) TerraformCloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerraformCloudOutput)
}

// TerraformCloudArrayInput is an input type that accepts TerraformCloudArray and TerraformCloudArrayOutput values.
// You can construct a concrete instance of `TerraformCloudArrayInput` via:
//
//	TerraformCloudArray{ TerraformCloudArgs{...} }
type TerraformCloudArrayInput interface {
	pulumi.Input

	ToTerraformCloudArrayOutput() TerraformCloudArrayOutput
	ToTerraformCloudArrayOutputWithContext(context.Context) TerraformCloudArrayOutput
}

type TerraformCloudArray []TerraformCloudInput

func (TerraformCloudArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TerraformCloud)(nil)).Elem()
}

func (i TerraformCloudArray) ToTerraformCloudArrayOutput() TerraformCloudArrayOutput {
	return i.ToTerraformCloudArrayOutputWithContext(context.Background())
}

func (i TerraformCloudArray) ToTerraformCloudArrayOutputWithContext(ctx context.Context) TerraformCloudArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerraformCloudArrayOutput)
}

// TerraformCloudMapInput is an input type that accepts TerraformCloudMap and TerraformCloudMapOutput values.
// You can construct a concrete instance of `TerraformCloudMapInput` via:
//
//	TerraformCloudMap{ "key": TerraformCloudArgs{...} }
type TerraformCloudMapInput interface {
	pulumi.Input

	ToTerraformCloudMapOutput() TerraformCloudMapOutput
	ToTerraformCloudMapOutputWithContext(context.Context) TerraformCloudMapOutput
}

type TerraformCloudMap map[string]TerraformCloudInput

func (TerraformCloudMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TerraformCloud)(nil)).Elem()
}

func (i TerraformCloudMap) ToTerraformCloudMapOutput() TerraformCloudMapOutput {
	return i.ToTerraformCloudMapOutputWithContext(context.Background())
}

func (i TerraformCloudMap) ToTerraformCloudMapOutputWithContext(ctx context.Context) TerraformCloudMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerraformCloudMapOutput)
}

type TerraformCloudOutput struct{ *pulumi.OutputState }

func (TerraformCloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TerraformCloud)(nil)).Elem()
}

func (o TerraformCloudOutput) ToTerraformCloudOutput() TerraformCloudOutput {
	return o
}

func (o TerraformCloudOutput) ToTerraformCloudOutputWithContext(ctx context.Context) TerraformCloudOutput {
	return o
}

func (o TerraformCloudOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *TerraformCloud) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

// The name of the integration
func (o TerraformCloudOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TerraformCloud) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type TerraformCloudArrayOutput struct{ *pulumi.OutputState }

func (TerraformCloudArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TerraformCloud)(nil)).Elem()
}

func (o TerraformCloudArrayOutput) ToTerraformCloudArrayOutput() TerraformCloudArrayOutput {
	return o
}

func (o TerraformCloudArrayOutput) ToTerraformCloudArrayOutputWithContext(ctx context.Context) TerraformCloudArrayOutput {
	return o
}

func (o TerraformCloudArrayOutput) Index(i pulumi.IntInput) TerraformCloudOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TerraformCloud {
		return vs[0].([]*TerraformCloud)[vs[1].(int)]
	}).(TerraformCloudOutput)
}

type TerraformCloudMapOutput struct{ *pulumi.OutputState }

func (TerraformCloudMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TerraformCloud)(nil)).Elem()
}

func (o TerraformCloudMapOutput) ToTerraformCloudMapOutput() TerraformCloudMapOutput {
	return o
}

func (o TerraformCloudMapOutput) ToTerraformCloudMapOutputWithContext(ctx context.Context) TerraformCloudMapOutput {
	return o
}

func (o TerraformCloudMapOutput) MapIndex(k pulumi.StringInput) TerraformCloudOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TerraformCloud {
		return vs[0].(map[string]*TerraformCloud)[vs[1].(string)]
	}).(TerraformCloudOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TerraformCloudInput)(nil)).Elem(), &TerraformCloud{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerraformCloudArrayInput)(nil)).Elem(), TerraformCloudArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerraformCloudMapInput)(nil)).Elem(), TerraformCloudMap{})
	pulumi.RegisterOutputType(TerraformCloudOutput{})
	pulumi.RegisterOutputType(TerraformCloudArrayOutput{})
	pulumi.RegisterOutputType(TerraformCloudMapOutput{})
}
