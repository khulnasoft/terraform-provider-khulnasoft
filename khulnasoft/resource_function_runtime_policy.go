package khulnasoft

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
)

func resourceFunctionRuntimePolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFunctionRuntimePolicyCreate,
		ReadContext:   resourceFunctionRuntimePolicyRead,
		UpdateContext: resourceFunctionRuntimePolicyUpdate,
		DeleteContext: resourceFunctionRuntimePolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the function runtime policy",
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "The description of the function runtime policy",
				Optional:    true,
			},
			"application_scopes": {
				Type:        schema.TypeList,
				Description: "Indicates the application scope of the service.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"scope_expression": {
				Type:        schema.TypeString,
				Description: "Logical expression of how to compute the dependency of the scope variables.",
				Optional:    true,
				Computed:    true,
			},
			"scope_variables": {
				Type:        schema.TypeList,
				Description: "List of scope attributes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute": {
							Type:        schema.TypeString,
							Description: "Class of supported scope.",
							Required:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name assigned to the attribute.",
							Optional:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "Value assigned to the attribute.",
							Required:    true,
						},
					},
				},
				Optional: true,
				Computed: true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Indicates if the runtime policy is enabled or not.",
				Default:     true,
				Optional:    true,
			},
			"enforce": {
				Type:        schema.TypeBool,
				Description: "Indicates that policy should effect container execution (not just for audit).",
				Default:     false,
				Optional:    true,
			},
			"author": {
				Type:        schema.TypeString,
				Description: "Username of the account that created the service.",
				Computed:    true,
			},
			"block_malicious_executables": {
				Type:        schema.TypeBool,
				Description: "If true, prevent creation of malicious executables in functions during their runtime post invocation.",
				Default:     false,
				Optional:    true,
			},
			"block_running_executables_in_tmp_folder": {
				Type:         schema.TypeBool,
				Description:  "If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.",
				RequiredWith: []string{"block_malicious_executables"},
				Default:      false,
				Optional:     true,
			},
			"block_malicious_executables_allowed_processes": {
				Type:         schema.TypeList,
				Description:  "List of processes that will be allowed",
				RequiredWith: []string{"block_malicious_executables"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"blocked_executables": {
				Type:        schema.TypeList,
				Description: "List of executables that are prevented from running in containers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"honeypot_access_key": {
				Type:        schema.TypeString,
				Description: "Honeypot User ID (Access Key)",
				Optional:    true,
			},
			"honeypot_secret_key": {
				Type:        schema.TypeString,
				Description: "Honeypot User Password (Secret Key)",
				Optional:    true,
				Sensitive:   true,
			},
			"honeypot_apply_on": {
				Type:        schema.TypeList,
				Description: "List of options to apply the honeypot on (Environment Vairable, Layer, File)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"honeypot_serverless_app_name": {
				Type:        schema.TypeString,
				Description: "Serverless application name",
				Optional:    true,
			},
		},
	}
}

func resourceFunctionRuntimePolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	name := d.Get("name").(string)

	crp := expandFunctionRuntimePolicy(d)
	err := c.CreateRuntimePolicy(crp)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(name)

	return resourceFunctionRuntimePolicyRead(ctx, d, m)

}

func resourceFunctionRuntimePolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	crp, err := c.GetRuntimePolicy(d.Id())

	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", err), "404") {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.Set("name", crp.Name)
	d.Set("description", crp.Description)
	d.Set("author", crp.Author)
	d.Set("application_scopes", crp.ApplicationScopes)
	d.Set("scope_variables", flattenScopeVariables(crp.Scope.Variables))
	d.Set("scope_expression", crp.Scope.Expression)
	d.Set("enabled", crp.Enabled)
	d.Set("enforce", crp.Enforce)
	d.Set("block_malicious_executables", crp.DriftPrevention.Enabled)
	d.Set("block_running_executables_in_tmp_folder", crp.DriftPrevention.ExecLockdown)
	d.Set("block_malicious_executables_allowed_processes", crp.DriftPrevention.ExecLockdownWhiteList)
	d.Set("blocked_executables", crp.ExecutableBlacklist.Executables)
	d.Set("honeypot_access_key", crp.Tripwire.UserID)
	d.Set("honeypot_secret_key", crp.Tripwire.UserPassword)
	d.Set("honeypot_apply_on", crp.Tripwire.ApplyOn)
	d.Set("honeypot_serverless_app_name", crp.Tripwire.ServerlessApp)

	d.SetId(crp.Name)

	return nil
}

func resourceFunctionRuntimePolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	name := d.Get("name").(string)
	if d.HasChanges("description",
		"author",
		"application_scopes",
		"scope_variables",
		"scope_expression",
		"enabled",
		"enforce",
		"block_malicious_executables",
		"block_running_executables_in_tmp_folder",
		"block_malicious_executables_allowed_processes",
		"blocked_executables",
		"honeypot_access_key",
		"honeypot_secret_key",
		"honeypot_apply_on",
		"honeypot_serverless_app_name") {

		crp := expandFunctionRuntimePolicy(d)
		err := c.UpdateRuntimePolicy(crp)
		if err == nil {
			d.SetId(name)
		} else {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceFunctionRuntimePolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	name := d.Get("name").(string)

	err := c.DeleteRuntimePolicy(name)
	if err == nil {
		d.SetId("")
	} else {
		return diag.FromErr(err)
	}

	//d.SetId("")

	return nil
}

func expandFunctionRuntimePolicy(d *schema.ResourceData) *client.RuntimePolicy {
	crp := client.RuntimePolicy{
		Name:        d.Get("name").(string),
		RuntimeType: "function",
	}

	description, ok := d.GetOk("description")
	if ok {
		crp.Description = description.(string)
	}

	applicationScopes, ok := d.GetOk("application_scopes")
	if ok {
		crp.ApplicationScopes = convertStringArr(applicationScopes.([]interface{}))
	}

	scopeExpression, ok := d.GetOk("scope_expression")
	if ok {
		crp.Scope.Expression = scopeExpression.(string)
	}

	variables := make([]client.Variable, 0)
	variableMap, ok := d.GetOk("scope_variables")
	if ok {
		for _, v := range variableMap.([]interface{}) {
			ifc := v.(map[string]interface{})
			variables = append(variables, client.Variable{
				Attribute: ifc["attribute"].(string),
				Name:      ifc["name"].(string),
				Value:     ifc["value"].(string),
			})
		}
	}
	crp.Scope.Variables = variables

	enabled, ok := d.GetOk("enabled")
	if ok {
		crp.Enabled = enabled.(bool)
	}

	enforce, ok := d.GetOk("enforce")
	if ok {
		crp.Enforce = enforce.(bool)
	}

	author, ok := d.GetOk("author")
	if ok {
		crp.Author = author.(string)
	}

	blockMalicious, ok := d.GetOk("block_malicious_executables")
	if ok {
		crp.DriftPrevention.Enabled = blockMalicious.(bool)
		blockRunningExecutablesInTmpFolder, ok := d.GetOk("block_running_executables_in_tmp_folder")
		if ok {
			crp.DriftPrevention.ExecLockdown = blockRunningExecutablesInTmpFolder.(bool)
		}
		blockMaliciousExecutablesAllowedProcesses, ok := d.GetOk("block_malicious_executables_allowed_processes")
		if ok {
			crp.DriftPrevention.ExecLockdownWhiteList = convertStringArr(blockMaliciousExecutablesAllowedProcesses.([]interface{}))
		}
	}

	blockedExecutables, ok := d.GetOk("blocked_executables")
	if ok {
		strArr := convertStringArr(blockedExecutables.([]interface{}))
		crp.ExecutableBlacklist.Enabled = len(strArr) != 0
		crp.ExecutableBlacklist.Executables = strArr
	} else {
		crp.ExecutableBlacklist.Enabled = false
	}

	accessKey, ok := d.GetOk("honeypot_access_key")
	if ok {
		crp.Tripwire.Enabled = true
		crp.Tripwire.UserID = accessKey.(string)
	}

	secretKey, ok := d.GetOk("honeypot_secret_key")
	if ok {
		crp.Tripwire.Enabled = true
		crp.Tripwire.UserPassword = secretKey.(string)
	}

	applyOn, ok := d.GetOk("honeypot_apply_on")
	if ok {
		crp.Tripwire.Enabled = true
		crp.Tripwire.ApplyOn = convertStringArr(applyOn.([]interface{}))
	}

	appName, ok := d.GetOk("honeypot_serverless_app_name")
	if ok {
		crp.Tripwire.Enabled = true
		crp.Tripwire.ServerlessApp = appName.(string)
	}

	return &crp
}
