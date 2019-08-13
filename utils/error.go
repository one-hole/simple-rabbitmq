package utils

func FailedOnError(err error) {
	if err != nil {
		panic(err)
	}
}
