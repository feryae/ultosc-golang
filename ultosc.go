
package main

import ("fmt"
		"os"
		"io"
		"math"
	)

type ArrData struct{
	date,time      		int
	opening        		float64
	high,low 			float64
	closing  	  		float64

}
type ArrBPTP struct{
	BP,TP	   	   		float64
}

    var file            *os.File
    var ArrInput[28]     ArrData
    var BuyRange[28]     ArrBPTP
    var fname            string
    var err              error
    var change  		 int
	var savechange       int

func main(){
// I.S : The main program that calls procedures and function. It will initialize the file first, fill the ArrInput and
//       calculating the UTLOSC Tech indicator.
// F.S : It will print the signal whether is buy or sell.

	var i int

	change = -1
	savechange = 0
	err = fileinit()
	initdata()
	determine(UTLOSC(),27)
	for err != io.EOF{
	  	i = 0
	  	for i < 28{
	  		Scandata(i)
			initBPTP(i)
	  		i = i + 1 
        }
        determine(UTLOSC(),i-1)
    }
    file.Close()    
}




func fileinit()error{
// I.S : Need a string fname for determining the file name. File must in the same directory.
// F.S : Will start the program if the file is found, and otherwise exit.
	 if len(os.Args) >1 {
		fname = os.Args[1]
	}else{
        fname = "EUM11709.DAT"
    }
    file, err = os.Open(fname)
    return err
}

func Scandata(i int){
// I.S : Need input variables (2 integers for date time, 1 real for opening price, and 3 ArrData for high,low,closing price.) 
// F.S : Array of ArrInput is filled with 28 data of real for formula purposes.

	_,err = fmt.Fscanf(file,"%d %d;%v;%v;%v;%v;0'\n'",&ArrInput[i].date,&ArrInput[i].time,&ArrInput[i].opening,&ArrInput[i].high,&ArrInput[i].low,&ArrInput[i].closing)
}

func initdata(){
// I.S : Need input variables (2 integers for date time, 1 real for opening price, and 3 ArrData for high,low,closing price.) 
// F.S : Array of ArrInput is filled with 28 data of real for formula purposes.
	
	var i int

	Scandata(0)
		// First input because initBPTP need.
	i = 1
	for i < 28{
		Scandata(i)
		i = i + 1
	}

}



func initBPTP(i int){	
// I.S : Need variables (3 ArrData of closing,high,and low price) and one parameter i for index of BuyRange array.
// F.S : Array of BuyRange is filled with 28 data of real for formula purposes.
	if i > 0{
	BuyRange[i-1].BP = ArrInput[i].closing - math.Min(ArrInput[i-1].closing,ArrInput[i].low)
	BuyRange[i-1].TP = math.Max(ArrInput[i].high,ArrInput[i-1].closing) - math.Min(ArrInput[i].low,ArrInput[i-1].closing)
}

}


func average(i int)(float64){
// I.S : Need variables (2 ArrData of Buying Pressure and True Range.) and one parameter i for index of BuyRange array.
// F.S : Will return real variable of the average by dividing both the sum.

	var sumBP,sumTP float64

	i = i - 1 
	for i > 0{
		sumBP = sumBP + BuyRange[i].BP
		sumTP = sumTP + BuyRange[i].TP
		i = i - 1
	}
	return sumBP/sumTP
}


func UTLOSC()float64{
// I.S : Need function average for return the value of the average data. 
// F.S : Will return real variable of the UTLOSC by weighted calculation.

	var UTLOSC    float64 
	UTLOSC = 100 * (((4 * average(7)) + (2 * average(14)) + average(28)) / 7)
	return UTLOSC
}

func determine(UTLOSC float64, i int){
// I.S : Need two parameters(UTLOSC for the formula, i for the output)
// F.S : Will print Sell output format if UTLOSC > 70 , or Buy output format if UTLOSC < 30
	if (UTLOSC > 70) {
		change = 0 
	}else if (UTLOSC < 30) {
		change = 1
	}

	if change != savechange {
		if change == 0{
			fmt.Println(ArrInput[i].date,ArrInput[i].time,ArrInput[i].opening,UTLOSC,"Sell")
			savechange = 0
		}else if change == 1{
			fmt.Println(ArrInput[i].date,ArrInput[i].time,ArrInput[i].opening,UTLOSC,"Buy")
			savechange = 1
		}
	}
}
	
	

