package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Roser struct {
	Html string
}








func ExampleScrape() {
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//band := s.Find("a").Text()
		band2,_ := s.Find("a").Html()
		//band := s.Find("a").Text()
		title := s.Find("i").Text()
		//title2,_ := s.Find("i").Html()
		fmt.Printf("Review %d: %s - %s\n", i, band2, title)
	})
}

func main() {
	ExampleScrape()
}






func main7() {
	html := `<body>
				<div lang="zh">DIV1</div>
				
				<span><div>DIV5</div></span>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}
//为什么多了一个nil  </div>  右<>结尾  为nil
	dom.Find("div,span").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}



func main6() {
	html := `<body>
				<div lang="zh">DIV1</div>
				<span>
					<div>DIV5</div>
				</span>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	//唯一的元素筛选   可以选择div  可以选择span
	//only-of-type  同类元素只筛选一个
	//:only-child   只筛选唯一的div
	dom.Find("div:only-child").Each(func(i int, selection *goquery.Selection) {
	//dom.Find(":only-child").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}







func main5() {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<p>P2</p>
					<div>DIV5</div>
				</span>
				<div></div>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	//nth-last-child(n) 和:nth-last-of-type(n) 过滤器  倒叙开始计算
	//:nth-of-type(n) 过滤器  一样的
	//todo 在父级的第三个元素
	dom.Find(":nth-child(3)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}




func main4() {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<p>P2</p>
					<div>DIV5</div>
				</span>
				<div></div>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	//todo  :last-child 和 :last-of-type过滤器  最后一个

	//todo 父级 类型下面的第一个
	dom.Find("div:first-of-type").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}


func main3() {
	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div style="display:none;">DIV4</div>
					<div>DIV5</div>
				</span>
				<p>P2</p>
				<div></div>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	//筛选出的元素要是他们的父元素的第一个子元素，如果不是，则不会被筛选出来。
	//TODO  筛选父级下面的第一个子元素
	dom.Find("div:first-child").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}




//下一个标签
func main2() {

	//几乎所有的都是增删改查

	html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1_P标签</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV1</div>
				<span>
					<div>DIV4</div>
				</span>
				<p>P2</p>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}


	//下一个标签的[lang=zh]
	//相邻的标签
	//兄弟标签  同一个父级  lang=zh的兄弟选择器 ~
	//dom.Find("div[lang=zh]~p").Each(func(i int, selection *goquery.Selection) {
	//	fmt.Println(selection.Text())
	//})

	//内容过滤器   指定内容进行的查找
	//筛选出的元素要包含指定的文本
	//dom.Find("div:contains(DIV1)").Each(func(i int, selection *goquery.Selection) {
	//	fmt.Println(selection.Text())
	//})

	//标签里面有标签的  常用
	dom.Find("span:has(div)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})


}







//func main() {
//	html := `<body>
//
//				<div lang="ZH">DIV1</div>
//				<div lang="zh-cn">DIV2</div>
//				<div lang="en">DIV3</div>
//				<span>
//					<div>DIV4</div>
//				</span>
//
//			</body>
//			`
//
//	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
//	if err!=nil{
//		log.Fatalln(err)
//	}
//      //body>div  body下的div  body父元素 下的符合的一级元素
//      //bodydiv   空格  body下的所有的元素
//      //div[lang=zh-cn] div下的lang=zh-cn
//	dom.Find("div[lang=zh-cn]").Each(func(i int, selection *goquery.Selection) {
//		fmt.Println(selection.Text())
//	})
//}

func main1() {
	html :=
		`<body>
				<div>

<div id="div1">内层DIV唯一</div>


</div>
				
				<div>DIV2</div>
				<span>SPAN</span>
			</body>
			` //必须是网页的文档结构  string   可以把文档转成string格式
	htmlOneDiv :=
		`<body>
				<div id="div1">百度内层DIV唯一</div>
				<div>DIV2</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html)) //string==>输入流er()
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text()) //text文本打印
	})

	var R Roser
	R.Html = html
	R.find() //string类型的文档 直接输入进来  是可以的
	fmt.Println("=============")
	var Rone Roser
	Rone.Html = htmlOneDiv //结构体入值操作  然后操作接受者
	Rone.FindOnlyLabel()
	fmt.Println("=============")
	var Rone1 Roser
	Rone1.Html = html //结构体入值操作  然后操作接受者
	Rone1.FindOnlyLabel1()
fmt.Println("class选择器====")
	//class选择器
	html_class := `<body>

				<div id="div1">
			<div class="name1">Class选择器测试</div>
</div>
				<div class="name">DIV2_Class选择器</div>  
				<span>SPAN</span>

			</body>
			`
	var Rone_class Roser
	Rone_class.Html = html_class //结构体入值操作  然后操作接受者
	Rone_class.Find_class()

}

//Todo id选择器
func (c *Roser) Find_class() {
	var (
		err error
		doc *goquery.Document
	)

	doc, err = goquery.NewDocumentFromReader(strings.NewReader(c.Html))
	if err != nil {
		//fmt.Println("数据读取出现异常!")
		//return
		log.Fatalln(err) //os.Exit
	}
	//find查找样式选择器   each 重复
	//Each  相当于一个for循环遍历每一个对象  不能直接给它赋值

	//TODO  方法一
	doc.Find("div.name1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	//TODO  方法二
	doc.Find("div[class]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}





//Todo id选择器
func (c *Roser) FindOnlyLabel1() {
	var (
		err error
		doc *goquery.Document
	)

	doc, err = goquery.NewDocumentFromReader(strings.NewReader(c.Html))
	if err != nil {
		//fmt.Println("数据读取出现异常!")
		//return
		log.Fatalln(err) //os.Exit
	}
	//find查找样式选择器   each 重复
	//Each  相当于一个for循环遍历每一个对象  不能直接给它赋值
	doc.Find("div#div1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}





//Todo id选择器
func (c *Roser) FindOnlyLabel() {
	var (
		err error
		doc *goquery.Document
	)

	doc, err = goquery.NewDocumentFromReader(strings.NewReader(c.Html))
	if err != nil {
		//fmt.Println("数据读取出现异常!")
		//return
		log.Fatalln(err) //os.Exit
	}
	//find查找样式选择器   each 重复
	//Each  相当于一个for循环遍历每一个对象  不能直接给它赋值
	doc.Find("#div1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

/*
爬虫爬取结果直接保存或者展示 不要把结果再带出来
带出来的话需要定义切片  涉及到多个参数
string类型的文档 直接输入进来  是可以的
*/
//TODO <>元素选择器
func (c *Roser) find() {
	var (
		err error
		doc *goquery.Document
	)

	doc, err = goquery.NewDocumentFromReader(strings.NewReader(c.Html))
	if err != nil {
		//fmt.Println("数据读取出现异常!")
		//return
		log.Fatalln(err) //os.Exit
	}
	//find查找样式选择器   each 重复
	//Each  相当于一个for循环遍历每一个对象  不能直接给它赋值
	doc.Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
