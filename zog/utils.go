package zog

func check[T any](v T, checks []func(T) error) error {
	for _, check := range checks {
		if err := check(v); err != nil {
			return err
		}
	}
	return nil
}
