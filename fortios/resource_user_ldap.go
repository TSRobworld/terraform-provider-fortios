// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure LDAP server entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserLdap() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserLdapCreate,
		Read:   resourceUserLdapRead,
		Update: resourceUserLdapUpdate,
		Delete: resourceUserLdapDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required: true,
			},
			"secondary_server": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"tertiary_server": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"server_identity_check": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cnid": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 20),
				Optional: true,
				Computed: true,
			},
			"dn": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required: true,
			},
			"type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional: true,
			},
			"group_member_check": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_search_base": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional: true,
				Computed: true,
			},
			"group_object_filter": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional: true,
				Computed: true,
			},
			"group_filter": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional: true,
				Computed: true,
			},
			"secure": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cert": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional: true,
				Computed: true,
			},
			"password_expiry_warning": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_renewal": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member_attr": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"account_key_processing": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_filter": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional: true,
				Computed: true,
			},
			"search_type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserLdapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserLdap(d)
	if err != nil {
		return fmt.Errorf("Error creating UserLdap resource while getting object: %v", err)
	}

	o, err := c.CreateUserLdap(obj)

	if err != nil {
		return fmt.Errorf("Error creating UserLdap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserLdap(d)
	if err != nil {
		return fmt.Errorf("Error updating UserLdap resource while getting object: %v", err)
	}

	o, err := c.UpdateUserLdap(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserLdap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteUserLdap(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserLdap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLdapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserLdap(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserLdap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserLdap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserLdap resource from API: %v", err)
	}
	return nil
}


func flattenUserLdapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapCnid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapDn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapGroupMemberCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapGroupSearchBase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapGroupObjectFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapGroupFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapPasswordExpiryWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapMemberAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapAccountKeyProcessing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapAccountKeyFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserLdapSearchType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectUserLdap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenUserLdapName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenUserLdapServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenUserLdapSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if !fortiAPIPatch(o["secondary-server"]) {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenUserLdapTertiaryServer(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if !fortiAPIPatch(o["tertiary-server"]) {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenUserLdapServerIdentityCheck(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if !fortiAPIPatch(o["server-identity-check"]) {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserLdapSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("cnid", flattenUserLdapCnid(o["cnid"], d, "cnid")); err != nil {
		if !fortiAPIPatch(o["cnid"]) {
			return fmt.Errorf("Error reading cnid: %v", err)
		}
	}

	if err = d.Set("dn", flattenUserLdapDn(o["dn"], d, "dn")); err != nil {
		if !fortiAPIPatch(o["dn"]) {
			return fmt.Errorf("Error reading dn: %v", err)
		}
	}

	if err = d.Set("type", flattenUserLdapType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username", flattenUserLdapUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("password", flattenUserLdapPassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("group_member_check", flattenUserLdapGroupMemberCheck(o["group-member-check"], d, "group_member_check")); err != nil {
		if !fortiAPIPatch(o["group-member-check"]) {
			return fmt.Errorf("Error reading group_member_check: %v", err)
		}
	}

	if err = d.Set("group_search_base", flattenUserLdapGroupSearchBase(o["group-search-base"], d, "group_search_base")); err != nil {
		if !fortiAPIPatch(o["group-search-base"]) {
			return fmt.Errorf("Error reading group_search_base: %v", err)
		}
	}

	if err = d.Set("group_object_filter", flattenUserLdapGroupObjectFilter(o["group-object-filter"], d, "group_object_filter")); err != nil {
		if !fortiAPIPatch(o["group-object-filter"]) {
			return fmt.Errorf("Error reading group_object_filter: %v", err)
		}
	}

	if err = d.Set("group_filter", flattenUserLdapGroupFilter(o["group-filter"], d, "group_filter")); err != nil {
		if !fortiAPIPatch(o["group-filter"]) {
			return fmt.Errorf("Error reading group_filter: %v", err)
		}
	}

	if err = d.Set("secure", flattenUserLdapSecure(o["secure"], d, "secure")); err != nil {
		if !fortiAPIPatch(o["secure"]) {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenUserLdapSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenUserLdapCaCert(o["ca-cert"], d, "ca_cert")); err != nil {
		if !fortiAPIPatch(o["ca-cert"]) {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("port", flattenUserLdapPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("password_expiry_warning", flattenUserLdapPasswordExpiryWarning(o["password-expiry-warning"], d, "password_expiry_warning")); err != nil {
		if !fortiAPIPatch(o["password-expiry-warning"]) {
			return fmt.Errorf("Error reading password_expiry_warning: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenUserLdapPasswordRenewal(o["password-renewal"], d, "password_renewal")); err != nil {
		if !fortiAPIPatch(o["password-renewal"]) {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("member_attr", flattenUserLdapMemberAttr(o["member-attr"], d, "member_attr")); err != nil {
		if !fortiAPIPatch(o["member-attr"]) {
			return fmt.Errorf("Error reading member_attr: %v", err)
		}
	}

	if err = d.Set("account_key_processing", flattenUserLdapAccountKeyProcessing(o["account-key-processing"], d, "account_key_processing")); err != nil {
		if !fortiAPIPatch(o["account-key-processing"]) {
			return fmt.Errorf("Error reading account_key_processing: %v", err)
		}
	}

	if err = d.Set("account_key_filter", flattenUserLdapAccountKeyFilter(o["account-key-filter"], d, "account_key_filter")); err != nil {
		if !fortiAPIPatch(o["account-key-filter"]) {
			return fmt.Errorf("Error reading account_key_filter: %v", err)
		}
	}

	if err = d.Set("search_type", flattenUserLdapSearchType(o["search-type"], d, "search_type")); err != nil {
		if !fortiAPIPatch(o["search-type"]) {
			return fmt.Errorf("Error reading search_type: %v", err)
		}
	}


	return nil
}

func flattenUserLdapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandUserLdapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCnid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupMemberCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupSearchBase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupObjectFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordExpiryWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapMemberAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyProcessing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSearchType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectUserLdap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserLdapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandUserLdapServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok {
		t, err := expandUserLdapSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok {
		t, err := expandUserLdapTertiaryServer(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok {
		t, err := expandUserLdapServerIdentityCheck(d, v, "server_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandUserLdapSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("cnid"); ok {
		t, err := expandUserLdapCnid(d, v, "cnid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cnid"] = t
		}
	}

	if v, ok := d.GetOk("dn"); ok {
		t, err := expandUserLdapDn(d, v, "dn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dn"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandUserLdapType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandUserLdapUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandUserLdapPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("group_member_check"); ok {
		t, err := expandUserLdapGroupMemberCheck(d, v, "group_member_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-member-check"] = t
		}
	}

	if v, ok := d.GetOk("group_search_base"); ok {
		t, err := expandUserLdapGroupSearchBase(d, v, "group_search_base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-search-base"] = t
		}
	}

	if v, ok := d.GetOk("group_object_filter"); ok {
		t, err := expandUserLdapGroupObjectFilter(d, v, "group_object_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-object-filter"] = t
		}
	}

	if v, ok := d.GetOk("group_filter"); ok {
		t, err := expandUserLdapGroupFilter(d, v, "group_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-filter"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok {
		t, err := expandUserLdapSecure(d, v, "secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		t, err := expandUserLdapSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok {
		t, err := expandUserLdapCaCert(d, v, "ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandUserLdapPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("password_expiry_warning"); ok {
		t, err := expandUserLdapPasswordExpiryWarning(d, v, "password_expiry_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expiry-warning"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok {
		t, err := expandUserLdapPasswordRenewal(d, v, "password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("member_attr"); ok {
		t, err := expandUserLdapMemberAttr(d, v, "member_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-attr"] = t
		}
	}

	if v, ok := d.GetOk("account_key_processing"); ok {
		t, err := expandUserLdapAccountKeyProcessing(d, v, "account_key_processing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-processing"] = t
		}
	}

	if v, ok := d.GetOk("account_key_filter"); ok {
		t, err := expandUserLdapAccountKeyFilter(d, v, "account_key_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-filter"] = t
		}
	}

	if v, ok := d.GetOk("search_type"); ok {
		t, err := expandUserLdapSearchType(d, v, "search_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-type"] = t
		}
	}


	return &obj, nil
}

