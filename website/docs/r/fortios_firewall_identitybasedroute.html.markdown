---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_identitybasedroute"
description: |-
  Configure identity based routing.
---

# fortios_firewall_identitybasedroute
Configure identity based routing.

## Example Usage

```hcl
resource "fortios_firewall_identitybasedroute" "trname" {
  comments = "test"
  name     = "route1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Comments.
* `rule` - Rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `rule` block supports:

* `id` - Rule ID.
* `gateway` - IPv4 address of the gateway (Format: xxx.xxx.xxx.xxx , Default: 0.0.0.0).
* `device` - Outgoing interface for the rule.
* `groups` - Select one or more group(s) from available groups that are allowed to use this route. Separate group names with a space. The structure of `groups` block is documented below.

The `groups` block supports:

* `name` - Group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall IdentityBasedRoute can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_identitybasedroute.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_identitybasedroute.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
