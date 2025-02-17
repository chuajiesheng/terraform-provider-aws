package acmpca_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccACMPCACertificateAuthorityDataSource_basic(t *testing.T) {
	resourceName := "aws_acmpca_certificate_authority.test"
	datasourceName := "data.aws_acmpca_certificate_authority.test"

	commonName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, acmpca.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCertificateAuthorityDataSourceConfig_nonExistent,
				ExpectError: regexp.MustCompile(`(AccessDeniedException|ResourceNotFoundException)`),
			},
			{
				Config: testAccCertificateAuthorityDataSourceConfig_arn(commonName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "certificate", resourceName, "certificate"),
					resource.TestCheckResourceAttrPair(datasourceName, "certificate_chain", resourceName, "certificate_chain"),
					resource.TestCheckResourceAttrPair(datasourceName, "certificate_signing_request", resourceName, "certificate_signing_request"),
					resource.TestCheckResourceAttrPair(datasourceName, "not_after", resourceName, "not_after"),
					resource.TestCheckResourceAttrPair(datasourceName, "not_before", resourceName, "not_before"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.#", resourceName, "revocation_configuration.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.#", resourceName, "revocation_configuration.0.crl_configuration.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.0.enabled", resourceName, "revocation_configuration.0.crl_configuration.0.enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "serial", resourceName, "serial"),
					resource.TestCheckResourceAttrPair(datasourceName, "status", resourceName, "status"),
					resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
					resource.TestCheckResourceAttrPair(datasourceName, "type", resourceName, "type"),
					resource.TestCheckResourceAttrPair(datasourceName, "usage_mode", resourceName, "usage_mode"),
				),
			},
		},
	})
}

func TestAccACMPCACertificateAuthorityDataSource_s3ObjectACL(t *testing.T) {
	resourceName := "aws_acmpca_certificate_authority.test"
	datasourceName := "data.aws_acmpca_certificate_authority.test"

	commonName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, acmpca.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCertificateAuthorityDataSourceConfig_nonExistent,
				ExpectError: regexp.MustCompile(`(AccessDeniedException|ResourceNotFoundException)`),
			},
			{
				Config: testAccCertificateAuthorityDataSourceConfig_s3ObjectACLARN(commonName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "certificate", resourceName, "certificate"),
					resource.TestCheckResourceAttrPair(datasourceName, "certificate_chain", resourceName, "certificate_chain"),
					resource.TestCheckResourceAttrPair(datasourceName, "certificate_signing_request", resourceName, "certificate_signing_request"),
					resource.TestCheckResourceAttrPair(datasourceName, "not_after", resourceName, "not_after"),
					resource.TestCheckResourceAttrPair(datasourceName, "not_before", resourceName, "not_before"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.#", resourceName, "revocation_configuration.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.#", resourceName, "revocation_configuration.0.crl_configuration.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.0.enabled", resourceName, "revocation_configuration.0.crl_configuration.0.enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.0.custom_cname", resourceName, "revocation_configuration.0.crl_configuration.0.custom_cname"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.0.expiration_in_days", resourceName, "revocation_configuration.0.crl_configuration.0.expiration_in_days"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.0.s3_bucket_name", resourceName, "revocation_configuration.0.crl_configuration.0.s3_bucket_name"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.0.s3_object_acl", resourceName, "revocation_configuration.0.crl_configuration.0.s3_object_acl"),
					resource.TestCheckResourceAttrPair(datasourceName, "serial", resourceName, "serial"),
					resource.TestCheckResourceAttrPair(datasourceName, "status", resourceName, "status"),
					resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
					resource.TestCheckResourceAttrPair(datasourceName, "type", resourceName, "type"),
					resource.TestCheckResourceAttrPair(datasourceName, "usage_mode", resourceName, "usage_mode"),
				),
			},
		},
	})
}

func TestAccACMPCACertificateAuthorityDataSource_ramShared(t *testing.T) {
	resourceName := "aws_acmpca_certificate_authority.test"
	datasourceName := "data.aws_acmpca_certificate_authority.test"

	commonName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, acmpca.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCertificateAuthorityDataSourceConfig_nonExistent,
				ExpectError: regexp.MustCompile(`(AccessDeniedException|ResourceNotFoundException)`),
			},
			{
				Config: testAccCertificateAuthorityDataSourceConfig_ramShared(commonName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "certificate", resourceName, "certificate"),
					resource.TestCheckResourceAttrPair(datasourceName, "certificate_chain", resourceName, "certificate_chain"),
					resource.TestCheckResourceAttr(resourceName, "certificate_signing_request", ""),
					resource.TestCheckResourceAttrPair(datasourceName, "not_after", resourceName, "not_after"),
					resource.TestCheckResourceAttrPair(datasourceName, "not_before", resourceName, "not_before"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.#", resourceName, "revocation_configuration.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.#", resourceName, "revocation_configuration.0.crl_configuration.#"),
					resource.TestCheckResourceAttrPair(datasourceName, "revocation_configuration.0.crl_configuration.0.enabled", resourceName, "revocation_configuration.0.crl_configuration.0.enabled"),
					resource.TestCheckResourceAttrPair(datasourceName, "serial", resourceName, "serial"),
					resource.TestCheckResourceAttrPair(datasourceName, "status", resourceName, "status"),
					resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
					resource.TestCheckResourceAttrPair(datasourceName, "type", resourceName, "type"),
					resource.TestCheckResourceAttrPair(datasourceName, "usage_mode", resourceName, "usage_mode"),
				),
			},
		},
	})
}

func testAccCertificateAuthorityDataSourceConfig_arn(commonName string) string {
	return fmt.Sprintf(`
resource "aws_acmpca_certificate_authority" "wrong" {
  permanent_deletion_time_in_days = 7

  certificate_authority_configuration {
    key_algorithm     = "RSA_4096"
    signing_algorithm = "SHA512WITHRSA"

    subject {
      common_name = %[1]q
    }
  }
}

resource "aws_acmpca_certificate_authority" "test" {
  permanent_deletion_time_in_days = 7

  certificate_authority_configuration {
    key_algorithm     = "RSA_4096"
    signing_algorithm = "SHA512WITHRSA"

    subject {
      common_name = %[1]q
    }
  }
}

data "aws_acmpca_certificate_authority" "test" {
  arn = aws_acmpca_certificate_authority.test.arn
}
`, commonName)
}

func testAccCertificateAuthorityDataSourceConfig_s3ObjectACLARN(commonName string) string {
	return fmt.Sprintf(`
resource "aws_acmpca_certificate_authority" "wrong" {
  permanent_deletion_time_in_days = 7

  certificate_authority_configuration {
    key_algorithm     = "RSA_4096"
    signing_algorithm = "SHA512WITHRSA"

    subject {
      common_name = %[1]q
    }
  }
}

resource "aws_acmpca_certificate_authority" "test" {
  permanent_deletion_time_in_days = 7

  certificate_authority_configuration {
    key_algorithm     = "RSA_4096"
    signing_algorithm = "SHA512WITHRSA"

    subject {
      common_name = %[1]q
    }
  }
}

data "aws_acmpca_certificate_authority" "test" {
  arn = aws_acmpca_certificate_authority.test.arn
}
`, commonName)
}

func testAccCertificateAuthorityDataSourceConfig_ramShared(commonName string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}

resource "aws_acmpca_certificate_authority" "alternate" {
  provider = "awsalternate"

  certificate_authority_configuration {
    key_algorithm     = "RSA_4096"
    signing_algorithm = "SHA512WITHRSA"

    subject {
      common_name = %[1]q
    }
  }
}

resource "aws_ram_resource_share" "alternate" {
  provider = "awsalternate"

  name                      = "alternate"
  allow_external_principals = true
  permission_arns           = ["arn:aws:ram::aws:permission/AWSRAMDefaultPermissionCertificateAuthority"]
}

resource "aws_ram_principal_association" "alternate" {
  provider = "awsalternate"

  resource_share_arn = aws_ram_resource_share.alternate.arn
  principal          = data.aws_caller_identity.current.account_id
}

resource "aws_ram_resource_association" "alternate" {
  provider = "awsalternate"

  resource_share_arn = aws_ram_resource_share.alternate.arn
  resource_arn       = aws_acmpca_certificate_authority.alternate.arn
}

data "aws_acmpca_certificate_authority" "test" {
  arn = aws_acmpca_certificate_authority.alternate.arn
}
`, commonName)
}

// lintignore:AWSAT003,AWSAT005
const testAccCertificateAuthorityDataSourceConfig_nonExistent = `
data "aws_acmpca_certificate_authority" "test" {
  arn = "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/tf-acc-test-does-not-exist"
}
`
