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

// Manage a Doppler group's memberships.
//
// **Note:** The `GroupMembers` resource will clear/replace all existing memberships.
// Multiple `GroupMembers` resources or combinations of `GroupMembers` and `GroupMember` will produce inconsistent behavior.
// To non-exclusively manage group memberships, use `GroupMember` only.
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
//			engineering, err := doppler.NewGroup(ctx, "engineering", &doppler.GroupArgs{
//				Name: pulumi.String("engineering"),
//			})
//			if err != nil {
//				return err
//			}
//			nic, err := doppler.GetUser(ctx, &doppler.GetUserArgs{
//				Email: "nic@doppler.com",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			andre, err := doppler.GetUser(ctx, &doppler.GetUserArgs{
//				Email: "andre@doppler.com",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = doppler.NewGroupMembers(ctx, "engineering", &doppler.GroupMembersArgs{
//				GroupSlug: engineering.Slug,
//				UserSlugs: pulumi.StringArray{
//					pulumi.String(nic.Slug),
//					pulumi.String(andre.Slug),
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
//
// ## Import
//
// import using the group slug from the URL:
//
// https://dashboard.doppler.com/workplace/[workplace-slug]/team/groups/[group-slug]
//
// and the user slugs from the URL:
//
// https://dashboard.doppler.com/workplace/[workplace-slug]/team/users/[user-slug]
//
// ```sh
// $ pulumi import doppler:index/groupMembers:GroupMembers default <group-slug>
// ```
type GroupMembers struct {
	pulumi.CustomResourceState

	// The slug of the group
	GroupSlug pulumi.StringOutput `pulumi:"groupSlug"`
	// A list of user slugs in the group
	UserSlugs pulumi.StringArrayOutput `pulumi:"userSlugs"`
}

// NewGroupMembers registers a new resource with the given unique name, arguments, and options.
func NewGroupMembers(ctx *pulumi.Context,
	name string, args *GroupMembersArgs, opts ...pulumi.ResourceOption) (*GroupMembers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupSlug == nil {
		return nil, errors.New("invalid value for required argument 'GroupSlug'")
	}
	if args.UserSlugs == nil {
		return nil, errors.New("invalid value for required argument 'UserSlugs'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupMembers
	err := ctx.RegisterResource("doppler:index/groupMembers:GroupMembers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMembers gets an existing GroupMembers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMembers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembersState, opts ...pulumi.ResourceOption) (*GroupMembers, error) {
	var resource GroupMembers
	err := ctx.ReadResource("doppler:index/groupMembers:GroupMembers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMembers resources.
type groupMembersState struct {
	// The slug of the group
	GroupSlug *string `pulumi:"groupSlug"`
	// A list of user slugs in the group
	UserSlugs []string `pulumi:"userSlugs"`
}

type GroupMembersState struct {
	// The slug of the group
	GroupSlug pulumi.StringPtrInput
	// A list of user slugs in the group
	UserSlugs pulumi.StringArrayInput
}

func (GroupMembersState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembersState)(nil)).Elem()
}

type groupMembersArgs struct {
	// The slug of the group
	GroupSlug string `pulumi:"groupSlug"`
	// A list of user slugs in the group
	UserSlugs []string `pulumi:"userSlugs"`
}

// The set of arguments for constructing a GroupMembers resource.
type GroupMembersArgs struct {
	// The slug of the group
	GroupSlug pulumi.StringInput
	// A list of user slugs in the group
	UserSlugs pulumi.StringArrayInput
}

func (GroupMembersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembersArgs)(nil)).Elem()
}

type GroupMembersInput interface {
	pulumi.Input

	ToGroupMembersOutput() GroupMembersOutput
	ToGroupMembersOutputWithContext(ctx context.Context) GroupMembersOutput
}

func (*GroupMembers) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembers)(nil)).Elem()
}

func (i *GroupMembers) ToGroupMembersOutput() GroupMembersOutput {
	return i.ToGroupMembersOutputWithContext(context.Background())
}

func (i *GroupMembers) ToGroupMembersOutputWithContext(ctx context.Context) GroupMembersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembersOutput)
}

// GroupMembersArrayInput is an input type that accepts GroupMembersArray and GroupMembersArrayOutput values.
// You can construct a concrete instance of `GroupMembersArrayInput` via:
//
//	GroupMembersArray{ GroupMembersArgs{...} }
type GroupMembersArrayInput interface {
	pulumi.Input

	ToGroupMembersArrayOutput() GroupMembersArrayOutput
	ToGroupMembersArrayOutputWithContext(context.Context) GroupMembersArrayOutput
}

type GroupMembersArray []GroupMembersInput

func (GroupMembersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembers)(nil)).Elem()
}

func (i GroupMembersArray) ToGroupMembersArrayOutput() GroupMembersArrayOutput {
	return i.ToGroupMembersArrayOutputWithContext(context.Background())
}

func (i GroupMembersArray) ToGroupMembersArrayOutputWithContext(ctx context.Context) GroupMembersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembersArrayOutput)
}

// GroupMembersMapInput is an input type that accepts GroupMembersMap and GroupMembersMapOutput values.
// You can construct a concrete instance of `GroupMembersMapInput` via:
//
//	GroupMembersMap{ "key": GroupMembersArgs{...} }
type GroupMembersMapInput interface {
	pulumi.Input

	ToGroupMembersMapOutput() GroupMembersMapOutput
	ToGroupMembersMapOutputWithContext(context.Context) GroupMembersMapOutput
}

type GroupMembersMap map[string]GroupMembersInput

func (GroupMembersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembers)(nil)).Elem()
}

func (i GroupMembersMap) ToGroupMembersMapOutput() GroupMembersMapOutput {
	return i.ToGroupMembersMapOutputWithContext(context.Background())
}

func (i GroupMembersMap) ToGroupMembersMapOutputWithContext(ctx context.Context) GroupMembersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembersMapOutput)
}

type GroupMembersOutput struct{ *pulumi.OutputState }

func (GroupMembersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembers)(nil)).Elem()
}

func (o GroupMembersOutput) ToGroupMembersOutput() GroupMembersOutput {
	return o
}

func (o GroupMembersOutput) ToGroupMembersOutputWithContext(ctx context.Context) GroupMembersOutput {
	return o
}

// The slug of the group
func (o GroupMembersOutput) GroupSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMembers) pulumi.StringOutput { return v.GroupSlug }).(pulumi.StringOutput)
}

// A list of user slugs in the group
func (o GroupMembersOutput) UserSlugs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupMembers) pulumi.StringArrayOutput { return v.UserSlugs }).(pulumi.StringArrayOutput)
}

type GroupMembersArrayOutput struct{ *pulumi.OutputState }

func (GroupMembersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembers)(nil)).Elem()
}

func (o GroupMembersArrayOutput) ToGroupMembersArrayOutput() GroupMembersArrayOutput {
	return o
}

func (o GroupMembersArrayOutput) ToGroupMembersArrayOutputWithContext(ctx context.Context) GroupMembersArrayOutput {
	return o
}

func (o GroupMembersArrayOutput) Index(i pulumi.IntInput) GroupMembersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMembers {
		return vs[0].([]*GroupMembers)[vs[1].(int)]
	}).(GroupMembersOutput)
}

type GroupMembersMapOutput struct{ *pulumi.OutputState }

func (GroupMembersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembers)(nil)).Elem()
}

func (o GroupMembersMapOutput) ToGroupMembersMapOutput() GroupMembersMapOutput {
	return o
}

func (o GroupMembersMapOutput) ToGroupMembersMapOutputWithContext(ctx context.Context) GroupMembersMapOutput {
	return o
}

func (o GroupMembersMapOutput) MapIndex(k pulumi.StringInput) GroupMembersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMembers {
		return vs[0].(map[string]*GroupMembers)[vs[1].(string)]
	}).(GroupMembersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembersInput)(nil)).Elem(), &GroupMembers{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembersArrayInput)(nil)).Elem(), GroupMembersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembersMapInput)(nil)).Elem(), GroupMembersMap{})
	pulumi.RegisterOutputType(GroupMembersOutput{})
	pulumi.RegisterOutputType(GroupMembersArrayOutput{})
	pulumi.RegisterOutputType(GroupMembersMapOutput{})
}