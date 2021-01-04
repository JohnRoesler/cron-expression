package cronexpr

type Expression struct {
}

// Parse takes in a cron expression and returns an Expression
// and an error if any.
//
// The following implementation of the linux cron expression is used:
//
// # ┌───────────── minute (0 - 59)
// # │ ┌───────────── hour (0 - 23)
// # │ │ ┌───────────── day of the month (1 - 31)
// # │ │ │ ┌───────────── month (1 - 12)
// # │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
// # │ │ │ │ │                                   7 is also Sunday on some systems)
// # │ │ │ │ │
// # │ │ │ │ │
// # * * * * * <command to execute>
func Parse(expression string) (Expression, error) {
	return Expression{}, nil
}
