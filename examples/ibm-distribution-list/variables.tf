variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for distribution_list_destination
variable "distribution_list_destination_account_id" {
  description = "The IBM Cloud account ID."
  type        = string
  default     = "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
}
variable "distribution_list_destination_destination_id" {
  description = "The GUID of the Event Notifications instance."
  type        = string
  default     = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
}
variable "distribution_list_destination_destination_type" {
  description = "The type of the destination."
  type        = string
  default     = "event_notifications"
}
variable "distribution_list_destination_email" {
  description = "The email address for the destination."
  type        = string
  default     = "user@example.com"
}
variable "distribution_list_destination_name" {
  description = "The email name for the destination."
  type        = string
  default     = "name"
}
