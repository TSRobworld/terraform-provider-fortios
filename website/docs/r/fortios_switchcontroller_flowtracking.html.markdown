---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_flowtracking"
description: |-
  Configure FortiSwitch flow tracking and export via ipfix/netflow.
---

# fortios_switchcontroller_flowtracking
Configure FortiSwitch flow tracking and export via ipfix/netflow. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `sample_mode` - Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
* `sample_rate` - Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
* `format` - Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
* `collectors` - Configure collectors for the flow. The structure of `collectors` block is documented below.
* `collector_ip` - Configure collector ip address.
* `collector_port` - Configure collector port number(0-65535, default=0).
* `transport` - Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
* `level` - Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
* `max_export_pkt_size` - Configure flow max export packet size (512-9216, default=512 bytes).
* `template_export_period` - Configure template export period (1-60, default=5 minutes).
* `timeout_general` - Configure flow session general timeout (60-604800, default=3600 seconds).
* `timeout_icmp` - Configure flow session ICMP timeout (60-604800, default=300 seconds).
* `timeout_max` - Configure flow session max timeout (60-604800, default=604800 seconds).
* `timeout_tcp` - Configure flow session TCP timeout (60-604800, default=3600 seconds).
* `timeout_tcp_fin` - Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
* `timeout_tcp_rst` - Configure flow session TCP RST timeout (60-604800, default=120 seconds).
* `timeout_udp` - Configure flow session UDP timeout (60-604800, default=300 seconds).
* `aggregates` - Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `collectors` block supports:

* `name` - Collector name.
* `ip` - Collector IP address.
* `port` - Collector port number(0-65535, default:0, netflow:2055, ipfix:4739).
* `transport` - Collector L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.

The `aggregates` block supports:

* `id` - Aggregate id.
* `ip` - IP address to group all matching traffic sessions to a flow.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController FlowTracking can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_flowtracking.labelname SwitchControllerFlowTracking

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_flowtracking.labelname SwitchControllerFlowTracking
$ unset "FORTIOS_IMPORT_TABLE"
```
