---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_userbookmark"
description: |-
  Configure SSL VPN user bookmark.
---

# fortios_vpnsslweb_userbookmark
Configure SSL VPN user bookmark.

## Example Usage

```hcl
resource "fortios_vpnsslweb_userbookmark" "trname" {
  custom_lang = "big5"
  name        = "userbookmarks1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - User and group name.
* `custom_lang` - Personal language.
* `bookmarks` - Bookmark table. The structure of `bookmarks` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `bookmarks` block supports:

* `name` - Bookmark name.
* `apptype` - Application type.
* `url` - URL parameter.
* `host` - Host name/IP parameter.
* `folder` - Network shared file folder parameter.
* `domain` - Login domain.
* `additional_params` - Additional parameters.
* `listening_port` - Listening port (0 - 65535).
* `remote_port` - Remote port (0 - 65535).
* `show_status_window` - Enable/disable showing of status window. Valid values: `enable`, `disable`.
* `description` - Description.
* `keyboard_layout` - Keyboard layout.
* `server_layout` - Server side keyboard layout.
* `security` - Security mode for RDP connection. Valid values: `rdp`, `nla`, `tls`, `any`.
* `send_preconnection_id` - Enable/disable sending of preconnection ID. Valid values: `enable`, `disable`.
* `preconnection_id` - The numeric ID of the RDP source (0-2147483648).
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `restricted_admin` - Enable/disable restricted admin mode for RDP. Valid values: `enable`, `disable`.
* `port` - Remote port.
* `logon_user` - Logon user.
* `logon_password` - Logon password.
* `color_depth` - Color depth per pixel. Valid values: `32`, `16`, `8`.
* `sso` - Single Sign-On. Valid values: `disable`, `static`, `auto`.
* `form_data` - Form data. The structure of `form_data` block is documented below.
* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.
* `sso_username` - SSO user name.
* `sso_password` - SSO password.
* `sso_credential_sent_once` - Single sign-on credentials are only sent once to remote server. Valid values: `enable`, `disable`.
* `width` - Screen width (range from 640 - 65535, default = 1024).
* `height` - Screen height (range from 480 - 65535, default = 768).

The `form_data` block supports:

* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnSslWeb UserBookmark can be imported using any of these accepted formats:
```
$ terraform import fortios_vpnsslweb_userbookmark.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnsslweb_userbookmark.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
