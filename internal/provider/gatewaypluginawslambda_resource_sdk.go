// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginAWSLambdaResourceModel) ToSharedCreateAWSLambdaPlugin() *shared.CreateAWSLambdaPlugin {
	var config *shared.CreateAWSLambdaPluginConfig
	if r.Config != nil {
		awsAssumeRoleArn := new(string)
		if !r.Config.AwsAssumeRoleArn.IsUnknown() && !r.Config.AwsAssumeRoleArn.IsNull() {
			*awsAssumeRoleArn = r.Config.AwsAssumeRoleArn.ValueString()
		} else {
			awsAssumeRoleArn = nil
		}
		awsImdsProtocolVersion := new(shared.CreateAWSLambdaPluginAWSImdsProtocolVersion)
		if !r.Config.AwsImdsProtocolVersion.IsUnknown() && !r.Config.AwsImdsProtocolVersion.IsNull() {
			*awsImdsProtocolVersion = shared.CreateAWSLambdaPluginAWSImdsProtocolVersion(r.Config.AwsImdsProtocolVersion.ValueString())
		} else {
			awsImdsProtocolVersion = nil
		}
		awsKey := new(string)
		if !r.Config.AwsKey.IsUnknown() && !r.Config.AwsKey.IsNull() {
			*awsKey = r.Config.AwsKey.ValueString()
		} else {
			awsKey = nil
		}
		awsRegion := new(string)
		if !r.Config.AwsRegion.IsUnknown() && !r.Config.AwsRegion.IsNull() {
			*awsRegion = r.Config.AwsRegion.ValueString()
		} else {
			awsRegion = nil
		}
		awsRoleSessionName := new(string)
		if !r.Config.AwsRoleSessionName.IsUnknown() && !r.Config.AwsRoleSessionName.IsNull() {
			*awsRoleSessionName = r.Config.AwsRoleSessionName.ValueString()
		} else {
			awsRoleSessionName = nil
		}
		awsSecret := new(string)
		if !r.Config.AwsSecret.IsUnknown() && !r.Config.AwsSecret.IsNull() {
			*awsSecret = r.Config.AwsSecret.ValueString()
		} else {
			awsSecret = nil
		}
		awsgatewayCompatible := new(bool)
		if !r.Config.AwsgatewayCompatible.IsUnknown() && !r.Config.AwsgatewayCompatible.IsNull() {
			*awsgatewayCompatible = r.Config.AwsgatewayCompatible.ValueBool()
		} else {
			awsgatewayCompatible = nil
		}
		base64EncodeBody := new(bool)
		if !r.Config.Base64EncodeBody.IsUnknown() && !r.Config.Base64EncodeBody.IsNull() {
			*base64EncodeBody = r.Config.Base64EncodeBody.ValueBool()
		} else {
			base64EncodeBody = nil
		}
		disableHTTPS := new(bool)
		if !r.Config.DisableHTTPS.IsUnknown() && !r.Config.DisableHTTPS.IsNull() {
			*disableHTTPS = r.Config.DisableHTTPS.ValueBool()
		} else {
			disableHTTPS = nil
		}
		forwardRequestBody := new(bool)
		if !r.Config.ForwardRequestBody.IsUnknown() && !r.Config.ForwardRequestBody.IsNull() {
			*forwardRequestBody = r.Config.ForwardRequestBody.ValueBool()
		} else {
			forwardRequestBody = nil
		}
		forwardRequestHeaders := new(bool)
		if !r.Config.ForwardRequestHeaders.IsUnknown() && !r.Config.ForwardRequestHeaders.IsNull() {
			*forwardRequestHeaders = r.Config.ForwardRequestHeaders.ValueBool()
		} else {
			forwardRequestHeaders = nil
		}
		forwardRequestMethod := new(bool)
		if !r.Config.ForwardRequestMethod.IsUnknown() && !r.Config.ForwardRequestMethod.IsNull() {
			*forwardRequestMethod = r.Config.ForwardRequestMethod.ValueBool()
		} else {
			forwardRequestMethod = nil
		}
		forwardRequestURI := new(bool)
		if !r.Config.ForwardRequestURI.IsUnknown() && !r.Config.ForwardRequestURI.IsNull() {
			*forwardRequestURI = r.Config.ForwardRequestURI.ValueBool()
		} else {
			forwardRequestURI = nil
		}
		functionName := new(string)
		if !r.Config.FunctionName.IsUnknown() && !r.Config.FunctionName.IsNull() {
			*functionName = r.Config.FunctionName.ValueString()
		} else {
			functionName = nil
		}
		host := new(string)
		if !r.Config.Host.IsUnknown() && !r.Config.Host.IsNull() {
			*host = r.Config.Host.ValueString()
		} else {
			host = nil
		}
		invocationType := new(shared.CreateAWSLambdaPluginInvocationType)
		if !r.Config.InvocationType.IsUnknown() && !r.Config.InvocationType.IsNull() {
			*invocationType = shared.CreateAWSLambdaPluginInvocationType(r.Config.InvocationType.ValueString())
		} else {
			invocationType = nil
		}
		isProxyIntegration := new(bool)
		if !r.Config.IsProxyIntegration.IsUnknown() && !r.Config.IsProxyIntegration.IsNull() {
			*isProxyIntegration = r.Config.IsProxyIntegration.ValueBool()
		} else {
			isProxyIntegration = nil
		}
		keepalive := new(float64)
		if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
			*keepalive, _ = r.Config.Keepalive.ValueBigFloat().Float64()
		} else {
			keepalive = nil
		}
		logType := new(shared.CreateAWSLambdaPluginLogType)
		if !r.Config.LogType.IsUnknown() && !r.Config.LogType.IsNull() {
			*logType = shared.CreateAWSLambdaPluginLogType(r.Config.LogType.ValueString())
		} else {
			logType = nil
		}
		port := new(int64)
		if !r.Config.Port.IsUnknown() && !r.Config.Port.IsNull() {
			*port = r.Config.Port.ValueInt64()
		} else {
			port = nil
		}
		proxyURL := new(string)
		if !r.Config.ProxyURL.IsUnknown() && !r.Config.ProxyURL.IsNull() {
			*proxyURL = r.Config.ProxyURL.ValueString()
		} else {
			proxyURL = nil
		}
		qualifier := new(string)
		if !r.Config.Qualifier.IsUnknown() && !r.Config.Qualifier.IsNull() {
			*qualifier = r.Config.Qualifier.ValueString()
		} else {
			qualifier = nil
		}
		skipLargeBodies := new(bool)
		if !r.Config.SkipLargeBodies.IsUnknown() && !r.Config.SkipLargeBodies.IsNull() {
			*skipLargeBodies = r.Config.SkipLargeBodies.ValueBool()
		} else {
			skipLargeBodies = nil
		}
		timeout := new(float64)
		if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
			*timeout, _ = r.Config.Timeout.ValueBigFloat().Float64()
		} else {
			timeout = nil
		}
		unhandledStatus := new(int64)
		if !r.Config.UnhandledStatus.IsUnknown() && !r.Config.UnhandledStatus.IsNull() {
			*unhandledStatus = r.Config.UnhandledStatus.ValueInt64()
		} else {
			unhandledStatus = nil
		}
		config = &shared.CreateAWSLambdaPluginConfig{
			AwsAssumeRoleArn:       awsAssumeRoleArn,
			AwsImdsProtocolVersion: awsImdsProtocolVersion,
			AwsKey:                 awsKey,
			AwsRegion:              awsRegion,
			AwsRoleSessionName:     awsRoleSessionName,
			AwsSecret:              awsSecret,
			AwsgatewayCompatible:   awsgatewayCompatible,
			Base64EncodeBody:       base64EncodeBody,
			DisableHTTPS:           disableHTTPS,
			ForwardRequestBody:     forwardRequestBody,
			ForwardRequestHeaders:  forwardRequestHeaders,
			ForwardRequestMethod:   forwardRequestMethod,
			ForwardRequestURI:      forwardRequestURI,
			FunctionName:           functionName,
			Host:                   host,
			InvocationType:         invocationType,
			IsProxyIntegration:     isProxyIntegration,
			Keepalive:              keepalive,
			LogType:                logType,
			Port:                   port,
			ProxyURL:               proxyURL,
			Qualifier:              qualifier,
			SkipLargeBodies:        skipLargeBodies,
			Timeout:                timeout,
			UnhandledStatus:        unhandledStatus,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var protocols []shared.CreateAWSLambdaPluginProtocols = []shared.CreateAWSLambdaPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateAWSLambdaPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateAWSLambdaPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateAWSLambdaPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateAWSLambdaPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateAWSLambdaPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateAWSLambdaPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateAWSLambdaPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateAWSLambdaPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateAWSLambdaPluginService{
			ID: id3,
		}
	}
	out := shared.CreateAWSLambdaPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginAWSLambdaResourceModel) RefreshFromSharedAWSLambdaPlugin(resp *shared.AWSLambdaPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAWSLambdaPluginConfig{}
			r.Config.AwsAssumeRoleArn = types.StringPointerValue(resp.Config.AwsAssumeRoleArn)
			if resp.Config.AwsImdsProtocolVersion != nil {
				r.Config.AwsImdsProtocolVersion = types.StringValue(string(*resp.Config.AwsImdsProtocolVersion))
			} else {
				r.Config.AwsImdsProtocolVersion = types.StringNull()
			}
			r.Config.AwsKey = types.StringPointerValue(resp.Config.AwsKey)
			r.Config.AwsRegion = types.StringPointerValue(resp.Config.AwsRegion)
			r.Config.AwsRoleSessionName = types.StringPointerValue(resp.Config.AwsRoleSessionName)
			r.Config.AwsSecret = types.StringPointerValue(resp.Config.AwsSecret)
			r.Config.AwsgatewayCompatible = types.BoolPointerValue(resp.Config.AwsgatewayCompatible)
			r.Config.Base64EncodeBody = types.BoolPointerValue(resp.Config.Base64EncodeBody)
			r.Config.DisableHTTPS = types.BoolPointerValue(resp.Config.DisableHTTPS)
			r.Config.ForwardRequestBody = types.BoolPointerValue(resp.Config.ForwardRequestBody)
			r.Config.ForwardRequestHeaders = types.BoolPointerValue(resp.Config.ForwardRequestHeaders)
			r.Config.ForwardRequestMethod = types.BoolPointerValue(resp.Config.ForwardRequestMethod)
			r.Config.ForwardRequestURI = types.BoolPointerValue(resp.Config.ForwardRequestURI)
			r.Config.FunctionName = types.StringPointerValue(resp.Config.FunctionName)
			r.Config.Host = types.StringPointerValue(resp.Config.Host)
			if resp.Config.InvocationType != nil {
				r.Config.InvocationType = types.StringValue(string(*resp.Config.InvocationType))
			} else {
				r.Config.InvocationType = types.StringNull()
			}
			r.Config.IsProxyIntegration = types.BoolPointerValue(resp.Config.IsProxyIntegration)
			if resp.Config.Keepalive != nil {
				r.Config.Keepalive = types.NumberValue(big.NewFloat(float64(*resp.Config.Keepalive)))
			} else {
				r.Config.Keepalive = types.NumberNull()
			}
			if resp.Config.LogType != nil {
				r.Config.LogType = types.StringValue(string(*resp.Config.LogType))
			} else {
				r.Config.LogType = types.StringNull()
			}
			r.Config.Port = types.Int64PointerValue(resp.Config.Port)
			r.Config.ProxyURL = types.StringPointerValue(resp.Config.ProxyURL)
			r.Config.Qualifier = types.StringPointerValue(resp.Config.Qualifier)
			r.Config.SkipLargeBodies = types.BoolPointerValue(resp.Config.SkipLargeBodies)
			if resp.Config.Timeout != nil {
				r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
			} else {
				r.Config.Timeout = types.NumberNull()
			}
			r.Config.UnhandledStatus = types.Int64PointerValue(resp.Config.UnhandledStatus)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
