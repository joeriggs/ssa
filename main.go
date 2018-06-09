// This is a sample program that shows how to use the various sub-packages in
// the ssa package.

package main

import (
	"fmt"
	"github.com/joeriggs/ssa/data"
)

func main() {

	statement := data.New("Your_Social_Security_Statement_Data.xml")

	fmt.Printf("Name:                         %v\n", statement.Name())
	fmt.Printf("DateOfBirth:                  %v\n", statement.DateOfBirth())

	var ageYears, ageMonths, amount1, amount2 int

	ageYears, amount1 = statement.EarlyRetirement()
	fmt.Printf("Early Retirement:                %9v.  %3v years old.\n", amount1, ageYears)

	ageYears, ageMonths, amount1 = statement.FullRetirement()
	fmt.Printf("Full Retirement:                 %9v.  %3v years %v months old.\n", amount1, ageYears, ageMonths)

	ageYears, amount1 = statement.DelayedRetirement()
	fmt.Printf("Delayed Retirement:              %9v.  %3v years old.\n", amount1, ageYears)

	amount1 = statement.DisabilityEstimate()
	fmt.Printf("DisabilityEstimate:              %9v.\n", amount1)

	amount1 = statement.OneTimeDeathBenefit()
	fmt.Printf("OneTimeDeathBenefit:             %9v.\n", amount1)

	amount1 = statement.SurvivorsEstimateChild()
	fmt.Printf("SurvivorsEstimateChild:          %9v.\n", amount1)

	amount1 = statement.SurvivorsEstimateFamily()
	fmt.Printf("SurvivorsEstimateFamily:         %9v.\n", amount1)

	amount1 = statement.SurvivorsEstimateRetired()
	fmt.Printf("SurvivorsEstimateRetired:        %9v.\n", amount1)

	amount1 = statement.SurvivorsEstimateSpouseChild()
	fmt.Printf("SurvivorsEstimateSpouseChild:    %9v.\n", amount1)

	amount1, amount2 = statement.FicaTaxTotals()
	fmt.Printf("FicaTaxTotalEmployer:            %9v.\n", amount1)
	fmt.Printf("FicaTaxTotalIndividual:          %9v.\n", amount2)

	amount1, amount2 = statement.MedicareTaxTotals()
	fmt.Printf("MedicareTaxTotalEmployer:        %9v.\n", amount1)
	fmt.Printf("MedicareTaxTotalIndividual:      %9v.\n", amount2)

	var numEarningsYears = statement.NumEarningsYears()
	fmt.Printf("Earnings Record (%v years):\n", numEarningsYears)
	fmt.Printf("                    FICA     Medicare\n")
	fmt.Printf("          Year    Earnings   Earnings\n")

	for i := 0; i < numEarningsYears; i++ {
		var year, ficaEarnings, medicareEarnings int
		year, ficaEarnings, medicareEarnings = statement.EarningsYear(i)
                fmt.Printf("          %4v   %9v  %9v\n", year, ficaEarnings, medicareEarnings)
	}
}

