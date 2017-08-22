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


//Calculate the Pay Period Start Date Based on a 7 day work week
func GetPayPeriodStartDate(processingDate time.Time, payPeriodStartDow time.Weekday) time.Time {
	payPeriodStartDate := processingDate;

	//if the payPeriodStartDate DOW = payPeriodStartDow then we are done
	//otherwise subtract a day from the payPeriodStartDate and try again
	for (payPeriodStartDate.Weekday() != payPeriodStartDow) {
		payPeriodStartDate = payPeriodStartDate.AddDate(0, 0, -1)
	}

	return payPeriodStartDate
}

func GetPayPeriodEndDate(payPeriodStartDate time.Time, daysInPayPeriod int) time.Time {
	return payPeriodStartDate.AddDate(0,0,daysInPayPeriod - 1)
}

func GetPayPeriodExpectedPayDate(payPeriodEndDate time.Time, payDateDow time.Weekday, delayed int) time.Time {

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

	return payPeriodPayDate
}


func GetPayPeriodWeekly(processingDate time.Time, payPeriodStartDow time.Weekday, payDateDow time.Weekday, delayed int) (time.Time, time.Time, time.Time) {

	//employee gets paid every week

	//calculate pay Period Start Date
	payPeriodStartDate := GetPayPeriodStartDate(processingDate, payPeriodStartDow);

	//calculate pay Period End Date
	payPeriodEndDate := GetPayPeriodEndDate(payPeriodStartDate, DAYS_IN_A_WEEK);

	//calculate pay Period Expected Pay date
	payPeriodPayDate := GetPayPeriodExpectedPayDate(payPeriodEndDate, payDateDow, delayed)

	return payPeriodStartDate, payPeriodEndDate, payPeriodPayDate
}

func GetPayPeriodBiWeekly(processingDate time.Time, payPeriodStartDow time.Weekday, hireDate time.Time, payDateDow time.Weekday, delayed int) (time.Time, time.Time, time.Time) {

	//employee gets paid every 2 weeks

	//calculate pay Period Start Date based on 7 days in pay period
	payPeriodStartDate := GetPayPeriodStartDate(processingDate, payPeriodStartDow);

	//if hireDate is ODD, and payPeriodStartDate is ODD then 1st week of pay period
	//if hireDate is ODD and payPeriodStartDate is EVEN then 2nd week of the pay period
	//conversly
	//if hireDate is EVEN, and payPeriodStartDate is EVEN then 1st week of pay period
	//if hireDate is EVEN and payPeriodStartDate is ODD, then 2nd week of pay period
	payPeriodStartDateOddEven := GetOddEvenDate(payPeriodStartDate)
	hireDateOddEven := GetOddEven(hireDate);

	if payPeriodStartDateOddEven != hireDateOddEven {
		payPeriodStartDate = payPeriodStartDate.AddDate(0,0,-DAYS_IN_A_WEEK)
	}

	payPeriodEndDate := GetPayPeriodEndDate(payPeriodStartDate, 2*DAYS_IN_A_WEEK);

	//calculate pay Period Expected Pay date
	payPeriodPayDate := GetPayPeriodExpectedPayDate(payPeriodEndDate, payDateDow, delayed)

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
