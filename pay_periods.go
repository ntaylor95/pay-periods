package main

import (
	"time"
	"fmt"
)

const DAYS_IN_A_WEEK = 7
const START_DAY_OF_WEEK = time.Monday
const END_DAY_OF_WEEK = time.Sunday

const (
	_ = iota
	ADVANCED
	DELAYED
)

func main() {

}


func GetPayPeriodWeekly(processingDate time.Time, payDateDow time.Weekday, delayed int, payPeriodEndDow time.Weekday) (time.Time, time.Time, time.Time) {

	//employee gets paid every week

	//default the pay period end date to the date passed in
	payPeriodEndDate := processingDate

	//if the payPeriodEndDate DOW = payPeriodEndDOW then we are done
	//otherwise add a day to the payPeriodEndDate and try again
	for (payPeriodEndDate.Weekday() != payPeriodEndDow) {
		payPeriodEndDate = payPeriodEndDate.AddDate(0,0,1)
	}

	//calculate pay Period Start Date
	payPeriodStartDate := payPeriodEndDate.AddDate(0,0,(1-DAYS_IN_A_WEEK))

	//set pay date to the pay period end date as the default
	payPeriodPayDate := payPeriodEndDate

	//if the default pay date day of the week does not match the pay date day of the week
	//we need to calculate a new pay date
	for (int(payDateDow) != int(payPeriodPayDate.Weekday())) {
		payPeriodPayDate = payPeriodPayDate.AddDate(0, 0, -1)
	}

	//if the pay date is delayed
	if (delayed != 0) {
		payPeriodPayDate = payPeriodPayDate.AddDate(0,0,delayed)
	}

	return payPeriodStartDate, payPeriodEndDate, payPeriodPayDate
}

func GetPayPeriodBiWeekly(payPeriodDate time.Time, delayed bool, biWeeklyOddEven string, payDateDow int) (time.Time, time.Time, time.Time) {

	//employee gets paid every 2 weeks

	payPeriodDaysStartFactor := DAYS_IN_A_WEEK
	payPeriodDaysEndFactor := DAYS_IN_A_WEEK

	payPeriodDateOddEven := GetOddEvenDate(payPeriodDate)

	//Determine the start day of the week and the end day of the week
	//for the pay period - default is Monday thru Sunday
	payPeriodOffset := -2 //Friday is the last day of the pay period
	//payPeriodOffset = 0 //Sunday is last day of the pay period week

	//if hireDate is ODD, and payPeriodDate is ODD then 1st week of pay period
	//if hireDate is ODD and payPeriodDate is EVEN then 2nd week of the pay period
	//conversly
	//if hireDate is EVEN, and payPeriodDate is EVEN then 1st week of pay period
	//if hireDate is EVEN and payPeriodDate is ODD, then 2nd week of pay period
	payPeriodDow := int(payPeriodDate.Weekday())

	if payPeriodDateOddEven == biWeeklyOddEven {
		payPeriodDaysStartFactor = (payPeriodDow - 1 - payPeriodOffset)
		payPeriodDaysEndFactor = 2 * DAYS_IN_A_WEEK - payPeriodDow + payPeriodOffset
	} else {
		payPeriodDaysStartFactor = (payPeriodDow - 1) + DAYS_IN_A_WEEK - payPeriodOffset
		payPeriodDaysEndFactor = DAYS_IN_A_WEEK - payPeriodDow + payPeriodOffset
	}

	payPeriodStartDate := payPeriodDate.AddDate(0, 0, -(payPeriodDaysStartFactor))
	payPeriodEndDate := payPeriodDate.AddDate(0, 0, payPeriodDaysEndFactor)

	//set pay date to the pay period end date as the default
	payPeriodPayDate := payPeriodEndDate

	//set the pay day day of the week - the default would be Sunday = 7
	//payDateDow := 5 //PayDate of Friday

	//if the default pay date day of the week does not match the pay date day of the week
	//we need to calculate a new pay date
	if (payDateDow != int(payPeriodEndDate.Weekday())) {
		payDateOffset := DAYS_IN_A_WEEK-payDateDow
		payPeriodPayDate = payPeriodEndDate.AddDate(0, 0, -(payDateOffset))
	}

	//if the pay date is delayed and NOT advanced
	if delayed {
		payPeriodPayDate = payPeriodPayDate.AddDate(0, 0, DAYS_IN_A_WEEK)
	}

	return payPeriodStartDate, payPeriodEndDate, payPeriodPayDate
}

func GetPayPeriodSemiMonthly(payPeriodDate time.Time, delayed bool, payDateDow int) (time.Time, time.Time, time.Time) {

	//employee gets paid on fixed days in the month, the 1st and 15th
	var payPeriodStartDate time.Time
	var payPeriodEndDate time.Time

	semiMonthlyDomStart := 1
	semiMonthlyDomEnd := 15

	payPeriodDom := payPeriodDate.Day()

	if (payPeriodDom > semiMonthlyDomEnd) {
		semiMonthlyDomStart = semiMonthlyDomEnd + 1
		semiMonthlyDomEnd = GetDaysInMonth(payPeriodDate)
	}

	payPeriodStartDate = time.Date(payPeriodDate.Year(), payPeriodDate.Month(), semiMonthlyDomStart, 0, 0, 0, 0, time.UTC)
	payPeriodEndDate = time.Date(payPeriodDate.Year(), payPeriodDate.Month(), semiMonthlyDomEnd, 0, 0, 0, 0, time.UTC)

	//---------------------------------
	payPeriodPayDate := payPeriodEndDate

	//set the pay day day of the week - the default would be Sunday = 7
	//payDateDow

	//if the default pay date day of the week does not match the pay date day of the week
	//we need to calculate a new pay date
	payPeriodEndDateDow := int(payPeriodEndDate.Weekday())
	var payDateOffset int
	fmt.Println("Pay dates DOWS", payPeriodEndDateDow, payDateDow)
	if (int(payDateDow) != payPeriodEndDateDow) {
		if (payDateDow < payPeriodEndDateDow) {
			//pay period ends on a Saturday but we want to pay them on a Friday
			payDateOffset = payPeriodEndDateDow - payDateDow
		} else {
			//pay period end date is on a Monday {1} but we want to pay them on Friday {5}
			payDateOffset = DAYS_IN_A_WEEK - payDateDow + payPeriodEndDateDow
		}

		fmt.Println("pay period offset", payDateOffset)
		payPeriodPayDate = payPeriodEndDate.AddDate(0, 0, -(payDateOffset))
	}

	//if the pay date is delayed and NOT advanced
	if delayed {
		payPeriodPayDate = payPeriodPayDate.AddDate(0, 0, DAYS_IN_A_WEEK)
	}


	return payPeriodStartDate, payPeriodEndDate, payPeriodPayDate
}

//TODO: How could I extend the time.Time type to include this method
func GetOddEvenDate(dt time.Time) string {
	_, week := dt.ISOWeek()
	return GetOddEven(week);
}

func GetOddEven(val int) string {
	if val%2 == 0 {
		return "EVEN"
	}
	return "ODD"
}


//TODO: How could I extend the time.Time type to include this method
func GetDaysInMonth(dt time.Time) int {
	dt1 := time.Date(dt.Year(), dt.Month(), 27, 0, 0, 0, 0, time.UTC)

	for i := 1; i < 6; i++ {
		dt2 := dt1.AddDate(0, 0, i)
		if (dt1.Month() != dt2.Month()) {
			return dt2.AddDate(0,0,-1).Day()
		}
	}

	return 0
}
