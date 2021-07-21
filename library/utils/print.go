package utils

func PrettyPrint(v interface{}) {
	println(MarshalString(v))
}
