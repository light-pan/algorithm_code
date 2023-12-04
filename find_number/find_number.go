package findnumber

import (
	"fmt"
	"strconv"
	"strings"
)

func getNumber(n int64) int64 {

	str := strconv.FormatInt(n, 2)
	fmt.Println("修改前:", str)
	count1 := strings.Count(str, "1")

	if count1 == 1 {
		// 只有1个1时，只要在后面加0就行了
		str += "0"
	} else {
		count0 := strings.Count(str, "0")
		if count0 == 0 {
			// 没有0时,只要把最后一个1换成01,必然是最小的数
			str = str[:len(str)-1] + "01"
		} else {
			// 如果有0时，只要判断最后一个1前面有没有0，如果有，对换一下离最后一个1最近的0,那么必然最小，否则只能加0了
			lastIndex1 := strings.LastIndex(str, "1")
			substr := str[:lastIndex1]
			lastIndex2 := strings.LastIndex(substr, "0")
			if lastIndex1 != -1 {
				strBytes := []byte(str)
				strBytes[lastIndex1] = '0'
				strBytes[lastIndex2] = '1'
				str = string(strBytes)
			} else {
				str += "0"
			}
		}
	}
	fmt.Println("修改后:", str)
	v, _ := strconv.ParseInt(str, 2, 64)
	return v
}
