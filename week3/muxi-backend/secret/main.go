package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"tool/getDecryptedPaper"
	"tool/savePaper"
)

func main() {
	// 目标根URL
	url := "http://121.43.151.190:8000/"
	// 发送 GET 请求,返回的结果还需要进行处理才能得到你需要的结果
	response, err := http.Get(url + "paper")
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: Status code %d\n", response.StatusCode)
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	keys, err := getData(url + "secret")

	article := getDecryptedPaper.GetDecryptedPaper(string(body), string(keys))

	savePaper.SavePaper("C:/Users/17863/Desktop/muxi-backend/paper/Academician Sun's papers.txt", article)
}
func getData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data: %s", response.Status)
	}

	return io.ReadAll(response.Body)
}
