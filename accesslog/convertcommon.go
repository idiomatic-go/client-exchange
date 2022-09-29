package accesslog

import (
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
	"strconv"
)

func convertCommon(log *md.AccessLogCommon, envoy *datav3.AccessLogCommon) {
	log.SampleRate = envoy.GetSampleRate()
	log.DownstreamRemoteAddress = convertAddress(envoy.GetDownstreamRemoteAddress())
	log.DownstreamLocalAddress = convertAddress(envoy.GetDownstreamLocalAddress())

	log.TlsProperties = convertTls(envoy.GetTlsProperties())
	if envoy.GetStartTime() != nil {
		t := envoy.GetStartTime().AsTime()
		log.StartTime = &t
	}

	if envoy.GetTimeToLastRxByte() != nil {
		t := envoy.GetTimeToLastRxByte().AsDuration()
		log.TimeToLastRxByte = &t
	}
	if envoy.GetTimeToFirstUpstreamTxByte() != nil {
		t := envoy.GetTimeToFirstUpstreamTxByte().AsDuration()
		log.TimeToFirstUpstreamTxByte = &t
	}
	if envoy.GetTimeToLastUpstreamTxByte() != nil {
		t := envoy.GetTimeToLastUpstreamTxByte().AsDuration()
		log.TimeToLastUpstreamTxByte = &t
	}
	if envoy.GetTimeToFirstUpstreamRxByte() != nil {
		t := envoy.GetTimeToFirstUpstreamRxByte().AsDuration()
		log.TimeToFirstUpstreamRxByte = &t
	}
	if envoy.GetTimeToLastUpstreamRxByte() != nil {
		t := envoy.GetTimeToLastUpstreamRxByte().AsDuration()
		log.TimeToLastUpstreamRxByte = &t
	}
	if envoy.GetTimeToFirstDownstreamTxByte() != nil {
		t := envoy.GetTimeToFirstDownstreamTxByte().AsDuration()
		log.TimeToFirstDownstreamTxByte = &t
	}
	if envoy.GetTimeToLastDownstreamTxByte() != nil {
		t := envoy.GetTimeToLastDownstreamTxByte().AsDuration()
		log.TimeToLastDownstreamTxByte = &t
	}

	log.UpstreamRemoteAddress = convertAddress(envoy.GetUpstreamRemoteAddress())
	log.UpstreamLocalAddress = convertAddress(envoy.GetUpstreamLocalAddress())
	log.UpstreamCluster = envoy.GetUpstreamCluster()

	convertResponseFlags(log, envoy.GetResponseFlags())
	log.UpstreamTransportFailureReason = envoy.GetUpstreamTransportFailureReason()
	log.RouteName = envoy.GetRouteName()
	log.DownstreamDirectRemoteAddress = convertAddress(envoy.GetDownstreamRemoteAddress())

	log.CustomTags = envoy.GetCustomTags()

	// BUG : v3 common still the same as v2
	//if envoy.GetDuration() != nil {
	//	t := envoy.GetDuration().AsTime()
	//	common.Duration = &t
	//}
	//log.UpstreamRequestAttemptCount  = envoy.UpstreamRequestAttemptCount
	//log.ConnectionTerminationDetails = envoy.GetConnectionTerminationDetails()

}

func convertAddress(envoy *corev3.Address) *md.Address {
	if envoy == nil {
		return nil
	}
	address := new(md.Address)
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
	return address
}

func convertTls(envoy *datav3.TLSProperties) *md.TLSProperties {
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
