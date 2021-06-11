package util

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ParseUint 转换字符串为uint
func ParseUint(key string) (uint, error) {
	if key == "" {
		return 0, nil
	}
	id, err := strconv.ParseUint(key, 10, strconv.IntSize)
	return uint(id), err
}

// ValiateMobileNumber 校验手机号码是否有效
func ValiateMobileNumber(mobileNumber string) error {
	match, _ := regexp.MatchString("^1[23456789][0-9]{9}$", mobileNumber)
	if match {
		return nil
	}
	return errors.New("无效的手机号码")
}

// TrimSpace 去空格
func TrimSpace(a string) string {
	a = strings.TrimSpace(a)
	a = strings.Replace(a, " ", "", -1)
	return a
}

// BubbleSort 冒泡排序
func BubbleSort(arr []uint) (res []uint) {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		res = append(res, arr[i])
	}
	return
}

// Intersect 交集计算
func Intersect(arr1 []uint, arr2 []uint) []uint {
	if arr1 == nil || arr2 == nil {
		return []uint{}
	}
	BubbleSort(arr1)
	BubbleSort(arr2)
	var x, y = 0, 0
	var result []uint
	for {
		if x < len(arr1) && y < len(arr2) {
			if arr1[x] == arr2[y] {
				result = append(result, arr1[x])
				x++
				y++
			} else if arr1[x] > arr2[y] {
				y++
			} else {
				x++
			}
		} else {
			break
		}
	}
	return result
}

// Difference 差集计算
func Difference(arr1 []uint, arr2 []uint) []uint {
	if arr1 == nil && arr2 == nil {
		return []uint{}
	}
	if len(arr2) > len(arr1) {
		arr1, arr2 = arr2, arr1
	}
	var result []uint
	var isEqual bool
	for x := 0; x < len(arr1); x++ {
		for y := 0; y < len(arr2); y++ {
			isEqual = false
			if arr1[x] == arr2[y] {
				isEqual = true
				break
			}
		}
		if isEqual {
			continue
		}
		result = append(result, arr1[x])
	}
	return result
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	if unPadding < 1 || unPadding > 32 {
		unPadding = 0
	}
	return plantText[:(length - unPadding)]
}
