package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type AccountInfo struct {
	CustomerId      int
	CustomerName    string
	CustomerBalance int
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=gosample password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {

	var in string

	fmt.Println("select (1:create 2:truncate 3: selectAll 4: selectAny 5: update 8: importCSV)")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in = scanner.Text()
		fmt.Println("in: ", in)
		switch in {
		case "1":
			SampleCreate()
			goto L
		case "2":
			SampleTruncate()
			goto L
		case "3":
			SampleSelectAll()
			goto L
		case "4":
			SampleSelectAny()
			goto L
		case "5":
			SampleUpdateAny()
			goto L
		case "6":
			SampleDeleteAny()
			goto L
		case "8":
			SampleImport()
			goto L
		default:
			fmt.Println("コマンドが不正なのでもう一度入力を促す")
			continue
		}
	}
L:
}
