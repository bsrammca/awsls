// Code is generated by terratools. DO NOT EDIT.

package gen

// ResourceIDs stores the name of the struct field of the AWS API used as ID for each Terraform resource type.
var ResourceIDs = map[string]string{
	"aws_acm_certificate":                             "CertificateArn",
	"aws_acmpca_certificate_authority":                "CertificateAuthorityArn",
	"aws_alb":                                         "LoadBalancerArn",
	"aws_alb_listener":                                "ListenerArn",
	"aws_alb_listener_rule":                           "RuleArn",
	"aws_alb_target_group":                            "TargetGroupArn",
	"aws_ami":                                         "ImageId",
	"aws_ami_copy":                                    "ImageId",
	"aws_ami_from_instance":                           "ImageId",
	"aws_api_gateway_api_key":                         "Id",
	"aws_api_gateway_authorizer":                      "Id",
	"aws_api_gateway_client_certificate":              "ClientCertificateId",
	"aws_api_gateway_deployment":                      "Id",
	"aws_api_gateway_domain_name":                     "DomainName",
	"aws_api_gateway_model":                           "Id",
	"aws_api_gateway_request_validator":               "Id",
	"aws_api_gateway_resource":                        "Id",
	"aws_api_gateway_rest_api":                        "Id",
	"aws_api_gateway_usage_plan":                      "Id",
	"aws_api_gateway_usage_plan_key":                  "Id",
	"aws_api_gateway_vpc_link":                        "Id",
	"aws_appmesh_route":                               "Uid",
	"aws_appmesh_virtual_node":                        "Uid",
	"aws_appmesh_virtual_router":                      "Uid",
	"aws_appmesh_virtual_service":                     "Uid",
	"aws_appsync_graphql_api":                         "ApiId",
	"aws_athena_named_query":                          "NamedQueryId",
	"aws_backup_plan":                                 "BackupPlanId",
	"aws_backup_selection":                            "SelectionId",
	"aws_batch_job_definition":                        "JobDefinitionArn",
	"aws_batch_job_queue":                             "JobQueueArn",
	"aws_cloud9_environment_ec2":                      "EnvironmentId",
	"aws_cloudformation_stack":                        "StackId",
	"aws_cloudfront_distribution":                     "Id",
	"aws_cloudfront_origin_access_identity":           "Id",
	"aws_cloudfront_public_key":                       "Id",
	"aws_cloudhsm_v2_cluster":                         "ClusterId",
	"aws_cloudhsm_v2_hsm":                             "HsmId",
	"aws_cloudtrail":                                  "Name",
	"aws_cloudwatch_event_rule":                       "Name",
	"aws_codebuild_project":                           "Arn",
	"aws_codebuild_source_credential":                 "Arn",
	"aws_codedeploy_deployment_group":                 "DeploymentGroupId",
	"aws_codepipeline":                                "Name",
	"aws_codepipeline_webhook":                        "Arn",
	"aws_cognito_identity_pool":                       "IdentityPoolId",
	"aws_cognito_user_pool":                           "Id",
	"aws_cognito_user_pool_client":                    "ClientId",
	"aws_cur_report_definition":                       "ReportName",
	"aws_customer_gateway":                            "CustomerGatewayId",
	"aws_datapipeline_pipeline":                       "PipelineId",
	"aws_datasync_agent":                              "AgentArn",
	"aws_datasync_location_efs":                       "LocationArn",
	"aws_datasync_location_nfs":                       "LocationArn",
	"aws_datasync_location_s3":                        "LocationArn",
	"aws_datasync_task":                               "TaskArn",
	"aws_db_event_subscription":                       "CustSubscriptionId",
	"aws_db_parameter_group":                          "DBParameterGroupName",
	"aws_db_subnet_group":                             "DBSubnetGroupName",
	"aws_default_route_table":                         "RouteTableId",
	"aws_default_security_group":                      "GroupId",
	"aws_default_subnet":                              "SubnetId",
	"aws_default_vpc":                                 "VpcId",
	"aws_default_vpc_dhcp_options":                    "DhcpOptionsId",
	"aws_devicefarm_project":                          "Arn",
	"aws_dlm_lifecycle_policy":                        "PolicyId",
	"aws_dms_certificate":                             "CertificateIdentifier",
	"aws_dms_endpoint":                                "EndpointIdentifier",
	"aws_dms_replication_subnet_group":                "ReplicationSubnetGroupIdentifier",
	"aws_dms_replication_task":                        "ReplicationTaskIdentifier",
	"aws_docdb_cluster_instance":                      "DBInstanceIdentifier",
	"aws_docdb_cluster_parameter_group":               "DBClusterParameterGroupName",
	"aws_dx_connection":                               "ConnectionId",
	"aws_dx_connection_association":                   "ConnectionId",
	"aws_dx_gateway":                                  "DirectConnectGatewayId",
	"aws_dx_gateway_association_proposal":             "ProposalId",
	"aws_dx_hosted_private_virtual_interface":         "VirtualInterfaceId",
	"aws_dx_hosted_public_virtual_interface":          "VirtualInterfaceId",
	"aws_dx_hosted_transit_virtual_interface":         "VirtualInterfaceId",
	"aws_dx_lag":                                      "LagId",
	"aws_dx_private_virtual_interface":                "VirtualInterfaceId",
	"aws_dx_public_virtual_interface":                 "VirtualInterfaceId",
	"aws_dx_transit_virtual_interface":                "VirtualInterfaceId",
	"aws_dynamodb_table":                              "TableName",
	"aws_ebs_default_kms_key":                         "KmsKeyId",
	"aws_ebs_snapshot":                                "SnapshotId",
	"aws_ebs_snapshot_copy":                           "SnapshotId",
	"aws_ebs_volume":                                  "VolumeId",
	"aws_ec2_capacity_reservation":                    "CapacityReservationId",
	"aws_ec2_client_vpn_endpoint":                     "ClientVpnEndpointId",
	"aws_ec2_client_vpn_network_association":          "AssociationId",
	"aws_ec2_fleet":                                   "FleetId",
	"aws_ec2_transit_gateway":                         "TransitGatewayId",
	"aws_ec2_transit_gateway_route_table":             "TransitGatewayRouteTableId",
	"aws_ec2_transit_gateway_vpc_attachment":          "TransitGatewayAttachmentId",
	"aws_ec2_transit_gateway_vpc_attachment_accepter": "TransitGatewayAttachmentId",
	"aws_ecr_lifecycle_policy":                        "RepositoryName",
	"aws_ecr_repository":                              "RepositoryName",
	"aws_ecr_repository_policy":                       "RepositoryName",
	"aws_ecs_capacity_provider":                       "CapacityProviderArn",
	"aws_ecs_cluster":                                 "ClusterArn",
	"aws_ecs_service":                                 "ServiceArn",
	"aws_ecs_task_definition":                         "Family",
	"aws_efs_file_system":                             "FileSystemId",
	"aws_efs_mount_target":                            "MountTargetId",
	"aws_egress_only_internet_gateway":                "EgressOnlyInternetGatewayId",
	"aws_eip":                                         "AllocationId",
	"aws_eip_association":                             "AssociationId",
	"aws_elastic_beanstalk_environment":               "EnvironmentId",
	"aws_elasticache_parameter_group":                 "CacheParameterGroupName",
	"aws_elasticache_replication_group":               "ReplicationGroupId",
	"aws_elasticsearch_domain":                        "ARN",
	"aws_elastictranscoder_pipeline":                  "Id",
	"aws_elastictranscoder_preset":                    "Id",
	"aws_emr_cluster":                                 "JobFlowId",
	"aws_emr_security_configuration":                  "Name",
	"aws_fsx_lustre_file_system":                      "FileSystemId",
	"aws_fsx_windows_file_system":                     "FileSystemId",
	"aws_gamelift_alias":                              "AliasId",
	"aws_gamelift_build":                              "BuildId",
	"aws_gamelift_fleet":                              "FleetId",
	"aws_gamelift_game_session_queue":                 "Name",
	"aws_globalaccelerator_accelerator":               "AcceleratorArn",
	"aws_globalaccelerator_endpoint_group":            "EndpointGroupArn",
	"aws_globalaccelerator_listener":                  "ListenerArn",
	"aws_guardduty_detector":                          "DetectorId",
	"aws_iam_access_key":                              "AccessKeyId",
	"aws_iam_group":                                   "GroupName",
	"aws_iam_instance_profile":                        "InstanceProfileName",
	"aws_iam_openid_connect_provider":                 "OpenIDConnectProviderArn",
	"aws_iam_policy":                                  "Arn",
	"aws_iam_role":                                    "RoleName",
	"aws_iam_saml_provider":                           "SAMLProviderArn",
	"aws_iam_server_certificate":                      "ServerCertificateId",
	"aws_iam_service_linked_role":                     "Arn",
	"aws_iam_user":                                    "UserName",
	"aws_iam_user_login_profile":                      "UserName",
	"aws_iam_user_ssh_key":                            "SSHPublicKeyId",
	"aws_inspector_assessment_target":                 "AssessmentTargetArn",
	"aws_inspector_assessment_template":               "AssessmentTemplateArn",
	"aws_inspector_resource_group":                    "ResourceGroupArn",
	"aws_instance":                                    "InstanceId",
	"aws_internet_gateway":                            "InternetGatewayId",
	"aws_iot_certificate":                             "CertificateId",
	"aws_iot_policy":                                  "PolicyName",
	"aws_iot_thing":                                   "ThingName",
	"aws_iot_thing_type":                              "ThingTypeName",
	"aws_key_pair":                                    "KeyName",
	"aws_kinesis_analytics_application":               "ApplicationARN",
	"aws_kinesis_firehose_delivery_stream":            "DeliveryStreamARN",
	"aws_kms_external_key":                            "KeyId",
	"aws_kms_key":                                     "KeyId",
	"aws_lambda_alias":                                "AliasArn",
	"aws_lambda_event_source_mapping":                 "UUID",
	"aws_lambda_layer_version":                        "LayerVersionArn",
	"aws_launch_template":                             "LaunchTemplateId",
	"aws_lb":                                          "LoadBalancerArn",
	"aws_lb_listener":                                 "ListenerArn",
	"aws_lb_listener_rule":                            "RuleArn",
	"aws_lb_target_group":                             "TargetGroupArn",
	"aws_licensemanager_license_configuration":        "LicenseConfigurationArn",
	"aws_main_route_table_association":                "NewAssociationId",
	"aws_media_convert_queue":                         "Name",
	"aws_media_package_channel":                       "Id",
	"aws_media_store_container":                       "Name",
	"aws_mq_broker":                                   "BrokerId",
	"aws_mq_configuration":                            "Id",
	"aws_msk_cluster":                                 "ClusterArn",
	"aws_msk_configuration":                           "Arn",
	"aws_nat_gateway":                                 "NatGatewayId",
	"aws_neptune_cluster_instance":                    "DBInstanceIdentifier",
	"aws_neptune_cluster_parameter_group":             "DBClusterParameterGroupName",
	"aws_neptune_event_subscription":                  "CustSubscriptionId",
	"aws_neptune_parameter_group":                     "DBParameterGroupName",
	"aws_neptune_subnet_group":                        "DBSubnetGroupName",
	"aws_network_acl":                                 "NetworkAclId",
	"aws_network_interface":                           "NetworkInterfaceId",
	"aws_network_interface_attachment":                "AttachmentId",
	"aws_opsworks_application":                        "AppId",
	"aws_opsworks_instance":                           "InstanceId",
	"aws_opsworks_stack":                              "StackId",
	"aws_opsworks_user_profile":                       "IamUserArn",
	"aws_organizations_organization":                  "Id",
	"aws_organizations_policy":                        "Id",
	"aws_pinpoint_app":                                "Id",
	"aws_qldb_ledger":                                 "Name",
	"aws_ram_resource_share":                          "ResourceShareArn",
	"aws_rds_cluster_instance":                        "DBInstanceIdentifier",
	"aws_rds_cluster_parameter_group":                 "DBClusterParameterGroupName",
	"aws_rds_global_cluster":                          "GlobalClusterIdentifier",
	"aws_redshift_cluster":                            "ClusterIdentifier",
	"aws_redshift_event_subscription":                 "CustSubscriptionId",
	"aws_redshift_parameter_group":                    "ParameterGroupName",
	"aws_redshift_snapshot_schedule":                  "ScheduleIdentifier",
	"aws_redshift_subnet_group":                       "ClusterSubnetGroupName",
	"aws_resourcegroups_group":                        "Name",
	"aws_route53_health_check":                        "Id",
	"aws_route53_query_log":                           "Id",
	"aws_route53_resolver_endpoint":                   "Id",
	"aws_route53_resolver_rule":                       "Id",
	"aws_route53_resolver_rule_association":           "Id",
	"aws_route_table":                                 "RouteTableId",
	"aws_route_table_association":                     "NewAssociationId",
	"aws_secretsmanager_secret":                       "ARN",
	"aws_security_group":                              "GroupId",
	"aws_securityhub_standards_subscription":          "StandardsSubscriptionArn",
	"aws_service_discovery_service":                   "Id",
	"aws_servicecatalog_portfolio":                    "Id",
	"aws_sfn_activity":                                "ActivityArn",
	"aws_sfn_state_machine":                           "StateMachineArn",
	"aws_shield_protection":                           "ProtectionId",
	"aws_sns_platform_application":                    "PlatformApplicationArn",
	"aws_sns_topic":                                   "TopicArn",
	"aws_sns_topic_subscription":                      "SubscriptionArn",
	"aws_spot_fleet_request":                          "SpotFleetRequestId",
	"aws_spot_instance_request":                       "SpotInstanceRequestId",
	"aws_sqs_queue":                                   "QueueUrl",
	"aws_ssm_activation":                              "ActivationId",
	"aws_ssm_association":                             "AssociationId",
	"aws_ssm_document":                                "Name",
	"aws_ssm_maintenance_window":                      "WindowId",
	"aws_ssm_maintenance_window_target":               "WindowTargetId",
	"aws_ssm_maintenance_window_task":                 "WindowTaskId",
	"aws_ssm_patch_baseline":                          "BaselineId",
	"aws_ssm_patch_group":                             "PatchGroup",
	"aws_storagegateway_cached_iscsi_volume":          "VolumeARN",
	"aws_storagegateway_gateway":                      "GatewayARN",
	"aws_storagegateway_nfs_file_share":               "FileShareARN",
	"aws_storagegateway_smb_file_share":               "FileShareARN",
	"aws_subnet":                                      "SubnetId",
	"aws_transfer_server":                             "ServerId",
	"aws_vpc":                                         "VpcId",
	"aws_vpc_dhcp_options":                            "DhcpOptionsId",
	"aws_vpc_endpoint":                                "VpcEndpointId",
	"aws_vpc_endpoint_connection_notification":        "ConnectionNotificationId",
	"aws_vpc_endpoint_service":                        "ServiceId",
	"aws_vpc_ipv4_cidr_block_association":             "AssociationId",
	"aws_vpc_peering_connection":                      "VpcPeeringConnectionId",
	"aws_vpn_connection":                              "VpnConnectionId",
	"aws_vpn_gateway":                                 "VpnGatewayId",
	"aws_waf_byte_match_set":                          "ByteMatchSetId",
	"aws_waf_geo_match_set":                           "GeoMatchSetId",
	"aws_waf_ipset":                                   "IPSetId",
	"aws_waf_rate_based_rule":                         "RuleId",
	"aws_waf_regex_match_set":                         "RegexMatchSetId",
	"aws_waf_regex_pattern_set":                       "RegexPatternSetId",
	"aws_waf_rule":                                    "RuleId",
	"aws_waf_rule_group":                              "RuleGroupId",
	"aws_waf_size_constraint_set":                     "SizeConstraintSetId",
	"aws_waf_sql_injection_match_set":                 "SqlInjectionMatchSetId",
	"aws_waf_web_acl":                                 "WebACLId",
	"aws_waf_xss_match_set":                           "XssMatchSetId",
	"aws_wafregional_byte_match_set":                  "ByteMatchSetId",
	"aws_wafregional_geo_match_set":                   "GeoMatchSetId",
	"aws_wafregional_ipset":                           "IPSetId",
	"aws_wafregional_rate_based_rule":                 "RuleId",
	"aws_wafregional_regex_match_set":                 "RegexMatchSetId",
	"aws_wafregional_regex_pattern_set":               "RegexPatternSetId",
	"aws_wafregional_rule":                            "RuleId",
	"aws_wafregional_rule_group":                      "RuleGroupId",
	"aws_wafregional_size_constraint_set":             "SizeConstraintSetId",
	"aws_wafregional_sql_injection_match_set":         "SqlInjectionMatchSetId",
	"aws_wafregional_web_acl":                         "WebACLId",
	"aws_wafregional_xss_match_set":                   "XssMatchSetId",
	"aws_worklink_fleet":                              "FleetArn",
	"aws_workspaces_ip_group":                         "GroupId",
	"aws_xray_sampling_rule":                          "RuleName",
}
