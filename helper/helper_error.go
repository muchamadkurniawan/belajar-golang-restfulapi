package helper

func PanicifError(err error) {
	if err != nil {
		panic(err)
	}
}
