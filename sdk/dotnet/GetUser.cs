// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Doppler
{
    public static class GetUser
    {
        /// <summary>
        /// Retrieve an existing Doppler user.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Doppler = Pulumi.Doppler;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var nic = Doppler.GetUser.Invoke(new()
        ///     {
        ///         Email = "nic@doppler.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("doppler:index/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve an existing Doppler user.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Doppler = Pulumi.Doppler;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var nic = Doppler.GetUser.Invoke(new()
        ///     {
        ///         Email = "nic@doppler.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("doppler:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The email address of the Doppler user
        /// </summary>
        [Input("email", required: true)]
        public string Email { get; set; } = null!;

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The email address of the Doppler user
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// The email address of the Doppler user
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The slug of the Doppler user
        /// </summary>
        public readonly string Slug;

        [OutputConstructor]
        private GetUserResult(
            string email,

            string id,

            string slug)
        {
            Email = email;
            Id = id;
            Slug = slug;
        }
    }
}