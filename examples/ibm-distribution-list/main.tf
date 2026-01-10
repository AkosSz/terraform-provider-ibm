provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision distribution_list_destination resource instance
resource "ibm_distribution_list_destination" "distribution_list_destination_instance" {
  account_id = var.distribution_list_destination_account_id
  destination_id = var.distribution_list_destination_destination_id
  destination_type = var.distribution_list_destination_destination_type
  email = var.distribution_list_destination_email
  name = var.distribution_list_destination_name
}
