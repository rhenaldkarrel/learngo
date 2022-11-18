package main

import "runtime"

func Version() string {
	return runtime.Version()
}

// push to github and than we can use the libary