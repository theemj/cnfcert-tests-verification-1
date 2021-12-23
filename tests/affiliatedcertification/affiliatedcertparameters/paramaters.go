package affiliatedcertparameters

var (
	AffiliatedCertificationTestSuiteName = "affiliated-certification"
	TestCaseContainerSkipRegEx           = "operator-is-certified"
	TestCaseOperatorSkipRegEx            = "container-is-certified"
	TestCaseContainerAffiliatedCertName  = "affiliated-certification affiliated-certification-container-is-certified"
	CertifiedContainerNodeJsUbi          = "nodejs-12/ubi8"
	CertifiedContainerRhel7OpenJdk       = "openjdk-11-rhel7/openjdk"
	UncertifiedContainerFooBar           = "foo/bar"
	TestCaseOperatorAffiliatedCertName   = "affiliated-certification affiliated-certification-operator-is-certified"
	CertifiedOperatorApicast             = "apicast-operator/redhat-operators"
	CertifiedOperatorKubeturbo           = "kubeturbo-certified/certified-operators"
	UncertifiedOperatorBarFoo            = "bar/foo"
)
