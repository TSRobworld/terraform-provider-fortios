---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_nat64"
description: |-
  Configure NAT64.
---

# fortios_system_nat64
Configure NAT64. Applies to FortiOS Version `<= 7.0.0`.

## Example Usage

```hcl
resource "fortios_system_nat64" "trname" {
  always_synthesize_aaaa_record      = "enable"
  generate_ipv6_fragment_header      = "disable"
  nat46_force_ipv4_packet_forwarding = "disable"
  nat64_prefix                       = "2001:1:2:3::/96"
  secondary_prefix_status            = "disable"
  status                             = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
* `nat64_prefix` - (Required) NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
* `secondary_prefix_status` - Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
* `secondary_prefix` - Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
* `always_synthesize_aaaa_record` - Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
* `generate_ipv6_fragment_header` - Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
* `nat46_force_ipv4_packet_forwarding` - Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `secondary_prefix` block supports:

* `name` - NAT64 prefix name.
* `nat64_prefix` - NAT64 prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Nat64 can be imported using any of these accepted formats:
```
$ terraform import fortios_system_nat64.labelname SystemNat64

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_nat64.labelname SystemNat64
$ unset "FORTIOS_IMPORT_TABLE"
```
