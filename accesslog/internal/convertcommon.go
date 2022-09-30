package internal

import (
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func ConvertCommon(envoy *datav3.AccessLogCommon) *md.AccessLogCommon {
	if envoy == nil {
		return nil
	}
	common := new(md.AccessLogCommon)
	common.SampleRate = envoy.GetSampleRate()
	common.DownstreamRemoteAddress = ConvertAddress(envoy.GetDownstreamRemoteAddress())
	common.DownstreamLocalAddress = ConvertAddress(envoy.GetDownstreamLocalAddress())

	common.TlsProperties = ConvertTls(envoy.GetTlsProperties())
	if envoy.GetStartTime() != nil {
		t := envoy.GetStartTime().AsTime()
		common.StartTime = &t
	}

	if envoy.GetTimeToLastRxByte() != nil {
		t := envoy.GetTimeToLastRxByte().AsDuration()
		common.TimeToLastRxByte = &t
	}
	if envoy.GetTimeToFirstUpstreamTxByte() != nil {
		t := envoy.GetTimeToFirstUpstreamTxByte().AsDuration()
		common.TimeToFirstUpstreamTxByte = &t
	}
	if envoy.GetTimeToLastUpstreamTxByte() != nil {
		t := envoy.GetTimeToLastUpstreamTxByte().AsDuration()
		common.TimeToLastUpstreamTxByte = &t
	}
	if envoy.GetTimeToFirstUpstreamRxByte() != nil {
		t := envoy.GetTimeToFirstUpstreamRxByte().AsDuration()
		common.TimeToFirstUpstreamRxByte = &t
	}
	if envoy.GetTimeToLastUpstreamRxByte() != nil {
		t := envoy.GetTimeToLastUpstreamRxByte().AsDuration()
		common.TimeToLastUpstreamRxByte = &t
	}
	if envoy.GetTimeToFirstDownstreamTxByte() != nil {
		t := envoy.GetTimeToFirstDownstreamTxByte().AsDuration()
		common.TimeToFirstDownstreamTxByte = &t
	}
	if envoy.GetTimeToLastDownstreamTxByte() != nil {
		t := envoy.GetTimeToLastDownstreamTxByte().AsDuration()
		common.TimeToLastDownstreamTxByte = &t
	}

	common.UpstreamRemoteAddress = ConvertAddress(envoy.GetUpstreamRemoteAddress())
	common.UpstreamLocalAddress = ConvertAddress(envoy.GetUpstreamLocalAddress())
	common.UpstreamCluster = envoy.GetUpstreamCluster()

	common.ResponseFlags = ConvertResponseFlags(envoy.GetResponseFlags())
	common.UpstreamTransportFailureReason = envoy.GetUpstreamTransportFailureReason()
	common.RouteName = envoy.GetRouteName()
	common.DownstreamDirectRemoteAddress = ConvertAddress(envoy.GetDownstreamRemoteAddress())

	common.CustomTags = envoy.GetCustomTags()

	// BUG : v3 common still the same as v2
	//if envoy.GetDuration() != nil {
	//	t := envoy.GetDuration().AsTime()
	//	common.Duration = &t
	//}
	//common.UpstreamRequestAttemptCount  = envoy.UpstreamRequestAttemptCount
	//common.ConnectionTerminationDetails = envoy.GetConnectionTerminationDetails()
	return common
}

func ConvertTls(envoy *datav3.TLSProperties) *md.TLSProperties {
	if envoy == nil {
		return nil
	}

	tls := new(md.TLSProperties)
	tls.TlsVersion = md.TLSProperties_TLSVersion(envoy.GetTlsVersion())
	tls.TlsCipherSuite = envoy.GetTlsCipherSuite().Value
	tls.TlsSniHostname = envoy.GetTlsSniHostname()
	if envoy.GetLocalCertificateProperties() != nil {
		//ea := envoy.GetLocalCertificateProperties().GetSubject()
	}
	tls.TlsSessionId = envoy.GetTlsSessionId()
	// BUG : missing Ja3Fingerprint
	//tls.Ja3Fingerprint = envoy.GetJa3Fingerprint()
	return tls
}

func ConvertResponseFlags(envoy *datav3.ResponseFlags) *md.ResponseFlags {
	if envoy == nil {
		return nil
	}
	resp := new(md.ResponseFlags)
	resp.Encoded = EncodeResponseFlags(envoy)
	if envoy.GetUnauthorizedDetails() != nil {
		resp.UnauthorizedDetails = new(md.ResponseFlags_Unauthorized)
		resp.UnauthorizedDetails.Reason = md.ResponseFlags_Unauthorized_Reason(envoy.GetUnauthorizedDetails().Reason)
	}
	return resp
}

func ConvertAddress(envoy *corev3.Address) *md.Address {
	if envoy == nil {
		return nil
	}
	address := new(md.Address)
	/*
		if envoy.GetSocketAddress() != nil {

			ea := envoy.GetSocketAddress()
			address.SocketAddress = new(md.SocketAddress)
			address.SocketAddress.Protocol = md.SocketAddress_Protocol(ea.GetProtocol())
			address.SocketAddress.Address = ea.GetAddress()
			if ea.GetNamedPort() != "" {
				address.SocketAddress.PortSpecifier = ea.GetNamedPort()
			} else {
				address.SocketAddress.PortSpecifier = strconv.Itoa(int(ea.GetPortValue()))
			}
			address.SocketAddress.ResolverName = ea.GetResolverName()
			address.SocketAddress.Ipv4Compat = ea.GetIpv4Compat()
			return address
		}
		if envoy.GetPipe() != nil {
			address.Type = md.Address_Pipe
			ea := envoy.GetPipe()
			address.Pipe = new(md.Pipe)
			address.Pipe.Path = ea.GetPath()
			address.Pipe.Mode = ea.GetMode()
			return address
		}
		if envoy.GetEnvoyInternalAddress() != nil {
			ea := envoy.GetEnvoyInternalAddress()
			address.EnvoyInternalAddress = new(md.EnvoyInternalAddress)
			address.EnvoyInternalAddress.ServerListenerName = ea.GetServerListenerName()
			// BUG : v3 address still same as v2
			//address.EnvoyInternalAddress.EndpointId = ea.EndpointId
		}

	*/
	return address
}
