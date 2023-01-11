package validate

import "github.com/hashicorp/go-multierror"

func CombineErrors(errs ...error) error {
	var merr *multierror.Error

	for _, err := range errs {
		if err == nil {
			return nil
		}

		if err != nil {
			merr = multierror.Append(merr, err)
		}
	}

	return merr.ErrorOrNil()
}
