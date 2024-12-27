package main

import (
	"GetBandori/model"
	"GetBandori/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {

	//fetchSongDataToLocal("songinfo.json")
	songdatas := getSongDataFromLocal("songinfo.json")

	fmt.Println(songdatas)
	return
}

// 从bestdori.com获取相关内容
func fetchSongDataToLocal(filepath string) model.Songdatas {
	results := make(model.Songdatas, model.MaxID) // 存储结果，按 ID 顺序

	errChan := make(chan error, model.MaxID)

	var wg sync.WaitGroup

	for i := 1; i <= model.MaxID; i++ {
		wg.Add(1)
		go fetchSongData(fmt.Sprintf("%s%d.json", model.UrlBase, i), i-1, results, &wg, errChan)
	}

	wg.Wait()
	close(errChan)

	// 收集错误并打印
	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}
	for _, err := range errors {
		fmt.Printf("Error: %v\n", err)
	}

	resultFiltered := utils.FilterEmptyEntries(results)
	err := utils.WriteJson(filepath, resultFiltered)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return results
}

// 获取相关内容的协程具体实现
func fetchSongData(url string, index int, results []model.SongData, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		errChan <- fmt.Errorf("songID %d: %v", index, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		//不存在该数据
		return
	}

	// io读取错误
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errChan <- fmt.Errorf("songID %d: %v", index, err)
		return
	}

	// 按索引存入结果
	err = json.Unmarshal(body, &results[index])

	if err != nil {
		fmt.Println("Error id :" + strconv.Itoa(index) + err.Error())
	}
}

// 从本地读取歌曲数据
func getSongDataFromLocal(filepath string) model.Songdatas {
	songdatas := make(model.Songdatas, model.MaxID)
	//这里需要传入指针
	err := utils.ReadJson(filepath, &songdatas)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return songdatas
}

func saveToDatabase() {

}
