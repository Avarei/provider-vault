/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"vault_ad_secret_backend":                            config.IdentifierFromProvider,
	"vault_ad_secret_backend_library":                    config.IdentifierFromProvider,
	"vault_ad_secret_role":                               config.IdentifierFromProvider,
	"vault_alicloud_auth_backend_role":                   config.IdentifierFromProvider,
	"vault_approle_auth_backend_login":                   config.IdentifierFromProvider,
	"vault_approle_auth_backend_role":                    config.IdentifierFromProvider,
	"vault_approle_auth_backend_role_secret_id":          config.IdentifierFromProvider,
	"vault_audit":                                        config.IdentifierFromProvider,
	"vault_audit_request_header":                         config.IdentifierFromProvider,
	"vault_auth_backend":                                 config.IdentifierFromProvider,
	"vault_aws_auth_backend_cert":                        config.IdentifierFromProvider,
	"vault_aws_auth_backend_client":                      config.IdentifierFromProvider,
	"vault_aws_auth_backend_config_identity":             config.IdentifierFromProvider,
	"vault_aws_auth_backend_identity_whitelist":          config.IdentifierFromProvider,
	"vault_aws_auth_backend_login":                       config.IdentifierFromProvider,
	"vault_aws_auth_backend_role":                        config.IdentifierFromProvider,
	"vault_aws_auth_backend_role_tag":                    config.IdentifierFromProvider,
	"vault_aws_auth_backend_roletag_blacklist":           config.IdentifierFromProvider,
	"vault_aws_auth_backend_sts_role":                    config.IdentifierFromProvider,
	"vault_aws_secret_backend":                           config.IdentifierFromProvider,
	"vault_aws_secret_backend_role":                      config.IdentifierFromProvider,
	"vault_azure_auth_backend_config":                    config.IdentifierFromProvider,
	"vault_azure_auth_backend_role":                      config.IdentifierFromProvider,
	"vault_azure_secret_backend":                         config.IdentifierFromProvider,
	"vault_azure_secret_backend_role":                    config.IdentifierFromProvider,
	"vault_cert_auth_backend_role":                       config.IdentifierFromProvider,
	"vault_config_ui_custom_message":                     config.IdentifierFromProvider,
	"vault_consul_secret_backend":                        config.IdentifierFromProvider,
	"vault_consul_secret_backend_role":                   config.IdentifierFromProvider,
	"vault_database_secret_backend_connection":           config.IdentifierFromProvider,
	"vault_database_secret_backend_role":                 config.IdentifierFromProvider,
	"vault_database_secret_backend_static_role":          config.IdentifierFromProvider,
	"vault_database_secrets_mount":                       config.IdentifierFromProvider,
	"vault_egp_policy":                                   config.IdentifierFromProvider,
	"vault_gcp_auth_backend":                             config.IdentifierFromProvider,
	"vault_gcp_auth_backend_role":                        config.IdentifierFromProvider,
	"vault_gcp_secret_backend":                           config.IdentifierFromProvider,
	"vault_gcp_secret_impersonated_account":              config.IdentifierFromProvider,
	"vault_gcp_secret_roleset":                           config.IdentifierFromProvider,
	"vault_gcp_secret_static_account":                    config.IdentifierFromProvider,
	"vault_generic_endpoint":                             config.IdentifierFromProvider,
	"vault_generic_secret":                               config.IdentifierFromProvider,
	"vault_github_auth_backend":                          config.IdentifierFromProvider,
	"vault_github_team":                                  config.IdentifierFromProvider,
	"vault_github_user":                                  config.IdentifierFromProvider,
	"vault_identity_entity":                              config.IdentifierFromProvider,
	"vault_identity_entity_alias":                        config.IdentifierFromProvider,
	"vault_identity_entity_policies":                     config.IdentifierFromProvider,
	"vault_identity_group":                               config.IdentifierFromProvider,
	"vault_identity_group_alias":                         config.IdentifierFromProvider,
	"vault_identity_group_member_entity_ids":             config.IdentifierFromProvider,
	"vault_identity_group_member_group_ids":              config.IdentifierFromProvider,
	"vault_identity_group_policies":                      config.IdentifierFromProvider,
	"vault_identity_mfa_duo":                             config.IdentifierFromProvider,
	"vault_identity_mfa_login_enforcement":               config.IdentifierFromProvider,
	"vault_identity_mfa_okta":                            config.IdentifierFromProvider,
	"vault_identity_mfa_pingid":                          config.IdentifierFromProvider,
	"vault_identity_mfa_totp":                            config.IdentifierFromProvider,
	"vault_identity_oidc":                                config.IdentifierFromProvider,
	"vault_identity_oidc_assignment":                     config.IdentifierFromProvider,
	"vault_identity_oidc_client":                         config.IdentifierFromProvider,
	"vault_identity_oidc_key":                            config.IdentifierFromProvider,
	"vault_identity_oidc_key_allowed_client_id":          config.IdentifierFromProvider,
	"vault_identity_oidc_provider":                       config.IdentifierFromProvider,
	"vault_identity_oidc_role":                           config.IdentifierFromProvider,
	"vault_identity_oidc_scope":                          config.IdentifierFromProvider,
	"vault_jwt_auth_backend":                             config.IdentifierFromProvider,
	"vault_jwt_auth_backend_role":                        config.IdentifierFromProvider,
	"vault_kmip_secret_backend":                          config.IdentifierFromProvider,
	"vault_kmip_secret_role":                             config.IdentifierFromProvider,
	"vault_kmip_secret_scope":                            config.IdentifierFromProvider,
	"vault_kubernetes_auth_backend_config":               config.IdentifierFromProvider,
	"vault_kubernetes_auth_backend_role":                 config.IdentifierFromProvider,
	"vault_kubernetes_secret_backend":                    config.IdentifierFromProvider,
	"vault_kubernetes_secret_backend_role":               config.IdentifierFromProvider,
	"vault_kv_secret":                                    config.IdentifierFromProvider,
	"vault_kv_secret_backend_v2":                         config.IdentifierFromProvider,
	"vault_kv_secret_v2":                                 config.IdentifierFromProvider,
	"vault_ldap_auth_backend":                            config.IdentifierFromProvider,
	"vault_ldap_auth_backend_group":                      config.IdentifierFromProvider,
	"vault_ldap_auth_backend_user":                       config.IdentifierFromProvider,
	"vault_managed_keys":                                 config.IdentifierFromProvider,
	"vault_mfa_duo":                                      config.IdentifierFromProvider,
	"vault_mfa_okta":                                     config.IdentifierFromProvider,
	"vault_mfa_pingid":                                   config.IdentifierFromProvider,
	"vault_mfa_totp":                                     config.IdentifierFromProvider,
	"vault_mongodbatlas_secret_backend":                  config.IdentifierFromProvider,
	"vault_mongodbatlas_secret_role":                     config.IdentifierFromProvider,
	"vault_mount":                                        config.IdentifierFromProvider,
	"vault_namespace":                                    config.IdentifierFromProvider,
	"vault_nomad_secret_backend":                         config.IdentifierFromProvider,
	"vault_nomad_secret_role":                            config.IdentifierFromProvider,
	"vault_okta_auth_backend":                            config.IdentifierFromProvider,
	"vault_okta_auth_backend_group":                      config.IdentifierFromProvider,
	"vault_okta_auth_backend_user":                       config.IdentifierFromProvider,
	"vault_password_policy":                              config.IdentifierFromProvider,
	"vault_plugin":                                       config.IdentifierFromProvider,
	"vault_plugin_pinned_version":                        config.IdentifierFromProvider,
	"vault_pki_secret_backend_cert":                      config.IdentifierFromProvider,
	"vault_pki_secret_backend_config_ca":                 config.IdentifierFromProvider,
	"vault_pki_secret_backend_config_urls":               config.IdentifierFromProvider,
	"vault_pki_secret_backend_crl_config":                config.IdentifierFromProvider,
	"vault_pki_secret_backend_intermediate_cert_request": config.IdentifierFromProvider,
	"vault_pki_secret_backend_intermediate_set_signed":   config.IdentifierFromProvider,
	"vault_pki_secret_backend_role":                      config.IdentifierFromProvider,
	"vault_pki_secret_backend_root_cert":                 config.IdentifierFromProvider,
	"vault_pki_secret_backend_root_sign_intermediate":    config.IdentifierFromProvider,
	"vault_pki_secret_backend_sign":                      config.IdentifierFromProvider,
	"vault_policy":                                       config.IdentifierFromProvider,
	"vault_quota_lease_count":                            config.IdentifierFromProvider,
	"vault_quota_rate_limit":                             config.IdentifierFromProvider,
	"vault_rabbitmq_secret_backend":                      config.IdentifierFromProvider,
	"vault_rabbitmq_secret_backend_role":                 config.IdentifierFromProvider,
	"vault_raft_autopilot":                               config.IdentifierFromProvider,
	"vault_raft_snapshot_agent_config":                   config.IdentifierFromProvider,
	"vault_rgp_policy":                                   config.IdentifierFromProvider,
	"vault_ssh_secret_backend_ca":                        config.IdentifierFromProvider,
	"vault_ssh_secret_backend_role":                      config.IdentifierFromProvider,
	"vault_terraform_cloud_secret_backend":               config.IdentifierFromProvider,
	"vault_terraform_cloud_secret_creds":                 config.IdentifierFromProvider,
	"vault_terraform_cloud_secret_role":                  config.IdentifierFromProvider,
	"vault_token":                                        config.IdentifierFromProvider,
	"vault_token_auth_backend_role":                      config.IdentifierFromProvider,
	"vault_transform_alphabet":                           config.IdentifierFromProvider,
	"vault_transform_role":                               config.IdentifierFromProvider,
	"vault_transform_template":                           config.IdentifierFromProvider,
	"vault_transform_transformation":                     config.IdentifierFromProvider,
	"vault_transit_secret_backend_cache_config":          config.IdentifierFromProvider,
	"vault_transit_secret_backend_key":                   config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}
