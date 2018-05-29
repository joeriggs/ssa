package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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
	XMLName                   xml.Name                  `xml:"EstimatedBenefits"`
	EarlyRetirementEstimate   EarlyRetirementEstimate   `xml:"EarlyRetirementEstimate"`
	FullRetirementEstimate    FullRetirementEstimate    `xml:"FullRetirementEstimate"`
	DelayedRetirementEstimate DelayedRetirementEstimate `xml:"DelayedRetirementEstimate"`
}

// This data structure defines the contents of osss:EarningsRecord.  It contains an
// array of entries for each year that the person worked.
type EarningsRecord struct {
	XMLName   xml.Name  `xml:"EarningsRecord"`
	Earnings []Earnings `xml:"Earnings"`
}

type Earnings struct {
	XMLName          xml.Name `xml:"Earnings"`
	Year             string   `xml:"startYear,attr"`
	FicaEarnings     string   `xml:"FicaEarnings"`
	MedicareEarnings string   `xml:"MedicareEarnings"`
}

type EarlyRetirementEstimate struct {
	XMLName       xml.Name      `xml:"EarlyRetirementEstimate"`
	RetirementAge RetirementAge `xml:"RetirementAge"`
	Estimate      string        `xml:"Estimate"`
}

type FullRetirementEstimate struct {
	XMLName       xml.Name      `xml:"FullRetirementEstimate"`
	RetirementAge RetirementAge `xml:"RetirementAge"`
	Estimate      string        `xml:"Estimate"`
}

type DelayedRetirementEstimate struct {
	XMLName       xml.Name      `xml:"DelayedRetirementEstimate"`
	RetirementAge RetirementAge `xml:"RetirementAge"`
	Estimate      string        `xml:"Estimate"`
}

type RetirementAge struct {
	XMLName xml.Name `xml:"RetirementAge"`
	Years   string   `xml:"Years"`
}

func main() {

	// Open our xmlFile.  Print an error message if we fail.
	xmlFile, err := os.Open("Your_Social_Security_Statement_Data.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened xml file")

	// Defer closing the file.
	defer xmlFile.Close()

	// Read the file as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// Initialize the Statement array.
	var userStatement Statement

	// Unmarshal the byteArray that contains the contents of the file
	// into 'userStatement' (which we defined above).
	xml.Unmarshal(byteValue, &userStatement)

	fmt.Println("Name:               " + userStatement.UserInfo.Name)
	fmt.Println("DateOfBirth:        " + userStatement.UserInfo.DateOfBirth)

	fmt.Println("Early Retirement:   " + userStatement.EstimatedBenefits.EarlyRetirementEstimate.RetirementAge.Years +
	                                     " years old.  $" +
	                                     userStatement.EstimatedBenefits.EarlyRetirementEstimate.Estimate)
	fmt.Println("Full Retirement:    " + userStatement.EstimatedBenefits.FullRetirementEstimate.RetirementAge.Years +
	                                     " years old.  $" +
	                                     userStatement.EstimatedBenefits.FullRetirementEstimate.Estimate)
	fmt.Println("Delayed Retirement: " + userStatement.EstimatedBenefits.DelayedRetirementEstimate.RetirementAge.Years +
	                                     " years old.  $" +
	                                     userStatement.EstimatedBenefits.DelayedRetirementEstimate.Estimate)

	for i := 0; i < len(userStatement.EarningsRecord.Earnings); i++ {
                fmt.Println("                    " +
		            userStatement.EarningsRecord.Earnings[i].Year +
		            ": FICA Earnings $" +
		            userStatement.EarningsRecord.Earnings[i].FicaEarnings +
		            ": Medicare Earnings $" +
		            userStatement.EarningsRecord.Earnings[i].MedicareEarnings)
	}
}

