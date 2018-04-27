package crawler

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// PathExists 判断文件夹是否存在.
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// SaveForMd .
func SaveForMd(title, text string, loc string) (err error) {
	exist, err := PathExists(loc)
	if err != nil {
		log.Fatal("PathExists error:", err)
		return err
	}
	if !exist {
		err := os.Mkdir(loc, os.ModePerm)
		if err != nil {
			log.Fatal("Mkdir error:", err, loc)
			return err
		}
	}
	name := loc + "/" + title + ".md"
	fmt.Println("file name: ", name)
	f, err := os.Create(name) //创建文件
	if err != nil {
		log.Fatal("Create error: ", err)
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
