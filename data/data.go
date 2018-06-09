package data

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// This struct contains the entire data file.  It's everything that is inside
// osss:OnlineSocialSecurityStatementData.
type statement struct {
	XMLName             xml.Name          `xml:"OnlineSocialSecurityStatementData"`
	UserInfo            UserInfo          `xml:"UserInformation"`
	EstimatedBenefits   EstimatedBenefits `xml:"EstimatedBenefits"`
	EarningsRecord      EarningsRecord    `xml:"EarningsRecord"`
}

// This data structure defines the contents of osss:UserInformation.
type UserInfo struct {
	XMLName     xml.Name `xml:"UserInformation"`
	Name        string   `xml:"Name"`
	DateOfBirth string   `xml:"DateOfBirth"`
}

// This data structure defines the contents of osss:EstimatedBenefits.
type EstimatedBenefits struct {
	XMLName                      xml.Name                  `xml:"EstimatedBenefits"`
	EarlyRetirementEstimate      EarlyRetirementEstimate   `xml:"EarlyRetirementEstimate"`
	FullRetirementEstimate       FullRetirementEstimate    `xml:"FullRetirementEstimate"`
	DelayedRetirementEstimate    DelayedRetirementEstimate `xml:"DelayedRetirementEstimate"`
	DisabilityEstimate           int                       `xml:"DisabilityEstimate"`
	OneTimeDeathBenefit          int                       `xml:"OneTimeDeathBenefit"`
	SurvivorsEstimateChild       int                       `xml:"SurvivorsEstimateChild"`
	SurvivorsEstimateFamily      int                       `xml:"SurvivorsEstimateFamily"`
	SurvivorsEstimateRetired     int                       `xml:"SurvivorsEstimateRetired"`
	SurvivorsEstimateSpouseChild int                       `xml:"SurvivorsEstimateSpouseChild"`
}

// This data structure defines the contents of osss:EarningsRecord.  It contains an
// array of entries for each year that the person worked.
type EarningsRecord struct {
	XMLName                     xml.Name  `xml:"EarningsRecord"`
	Earnings                   []Earnings `xml:"Earnings"`
	FicaTaxTotalEmployer       int        `xml:"FicaTaxTotalEmployer"`
	FicaTaxTotalIndividual     int        `xml:"FicaTaxTotalIndividual"`
	MedicareTaxTotalEmployer   int        `xml:"MedicareTaxTotalEmployer"`
	MedicareTaxTotalIndividual int        `xml:"MedicareTaxTotalIndividual"`
}

type Earnings struct {
	XMLName          xml.Name `xml:"Earnings"`
	Year             int      `xml:"startYear,attr"`
	FicaEarnings     int      `xml:"FicaEarnings"`
	MedicareEarnings int      `xml:"MedicareEarnings"`
}

type EarlyRetirementEstimate struct {
	XMLName       xml.Name      `xml:"EarlyRetirementEstimate"`
	RetirementAge RetirementAge `xml:"RetirementAge"`
	Estimate      int           `xml:"Estimate"`
}

type FullRetirementEstimate struct {
	XMLName       xml.Name      `xml:"FullRetirementEstimate"`
	RetirementAge RetirementAge `xml:"RetirementAge"`
	Estimate      int           `xml:"Estimate"`
}

type DelayedRetirementEstimate struct {
	XMLName       xml.Name      `xml:"DelayedRetirementEstimate"`
	RetirementAge RetirementAge `xml:"RetirementAge"`
	Estimate      int           `xml:"Estimate"`
}

type RetirementAge struct {
	XMLName xml.Name `xml:"RetirementAge"`
	Years   int      `xml:"Years"`
	Months  int      `xml:"Months"`
}

/*******************************************************************************
 * This is the beginning of the public API.
 ******************************************************************************/

func New(fileName string) statement {
	// Initialize the Statement structure.
	statement := statement { }

	// Open our xmlFile.  Print an error message if we fail.
	xmlFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Printf("Successfully Opened xml file '%v'.\n", fileName)

		// Defer closing the file.
		defer xmlFile.Close()

		// Read the file as a byte array.
		byteValue, _ := ioutil.ReadAll(xmlFile)

		// Unmarshal the byteArray that contains the contents of the file
		// into 'Statement' (which we defined above).
		xml.Unmarshal(byteValue, &statement)
	}

	return statement
}

// Return the person's name.
func (s statement) Name() string {
	return s.UserInfo.Name
}

// Return the person's date of birth.
func (s statement) DateOfBirth() string {
	return s.UserInfo.DateOfBirth
}

// Return the person's early retirement data.
func (s statement) EarlyRetirement() (int, int) {
	return s.EstimatedBenefits.EarlyRetirementEstimate.RetirementAge.Years,
	       s.EstimatedBenefits.EarlyRetirementEstimate.Estimate
}

// Return the person's full retirement data.  Note that it returns 3 values,
// because the "age" is given as years and months.
func (s statement) FullRetirement() (int, int, int) {
	return s.EstimatedBenefits.FullRetirementEstimate.RetirementAge.Years,
	       s.EstimatedBenefits.FullRetirementEstimate.RetirementAge.Months,
	       s.EstimatedBenefits.FullRetirementEstimate.Estimate
}

// Return the person's delayed retirement data.
func (s statement) DelayedRetirement() (int, int) {
	return s.EstimatedBenefits.DelayedRetirementEstimate.RetirementAge.Years,
	       s.EstimatedBenefits.DelayedRetirementEstimate.Estimate
}

// Return the person's disability benefit.
func (s statement) DisabilityEstimate() int {
	return s.EstimatedBenefits.DisabilityEstimate
}

// Return the person's one-time death benefit.
func (s statement) OneTimeDeathBenefit() int {
	return s.EstimatedBenefits.OneTimeDeathBenefit
}

// Return the surviving child estimated benefit.
func (s statement) SurvivorsEstimateChild() int {
	return s.EstimatedBenefits.SurvivorsEstimateChild
}

// Return the surviving family estimated benefit.
func (s statement) SurvivorsEstimateFamily() int {
	return s.EstimatedBenefits.SurvivorsEstimateFamily
}

// Return the surviving spouse's estimated retirement benefit.
func (s statement) SurvivorsEstimateRetired() int {
	return s.EstimatedBenefits.SurvivorsEstimateRetired
}

// Return the surviving spouse and surviving child benefit.
func (s statement) SurvivorsEstimateSpouseChild() int {
	return s.EstimatedBenefits.SurvivorsEstimateSpouseChild
}

// Return the person's FICA totals.
func (s statement) FicaTaxTotals() (int, int) {
	return s.EarningsRecord.FicaTaxTotalEmployer,
	       s.EarningsRecord.FicaTaxTotalIndividual
}

// Return the person's Medicare totals.
func (s statement) MedicareTaxTotals() (int, int) {
	return s.EarningsRecord.MedicareTaxTotalEmployer,
	       s.EarningsRecord.MedicareTaxTotalIndividual
}

// Return the number of earnings years that the person has.
func (s statement) NumEarningsYears() int {
	return len(s.EarningsRecord.Earnings)
}

// Return the earnings information for the specified year.  Note that the
// year is specified as an index.
func (s statement) EarningsYear(index int) (int, int, int) {
	return s.EarningsRecord.Earnings[index].Year,
	       s.EarningsRecord.Earnings[index].FicaEarnings,
	       s.EarningsRecord.Earnings[index].MedicareEarnings
}

