package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func SampleCreate() {
	fmt.Println("サンプル追加の処理を行います")
	for i := 0; i < 10; i++ {
		accountInfo := AccountInfo{CustomerId: i, CustomerName: "test", CustomerBalance: 10_000}
		err := accountInfo.Create()
		if err != nil {
			fmt.Printf("i=%v err=%v \n", strconv.Itoa(i), err)
		}
	}
}
func SampleTruncate() {
	fmt.Println("テーブル初期化の処理を行います")
	accountInfo := AccountInfo{}
	accountInfo.Truncate()
}
func SampleDeleteAny() {
	fmt.Println("削除したい顧客のidを入力してください>")

	//scan
	scanner.Scan()
	st := scanner.Text()

	//文字列を数字に変換
	i, _ := strconv.Atoi(st)

	//変換した数字を使って検索。エラーは受け取っていない
	accountInfo := AccountInfo{}
	accountInfo.DeleteAny(i)
	return
}
func SampleSelectAll() {
	fmt.Println("全件抽出します")
	accountInfo := AccountInfo{}
	ac, _ := accountInfo.SelectAll()
	for _, account := range ac {
		fmt.Println(account)
	}
}
func SampleSelectAny() (accountInfo AccountInfo) {
	fmt.Println("呼び出したい顧客のidを入力してください>")

	//scan
	scanner.Scan()
	st := scanner.Text()

	//文字列を数字に変換
	i, _ := strconv.Atoi(st)

	//変換した数字を使って検索。エラーは受け取っていない
	ac, _ := accountInfo.SelectAny(i)
	fmt.Println(ac[0])
	return ac[0]
}
func SampleUpdateAny() {
	SampleSelectAll()

	fmt.Println("idを入力してください。")
	ac := SampleSelectAny()
	fmt.Println()
	ac.CustomerName = "update"
	ac.CustomerBalance = ac.CustomerBalance * 2
	ac.Update()

}

func SampleImport() {
	file, err := os.Open("dummy_accounts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	accounts := []AccountInfo{}

	// 一行目を読み飛ばす
	_, err = csvReader.Read()
	if err == io.EOF {
		log.Fatalln(err)
	}

	for {
		line, err := csvReader.Read()
		if err != nil {
			break
		}

		customerId, err := strconv.Atoi(line[0])
		customerName := line[1]
		customerBalance, err := strconv.Atoi(line[2])

		if err != nil {
			continue
		}

		accounts = append(accounts, AccountInfo{
			CustomerId:      customerId,
			CustomerName:    customerName,
			CustomerBalance: customerBalance,
		})
	}
	//fmt.Println(accounts)

	for _, a := range accounts {
		a.Import()
	}
}
