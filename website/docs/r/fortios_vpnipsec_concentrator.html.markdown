---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_concentrator"
description: |-
  Concentrator configuration.
---

# fortios_vpnipsec_concentrator
Concentrator configuration.

## Example Usage

```hcl
resource "fortios_vpnipsec_concentrator" "trname" {
  name      = "1"
  src_check = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Concentrator ID. (1-65535)
* `name` - Concentrator name.
* `src_check` - Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
* `member` - Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Member name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnIpsec Concentrator can be imported using any of these accepted formats:
```
$ terraform import fortios_vpnipsec_concentrator.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnipsec_concentrator.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
