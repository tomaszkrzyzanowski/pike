package pike

import (
	"log"
)

// GetGCPPermissions for GCP resources
func GetGCPPermissions(result ResourceV2) []string {
	var Permissions []string
	if result.TypeName == "resource" {
		Permissions = GetGCPResourcePermissions(result)
	} else {
		Permissions = GetGCPDataPermissions(result)
	}

	return Permissions
}

// GetGCPResourcePermissions looks up permissions required for resources
func GetGCPResourcePermissions(result ResourceV2) []string {
	TFLookup := map[string]interface{}{
		"google_compute_instance":                   googleComputeInstance,
		"google_storage_bucket":                     googleStorageBucket,
		"google_storage_bucket_object":              googleStorageBucketObject,
		"google_compute_network":                    googleComputeNetwork,
		"google_compute_subnetwork":                 googleComputeSubnetwork,
		"google_compute_address":                    googleComputeAddress,
		"google_compute_firewall":                   googleComputeFirewall,
		"google_project_iam_custom_role":            googleProjectIamCustomRole,
		"google_service_account":                    googleServiceAccount,
		"google_kms_crypto_key":                     googleKmsCryptoKey,
		"google_kms_key_ring":                       googleKmsKeyRing,
		"google_storage_bucket_acl":                 googleStorageBucketACL,
		"google_storage_bucket_iam_binding":         googleStorageBucketIamBinding,
		"google_sql_database_instance":              googleSQLDatabaseInstance,
		"google_sql_database":                       googleSQLDatabase,
		"google_sql_user":                           googleSQLUser,
		"google_compute_global_address":             googleComputeGlobalAddress,
		"google_service_networking_connection":      googleServiceNetworkingConnection,
		"google_cloudfunctions_function":            googleCloudfunctionsFunction,
		"google_cloudfunctions_function_iam_policy": googleCloudfunctionsFunctionIamPolicy,
		"google_cloudfunctions_function_iam_member": googleCloudfunctionsFunctionIamPolicy,
		"google_pubsub_topic":                       googlePubsubTopic,
		"google_sourcerepo_repository":              googleSourcerepoRepository,
		"google_project_iam_binding":                googleProjectIamBinding,
		"google_project_iam_member":                 googleProjectIamBinding,
		"google_service_account_iam_policy":         googleServiceAccountIamPolicy,
		"google_service_account_key":                googleServiceAccountKey,
		"google_compute_instance_template":          googleComputeInstanceTemplate,
		"google_compute_project_metadata_item":      googleComputeProjectMetadataItem,
		"google_container_cluster":                  googleContainerCluster,
		"google_container_node_pool":                googleContainerNodePool,
	}

	var Permissions []string

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		log.Printf("%s not implemented", result.Name)
	}

	return Permissions
}
