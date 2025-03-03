package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/gcp/data/iam/google_service_account.json
var dataGoogleServiceAccount []byte

//go:embed mapping/gcp/data/compute/google_compute_network.json
var dataGoogleComputeNetwork []byte

//go:embed mapping/gcp/data/compute/google_compute_subnetwork.json
var dataGoogleComputeSubnetwork []byte

//go:embed mapping/gcp/data/compute/google_compute_zones.json
var dataGoogleComputeZones []byte

//go:embed mapping/gcp/data/resourcemanager/google_project.json
var dataGoogleProject []byte

//go:embed mapping/gcp/data/cloudkms/google_kms_key_ring.json
var dataGoogleKmsKeyRing []byte

//go:embed  mapping/gcp/data/cloudkms/google_kms_crypto_key.json
var dataGoogleKmsCryptoKey []byte
