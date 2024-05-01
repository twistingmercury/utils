package utils

func SetExitFunc(f func(int)) {
	exitFunc = f
}
