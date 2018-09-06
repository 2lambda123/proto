// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: iam/grpcapi.proto
// </auto-generated>
#pragma warning disable 0414, 1591
#region Designer generated code

using grpc = global::Grpc.Core;

namespace SKF.Enlight.API.IAM {
  public static partial class IAM
  {
    static readonly string __ServiceName = "grpcapi.IAM";

    static readonly grpc::Marshaller<global::SKF.Enlight.API.IAM.PrimitiveVoid> __Marshaller_grpcapi_PrimitiveVoid = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::SKF.Enlight.API.IAM.PrimitiveVoid.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::SKF.Enlight.API.IAM.PrimitiveString> __Marshaller_grpcapi_PrimitiveString = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::SKF.Enlight.API.IAM.PrimitiveString.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::SKF.Enlight.API.IAM.CheckAuthenticationInput> __Marshaller_grpcapi_CheckAuthenticationInput = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::SKF.Enlight.API.IAM.CheckAuthenticationInput.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::SKF.Enlight.API.IAM.User> __Marshaller_grpcapi_User = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::SKF.Enlight.API.IAM.User.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::SKF.Enlight.API.IAM.GetHierarchyRelationsInput> __Marshaller_grpcapi_GetHierarchyRelationsInput = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::SKF.Enlight.API.IAM.GetHierarchyRelationsInput.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::SKF.Enlight.API.IAM.GetHierarchyRelationsOutput> __Marshaller_grpcapi_GetHierarchyRelationsOutput = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::SKF.Enlight.API.IAM.GetHierarchyRelationsOutput.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::SKF.Enlight.API.IAM.GetEventRecordsInput> __Marshaller_grpcapi_GetEventRecordsInput = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::SKF.Enlight.API.IAM.GetEventRecordsInput.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::SKF.Enlight.API.IAM.GetEventRecordsOutput> __Marshaller_grpcapi_GetEventRecordsOutput = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::SKF.Enlight.API.IAM.GetEventRecordsOutput.Parser.ParseFrom);

    static readonly grpc::Method<global::SKF.Enlight.API.IAM.PrimitiveVoid, global::SKF.Enlight.API.IAM.PrimitiveString> __Method_DeepPing = new grpc::Method<global::SKF.Enlight.API.IAM.PrimitiveVoid, global::SKF.Enlight.API.IAM.PrimitiveString>(
        grpc::MethodType.Unary,
        __ServiceName,
        "DeepPing",
        __Marshaller_grpcapi_PrimitiveVoid,
        __Marshaller_grpcapi_PrimitiveString);

    static readonly grpc::Method<global::SKF.Enlight.API.IAM.CheckAuthenticationInput, global::SKF.Enlight.API.IAM.User> __Method_CheckAuthentication = new grpc::Method<global::SKF.Enlight.API.IAM.CheckAuthenticationInput, global::SKF.Enlight.API.IAM.User>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CheckAuthentication",
        __Marshaller_grpcapi_CheckAuthenticationInput,
        __Marshaller_grpcapi_User);

    static readonly grpc::Method<global::SKF.Enlight.API.IAM.GetHierarchyRelationsInput, global::SKF.Enlight.API.IAM.GetHierarchyRelationsOutput> __Method_GetHierarchyRelations = new grpc::Method<global::SKF.Enlight.API.IAM.GetHierarchyRelationsInput, global::SKF.Enlight.API.IAM.GetHierarchyRelationsOutput>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetHierarchyRelations",
        __Marshaller_grpcapi_GetHierarchyRelationsInput,
        __Marshaller_grpcapi_GetHierarchyRelationsOutput);

    static readonly grpc::Method<global::SKF.Enlight.API.IAM.GetEventRecordsInput, global::SKF.Enlight.API.IAM.GetEventRecordsOutput> __Method_GetEventRecords = new grpc::Method<global::SKF.Enlight.API.IAM.GetEventRecordsInput, global::SKF.Enlight.API.IAM.GetEventRecordsOutput>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetEventRecords",
        __Marshaller_grpcapi_GetEventRecordsInput,
        __Marshaller_grpcapi_GetEventRecordsOutput);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::SKF.Enlight.API.IAM.GrpcapiReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of IAM</summary>
    public abstract partial class IAMBase
    {
      public virtual global::System.Threading.Tasks.Task<global::SKF.Enlight.API.IAM.PrimitiveString> DeepPing(global::SKF.Enlight.API.IAM.PrimitiveVoid request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::SKF.Enlight.API.IAM.User> CheckAuthentication(global::SKF.Enlight.API.IAM.CheckAuthenticationInput request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::SKF.Enlight.API.IAM.GetHierarchyRelationsOutput> GetHierarchyRelations(global::SKF.Enlight.API.IAM.GetHierarchyRelationsInput request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::SKF.Enlight.API.IAM.GetEventRecordsOutput> GetEventRecords(global::SKF.Enlight.API.IAM.GetEventRecordsInput request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for IAM</summary>
    public partial class IAMClient : grpc::ClientBase<IAMClient>
    {
      /// <summary>Creates a new client for IAM</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public IAMClient(grpc::Channel channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for IAM that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public IAMClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected IAMClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected IAMClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::SKF.Enlight.API.IAM.PrimitiveString DeepPing(global::SKF.Enlight.API.IAM.PrimitiveVoid request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return DeepPing(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::SKF.Enlight.API.IAM.PrimitiveString DeepPing(global::SKF.Enlight.API.IAM.PrimitiveVoid request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_DeepPing, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::SKF.Enlight.API.IAM.PrimitiveString> DeepPingAsync(global::SKF.Enlight.API.IAM.PrimitiveVoid request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return DeepPingAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::SKF.Enlight.API.IAM.PrimitiveString> DeepPingAsync(global::SKF.Enlight.API.IAM.PrimitiveVoid request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_DeepPing, null, options, request);
      }
      public virtual global::SKF.Enlight.API.IAM.User CheckAuthentication(global::SKF.Enlight.API.IAM.CheckAuthenticationInput request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CheckAuthentication(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::SKF.Enlight.API.IAM.User CheckAuthentication(global::SKF.Enlight.API.IAM.CheckAuthenticationInput request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CheckAuthentication, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::SKF.Enlight.API.IAM.User> CheckAuthenticationAsync(global::SKF.Enlight.API.IAM.CheckAuthenticationInput request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CheckAuthenticationAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::SKF.Enlight.API.IAM.User> CheckAuthenticationAsync(global::SKF.Enlight.API.IAM.CheckAuthenticationInput request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CheckAuthentication, null, options, request);
      }
      public virtual global::SKF.Enlight.API.IAM.GetHierarchyRelationsOutput GetHierarchyRelations(global::SKF.Enlight.API.IAM.GetHierarchyRelationsInput request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetHierarchyRelations(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::SKF.Enlight.API.IAM.GetHierarchyRelationsOutput GetHierarchyRelations(global::SKF.Enlight.API.IAM.GetHierarchyRelationsInput request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetHierarchyRelations, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::SKF.Enlight.API.IAM.GetHierarchyRelationsOutput> GetHierarchyRelationsAsync(global::SKF.Enlight.API.IAM.GetHierarchyRelationsInput request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetHierarchyRelationsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::SKF.Enlight.API.IAM.GetHierarchyRelationsOutput> GetHierarchyRelationsAsync(global::SKF.Enlight.API.IAM.GetHierarchyRelationsInput request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetHierarchyRelations, null, options, request);
      }
      public virtual global::SKF.Enlight.API.IAM.GetEventRecordsOutput GetEventRecords(global::SKF.Enlight.API.IAM.GetEventRecordsInput request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetEventRecords(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::SKF.Enlight.API.IAM.GetEventRecordsOutput GetEventRecords(global::SKF.Enlight.API.IAM.GetEventRecordsInput request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetEventRecords, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::SKF.Enlight.API.IAM.GetEventRecordsOutput> GetEventRecordsAsync(global::SKF.Enlight.API.IAM.GetEventRecordsInput request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetEventRecordsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::SKF.Enlight.API.IAM.GetEventRecordsOutput> GetEventRecordsAsync(global::SKF.Enlight.API.IAM.GetEventRecordsInput request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetEventRecords, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override IAMClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new IAMClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static grpc::ServerServiceDefinition BindService(IAMBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_DeepPing, serviceImpl.DeepPing)
          .AddMethod(__Method_CheckAuthentication, serviceImpl.CheckAuthentication)
          .AddMethod(__Method_GetHierarchyRelations, serviceImpl.GetHierarchyRelations)
          .AddMethod(__Method_GetEventRecords, serviceImpl.GetEventRecords).Build();
    }

  }
}
#endregion
