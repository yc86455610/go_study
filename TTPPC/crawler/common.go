package crawler

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// SaveForMd .
func SaveForMd(title, text string, loc string) (err error) {
	name := loc + "/" + title + ".md"
	fmt.Println("file name: ", name)
	f, err := os.Create(name) //创建文件
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n, err := w.WriteString(text)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("写入 %d 个字节", n)
	w.Flush()
	return nil
}
