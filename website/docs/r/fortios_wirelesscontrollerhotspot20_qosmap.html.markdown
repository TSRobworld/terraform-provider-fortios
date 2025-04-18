---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_qosmap"
description: |-
  Configure QoS map set.
---

# fortios_wirelesscontrollerhotspot20_qosmap
Configure QoS map set.

## Argument Reference

The following arguments are supported:

* `name` - QOS-MAP name.
* `dscp_except` - Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.
* `dscp_range` - Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `dscp_except` block supports:

* `index` - DSCP exception index.
* `dscp` - DSCP value.
* `up` - User priority.

The `dscp_range` block supports:

* `index` - DSCP range index.
* `up` - User priority.
* `low` - DSCP low value.
* `high` - DSCP high value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 QosMap can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_qosmap.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_qosmap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
