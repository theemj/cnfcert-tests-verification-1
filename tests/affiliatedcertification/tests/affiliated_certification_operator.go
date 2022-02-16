package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/test-network-function/cnfcert-tests-verification/tests/affiliatedcertification/affiliatedcerthelper"
	"github.com/test-network-function/cnfcert-tests-verification/tests/affiliatedcertification/affiliatedcertparameters"
	"github.com/test-network-function/cnfcert-tests-verification/tests/globalhelper"
	"github.com/test-network-function/cnfcert-tests-verification/tests/globalparameters"
	"github.com/test-network-function/cnfcert-tests-verification/tests/utils/execute"
)

var _ = Describe("Affiliated-certification operator certification,", func() {

	execute.BeforeAll(func() {

	})

	BeforeEach(func() {

	})

	// 46582
	It("one operator to test, operator is certified", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			[]string{affiliatedcertparameters.CertifiedOperatorApicast}, globalparameters.TestCasePassed)
		Expect(err).ToNot(HaveOccurred())
	})

	// 46695
	It("one operator to test, operator is not certified [negative]", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			[]string{affiliatedcertparameters.UncertifiedOperatorBarFoo}, globalparameters.TestCaseFailed)
		Expect(err).ToNot(HaveOccurred())
	})

	// 46696
	It("two operators to test, both are certified", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			[]string{affiliatedcertparameters.CertifiedOperatorApicast,
				affiliatedcertparameters.CertifiedOperatorKubeturbo}, globalparameters.TestCasePassed)
		Expect(err).ToNot(HaveOccurred())
	})

	// 46697
	It("two operators to test, one is certified, one is not [negative]", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			[]string{affiliatedcertparameters.CertifiedOperatorApicast,
				affiliatedcertparameters.UncertifiedOperatorBarFoo}, globalparameters.TestCaseFailed)
		Expect(err).ToNot(HaveOccurred())
	})

	// 46698
	It("certifiedoperatorinfo field exists in tnf_config but has no value [skip]", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			[]string{""}, globalparameters.TestCaseSkipped)
		Expect(err).ToNot(HaveOccurred())
	})

	// 46699
	It("Affiliated-certification one operator to test, operator does not belong to certified-operators organization"+
		" in Red Hat catalog [skip]", func() {
		By("Add container information to " + globalparameters.DefaultTnfConfigFileName)
		err := globalhelper.DefineTnfConfig(
			[]string{"tnf"},
			[]string{affiliatedcertparameters.TestPodLabel},
			[]string{},
			[]string{})
		Expect(err).ToNot(HaveOccurred())

		m := make(map[string]string)
		m["test-network-function.com/operator"] = "target"

		By("Label operator to be certified")
		err = affiliatedcerthelper.AddLabelToInstalledCSV(
			"performance-addon-operator",
			"tnf",
			m)

		Expect(err).ToNot(HaveOccurred())

		err = globalhelper.LaunchTests(
			[]string{affiliatedcertparameters.AffiliatedCertificationTestSuiteName},
			affiliatedcertparameters.TestCaseOperatorSkipRegEx,
		)
		Expect(err).ToNot(HaveOccurred())

		By("Verify test case status in Junit and Claim reports")
		err = globalhelper.ValidateIfReportsAreValid(
			affiliatedcertparameters.TestCaseOperatorAffiliatedCertName,
			globalparameters.TestCaseSkipped)
		Expect(err).ToNot(HaveOccurred(), "Error validating test reports")

	})

	// 46700
	It("name and organization fields exist in certifiedoperatorinfo but are empty [skip]", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			[]string{affiliatedcertparameters.EmptyFieldsContainerOrOperator}, globalparameters.TestCaseSkipped)
		Expect(err).ToNot(HaveOccurred())
	})

	// 46702
	It("name field in certifiedoperatorinfo field is populated but organization field is not [skip]", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			[]string{affiliatedcertparameters.OperatorNameOnlyKubeturbo}, globalparameters.TestCaseSkipped)
		Expect(err).ToNot(HaveOccurred())
	})

	// 46704
	It("organization field in certifiedoperatorinfo field is populated but name field is not [skip]", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			[]string{affiliatedcertparameters.OperatorOrgOnlyCertifiedOperators}, globalparameters.TestCaseSkipped)
		Expect(err).ToNot(HaveOccurred())
	})

	// 46706
	It("two operators to test, one is certified, one has empty name and organization fields", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			[]string{affiliatedcertparameters.CertifiedOperatorApicast,
				affiliatedcertparameters.EmptyFieldsContainerOrOperator}, globalparameters.TestCasePassed)
		Expect(err).ToNot(HaveOccurred())
	})

})
