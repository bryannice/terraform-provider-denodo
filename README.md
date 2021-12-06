# Terraform Provider Denodo

## Development Environment

The development environment used is a docker-compose stack in the deployments folder. It spins up containers for postgresql, terraform development environment, denodo virtual data port, and denodo design studio. This allows testing for the provider running terraform apply and destroy commands. The postgresql container acts as a data source where it holds sample data within the northwinds database. The denodo version is 8.0 using the express license.

+ Postgresql 14.1
+ Terraform
  + Go Lang
  + Gopher Notes
+ Denodo: Virtual Data Port
+ Denodo: Design Studio

## Make Targets

+ `clean-build` clean up terraform provider build artifacts
+ `clean-examples` clean up terraform examples used for testing the provider
+ `destory-examples` tear down test examples
+ `dev-env-up` spin up docker compose development environment
+ `dev-env-down` tear down docker compose development environment
+ `docs` generate terraform markdown documents
+ `go-fmt` format go lang code
+ `install` build provider and install it
+ `test` test go lang code
+ `test-examples` use examples to test provider
+ `tf-fmt` format terraform files used in test folder

## Permissions Documentation

+ [User and Access Right in Virtual DataPort](https://community.denodo.com/docs/html/browse/latest/en/vdp/administration/databases_users_and_access_rights_in_virtual_dataport/user_and_access_right_in_virtual_dataport/user_and_access_right_in_virtual_dataport#insert-update-and-delete-privileges)
+ [Data Catalog Permissions](https://community.denodo.com/docs/html/browse/latest/en/vdp/data_catalog/authorization/authorization#dc-authorization)
+ [Diagnostics and Monitoring Permissions](https://community.denodo.com/docs/html/browse/latest/en/vdp/dmt/authorization/authorization#dmt-authorization)

## References

+ [Call APIs with Terraform Providers](https://learn.hashicorp.com/collections/terraform/providers)
+ [Creating A DSN-Less Connection](https://community.denodo.com/docs/html/browse/8.0/en/vdp/developer/access_through_odbc/creating_a_dsn_less_connection/creating_a_dsn_less_connection)
+ [Creating A Terraform Provider - Part 1](https://medium.com/spaceapetech/creating-a-terraform-provider-part-1-ed12884e06d7#:~:text=To%20create%20a%20Terraform%20provider,the%20lifecycle%20of%20the%20resources.)
+ [Creating A Terraform Provider - Part 2](https://medium.com/spaceapetech/creating-a-terraform-provider-part-2-1346f89f082c)
+ [Writing Custom Terraform Providers](https://www.hashicorp.com/blog/writing-custom-terraform-providers)
+ [Denodo Creating a DSN-Less Connection](https://community.denodo.com/docs/html/browse/8.0/en/vdp/developer/access_through_odbc/creating_a_dsn_less_connection/creating_a_dsn_less_connection)
+ [Denodo Virtual DataPort VQL Guide v8.0](https://community.denodo.com/docs/html/browse/8.0/en/vdp/vql/index)
+ [Getting Started With pgx Through database/sql](https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql)
+ [pgx - PostgreSQL Driver and Toolkit](https://pkg.go.dev/github.com/jackc/pgx/v4)
+ [Retrieving Result Sets](http://go-database-sql.org/retrieving.html)
+ [Package SQL](https://golang.org/pkg/database/sql/)
+ [Terraform Provider Hashicups](https://github.com/hashicorp/terraform-provider-hashicups)