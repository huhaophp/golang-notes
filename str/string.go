package str

import "fmt"

/**

一个字符串是一个不可改变的字节序列，字符串通常是用来包含人类可读的文本数据。
和数组不同的是，字符串的元素不可修改，是一个只读的字节数组。
每个字符串的长度虽然也是固定的，但是字符串的长度并不是字符串类型的一部分。
由于Go语言的源代码要求是UTF8编码，导致Go源代码中出现的字符串面值常量一般也是UTF8编码的。
源代码中的文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。
因为字节序列对应的是只读的字节序列，因此字符串可以包含任意的数据，包括byte值0。
我们也可以用字符串表示GBK等非UTF8编码的数据，
不过这种时候将字符串看作是一个只读的二进制数组更准确，因为for range等语法并不能支持非UTF8编码的字符串的遍历。

*/

func Demo() {
	str1 := "你好世界"
	str2 := "hello word"

	// 字节长度非字符长度 中文占3个字节 英文占1个字节
	fmt.Println(len(str1), len(str2)) // len(str1): 12, len(str2): 10

	// 将字符串转字节
	str1Bytes := []byte(str1) // [228 189 160 229 165 189 228 184 150 231 149 140]
	str2Bytes := []byte(str2) // [104 101 108 108 111 32 119 111 114 100]
	fmt.Println(str1Bytes, str2Bytes)

	// 修改字符串中的字节
	str2Bytes[1] = 'n' // 仅仅支持 uint8 类型字符 uint32 不支持这样写
	newStr2 := string(str2Bytes)
	fmt.Println(newStr2) // hnllo word
}
