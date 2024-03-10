package main

import "fmt"
import "strconv"

//---Convert Type 

func main() {

	
		var i int  = 256
		fmt.Printf("type:%d\n", i,i)

		//var f float64 =  i

		var f float64  = float64(i)
	
		fmt.Printf("type:%T,val : %T\n",f,f  )


		//--- convert string
		var x := "72"
		 s, err  strconv.Atoi(x)
		 
		 
		 if err != nil {
			fmt.Println("err: ", err)
			return
		 }
		 
		 fmt.Println(s)



	/*   

		var f float64  = float64(i)
		fmt.Printf("type :%d\n",f,f)

		var i int
		i = int(v)
		fmt.Println(i)

	


	/*
	v := "72"
	m, err := strconv.Atoi(v)

	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(m)

	*/



}

/*
//-- Map

func main(){
	//var ss []string = []string{}
	var m map[string] int   = map[string]int{}

	fmt.Println("the value :", m)
	m["Answer"] = 42
	fmt.Printf("The value :%#v\n",m)

	v:=m["Answer"]
	fmt.Println("The value :", v)

    m["Jeab"]  =  15
	fmt.Println("The value :", m)

	delete(m,"Answer")

	fmt.Println("The value:",m)

}


//----- Error


type MyError struct{}

func  (e MyError) Error() string {
	return "MyError"
}

var myErr = errors.New("my costom error message")


func Divide(a, b float64) (float64 , error) {
	if b ==0 {
		// err := MyError{}
		//  err := fmr.Errorf("can't divide '%f' by zero",a)

		return 0,myErr

		//return 0, fmt.Errorf("Divide by zero")
	}


	r := a/b
	return r, nil
}

func main(){
	r,err := Divide(1,0)

	  if err !=  nil {
		fmt.Println("handler err: ", err)
		return

	  }



	fmt.Println(r,err)

}





//-------- Interface --------


func show ( val any){

	i , ok := val.(int)

	if ok {
		i = i+2
	   fmt.Println(i)
	}

	switch v:= val.(type){
	case int :
		i = v+2
		fmt.Println("int:" , i)
	case string:
		a:= v+ "Hello"
		fmt.Println("string :",  a)
	default:
		fmt.Println("not handle type.:")

	}

}

type promotion interface{
	discount()int
}

func sale(val promotion){
	fmt.Println("sale :",val)
}

type course struct{}

func (c course)  discount() int {
	return 599
}


func (c course)  info()   {

	fmt.Println("info :",c)
}



func main(){
	v:= course{}
	v.info()
	sale(v)


}

//--- end if interface



// ----- POINTER
func changeprice ( p *int){
	*p = *p - 599
	fmt.Println("change",p, &p)

}


type course struct{
	name, instructor string
	price int
}



func (c course)discount() int  {
	c.price  = c.price - 599
	fmt.Println("discount : ",c.price )
	return c.price
}



func main(){

	c := course("Basic Go ", "Jeab", 9999)
	//c.price = 5999

	d := c.discount()
	fmt.Println("discount price : ", d)
	fmt.Println ("price : ", c.price)


	/*
	var price int = 9999
	var addr *int =  &price

	fmt.Println(price, addr)
	*addr = 9400 // write  in to address
	v := *addr //read

	fmt.Println(price, addr)
	fmt.Println(v)
	//fmt.Println("%T\n", addr)

	changeprice(&price)
	fmt.Println("after ", price, addr)

}



type course struct{
	name, instructor string
	price  int
}

func (c course)discount( d int) int{
	p:= c.price -599 -d
	fmt.Println("discount:",p )
	return p
}


func (c course)info() {
	fmt.Println("name :", c.name)
	fmt.Println("instructor :", c.instructor)
	fmt.Println("price :", c.price)
}


func main(){
	c:= course{"Basic Go","Anuchito", 9999}
	fmt.Printf("%#v\n",c)

	d:= c.discount(1000)
	fmt.Println("discount:",d)

	c.info()
}

//----Method

//---end study method


//---- strct work shop

type movie struct{
	title   string
	year    int
	raiting float64
	mtype  []string
	isSuperHero  bool

}

func main(){

	//skills := []string{"js", "go" , "python"}

var m1  =  movie{
	title : "Endgame" ,
	year : 2019,
	raiting :8.4,
	mtype :[]string{"Action" ,"Drama"} ,
   isSuperHero : true,
}

var m2 =  movie{
	title : "Infinity war" ,
	year : 2018, raiting :8.4,
	mtype :[]string{"Action" ,"Sci-FI"} ,
	isSuperHero : true,

}
var mvs []movie

mvs = append(mvs,m1,m2)

fmt.Printf( "%#v\n",m1)
fmt.Printf( "%#v\n",m2)
fmt.Printf( "%#v\n",mvs)

}




//--- Struct   same as java class

type course struct{
	name   string
	instructor  string
	price  float64
}

func main(){

	var c1 = course{	name :"Basic Go" ,instructor : "Anuchito",	price : 9999 ,}
    var c2 = course{"Basic Go","Anuchito",9999}
	var c3= course{ instructor :"Anuchito"}

	fmt.Printf("%#v\n",c1)
	fmt.Printf("%#v\n",c2)
	fmt.Printf("%#v\n",c3)



	c:= course{
		name :"Basic Go" ,
		instructor : "Anuchito",
		price : 9999 ,

	}


		n:= c.name
		c.instructor =  "jeab"




	fmt.Println("name :",c.name)
	fmt.Println("name :",n)
	fmt.Println("instructor :", c.instructor)
	fmt.Println("price :",c.price)

*/

/*
	name := "Basic Go"
	instructor := "Anuchito"
	price := 9999


	fmt.Println("name :",name)
	fmt.Println("instructor :",instructor)
	fmt.Println("price :",price)

}




//---slice learning
 /*
func show(sk[]string){
	fmt.Printf("show : %#v\n",sk)
}

func main() {
	//---workshop

	//mini workshop
	xs:=[]float64{4,5,7,8,3,8,0}
	ys:=[]float64{7,2,10,9,7}

	var vote[]float64
	vote =  append(xs,ys...)
	fmt.Printf("vote : %#v\n",vote)
	fmt.Printf("vote 5-8 : %v\n",vote[5:9])

 //---study

	skills := []string{"js", "go" , "python"}
	fmt.Printf("%T => %#v\n",skills,skills )

	skills[1] = "golang"
	fmt.Printf("%#v\n", skills)

	skills = append(skills,"ruby","js", "go" , "python")
	//fmt.Printf("length: %d---val:%#v\n",len(skills),skills)
    show(skills)

	s1 := skills[0:2]
	show(s1)

	s2 := skills[1:3]
	show(s2)

	//---create slice

	ss:= make([]int ,3)
	fmt.Printf("%#v\n",ss)


	s := skills[0]
	fmt.Printf("%#v\n",s)

*/

/*
	//-----------  Workshop  Loop //-----------

	genras := [3]string{"Action", "Adventure", "Fantasy"}

	fmt.Printf("before for loop :%#v\n", genras)


	for i := 0; i < len(genras); i++ {
		genras[i] = "Movie:" + genras[i]
	}

	fmt.Printf("after for loop :%#v\n", genras)
    //--- For array
	for i, genras := range genras {
		fmt.Printf("genras[%d] :%s\n", i, genras)
	}


	 //------------End Loop --------



	sum :=  1
	for i := 0; i <5 ; i++ {
		sum += sum
		fmt.Println("sum:",sum)
	}

	 fmt.Println("Sum done:",sum)





		for i,val  := range genras{
		   fmt.Println( "genras :", i, "value:" , val)
		}


*/
/* --------- Study ---
	sum :=1
	//for sum < 5  {
    for i:=1; sum < 5 ;i++ {
		sum =  sum +1
	fmt.Println("sum:",sum)
	}

	fmt.Println("sum:",sum)



	for {
	   fmt.Println("print forever")
	} *

	 skills:= [3]string{"js","go", "python"}

	 for i:= 0; i< len(skills); i++{
		fmt.Println(skills[i])
	 }

	 for i,val  := range skills{
		fmt.Println( "index :", i, "value:" , val)
	 }

	}
*/

//------ Array -----

/*
func show(sk[3]string){
	fmt.Printf("%#v\n",sk)

}*/
/*
func main() {

//------ Workshop ------
var genras[3] string  = [3]string {"Action","Adventure","Fantasy"}

fmt.Println("genral[0]:" ,genras[0])
fmt.Println("genral[1]:" ,genras[1])
fmt.Println("genral[2]:" ,genras[2])
genras[1] =  "Sci-Fi"

fmt.Println("genral[1]:" ,genras[1])

 /* ---- Study
 //var skills[3]string
 //fmt.Printf("%#v\n",skills)

  var skills[3]string  =  [3]string{"js","go", "python"}
  fmt.Printf("%#v\n",skills)

  //--call funcion
  var xs[3]string
  show(xs)

  //--- Index ----
  s:= skills[2]
  fmt.Printf("%#v\n",s)

  l :=  len(skills)
  fmt.Printf("%#v\n",l)

  } */

/*----FUnction

func add(x float64, y float64) (float64, string) {
	fmt.Println("Work in function", x, y)
	v := x + y
	return v, "hello"
}


func swap(x,y string)(string,string){
	return y,x

}

func main() {
	//a, b := add(42,13)
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}


//import "math"

// var msg string = "Hello Gopher!!!"

//func main() {


		//var msg string = "Hello Gopher!!!"
		msg := "Jeab"
		var age int  = 55
		//var price float64  = 99.99
	     age, price , check :=55, 99.99 , true

		//fmt.Println"Hello Gopher!!!"
		fmt.Println("msg :",msg)
		msg  = "Hello"+"Go"
		fmt.Println(msg)
		fmt.Println(age)
		fmt.Println("price :",price)
		fmt.Println("check :",check)

*/
//------ Workshop Println ---------
/*
	var Name1 = "Avengers"
	var Name2 = "Endgame"
	var Year = "2019"
	var Rating = 8.4
	var type1  =  "Sci-Fi"
	var chk = true
*/
/*

		Name1 := "Avengers"
		Name2 := "Endgame"
		Year := "2019"
		Rating := 8.4
		type1  :=  "Sci-Fi"
		chk := true

		fmt.Println("เรื่อง : ",Name1," : ",Name2)
	    fmt.Println( "ปี : ",Year)
		fmt.Println("เรตติ้ง : ",Rating)
		fmt.Println("ประเภท : ",type1)
		fmt.Println("ซุปเปอร์อีโร่ : " ,chk)
*/

/*------ Formating --------

	 var age int  = 55
	 var r rune = 'A'
	 var price float64  = 99.99
	 var chk = true
	 msg := "Hello Gopher!!!"
	 fmt.Println(r)
	 fmt.Printf("r: %c\n", r)
	 fmt.Printf("age: %d\n", age)
	 fmt.Printf("price: %2f\n", price)
	 fmt.Printf("chk: %2f\n", chk)

	//---- Print ALl format
	 fmt.Printf("r: %#v\n", r)
	 fmt.Printf("type: %T -- msg : %#v\n", msg, msg)

	//----- Work Shop2

	Title := "Avengers : Endgame"
	Year := "2019"
	Rating := 8.4
	type1 := "Sci-Fi"
	//chk := true
	var age2 int
	var msg2  string

	fmt.Printf("เรื่อง : %s\n", Title)
	fmt.Printf("ปี : %s\n", Year)
	fmt.Printf("เรตติ้ง :%1f\n,", Rating)
	fmt.Printf("ประเภท : %s\n", type1)
	fmt.Println(age2)
	fmt.Println(msg2)


	//-----If  else
	num :=34

	if num == 34 && (num >30 || num <39 ) {
		fmt.Println("Yes!! It's Thirty four")
	}

	if num%2==0{
		fmt.Println("เลขคู่")
	}else if num == 35{
		fmt.Println("คนมีคู่ไม่รู้หรอก")

	}else{
		fmt.Println("เลขโสด")
	}
//-----------
	v := math.Pow(18,2)
	lim := 255.0


	if v:= math.Pow(18,2); v < lim {
		fmt.Println("x power n is : ",v)
	}else{
		fmt.Println("x power n is %g over limit %g\n",v,lim)
	}

  fmt.Println(v)


  //------ WOrk shop If - Else

  ratings :=8.4

  if ratings < 5.0 {
	fmt.Println("Disappointed")
  }else if ratings >= 5.0 && ratings < 7.0{
	fmt.Println("Normal")
  }else if ratings >= 7.0 && ratings < 10.0 {
	fmt.Println("Good")
  }else{
	fmt.Println("Out ot case")
  }


//-----Switch Case
today := "Sunday"
switch today {
case "Sunday" :
	fmt.Println("today is Sunday")
case "Monday" :
	fmt.Println("today is Monday")
case "Tueday" :
	fmt.Println("today is Tueday")
default :

fmt.Println("สวัสดีวัน %v  Happy day\n",today)

 }


//--- work shop swich case

ratings := 1

switch    {
case  ratings < 5.0 :
	fmt.Println("Disappointed")
case   ratings >= 5.0 && ratings < 7.0:
	fmt.Println("Normal")
case  ratings >= 7.0 && ratings < 10.0 :
	fmt.Println("Good")
default :
	fmt.Println("Out ot case")

}

}*/
