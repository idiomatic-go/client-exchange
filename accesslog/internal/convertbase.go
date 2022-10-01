package internal

import (
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

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

func ConvertTls(envoy *datav3.TLSProperties) *md.TLSProperties {
	if envoy == nil {
		return nil
	}

	tls := new(md.TLSProperties)
	tls.TlsVersion = md.TLSProperties_TLSVersion(envoy.GetTlsVersion())
	if envoy.GetTlsCipherSuite() != nil {
		tls.TlsCipherSuite = envoy.GetTlsCipherSuite().Value
	}
	tls.TlsSniHostname = envoy.GetTlsSniHostname()
	if envoy.GetLocalCertificateProperties() != nil {
		//ea := envoy.GetLocalCertificateProperties().GetSubject()
	}
	tls.TlsSessionId = envoy.GetTlsSessionId()
	// BUG : missing Ja3Fingerprint
	//tls.Ja3Fingerprint = envoy.GetJa3Fingerprint()
	return tls
}
