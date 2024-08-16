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

type Webhook struct {
	pulumi.CustomResourceState

	// Authentication method used by the webhook
	Authentication WebhookAuthenticationPtrOutput `pulumi:"authentication"`
	// Whether the webhook is enabled or disabled.  Default to true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Configs this webhook will trigger for
	EnabledConfigs pulumi.StringArrayOutput `pulumi:"enabledConfigs"`
	// The webhook's payload as a JSON string.  Leave empty to use the default webhook payload
	Payload pulumi.StringPtrOutput `pulumi:"payload"`
	// The name of the Doppler project where the webhook is located
	Project pulumi.StringOutput `pulumi:"project"`
	// Secret used for request signing
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// The slug of the Webhook
	Slug pulumi.StringOutput `pulumi:"slug"`
	// The URL of the webhook endpoint
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Payload != nil {
		args.Payload = pulumi.ToSecret(args.Payload).(pulumi.StringPtrInput)
	}
	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"payload",
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Webhook
	err := ctx.RegisterResource("doppler:index/webhook:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("doppler:index/webhook:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
	// Authentication method used by the webhook
	Authentication *WebhookAuthentication `pulumi:"authentication"`
	// Whether the webhook is enabled or disabled.  Default to true.
	Enabled *bool `pulumi:"enabled"`
	// Configs this webhook will trigger for
	EnabledConfigs []string `pulumi:"enabledConfigs"`
	// The webhook's payload as a JSON string.  Leave empty to use the default webhook payload
	Payload *string `pulumi:"payload"`
	// The name of the Doppler project where the webhook is located
	Project *string `pulumi:"project"`
	// Secret used for request signing
	Secret *string `pulumi:"secret"`
	// The slug of the Webhook
	Slug *string `pulumi:"slug"`
	// The URL of the webhook endpoint
	Url *string `pulumi:"url"`
}

type WebhookState struct {
	// Authentication method used by the webhook
	Authentication WebhookAuthenticationPtrInput
	// Whether the webhook is enabled or disabled.  Default to true.
	Enabled pulumi.BoolPtrInput
	// Configs this webhook will trigger for
	EnabledConfigs pulumi.StringArrayInput
	// The webhook's payload as a JSON string.  Leave empty to use the default webhook payload
	Payload pulumi.StringPtrInput
	// The name of the Doppler project where the webhook is located
	Project pulumi.StringPtrInput
	// Secret used for request signing
	Secret pulumi.StringPtrInput
	// The slug of the Webhook
	Slug pulumi.StringPtrInput
	// The URL of the webhook endpoint
	Url pulumi.StringPtrInput
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	// Authentication method used by the webhook
	Authentication *WebhookAuthentication `pulumi:"authentication"`
	// Whether the webhook is enabled or disabled.  Default to true.
	Enabled *bool `pulumi:"enabled"`
	// Configs this webhook will trigger for
	EnabledConfigs []string `pulumi:"enabledConfigs"`
	// The webhook's payload as a JSON string.  Leave empty to use the default webhook payload
	Payload *string `pulumi:"payload"`
	// The name of the Doppler project where the webhook is located
	Project string `pulumi:"project"`
	// Secret used for request signing
	Secret *string `pulumi:"secret"`
	// The URL of the webhook endpoint
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// Authentication method used by the webhook
	Authentication WebhookAuthenticationPtrInput
	// Whether the webhook is enabled or disabled.  Default to true.
	Enabled pulumi.BoolPtrInput
	// Configs this webhook will trigger for
	EnabledConfigs pulumi.StringArrayInput
	// The webhook's payload as a JSON string.  Leave empty to use the default webhook payload
	Payload pulumi.StringPtrInput
	// The name of the Doppler project where the webhook is located
	Project pulumi.StringInput
	// Secret used for request signing
	Secret pulumi.StringPtrInput
	// The URL of the webhook endpoint
	Url pulumi.StringInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(ctx context.Context) WebhookOutput
}

func (*Webhook) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (i *Webhook) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

// WebhookArrayInput is an input type that accepts WebhookArray and WebhookArrayOutput values.
// You can construct a concrete instance of `WebhookArrayInput` via:
//
//	WebhookArray{ WebhookArgs{...} }
type WebhookArrayInput interface {
	pulumi.Input

	ToWebhookArrayOutput() WebhookArrayOutput
	ToWebhookArrayOutputWithContext(context.Context) WebhookArrayOutput
}

type WebhookArray []WebhookInput

func (WebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Webhook)(nil)).Elem()
}

func (i WebhookArray) ToWebhookArrayOutput() WebhookArrayOutput {
	return i.ToWebhookArrayOutputWithContext(context.Background())
}

func (i WebhookArray) ToWebhookArrayOutputWithContext(ctx context.Context) WebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookArrayOutput)
}

// WebhookMapInput is an input type that accepts WebhookMap and WebhookMapOutput values.
// You can construct a concrete instance of `WebhookMapInput` via:
//
//	WebhookMap{ "key": WebhookArgs{...} }
type WebhookMapInput interface {
	pulumi.Input

	ToWebhookMapOutput() WebhookMapOutput
	ToWebhookMapOutputWithContext(context.Context) WebhookMapOutput
}

type WebhookMap map[string]WebhookInput

func (WebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Webhook)(nil)).Elem()
}

func (i WebhookMap) ToWebhookMapOutput() WebhookMapOutput {
	return i.ToWebhookMapOutputWithContext(context.Background())
}

func (i WebhookMap) ToWebhookMapOutputWithContext(ctx context.Context) WebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookMapOutput)
}

type WebhookOutput struct{ *pulumi.OutputState }

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

// Authentication method used by the webhook
func (o WebhookOutput) Authentication() WebhookAuthenticationPtrOutput {
	return o.ApplyT(func(v *Webhook) WebhookAuthenticationPtrOutput { return v.Authentication }).(WebhookAuthenticationPtrOutput)
}

// Whether the webhook is enabled or disabled.  Default to true.
func (o WebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Configs this webhook will trigger for
func (o WebhookOutput) EnabledConfigs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringArrayOutput { return v.EnabledConfigs }).(pulumi.StringArrayOutput)
}

// The webhook's payload as a JSON string.  Leave empty to use the default webhook payload
func (o WebhookOutput) Payload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.Payload }).(pulumi.StringPtrOutput)
}

// The name of the Doppler project where the webhook is located
func (o WebhookOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Secret used for request signing
func (o WebhookOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.Secret }).(pulumi.StringPtrOutput)
}

// The slug of the Webhook
func (o WebhookOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// The URL of the webhook endpoint
func (o WebhookOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type WebhookArrayOutput struct{ *pulumi.OutputState }

func (WebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Webhook)(nil)).Elem()
}

func (o WebhookArrayOutput) ToWebhookArrayOutput() WebhookArrayOutput {
	return o
}

func (o WebhookArrayOutput) ToWebhookArrayOutputWithContext(ctx context.Context) WebhookArrayOutput {
	return o
}

func (o WebhookArrayOutput) Index(i pulumi.IntInput) WebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Webhook {
		return vs[0].([]*Webhook)[vs[1].(int)]
	}).(WebhookOutput)
}

type WebhookMapOutput struct{ *pulumi.OutputState }

func (WebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Webhook)(nil)).Elem()
}

func (o WebhookMapOutput) ToWebhookMapOutput() WebhookMapOutput {
	return o
}

func (o WebhookMapOutput) ToWebhookMapOutputWithContext(ctx context.Context) WebhookMapOutput {
	return o
}

func (o WebhookMapOutput) MapIndex(k pulumi.StringInput) WebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Webhook {
		return vs[0].(map[string]*Webhook)[vs[1].(string)]
	}).(WebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookInput)(nil)).Elem(), &Webhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookArrayInput)(nil)).Elem(), WebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookMapInput)(nil)).Elem(), WebhookMap{})
	pulumi.RegisterOutputType(WebhookOutput{})
	pulumi.RegisterOutputType(WebhookArrayOutput{})
	pulumi.RegisterOutputType(WebhookMapOutput{})
}
