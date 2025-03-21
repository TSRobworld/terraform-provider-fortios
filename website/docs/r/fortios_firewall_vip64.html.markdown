---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vip64"
description: |-
  Configure IPv6 to IPv4 virtual IPs.
---

# fortios_firewall_vip64
Configure IPv6 to IPv4 virtual IPs. Applies to FortiOS Version `<= 7.0.0`.

## Example Usage

```hcl
resource "fortios_firewall_vip64" "trname" {
  arp_reply   = "enable"
  color       = 0
  extip       = "2001:db8:99:203::22"
  extport     = "0-65535"
  fosid       = 0
  ldb_method  = "static"
  mappedip    = "1.1.1.1"
  mappedport  = "0-65535"
  name        = "vip64s1"
  portforward = "disable"
  protocol    = "tcp"
  type        = "static-nat"
}
```

## Argument Reference

The following arguments are supported:

* `name` - VIP64 name.
* `fosid` - Custom defined id.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `comment` - Comment.
* `type` - VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
* `src_filter` - Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `src_filter` block is documented below.
* `extip` - (Required) Start-external-IP [-end-external-IP].
* `mappedip` - (Required) Start-mapped-IP [-end-mapped-IP].
* `arp_reply` - Enable ARP reply. Valid values: `disable`, `enable`.
* `portforward` - Enable port forwarding. Valid values: `disable`, `enable`.
* `protocol` - Mapped port protocol. Valid values: `tcp`, `udp`.
* `extport` - External service port.
* `mappedport` - Mapped service port.
* `color` - Color of icon on the GUI.
* `ldb_method` - Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
* `server_type` - Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
* `realservers` - Real servers. The structure of `realservers` block is documented below.
* `monitor` - Health monitors. The structure of `monitor` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `src_filter` block supports:

* `range` - Src-filter range.

The `realservers` block supports:

* `id` - Real server ID.
* `ip` - Mapped server IP.
* `port` - Mapped server port.
* `status` - Server administrative status. Valid values: `active`, `standby`, `disable`.
* `weight` - weight
* `holddown_interval` - Hold down interval.
* `healthcheck` - Per server health check. Valid values: `disable`, `enable`, `vip`.
* `max_connections` - Maximum number of connections allowed to server.
* `monitor` - Health monitors.
* `client_ip` - Restrict server to a client IP in this range.

The `monitor` block supports:

* `name` - Health monitor name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vip64 can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_vip64.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_vip64.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
