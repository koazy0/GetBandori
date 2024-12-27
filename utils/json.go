package utils

import (
	"GetBandori/model"
	"encoding/json"
	"os"
)

// FilterEmptyEntries 过滤数组空元素
func FilterEmptyEntries(dataArray model.Array) model.Array {
	return dataArray.FilterEmptyEntries()
}

func WriteJson1[T any](filename string, dataArray []T) error {
	// 创建一个文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 设置编码为JSON，并且缩进为4空格，使得输出的JSON文件易于阅读
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	// 将过滤后的SongData数组编码为JSON并写入文件
	err = encoder.Encode(dataArray)
	if err != nil {
		return err
	}
	return nil
}

// WriteJson 函数用于将数组接口写入JSON文件，跳过空元素
func WriteJson(filename string, dataArray model.Array) error {

	// 创建一个文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 设置编码为JSON，并且缩进为4空格，使得输出的JSON文件易于阅读
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	// 将过滤后的数组编码为JSON并写入文件
	err = encoder.Encode(dataArray)
	if err != nil {
		return err
	}
	return nil
}

func ReadJson(filename string, dataArray model.Array) error {
	// 读取文件内容到byte数组
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	// 反序列化JSON数据到dataArray
	err = json.Unmarshal(bytes, dataArray)
	return err
}
