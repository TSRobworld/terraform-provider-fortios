// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Replacement messages.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemReplacemsgSpam() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgSpamUpdate,
		Read:   resourceSystemReplacemsgSpamRead,
		Update: resourceSystemReplacemsgSpamUpdate,
		Delete: resourceSystemReplacemsgSpamDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"msg_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 28),
				ForceNew:     true,
				Required:     true,
			},
			"buffer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),
				Optional:     true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemReplacemsgSpamUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey = d.Get("msg_type").(string)
	obj, err := getObjectSystemReplacemsgSpam(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgSpam resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemReplacemsgSpam(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgSpam resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgSpam")
	}

	return resourceSystemReplacemsgSpamRead(d, m)
}

func resourceSystemReplacemsgSpamDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemReplacemsgSpam(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgSpam resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemReplacemsgSpam(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemReplacemsgSpam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgSpamRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemReplacemsgSpam(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgSpam resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgSpam(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgSpam resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgSpamMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgSpamBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgSpamHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgSpamFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgSpam(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("msg_type", flattenSystemReplacemsgSpamMsgType(o["msg-type"], d, "msg_type", sv)); err != nil {
		if !fortiAPIPatch(o["msg-type"]) {
			return fmt.Errorf("Error reading msg_type: %v", err)
		}
	}

	if err = d.Set("buffer", flattenSystemReplacemsgSpamBuffer(o["buffer"], d, "buffer", sv)); err != nil {
		if !fortiAPIPatch(o["buffer"]) {
			return fmt.Errorf("Error reading buffer: %v", err)
		}
	}

	if err = d.Set("header", flattenSystemReplacemsgSpamHeader(o["header"], d, "header", sv)); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("format", flattenSystemReplacemsgSpamFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	return nil
}

func flattenSystemReplacemsgSpamFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemReplacemsgSpamMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgSpamBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgSpamHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgSpamFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgSpam(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("msg_type"); ok {
		if setArgNil {
			obj["msg-type"] = nil
		} else {

			t, err := expandSystemReplacemsgSpamMsgType(d, v, "msg_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["msg-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("buffer"); ok {
		if setArgNil {
			obj["buffer"] = nil
		} else {

			t, err := expandSystemReplacemsgSpamBuffer(d, v, "buffer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["buffer"] = t
			}
		}
	}

	if v, ok := d.GetOk("header"); ok {
		if setArgNil {
			obj["header"] = nil
		} else {

			t, err := expandSystemReplacemsgSpamHeader(d, v, "header", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["header"] = t
			}
		}
	}

	if v, ok := d.GetOk("format"); ok {
		if setArgNil {
			obj["format"] = nil
		} else {

			t, err := expandSystemReplacemsgSpamFormat(d, v, "format", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["format"] = t
			}
		}
	}

	return &obj, nil
}
