package utils

import (
	"encoding/json"
	"log"
	"math"
	"reflect"
	"strconv"
	"strings"
)

var tenToAny = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
	16: "g",
	17: "h",
	18: "i",
	19: "j",
	20: "k",
	21: "l",
	22: "m",
	23: "n",
	24: "o",
	25: "p",
	26: "q",
	27: "r",
	28: "s",
	29: "t",
	30: "u",
	31: "v",
	32: "w",
	33: "x",
	34: "y",
	35: "z",
	36: "A",
	37: "B",
	38: "C",
	39: "D",
	40: "E",
	41: "F",
	42: "G",
	43: "H",
	44: "I",
	45: "J",
	46: "K",
	47: "L",
	48: "M",
	49: "N",
	50: "O",
	51: "P",
	52: "Q",
	53: "R",
	54: "S",
	55: "T",
	56: "U",
	57: "V",
	58: "W",
	59: "X",
	60: "Y",
	61: "Z"}

// 10进制转任意进制
func decimalToAny(num, n int) string {
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		if 76 > remainder && remainder > 9 {
			remainder_string = tenToAny[remainder]
		} else {
			remainder_string = strconv.Itoa(remainder)
		}
		new_num_str = remainder_string + new_num_str
		num = num / n
	}
	return new_num_str
}

// map根据value找key
func findKey(in string) int {
	result := -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func AnyToDecimal(num string, n int) int {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findKey(value))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(new_num)
}

// 反转字符串
func Reverse(s string) string {
	runs := []rune(s)
	for i, j := 0, len(runs)-1; i < j; i, j = i+1, j-1 {
		runs[i], runs[j] = runs[j], runs[i]
	}
	return string(runs)
}

// 十进制转62
func DecimalTo62(num int) string {
	return decimalToAny(num, 62)
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}

	t := reflect.TypeOf(v)
	switch t.Name() {
	case "int":
		return strconv.Itoa(v.(int))
	case "int64":
		return strconv.FormatInt(v.(int64), 10)
	case "float64":
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case "string":
		return v.(string)
	}
	return ""
}

func ToFloat64(v interface{}) float64 {
	t := reflect.TypeOf(v)
	switch t.Name() {
	case "int":
		return float64(v.(int))
	case "int64":
		return float64(v.(int64))
	case "float64":
		return v.(float64)
	case "string":
		return StrToFloat(v.(string))
	}

	return 0
}

func StrToInt(s string) int64 {
	if s == "" {
		return 0
	}

	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}

	return 0
}

func StrToFloat(s string) float64 {
	if s == "" {
		return 0
	}

	if i, err := strconv.ParseFloat(s, 64); err == nil {
		return i
	}

	return 0
}

func BuildJson(s interface{}) string {
	tmp, err := json.Marshal(s)
	if err != nil {
		log.Println("build json err:", err)
		return ""
	}

	return string(tmp)
}
