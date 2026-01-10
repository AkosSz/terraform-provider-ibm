// This output allows distribution_list_destination data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_distribution_list_destination" {
  value       = ibm_distribution_list_destination.distribution_list_destination_instance
  description = "distribution_list_destination resource instance"
}
