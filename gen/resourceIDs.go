// Code is generated by terratools. DO NOT EDIT.

package gen

// ResourceIDs stores the name of the struct field of the AWS API used as ID for each Terraform resource type.
var ResourceIDs = map[string]string{
	"aws_acm_certificate":                      "CertificateArn",
	"aws_alb_listener":                         "ListenerArn",
	"aws_ami":                                  "ImageId",
	"aws_ami_copy":                             "ImageId",
	"aws_ami_from_instance":                    "ImageId",
	"aws_api_gateway_authorizer":               "Id",
	"aws_api_gateway_client_certificate":       "ClientCertificateId",
	"aws_api_gateway_deployment":               "Id",
	"aws_api_gateway_domain_name":              "DomainName",
	"aws_api_gateway_model":                    "Id",
	"aws_api_gateway_request_validator":        "Id",
	"aws_api_gateway_resource":                 "Id",
	"aws_api_gateway_rest_api":                 "Id",
	"aws_api_gateway_usage_plan":               "Id",
	"aws_api_gateway_usage_plan_key":           "Id",
	"aws_api_gateway_vpc_link":                 "Id",
	"aws_appsync_graphql_api":                  "ApiId",
	"aws_athena_named_query":                   "NamedQueryId",
	"aws_batch_job_definition":                 "JobDefinitionArn",
	"aws_batch_job_queue":                      "JobQueueArn",
	"aws_cloud9_environment_ec2":               "EnvironmentId",
	"aws_cloudformation_stack":                 "StackId",
	"aws_cloudfront_distribution":              "Id",
	"aws_cloudfront_origin_access_identity":    "Id",
	"aws_cloudtrail":                           "Name",
	"aws_cloudwatch_event_rule":                "Name",
	"aws_codebuild_project":                    "Arn",
	"aws_codedeploy_deployment_group":          "DeploymentGroupId",
	"aws_cognito_identity_pool":                "IdentityPoolId",
	"aws_cognito_user_pool":                    "Id",
	"aws_cognito_user_pool_client":             "ClientId",
	"aws_customer_gateway":                     "CustomerGatewayId",
	"aws_db_subnet_group":                      "DBSubnetGroupName",
	"aws_default_route_table":                  "RouteTableId",
	"aws_default_security_group":               "GroupId",
	"aws_devicefarm_project":                   "Arn",
	"aws_dlm_lifecycle_policy":                 "PolicyId",
	"aws_dms_certificate":                      "CertificateIdentifier",
	"aws_dms_endpoint":                         "EndpointIdentifier",
	"aws_dms_replication_subnet_group":         "ReplicationSubnetGroupIdentifier",
	"aws_dms_replication_task":                 "ReplicationTaskIdentifier",
	"aws_docdb_cluster_instance":               "DBInstanceIdentifier",
	"aws_dx_connection_association":            "ConnectionId",
	"aws_dynamodb_table":                       "TableName",
	"aws_ebs_snapshot":                         "SnapshotId",
	"aws_ebs_snapshot_copy":                    "SnapshotId",
	"aws_ebs_volume":                           "VolumeId",
	"aws_ec2_capacity_reservation":             "CapacityReservationId",
	"aws_ec2_client_vpn_endpoint":              "ClientVpnEndpointId",
	"aws_ec2_client_vpn_network_association":   "AssociationId",
	"aws_ecr_lifecycle_policy":                 "RepositoryName",
	"aws_ecr_repository_policy":                "RepositoryName",
	"aws_efs_file_system":                      "FileSystemId",
	"aws_efs_mount_target":                     "MountTargetId",
	"aws_eip":                                  "AllocationId",
	"aws_eip_association":                      "AssociationId",
	"aws_elastic_beanstalk_environment":        "EnvironmentId",
	"aws_elasticache_parameter_group":          "CacheParameterGroupName",
	"aws_elasticache_replication_group":        "ReplicationGroupId",
	"aws_elastictranscoder_pipeline":           "Id",
	"aws_elastictranscoder_preset":             "Id",
	"aws_emr_cluster":                          "JobFlowId",
	"aws_emr_security_configuration":           "Name",
	"aws_fsx_lustre_file_system":               "FileSystemId",
	"aws_fsx_windows_file_system":              "FileSystemId",
	"aws_gamelift_alias":                       "AliasId",
	"aws_gamelift_build":                       "BuildId",
	"aws_gamelift_fleet":                       "FleetId",
	"aws_gamelift_game_session_queue":          "Name",
	"aws_globalaccelerator_accelerator":        "AcceleratorArn",
	"aws_globalaccelerator_endpoint_group":     "EndpointGroupArn",
	"aws_globalaccelerator_listener":           "ListenerArn",
	"aws_guardduty_detector":                   "DetectorId",
	"aws_iam_access_key":                       "AccessKeyId",
	"aws_iam_group":                            "GroupName",
	"aws_iam_instance_profile":                 "InstanceProfileName",
	"aws_iam_openid_connect_provider":          "OpenIDConnectProviderArn",
	"aws_iam_policy":                           "Arn",
	"aws_iam_role":                             "RoleName",
	"aws_iam_saml_provider":                    "SAMLProviderArn",
	"aws_iam_server_certificate":               "ServerCertificateId",
	"aws_iam_service_linked_role":              "Arn",
	"aws_iam_user_login_profile":               "UserName",
	"aws_iam_user_ssh_key":                     "SSHPublicKeyId",
	"aws_inspector_assessment_target":          "AssessmentTargetArn",
	"aws_inspector_assessment_template":        "AssessmentTemplateArn",
	"aws_instance":                             "InstanceId",
	"aws_internet_gateway":                     "InternetGatewayId",
	"aws_iot_certificate":                      "CertificateId",
	"aws_iot_policy":                           "PolicyName",
	"aws_iot_thing":                            "ThingName",
	"aws_iot_thing_type":                       "ThingTypeName",
	"aws_key_pair":                             "KeyName",
	"aws_kms_key":                              "KeyId",
	"aws_lambda_alias":                         "AliasArn",
	"aws_lambda_event_source_mapping":          "UUID",
	"aws_launch_template":                      "LaunchTemplateId",
	"aws_lb_listener":                          "ListenerArn",
	"aws_licensemanager_license_configuration": "LicenseConfigurationArn",
	"aws_main_route_table_association":         "NewAssociationId",
	"aws_mq_broker":                            "BrokerId",
	"aws_mq_configuration":                     "Id",
	"aws_nat_gateway":                          "NatGatewayId",
	"aws_neptune_parameter_group":              "DBParameterGroupName",
	"aws_network_acl":                          "NetworkAclId",
	"aws_network_interface":                    "NetworkInterfaceId",
	"aws_network_interface_attachment":         "AttachmentId",
	"aws_opsworks_application":                 "AppId",
	"aws_opsworks_instance":                    "InstanceId",
	"aws_opsworks_stack":                       "StackId",
	"aws_opsworks_user_profile":                "IamUserArn",
	"aws_organizations_organization":           "Id",
	"aws_organizations_policy":                 "Id",
	"aws_qldb_ledger":                          "Name",
	"aws_rds_cluster_instance":                 "DBInstanceIdentifier",
	"aws_rds_cluster_parameter_group":          "DBClusterParameterGroupName",
	"aws_redshift_cluster":                     "ClusterIdentifier",
	"aws_redshift_parameter_group":             "ParameterGroupName",
	"aws_redshift_subnet_group":                "ClusterSubnetGroupName",
	"aws_route53_health_check":                 "Id",
	"aws_route53_query_log":                    "Id",
	"aws_route_table":                          "RouteTableId",
	"aws_route_table_association":              "NewAssociationId",
	"aws_security_group":                       "GroupId",
	"aws_securityhub_standards_subscription":   "StandardsSubscriptionArn",
	"aws_service_discovery_service":            "Id",
	"aws_servicecatalog_portfolio":             "Id",
	"aws_sfn_activity":                         "ActivityArn",
	"aws_sfn_state_machine":                    "StateMachineArn",
	"aws_shield_protection":                    "ProtectionId",
	"aws_sns_platform_application":             "PlatformApplicationArn",
	"aws_sns_topic":                            "TopicArn",
	"aws_sns_topic_subscription":               "SubscriptionArn",
	"aws_spot_fleet_request":                   "SpotFleetRequestId",
	"aws_spot_instance_request":                "SpotInstanceRequestId",
	"aws_ssm_activation":                       "ActivationId",
	"aws_ssm_association":                      "AssociationId",
	"aws_ssm_document":                         "Name",
	"aws_ssm_maintenance_window":               "WindowId",
	"aws_ssm_maintenance_window_target":        "WindowTargetId",
	"aws_ssm_maintenance_window_task":          "WindowTaskId",
	"aws_ssm_patch_baseline":                   "BaselineId",
	"aws_ssm_patch_group":                      "PatchGroup",
	"aws_subnet":                               "SubnetId",
	"aws_transfer_server":                      "ServerId",
	"aws_vpc":                                  "VpcId",
	"aws_vpc_peering_connection":               "VpcPeeringConnectionId",
	"aws_vpn_gateway":                          "VpnGatewayId",
	"aws_waf_byte_match_set":                   "ByteMatchSetId",
	"aws_waf_geo_match_set":                    "GeoMatchSetId",
	"aws_waf_ipset":                            "IPSetId",
	"aws_waf_rate_based_rule":                  "RuleId",
	"aws_waf_regex_match_set":                  "RegexMatchSetId",
	"aws_waf_regex_pattern_set":                "RegexPatternSetId",
	"aws_waf_rule":                             "RuleId",
	"aws_waf_rule_group":                       "RuleGroupId",
	"aws_waf_size_constraint_set":              "SizeConstraintSetId",
	"aws_waf_sql_injection_match_set":          "SqlInjectionMatchSetId",
	"aws_waf_web_acl":                          "WebACLId",
	"aws_waf_xss_match_set":                    "XssMatchSetId",
	"aws_wafregional_byte_match_set":           "ByteMatchSetId",
	"aws_wafregional_geo_match_set":            "GeoMatchSetId",
	"aws_wafregional_ipset":                    "IPSetId",
	"aws_wafregional_rate_based_rule":          "RuleId",
	"aws_wafregional_rule":                     "RuleId",
	"aws_wafregional_rule_group":               "RuleGroupId",
	"aws_wafregional_web_acl":                  "WebACLId",
	"aws_wafregional_xss_match_set":            "XssMatchSetId",
	"aws_workspaces_ip_group":                  "GroupId",
	"aws_xray_sampling_rule":                   "RuleName",
}
