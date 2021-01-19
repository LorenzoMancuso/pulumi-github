module github.com/pulumi/pulumi-github/provider/v3

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/terraform-providers/terraform-provider-github => github.com/integrations/terraform-provider-github v1.3.1-0.20210116231816-872bb967b7a1
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.17.0
	github.com/pulumi/pulumi/sdk/v2 v2.17.0
	github.com/terraform-providers/terraform-provider-github v1.3.1-0.20201201134554-f6add3ed515b
)
