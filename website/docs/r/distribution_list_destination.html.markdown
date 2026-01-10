---
layout: "ibm"
page_title: "IBM : ibm_distribution_list_destination"
description: |-
  Manages distribution_list_destination.
subcategory: "Distribution List"
---

# ibm_distribution_list_destination

Create, update, and delete distribution_list_destinations with this resource.

## Example Usage

```hcl
resource "ibm_distribution_list_destination" "distribution_list_destination_instance" {
  account_id = "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
  destination_id = 12345678-1234-1234-1234-123456789012
  destination_type = "event_notifications"
  email = "user@example.com"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, Forces new resource, String) The IBM Cloud account ID.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-zA-Z]{1,100}$/`.
* `destination_id` - (Optional, Forces new resource, String) The GUID of the Event Notifications instance.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$/`.
* `destination_type` - (Required, Forces new resource, String) The type of the destination.
  * Constraints: Allowable values are: `event_notifications`, `email`.
* `email` - (Optional, Forces new resource, String) The email address for the destination.
  * Constraints: The maximum length is `320` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$/`.
* `name` - (Optional, Forces new resource, String) The email name for the destination.
  * Constraints: The maximum length is `320` characters. The minimum length is `3` characters.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the distribution_list_destination.


## Import

You can import the `ibm_distribution_list_destination` resource by using `destination_id`.
The `destination_id` property can be formed from `account_id`, and `destination_id` in the following format:

<pre>
&lt;account_id&gt;/&lt;destination_id&gt;
</pre>
* `account_id`: A string in the format `a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6`. The IBM Cloud account ID.
* `destination_id`: A strfmt.UUID in the format `12345678-1234-1234-1234-123456789012`. The GUID of the Event Notifications instance.

# Syntax
<pre>
$ terraform import ibm_distribution_list_destination.distribution_list_destination &lt;account_id&gt;/&lt;destination_id&gt;
</pre>
