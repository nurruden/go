package main

import (
	// "flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
	"bytes"
	"github.com/nfnt/resize"
	"sync"
)

const DEFAULT_MAX_WIDTH float64 = 64
const DEFAULT_MAX_HEIGHT float64 = 64
const DESTINATIONDIR string = "/Users/allanyang/Documents/despicdir/afterConvert_"

// 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

func covertToDesName(sourthname string) (desname string){
	v := strings.Split(sourthname,"/")
	desname = benchmarkStringFunction(DESTINATIONDIR,v[len(v)-1])
	// fmt.Printf("%v\n",v)
	return desname
}
func benchmarkStringFunction(desdir string, desname string) (fullname string){
	var buf bytes.Buffer
	buf.WriteString(desdir)
	buf.WriteString(desname)
	fullname = buf.String()
	return fullname
}
// 生成缩略图
func makeThumbnail(imagePath, savePath string) error {

	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// 需要保存的文件
	imgfile, err := os.Create(savePath)
	if err != nil {
		fmt.Printf("os create file, err:%v\n", err)
		return err
	}
	defer imgfile.Close()

	// 以PNG格式保存文件
	err = png.Encode(imgfile, m)
	if err != nil {
		return err
	}

	return nil
}

func WalkDir(dirPth, suffix string) (files []string,err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
   
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
	 //if err != nil { //忽略错误
	 // return err
	 //}
   
	 if fi.IsDir() { // 忽略目录
	  return nil
	 }
   
	 if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
	  files = append(files, filename)
	 }
   
	 return nil
	})
   
	return files, err
   }
   
   



func main() {
	files,_ := WalkDir("/Users/allanyang/Documents/", ".jpg")
	for _, imageFile := range files {
		fmt.Println(imageFile)
		
		
		wg.Add(1)
		go func() {
			v := covertToDesName(imageFile)
			err := makeThumbnail(imageFile, v)
			defer func() {
				if err != nil {
					fmt.Printf("catch panic exception err:%v\n", err)
				}
			}()

			wg.Done()
		}()

	}

	wg.Wait()

	//fmt.Printf("make thumbnail succ, path:%v\n", saveFile)

}
	// defer func(){
	// 	err := recover()
	// 	if err != nil{
	// 		fmt.Printf("catch panic exception.err:%v\n",err)
	// 	}
	// }()
	// var w sync.WaitGroup
	// // var m sync.Mutex
	// start := time.Now().UnixNano()
	
	
	// files,_ := WalkDir("/Users/allanyang/Documents/", ".jpg")
	
	
	// for _,v :=range(files){
	// 	w.Add(1)
	// 	go execute(v)
	// w.Wait()
			
	// }
	
	// end := time.Now().UnixNano()

	// fmt.Printf("cost=%d us\n",(end - start)/1000)
	
