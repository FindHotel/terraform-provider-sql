# terraform-provider-sql

Terraform provider for managing SQL schemas using migrations.

This plugin uses [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate),
it is recommended to go read how it works before using this provider.

## Usage

### Installation


Build the provider and put it in Terraform's third-party providers directory in `~/.terraform.d/plugins`:

#### Terraform <0.12

```bash
go get github.com/FindHotel/terraform-provider-sql
mkdir -p ~/.terraform.d/plugins
go build -o ~/.terraform.d/plugins/terraform-provider-sql github.com/FindHotel/terraform-provider-sql
```

#### Terraform >0.13


```bash
go get github.com/FindHotel/terraform-provider-sql
mkdir -p ~/.terraform.d/plugins
go build -o ~/.terraform.d/plugins/<your-module-url-specified-in-tf-file>/<version-specified-in-tf-file>/linux_amd64/terraform-provider-sql github.com/FindHotel/terraform-provider-sql
```

I recommend using [Go modules](https://github.com/golang/go/wiki/Modules) to ensure
using the same version in development and production.

### Configuration

In your Terraform configuration:

```terraform
resource "sql_schema" "this" {
  driver     = "<database driver>" # mysql/postgres/cloudsql/cloudsqlpostgres
  datasource = "<database connection string>"
  directory  = "migrations" # optional
  table      = "schema_migrations" # optional
}
```
