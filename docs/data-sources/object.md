---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "denodo_object Data Source - terraform-provider-denodo"
subcategory: "Data Source"
description: |-
Provides metadata information on objects in Denodo platform.
---

# denodo_object (Data Source)


## Example Usage

```terraform
resource "denodo_object" "dv" {
  database  = var.database
  object_name = var.object_name
  object_type = var.object_type
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **database** (String) Database name.

### Optional

- **id** (String) The ID of this resource.
- **object_name** (String) Object name.
- **object_type** (String) Type of object to retrieve. Folders, DataSources, StoredProcedures, Wrappers, Views, WebServices, Widgets, Associations, JMSListeners

### Read-Only

- **objects** (List of Object) (see [below for nested schema](#nestedatt--objects))

<a id="nestedatt--objects"></a>
### Nested Schema for `objects`

Read-Only:

- **catalog_id** (String)
- **create_date** (String)
- **database_name** (String)
- **description** (String)
- **folder** (String)
- **last_modification_date** (String)
- **last_user_modifier** (String)
- **object_name** (String)
- **sub_type** (String)
- **type** (String)
- **user_creator** (String)


