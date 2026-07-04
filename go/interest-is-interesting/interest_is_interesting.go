package interest

const (
	thr0    float64 = 0.0
	thr1000 float64 = 1000.0
	thr5000 float64 = 5000.0
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < thr0:
		return 3.213
	case balance >= thr0 && balance < thr1000:
		return 0.5
	case balance >= thr1000 && balance < thr5000:
		return 1.621
	default:
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100.0)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	cumulated := balance
	years := 0
	for ; cumulated < targetBalance; years++ {
		cumulated = AnnualBalanceUpdate(cumulated)
	}

	return years
}
