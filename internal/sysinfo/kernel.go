package sysinfo

func collectKernel() (string, error) {
	return execute("uname", "-sr")
}
