package main_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
	. "github.com/ntaylor95/pay-periods"
)

var _ = Describe("pay_periods", func() {

	BeforeEach(func() {
		fmt.Println("Starting Tests")
	})

	Describe("Given the Pay Period Start Day of the Week (DOW)", func() {
		Context("When a date is provided", func() {
			It("Then return pay period start date", func() {
				var flagtests = []struct {
					processingDate time.Time
					payPeriodStartDow time.Weekday
					payPeriodStartDate time.Time
				}{
					{time.Date(int(2017), time.May, int(3), int(0), int(0), int(0), int(0), time.UTC), time.Saturday, time.Date(int(2017), time.April, int(29), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.May, int(3), int(0), int(0), int(0), int(0), time.UTC), time.Tuesday, time.Date(int(2017), time.May, int(2), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(4), int(0), int(0), int(0), int(0), time.UTC), time.Saturday, time.Date(int(2017), time.July, int(1), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(4), int(0), int(0), int(0), int(0), time.UTC), time.Tuesday, time.Date(int(2017), time.July, int(4), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(4), int(0), int(0), int(0), int(0), time.UTC), time.Sunday, time.Date(int(2017), time.July, int(2), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(14), int(0), int(0), int(0), int(0), time.UTC), time.Sunday, time.Date(int(2017), time.July, int(9), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(1), int(0), int(0), int(0), int(0), time.UTC), time.Tuesday, time.Date(int(2017), time.June, int(27), int(0), int(0), int(0), int(0), time.UTC)},
				}

				for _, tt := range flagtests {
					_startDate := GetPayPeriodStartDate(tt.processingDate, tt.payPeriodStartDow)

					fmt.Printf("\nInput: \n   Pay Period Date: %d \n   Pay Period Start DOW: %d", tt.processingDate.Format(time.RFC3339), tt.payPeriodStartDow)
					fmt.Printf("\nOutput: \n   Pay Period Start Date: %d", _startDate.Format(time.RFC3339))

					Expect(_startDate).To(Equal(tt.payPeriodStartDate))
				}
			})
		})
	})

	Describe("Given Employee is paid weekly", func() {
		Context("When a date is provided", func() {
			It("Then return pay period end date and start date", func() {
				var flagtests = []struct {
					payPeriodDate time.Time
					payDateDow time.Weekday
					delay int
					payPeriodEndDow time.Weekday
					payPeriodStartDate time.Time
					payPeriodEndDate time.Time
					payDate time.Time
				}{
					{time.Date(int(2017), time.May, int(3), int(0), int(0), int(0), int(0), time.UTC), time.Friday, 0, time.Friday, time.Date(int(2017), time.April, int(29), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.May, int(5), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.May, int(5), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.May, int(3), int(0), int(0), int(0), int(0), time.UTC), time.Friday, 7, time.Friday, time.Date(int(2017), time.April, int(29), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.May, int(5), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.May, int(12), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(4), int(0), int(0), int(0), int(0), time.UTC), time.Friday, 0, time.Friday, time.Date(int(2017), time.July, int(1), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.July, int(7), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.July, int(7), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(4), int(0), int(0), int(0), int(0), time.UTC), time.Friday, 7, time.Friday, time.Date(int(2017), time.July, int(1), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.July, int(7), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.July, int(14), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(4), int(0), int(0), int(0), int(0), time.UTC), time.Friday, 0, time.Saturday, time.Date(int(2017), time.July, int(2), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.July, int(8), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.July, int(7), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(4), int(0), int(0), int(0), int(0), time.UTC), time.Friday, 7, time.Saturday, time.Date(int(2017), time.July, int(2), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.July, int(8), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.July, int(14), int(0), int(0), int(0), int(0), time.UTC)},
					{time.Date(int(2017), time.July, int(1), int(0), int(0), int(0), int(0), time.UTC), time.Wednesday, 0, time.Monday, time.Date(int(2017), time.June, int(27), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.July, int(3), int(0), int(0), int(0), int(0), time.UTC), time.Date(int(2017), time.June, int(28), int(0), int(0), int(0), int(0), time.UTC)},
				}

				for _, tt := range flagtests {
					_startDate, _endDate, _payDate := GetPayPeriodWeekly(tt.payPeriodDate, tt.payDateDow, tt.delay, tt.payPeriodEndDow)

					fmt.Printf("\nInput: \n   Pay Period Date: %d \n   Pay Date DOW: %d \n   Delay: %d \n   Pay Period End DOW: %d", tt.payPeriodDate.Format(time.RFC3339), tt.payDateDow, tt.delay, tt.payPeriodEndDow)
					fmt.Printf("\nOutput: \n   Pay Period Start Date: %d \n    Pay Period End Date: %d \n   Pay Date: %d", _startDate.Format(time.RFC3339), _endDate.Format(time.RFC3339), _payDate.Format(time.RFC3339) )

					Expect(_startDate).To(Equal(tt.payPeriodStartDate))
					Expect(_endDate).To(Equal(tt.payPeriodEndDate))
					Expect(_payDate).To(Equal(tt.payDate))
				}
			})
		})
	})

	AfterEach(func() {
		fmt.Println("\n\nEnding Tests")
		//os.RemoveAll("tmp/")
	})
})