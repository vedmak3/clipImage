package main

import (
	"os"
	"image/jpeg"
	"fmt"
	"github.com/oliamb/cutter"
	"image"
	"strconv"
	"bufio"
)

func main()  {
	var val2,val1,i,j,x,y int
	var cImg image.Image
	var name string
	fmt.Println("Имя файла")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	name = myscanner.Text()
	fmt.Println("Строчек")
	myscanner = bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	val1,_ = strconv.Atoi(myscanner.Text())
	fmt.Println("Столбцов")
	myscanner = bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	val2,_ = strconv.Atoi(myscanner.Text())
	fmt.Println("Ждите")
	input, _ := os.Open(name)
	img,_:=jpeg.Decode(input)
	wi:=img.Bounds().Max.X/val2
	hi:=img.Bounds().Max.Y/val1
	for i=0;i<val2;i++{
		for j=0;j<val1;j++{
			timg, _ := os.Create("pic "+strconv.Itoa(j)+" "+strconv.Itoa(i)+".jpg")
			x=wi*i
			y=hi*j
			cImg, _ = cutter.Crop(img, cutter.Config{
				Width:  wi,
				Height: hi,
				Anchor: image.Point{x, y},
			})
			jpeg.Encode(timg,cImg, &jpeg.Options{100})
			timg.Close()
		}
	}
	fmt.Println("Все!!!")
}