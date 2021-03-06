---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "denodo_jdbc_data_source_object Data Source - terraform-provider-denodo"
subcategory: "Data Source"
description: |-
  Fetch objects from a JDBC data source.
---

# denodo_jdbc_data_source_object (Data Source)

Uses Denodo's `GET_JDBC_DATASOURCE_TABLES` to fetch a list of objects within a given data source.

## Example Usage

```terraform
data "denodo_jdbc_data_source_table" "jdst" {
  catalog_name = var.data_source_catalog_name
  database     = var.data_source_database
  name         = var.data_source_name
  schema_name  = var.data_source_schema_name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **database** (String) Name of the database to which the data source belongs. If null, the procedure will use the current database.
- **name** (String) Name of the data source for which you want to get the list of tables.

### Optional

- **catalog_name** (String) Name of the catalog for which you want to get the list of tables. If the data source does not support catalogs, set to null. If null and the data source does support catalogs, the procedure will return all the matching tables across all catalogs.
- **id** (String) The ID of this resource.
- **schema_name** (String) When the data source has to insert several rows into the database of this data source, it can insert them in batches. This number sets the number of queries per batch.

### Read-Only

- **objects** (List of Object) (see [below for nested schema](#nestedatt--objects))

<a id="nestedatt--objects"></a>
### Nested Schema for `objects`

Read-Only:

- **catalog_name** (String)
- **object_name** (String)
- **object_type** (String)
- **schema_name** (String)

## Reference

+ [GET_JDBC_DATASOURCE_TABLES](https://community.denodo.com/docs/html/browse/8.0/en/vdp/vql/stored_procedures/predefined_stored_procedures/get_jdbc_datasource_tables)
