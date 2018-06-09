package data

import (
	"encoding/xml"
)

// This struct contains the entire data file.  It's everything that is inside
// osss:OnlineSocialSecurityStatementData.
type Statement struct {
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
}

