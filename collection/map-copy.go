package main

func main() {
	m := map[string]int{
		"a": 2,
		"b": 3,
	}

	// 遍历拷贝
	mCopy := map[string]int{}
	for k, v := range m {
		mCopy[k] = v
	}
}
