package khulnasoft

import (
	"context"
	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRolesMapping() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataRolesMappingRead,
		Schema: map[string]*schema.Schema{
			"saml": {
				Type:        schema.TypeSet,
				Description: "SAML Authentication",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_mapping": {
							Type:        schema.TypeMap,
							Description: "Role Mapping is used to define the IdP role that the user will assume in Khulnasoft",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"oauth2": {
				Type:        schema.TypeSet,
				Description: "Oauth2 Authentication",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_mapping": {
							Type:        schema.TypeMap,
							Description: "Role Mapping is used to define the IdP role that the user will assume in Khulnasoft",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"openid": {
				Type:        schema.TypeSet,
				Description: "OpenId Authentication",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_mapping": {
							Type:        schema.TypeMap,
							Description: "Role Mapping is used to define the IdP role that the user will assume in Khulnasoft",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"ldap": {
				Type:        schema.TypeSet,
				Description: "LDAP Authentication",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_mapping": {
							Type:        schema.TypeMap,
							Description: "Role Mapping is used to define the IdP role that the user will assume in Khulnasoft",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func dataRolesMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	sso, err := c.GetSSO()
	if err == nil {
		d.Set("saml", flattenSamlRoleMapping(sso.Saml))
		d.Set("oauth2", flattenOAuth2RoleMapping(sso.OAuth2))
		d.Set("openid", flattenOpenIdRoleMapping(sso.OpenId))
	} else {
		return diag.FromErr(err)
	}

	ldap, err := c.GetLdap()

	if err == nil {
		d.Set("ldap", flattenLdapRoleMapping(ldap))
	} else {
		return diag.FromErr(err)
	}
	d.SetId("khulnasoft-rolesMapping")

	return nil
}

func flattenSamlRoleMapping(saml client.Saml) []map[string]interface{} {
	return []map[string]interface{}{
		{
			"role_mapping": flattenRoleMap(saml.RoleMapping),
		},
	}
}

func flattenOAuth2RoleMapping(oAuth2 client.OAuth2) []map[string]interface{} {
	return []map[string]interface{}{
		{
			"role_mapping": flattenRoleMap(oAuth2.RoleMapping),
		},
	}
}

func flattenOpenIdRoleMapping(openId client.OpenId) []map[string]interface{} {
	return []map[string]interface{}{
		{
			"role_mapping": flattenRoleMap(openId.RoleMapping),
		},
	}
}

func flattenLdapRoleMapping(ldap *client.Ldap) []map[string]interface{} {
	return []map[string]interface{}{
		{
			"role_mapping": flattenRoleMap(ldap.RoleMapping),
		},
	}
}
