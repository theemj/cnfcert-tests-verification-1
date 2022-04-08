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

	var (
		installedLabeledOperators []affiliatedcertparameters.OperatorLabelInfo
		testingNamespaces         []string
	)

	execute.BeforeAll(func() {
		By("Deploy operators for testing if not already deployed")
		// falcon-operator: not in certified-operators group in catalog, for negative test cases
		if affiliatedcerthelper.IsOperatorInstalled(affiliatedcertparameters.UncertifiedOperatorPrefixFalcon,
			affiliatedcertparameters.UncertifiedOperatorDeploymentFalcon) != nil {
			err := affiliatedcerthelper.DeployOperatorSubscription(
				"falcon-operator",
				"alpha",
				affiliatedcertparameters.UncertifiedOperatorPrefixFalcon,
				affiliatedcertparameters.CommunityOperatorGroup,
				affiliatedcertparameters.OperatorSourceNamespace,
			)
			Expect(err).ToNot(HaveOccurred(), "Error deploying operator "+
				affiliatedcertparameters.UncertifiedOperatorPrefixFalcon)
			// confirm that operator is installed and ready
			Eventually(func() bool {
				err = affiliatedcerthelper.IsOperatorInstalled(affiliatedcertparameters.UncertifiedOperatorPrefixFalcon,
					affiliatedcertparameters.UncertifiedOperatorDeploymentFalcon)

				return err == nil
			}, affiliatedcertparameters.Timeout, affiliatedcertparameters.PollingInterval).Should(Equal(true),
				affiliatedcertparameters.UncertifiedOperatorPrefixFalcon+" is not ready.")
		}
		// add namespace to array to be added to tnf config file
		testingNamespaces = append(testingNamespaces, affiliatedcertparameters.UncertifiedOperatorPrefixFalcon)

		// add falcon operator info to array for cleanup in AfterEach
		installedLabeledOperators = append(installedLabeledOperators, affiliatedcertparameters.OperatorLabelInfo{
			OperatorPrefix: affiliatedcertparameters.UncertifiedOperatorPrefixFalcon,
			Namespace:      affiliatedcertparameters.UncertifiedOperatorPrefixFalcon,
			Label:          affiliatedcertparameters.OperatorLabel,
		})

		// crunchy-postgres-operator: in certified-operators group and version is certified
		if affiliatedcerthelper.IsOperatorInstalled(affiliatedcertparameters.CertifiedOperatorPrefixPostgres,
			affiliatedcertparameters.CertifiedOperatorDeploymentPostgres) != nil {
			err := affiliatedcerthelper.DeployOperatorSubscription(
				"crunchy-postgres-operator",
				"v5",
				affiliatedcertparameters.CertifiedOperatorPrefixPostgres,
				affiliatedcertparameters.CertifiedOperatorGroup,
				affiliatedcertparameters.OperatorSourceNamespace,
			)
			Expect(err).ToNot(HaveOccurred(), "Error deploying operator "+
				affiliatedcertparameters.CertifiedOperatorPrefixPostgres)
			// confirm that operator is installed and ready
			Eventually(func() bool {
				err = affiliatedcerthelper.IsOperatorInstalled(affiliatedcertparameters.CertifiedOperatorPrefixPostgres,
					affiliatedcertparameters.CertifiedOperatorDeploymentPostgres)

				return err == nil
			}, affiliatedcertparameters.Timeout, affiliatedcertparameters.PollingInterval).Should(Equal(true),
				affiliatedcertparameters.CertifiedOperatorPrefixPostgres+" is not ready.")
		}
		// add namespace to array to be added to tnf config file
		testingNamespaces = append(testingNamespaces, affiliatedcertparameters.CertifiedOperatorPrefixPostgres)

		// add postgres operator info to array for cleanup in AfterEach
		installedLabeledOperators = append(installedLabeledOperators, affiliatedcertparameters.OperatorLabelInfo{
			OperatorPrefix: affiliatedcertparameters.CertifiedOperatorPrefixPostgres,
			Namespace:      affiliatedcertparameters.CertifiedOperatorPrefixPostgres,
			Label:          affiliatedcertparameters.OperatorLabel,
		})

		// datadog-operator-certified: in certified-operatos group and version is certified
		if affiliatedcerthelper.IsOperatorInstalled(affiliatedcertparameters.CertifiedOperatorPrefixDatadog,
			affiliatedcertparameters.CertifiedOperatorDeploymentDatadog) != nil {
			err := affiliatedcerthelper.DeployOperatorSubscription(
				"datadog-operator-certified",
				"alpha",
				affiliatedcertparameters.CertifiedOperatorPrefixDatadog,
				affiliatedcertparameters.CertifiedOperatorGroup,
				affiliatedcertparameters.OperatorSourceNamespace,
			)
			Expect(err).ToNot(HaveOccurred(), "Error deploying operator "+
				affiliatedcertparameters.CertifiedOperatorPrefixDatadog)
			// confirm that operator is installed and ready
			Eventually(func() bool {
				err = affiliatedcerthelper.IsOperatorInstalled(affiliatedcertparameters.CertifiedOperatorPrefixDatadog,
					affiliatedcertparameters.CertifiedOperatorDeploymentDatadog)

				return err == nil
			}, affiliatedcertparameters.Timeout, affiliatedcertparameters.PollingInterval).Should(Equal(true),
				affiliatedcertparameters.CertifiedOperatorPrefixDatadog+" is not ready.")
		}
		// add namespace to array to be added to tnf config file
		testingNamespaces = append(testingNamespaces, affiliatedcertparameters.CertifiedOperatorPrefixDatadog)

		// add datadog operator info to array for cleanup in AfterEach
		installedLabeledOperators = append(installedLabeledOperators, affiliatedcertparameters.OperatorLabelInfo{
			OperatorPrefix: affiliatedcertparameters.CertifiedOperatorPrefixDatadog,
			Namespace:      affiliatedcertparameters.CertifiedOperatorPrefixDatadog,
			Label:          affiliatedcertparameters.OperatorLabel,
		})

		By("Add container information to " + globalparameters.DefaultTnfConfigFileName)

		err := globalhelper.DefineTnfConfig(
			testingNamespaces,
			[]string{affiliatedcertparameters.TestPodLabel},
			[]string{},
			[]string{})

		Expect(err).ToNot(HaveOccurred(), "Error defining tnf config file")
	})

	AfterEach(func() {
		By("Remove labels from operators")
		for _, info := range installedLabeledOperators {
			err := affiliatedcerthelper.DeleteLabelFromInstalledCSV(
				info.OperatorPrefix,
				info.Namespace,
				info.Label)
			Expect(err).ToNot(HaveOccurred(), "Error removing label from operator "+info.OperatorPrefix)
		}
	})

	// 46699
	It("one operator to test, operator is not in certified-operators organization [negative]",
		func() {
			By("Label operator to be certified")

			err := affiliatedcerthelper.AddLabelToInstalledCSV(
				affiliatedcertparameters.UncertifiedOperatorPrefixFalcon,
				affiliatedcertparameters.TestCertificationNameSpace,
				affiliatedcertparameters.OperatorLabel)
			Expect(err).ToNot(HaveOccurred(), "Error labeling operator "+
				affiliatedcertparameters.UncertifiedOperatorPrefixFalcon)

			By("Start test")

			err = globalhelper.LaunchTests(
				affiliatedcertparameters.AffiliatedCertificationTestSuiteName,
				globalhelper.ConvertSpecNameToFileName(CurrentGinkgoTestDescription().FullTestText),
				affiliatedcertparameters.TestCaseOperatorSkipRegEx,
			)
			Expect(err).To(HaveOccurred(), "Error running "+
				affiliatedcertparameters.AffiliatedCertificationTestSuiteName+" test")

			By("Verify test case status in Junit and Claim reports")

			err = globalhelper.ValidateIfReportsAreValid(
				affiliatedcertparameters.TestCaseOperatorAffiliatedCertName,
				globalparameters.TestCaseFailed)
			Expect(err).ToNot(HaveOccurred(), "Error validating test reports")
		})

	// 46697
	It("two operators to test, one is in certified-operators organization and its version is certified,"+
		" one is not in certified-operators organization [negative]", func() {
		By("Label operators to be certified")

		err := affiliatedcerthelper.AddLabelToInstalledCSV(
			affiliatedcertparameters.CertifiedOperatorPrefixPostgres,
			affiliatedcertparameters.TestCertificationNameSpace,
			affiliatedcertparameters.OperatorLabel)
		Expect(err).ToNot(HaveOccurred(), "Error labeling operator "+
			affiliatedcertparameters.CertifiedOperatorPrefixPostgres)

		err = affiliatedcerthelper.AddLabelToInstalledCSV(
			affiliatedcertparameters.UncertifiedOperatorPrefixFalcon,
			affiliatedcertparameters.TestCertificationNameSpace,
			affiliatedcertparameters.OperatorLabel)
		Expect(err).ToNot(HaveOccurred(), "Error labeling operator "+
			affiliatedcertparameters.UncertifiedOperatorPrefixFalcon)

		By("Start test")

		err = globalhelper.LaunchTests(
			affiliatedcertparameters.AffiliatedCertificationTestSuiteName,
			globalhelper.ConvertSpecNameToFileName(CurrentGinkgoTestDescription().FullTestText),
			affiliatedcertparameters.TestCaseOperatorSkipRegEx,
		)
		Expect(err).To(HaveOccurred(), "Error running "+
			affiliatedcertparameters.AffiliatedCertificationTestSuiteName+" test")

		By("Verify test case status in Junit and Claim reports")

		err = globalhelper.ValidateIfReportsAreValid(
			affiliatedcertparameters.TestCaseOperatorAffiliatedCertName,
			globalparameters.TestCaseFailed)
		Expect(err).ToNot(HaveOccurred(), "Error validating test reports")
	})

	// 46582
	It("one operator to test, operator is in certified-operators organization"+
		" and its version is certified", func() {
		By("Label operator to be certified")

		err := affiliatedcerthelper.AddLabelToInstalledCSV(
			affiliatedcertparameters.CertifiedOperatorPrefixPostgres,
			affiliatedcertparameters.TestCertificationNameSpace,
			affiliatedcertparameters.OperatorLabel)
		Expect(err).ToNot(HaveOccurred(), "Error labeling operator "+
			affiliatedcertparameters.CertifiedOperatorPrefixPostgres)

		By("Start test")

		err = globalhelper.LaunchTests(
			affiliatedcertparameters.AffiliatedCertificationTestSuiteName,
			globalhelper.ConvertSpecNameToFileName(CurrentGinkgoTestDescription().FullTestText),
			affiliatedcertparameters.TestCaseOperatorSkipRegEx,
		)
		Expect(err).ToNot(HaveOccurred(), "Error running "+
			affiliatedcertparameters.AffiliatedCertificationTestSuiteName+" test")

		By("Verify test case status in Junit and Claim reports")

		err = globalhelper.ValidateIfReportsAreValid(
			affiliatedcertparameters.TestCaseOperatorAffiliatedCertName,
			globalparameters.TestCasePassed)
		Expect(err).ToNot(HaveOccurred(), "Error validating test reports")
	})

	// 46696
	It("two operators to test, both are in certified-operators organization and their"+
		" versions are certified", func() {
		By("Label operators to be certified")

		err := affiliatedcerthelper.AddLabelToInstalledCSV(
			affiliatedcertparameters.CertifiedOperatorPrefixDatadog,
			affiliatedcertparameters.TestCertificationNameSpace,
			affiliatedcertparameters.OperatorLabel)
		Expect(err).ToNot(HaveOccurred(), "Error labeling operator "+
			affiliatedcertparameters.CertifiedOperatorPrefixDatadog)

		err = affiliatedcerthelper.AddLabelToInstalledCSV(
			affiliatedcertparameters.CertifiedOperatorPrefixPostgres,
			affiliatedcertparameters.TestCertificationNameSpace,
			affiliatedcertparameters.OperatorLabel)
		Expect(err).ToNot(HaveOccurred(), "Error labeling operator "+
			affiliatedcertparameters.CertifiedOperatorPrefixPostgres)

		By("Start test")

		err = globalhelper.LaunchTests(
			affiliatedcertparameters.AffiliatedCertificationTestSuiteName,
			globalhelper.ConvertSpecNameToFileName(CurrentGinkgoTestDescription().FullTestText),
			affiliatedcertparameters.TestCaseOperatorSkipRegEx,
		)
		Expect(err).ToNot(HaveOccurred(), "Error running "+
			affiliatedcertparameters.AffiliatedCertificationTestSuiteName+" test")

		By("Verify test case status in Junit and Claim reports")

		err = globalhelper.ValidateIfReportsAreValid(
			affiliatedcertparameters.TestCaseOperatorAffiliatedCertName,
			globalparameters.TestCasePassed)
		Expect(err).ToNot(HaveOccurred(), "Error validating test reports")
	})

	// 46695
	It("one operator to test, operator is not certified [negative]", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			globalhelper.ConvertSpecNameToFileName(CurrentGinkgoTestDescription().FullTestText),
			[]string{affiliatedcertparameters.UncertifiedOperatorBarFoo}, globalparameters.TestCaseFailed)
		Expect(err).ToNot(HaveOccurred())
	})

	// 46698
	It("no operators are labeled for testing [skip]", func() {
		By("Start test")

		err := globalhelper.LaunchTests(
			affiliatedcertparameters.AffiliatedCertificationTestSuiteName,
			globalhelper.ConvertSpecNameToFileName(CurrentGinkgoTestDescription().FullTestText),
			affiliatedcertparameters.TestCaseOperatorSkipRegEx,
		)
		Expect(err).ToNot(HaveOccurred(), "Error running "+
			affiliatedcertparameters.AffiliatedCertificationTestSuiteName+" test")

		By("Verify test case status in Junit and Claim reports")

		err = globalhelper.ValidateIfReportsAreValid(
			affiliatedcertparameters.TestCaseOperatorAffiliatedCertName,
			globalparameters.TestCaseSkipped)
		Expect(err).ToNot(HaveOccurred(), "Error validating test reports")
	})

	// 46700
	It("name and organization fields exist in certifiedoperatorinfo but are empty [skip]", func() {
		err := affiliatedcerthelper.SetUpAndRunOperatorCertTest(
			globalhelper.ConvertSpecNameToFileName(CurrentGinkgoTestDescription().FullTestText),
			[]string{affiliatedcertparameters.EmptyFieldsContainerOrOperator}, globalparameters.TestCaseSkipped)
		Expect(err).ToNot(HaveOccurred())
	})

})
