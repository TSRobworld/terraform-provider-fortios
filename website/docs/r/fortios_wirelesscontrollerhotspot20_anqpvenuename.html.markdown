---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpvenuename"
description: |-
  Configure venue name duple.
---

# fortios_wirelesscontrollerhotspot20_anqpvenuename
Configure venue name duple.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_anqpvenuename" "trname" {
  name = "1"
  value_list {
    index = 1
    lang  = "CN"
    value = "3"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of venue name duple.
* `value_list` - Name list. The structure of `value_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `value_list` block supports:

* `index` - Value index.
* `lang` - Language code.
* `value` - Venue name value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpVenueName can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_anqpvenuename.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_anqpvenuename.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
