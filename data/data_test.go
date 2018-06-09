package data

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestData(t *testing.T) {

	// Open our xmlFile.  Print an error message if we fail.
	xmlFile, err := os.Open("Your_Social_Security_Statement_Data_1.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened xml file")

	// Defer closing the file.
	defer xmlFile.Close()

	// Read the file as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// Initialize the Statement structure.
	var Statement Statement

	// Unmarshal the byteArray that contains the contents of the file
	// into 'Statement' (which we defined above).
	xml.Unmarshal(byteValue, &Statement)

	fmt.Printf("Name:                         %v\n", Statement.UserInfo.Name)
	fmt.Printf("DateOfBirth:                  %v\n", Statement.UserInfo.DateOfBirth)

	fmt.Printf("Early Retirement:             %9v.  %3v years old.\n",
	            Statement.EstimatedBenefits.EarlyRetirementEstimate.Estimate,
	            Statement.EstimatedBenefits.EarlyRetirementEstimate.RetirementAge.Years)
	fmt.Printf("Full Retirement:              %9v.  %3v years old.\n",
	            Statement.EstimatedBenefits.FullRetirementEstimate.Estimate,
	            Statement.EstimatedBenefits.FullRetirementEstimate.RetirementAge.Years)
	fmt.Printf("Delayed Retirement:           %9v.  %3v years old.\n\n",
	            Statement.EstimatedBenefits.DelayedRetirementEstimate.Estimate,
	            Statement.EstimatedBenefits.DelayedRetirementEstimate.RetirementAge.Years)

	fmt.Printf("DisabilityEstimate:           %9v.\n",
	            Statement.EstimatedBenefits.DisabilityEstimate)
	fmt.Printf("OneTimeDeathBenefit:          %9v.\n",
	            Statement.EstimatedBenefits.OneTimeDeathBenefit)
	fmt.Printf("SurvivorsEstimateChild:       %9v.\n",
	            Statement.EstimatedBenefits.SurvivorsEstimateChild)
	fmt.Printf("SurvivorsEstimateFamily:      %9v.\n",
	            Statement.EstimatedBenefits.SurvivorsEstimateFamily)
	fmt.Printf("SurvivorsEstimateRetired:     %9v.\n",
	            Statement.EstimatedBenefits.SurvivorsEstimateRetired)
	fmt.Printf("SurvivorsEstimateSpouseChild: %9v.\n\n",
	            Statement.EstimatedBenefits.SurvivorsEstimateSpouseChild)

	fmt.Printf("FicaTaxTotalEmployer:         %9v.\n",
	            Statement.EarningsRecord.FicaTaxTotalEmployer)
	fmt.Printf("FicaTaxTotalIndividual:       %9v.\n",
	            Statement.EarningsRecord.FicaTaxTotalIndividual)
	fmt.Printf("MedicareTaxTotalEmployer:     %9v.\n",
	            Statement.EarningsRecord.MedicareTaxTotalEmployer)
	fmt.Printf("MedicareTaxTotalIndividual:   %9v.\n\n",
	            Statement.EarningsRecord.MedicareTaxTotalIndividual)

	fmt.Printf("Earnings Record (%v years):\n", len(Statement.EarningsRecord.Earnings))
	fmt.Printf("                    FICA     Medicare\n")
	fmt.Printf("          Year    Earnings   Earnings\n")
	for i := 0; i < len(Statement.EarningsRecord.Earnings); i++ {
                fmt.Printf("          %4v   %9v  %9v\n",
		            Statement.EarningsRecord.Earnings[i].Year,
		            Statement.EarningsRecord.Earnings[i].FicaEarnings,
		            Statement.EarningsRecord.Earnings[i].MedicareEarnings)
	}
}

