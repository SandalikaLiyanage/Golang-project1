package main //important
import "fmt" //import a library so that we can be able to output stuff to the console

//project- quiz game
func main(){
	fmt.Println("Welcome to my quiz game!")
	//<<data types>>uint-unsigned integer-anypositive whole number,float64-we use 64 bits of our computer memory to store this float
	//name:="Tim"->this is like ,var name string="Tim"
	//fmt.Printf("Hello %v", name), here this is a print function , it will take an argument and assing it to the value(%v)
	
	// let's now take user inputs
	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)//&-go to the memory adress location of tis variable, scan requires th memory loction to assign the value to the variable
	fmt.Printf("Hello, %v , welcome to the game!\n",name)
	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >=10{
		fmt.Println("Yay you can play!")
	}else{
		fmt.Println("You cannot play!")
		return //break out of this function
	}

	score :=0
	num_questions:=2

	fmt.Printf("What is better ,the RTX 3000 or RTX 3090? ")
	var answer string
	var answer2 string//to get user inputs seperated by space. store them in 2 seperate variables
	fmt.Scan(&answer, &answer2)
	//check the answer is correct or not
	if(answer + " " + answer2=="RTX 3090"|| answer + " " + answer2=="rtx 3090"){
		fmt.Println("Correct!")
		score+=1
	}else{
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores int
	fmt.Scan(&cores)

	if cores==12{
		fmt.Println("Correct!")
		score++
	}else{
		fmt.Println("Incorrect!")
	}
	fmt.Printf("You scored %v out of %v.\n",score,num_questions)
	
	//percentage score that the user gets
	percent :=(float64(score)/float64(num_questions))*100
	fmt.Printf("You scored: %v%%.",percent )

}

