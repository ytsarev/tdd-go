package main

func walk(x interface{}, fn func(input string)) {
	fn("call it up")
}
