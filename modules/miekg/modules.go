package miekg

import (
	"github.com/miekg/dns"
	"github.com/zmap/zdns"
)

// let's register some modules!
func init() {
	a := new(GlobalLookupFactory)
	a.SetDNSType(dns.TypeA)
	zdns.RegisterLookup("A", a)

	aaaa := new(GlobalLookupFactory)
	aaaa.SetDNSType(dns.TypeAAAA)
	zdns.RegisterLookup("AAAA", aaaa)

	any := new(GlobalLookupFactory)
	any.SetDNSType(dns.TypeANY)
	zdns.RegisterLookup("ANY", any)

	afsdb := new(GlobalLookupFactory)
	afsdb.SetDNSType(dns.TypeAFSDB)
	zdns.RegisterLookup("AFSDB", afsdb)

	atma := new(GlobalLookupFactory)
	atma.SetDNSType(dns.TypeATMA)
	zdns.RegisterLookup("ATMA", atma)

	avc := new(GlobalLookupFactory)
	avc.SetDNSType(dns.TypeAVC)
	zdns.RegisterLookup("AVC", avc)

	caa := new(GlobalLookupFactory)
	caa.SetDNSType(dns.TypeCAA)
	zdns.RegisterLookup("CAA", caa)

	cert := new(GlobalLookupFactory)
	cert.SetDNSType(dns.TypeCERT)
	zdns.RegisterLookup("CERT", cert)

	cds := new(GlobalLookupFactory)
	cds.SetDNSType(dns.TypeCDS)
	zdns.RegisterLookup("CDS", cds)

	cdnskey := new(GlobalLookupFactory)
	cdnskey.SetDNSType(dns.TypeCDNSKEY)
	zdns.RegisterLookup("CDNSKEY", cdnskey)

	cname := new(GlobalLookupFactory)
	cname.SetDNSType(dns.TypeCNAME)
	zdns.RegisterLookup("CNAME", cname)

	csync := new(GlobalLookupFactory)
	csync.SetDNSType(dns.TypeCSYNC)
	zdns.RegisterLookup("CSYNC", csync)

	dhcid := new(GlobalLookupFactory)
	dhcid.SetDNSType(dns.TypeDHCID)
	zdns.RegisterLookup("DHCID", dhcid)

	dnskey := new(GlobalLookupFactory)
	dnskey.SetDNSType(dns.TypeDNSKEY)
	zdns.RegisterLookup("DNSKEY", dnskey)

	ds := new(GlobalLookupFactory)
	ds.SetDNSType(dns.TypeDS)
	zdns.RegisterLookup("DS", ds)

	eid := new(GlobalLookupFactory)
	eid.SetDNSType(dns.TypeEID)
	zdns.RegisterLookup("EID", eid)

	eui48 := new(GlobalLookupFactory)
	eui48.SetDNSType(dns.TypeEUI48)
	zdns.RegisterLookup("EUI48", eui48)

	eui64 := new(GlobalLookupFactory)
	eui64.SetDNSType(dns.TypeEUI64)
	zdns.RegisterLookup("EUI64", eui64)

	gid := new(GlobalLookupFactory)
	gid.SetDNSType(dns.TypeGID)
	zdns.RegisterLookup("GID", gid)

	gpos := new(GlobalLookupFactory)
	gpos.SetDNSType(dns.TypeGPOS)
	zdns.RegisterLookup("GPOS", gpos)

	hinfo := new(GlobalLookupFactory)
	hinfo.SetDNSType(dns.TypeHINFO)
	zdns.RegisterLookup("HINFO", hinfo)

	hip := new(GlobalLookupFactory)
	hip.SetDNSType(dns.TypeHIP)
	zdns.RegisterLookup("HIP", hip)

	isdn := new(GlobalLookupFactory)
	isdn.SetDNSType(dns.TypeISDN)
	zdns.RegisterLookup("ISDN", isdn)

	key := new(GlobalLookupFactory)
	key.SetDNSType(dns.TypeKEY)
	zdns.RegisterLookup("KEY", key)

	kx := new(GlobalLookupFactory)
	kx.SetDNSType(dns.TypeKX)
	zdns.RegisterLookup("KX", kx)

	l32 := new(GlobalLookupFactory)
	l32.SetDNSType(dns.TypeL32)
	zdns.RegisterLookup("L32", l32)

	l64 := new(GlobalLookupFactory)
	l64.SetDNSType(dns.TypeL64)
	zdns.RegisterLookup("L64", l64)

	loc := new(GlobalLookupFactory)
	loc.SetDNSType(dns.TypeLOC)
	zdns.RegisterLookup("LOC", loc)

	lp := new(GlobalLookupFactory)
	lp.SetDNSType(dns.TypeLP)
	zdns.RegisterLookup("LP", lp)

	md := new(GlobalLookupFactory)
	md.SetDNSType(dns.TypeMD)
	zdns.RegisterLookup("MD", md)

	mf := new(GlobalLookupFactory)
	mf.SetDNSType(dns.TypeMF)
	zdns.RegisterLookup("MF", mf)

	mb := new(GlobalLookupFactory)
	mb.SetDNSType(dns.TypeMB)
	zdns.RegisterLookup("MB", mb)

	mg := new(GlobalLookupFactory)
	mg.SetDNSType(dns.TypeMG)
	zdns.RegisterLookup("MG", mg)

	mr := new(GlobalLookupFactory)
	mr.SetDNSType(dns.TypeMR)
	zdns.RegisterLookup("MR", mr)

	mx := new(GlobalLookupFactory)
	mx.SetDNSType(dns.TypeMX)
	zdns.RegisterLookup("MX", mx)

	naptr := new(GlobalLookupFactory)
	naptr.SetDNSType(dns.TypeNAPTR)
	zdns.RegisterLookup("NAPTR", naptr)

	nimloc := new(GlobalLookupFactory)
	nimloc.SetDNSType(dns.TypeNIMLOC)
	zdns.RegisterLookup("NS", nimloc)

	nid := new(GlobalLookupFactory)
	nid.SetDNSType(dns.TypeNID)
	zdns.RegisterLookup("NID", nid)

	ninfo := new(GlobalLookupFactory)
	ninfo.SetDNSType(dns.TypeNINFO)
	zdns.RegisterLookup("NINFO", ninfo)

	nsapptr := new(GlobalLookupFactory)
	nsapptr.SetDNSType(dns.TypeNSAPPTR)
	zdns.RegisterLookup("NSAPPTR", nsapptr)

	ns := new(GlobalLookupFactory)
	ns.SetDNSType(dns.TypeNS)
	zdns.RegisterLookup("NS", ns)

	nxt := new(GlobalLookupFactory)
	nxt.SetDNSType(dns.TypeNXT)
	zdns.RegisterLookup("NXT", nxt)

	nsec := new(GlobalLookupFactory)
	nsec.SetDNSType(dns.TypeNSEC)
	zdns.RegisterLookup("NSEC", nsec)

	nsec3 := new(GlobalLookupFactory)
	nsec3.SetDNSType(dns.TypeNSEC3)
	zdns.RegisterLookup("NSEC3", nsec3)

	nsec3param := new(GlobalLookupFactory)
	nsec3param.SetDNSType(dns.TypeNSEC3PARAM)
	zdns.RegisterLookup("NSEC3PARAM", nsec3param)

	null := new(GlobalLookupFactory)
	null.SetDNSType(dns.TypeNULL)
	zdns.RegisterLookup("NULL", null)

	openpgpkey := new(GlobalLookupFactory)
	openpgpkey.SetDNSType(dns.TypeOPENPGPKEY)
	zdns.RegisterLookup("OPENPGPKEY", openpgpkey)

	opt := new(GlobalLookupFactory)
	opt.SetDNSType(dns.TypeOPT)
	zdns.RegisterLookup("OPT", opt)

	ptr := new(GlobalLookupFactory)
	ptr.SetDNSType(dns.TypePTR)
	zdns.RegisterLookup("PTR", ptr)

	px := new(GlobalLookupFactory)
	px.SetDNSType(dns.TypePX)
	zdns.RegisterLookup("PX", px)

	rp := new(GlobalLookupFactory)
	rp.SetDNSType(dns.TypeRP)
	zdns.RegisterLookup("RP", rp)

	rrsig := new(GlobalLookupFactory)
	rrsig.SetDNSType(dns.TypeRRSIG)
	zdns.RegisterLookup("RRSIG", rrsig)

	rt := new(GlobalLookupFactory)
	rt.SetDNSType(dns.TypeRT)
	zdns.RegisterLookup("RT", rt)

	smimea := new(GlobalLookupFactory)
	smimea.SetDNSType(dns.TypeSMIMEA)
	zdns.RegisterLookup("SMIMEA", smimea)

	sshfp := new(GlobalLookupFactory)
	sshfp.SetDNSType(dns.TypeSSHFP)
	zdns.RegisterLookup("SSHFP", sshfp)

	soa := new(GlobalLookupFactory)
	soa.SetDNSType(dns.TypeSOA)
	zdns.RegisterLookup("SOA", soa)

	spf := new(GlobalLookupFactory)
	spf.SetDNSType(dns.TypeSPF)
	zdns.RegisterLookup("SPF", spf)

	srv := new(GlobalLookupFactory)
	srv.SetDNSType(dns.TypeSRV)
	zdns.RegisterLookup("SRV", srv)

	talink := new(GlobalLookupFactory)
	talink.SetDNSType(dns.TypeTALINK)
	zdns.RegisterLookup("TALINK", talink)

	tlsa := new(GlobalLookupFactory)
	tlsa.SetDNSType(dns.TypeTLSA)
	zdns.RegisterLookup("TLSA", tlsa)

	txt := new(GlobalLookupFactory)
	txt.SetDNSType(dns.TypeTXT)
	zdns.RegisterLookup("TXT", txt)

	uid := new(GlobalLookupFactory)
	uid.SetDNSType(dns.TypeUID)
	zdns.RegisterLookup("UID", uid)

	uinfo := new(GlobalLookupFactory)
	uinfo.SetDNSType(dns.TypeUINFO)
	zdns.RegisterLookup("UINFO", uinfo)

	unspec := new(GlobalLookupFactory)
	unspec.SetDNSType(dns.TypeUNSPEC)
	zdns.RegisterLookup("UNSPEC", unspec)

	uri := new(GlobalLookupFactory)
	uri.SetDNSType(dns.TypeURI)
	zdns.RegisterLookup("URI", uri)

}
