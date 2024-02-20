package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	countCharFromFile()
}

func openAndCloseFile() {

	// 打开文件. file的叫法: file指针  file对象  file文件句柄
	file, err := os.Open("D:\\readme.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	// 输入下文件,看看文件是什么
	fmt.Printf("file=%v", file) // file=&{0xc00009ea00}

	// 关闭文件
	err1 := file.Close()
	if err1 != nil {
		fmt.Println("close file err1=", err1)
	}

}

func readFileContent() {
	file, err := os.Open("D:\\readme.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	// 当函数退出时,要及时关闭file
	defer file.Close() // 要及时关闭file句柄,否则会有内存泄漏

	// 创建一个 &Reader , 是带缓冲的.   defaultBufSize = 4096  // 默认的缓冲区4096
	reader := bufio.NewReader(file)
	// 循环的读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读取到一个换行就结束一次
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		// 输出内容
		fmt.Println(str)
	}
	fmt.Println("文件读取结束")

}

func readFileOnce() {
	file := "D:\\readme.txt"
	content, err := ioutil.ReadFile(file) // 文件的open和close都被封装到 readFile函数内部了
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	// 把读取到的内容显示到终端
	fmt.Printf("%v", string(content))
}

/*
*
const (

	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件

)

Create采用模式0666（任何人都可读写，不可执行）
创建一个名为name的文件，如果文件已存在会截断它（为空文件）。
如果成功，返回的文件对象可用于I/O；对应的文件描述符具有O_RDWR模式。
如果出错，错误底层类型是*PathError。
*/
func writeFile() {
	// 创建一个新文件,写入内容5句 "hello go"
	filePath := "D:\\readme.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("open file error = %v", err)
		return
	}
	// 及时关闭file句柄
	defer file.Close()
	// 1. 写入5句 "hello go"
	str := "hello go \n"
	// 写入时,使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为write是带缓存,因此在调用writeString方法时。
	// 内容是先写入缓存 需要调用Flush刷入磁盘
	writer.Flush()

}

func fileContentMove() {
	// 将D:/abc.txt 文件内容导入到 e:/kkk.txt
	file1Path := "D:/abc.txt"
	file2Path := "D:/kkk.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		// 说明文件有错误
		fmt.Printf("open file error = %v", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("open file error = %v", err)
	}
}

func pathExits(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		// 文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 自己编写一个函数,接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file error = %v", err)
	}
	// 通过srcFile, 获取Reader
	reader := bufio.NewReader(srcFile)

	// 打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file error = %v", err)
		return
	}
	// 通过dstFile,获取到write
	write := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(write, reader) // copy方法带缓存
}

type charCount struct {
	ChCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其它字符的个数
}

func countCharFromFile() {
	file, err := os.Open("D:\\readme.txt")
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()
	var count charCount
	reader := bufio.NewReader(file)

	// 开始循环的读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// 遍历str, 进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Println(count)
}
