package main
import ("fmt")

func main() {
  var (
    sum1 = 100 + 50 // 150 (100 + 50)
    sum2 = sum1 + 250 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)

	/*Assignment Operators
	Oper | Exam     | Same as
	=	   | x = 5	  | x = 5	
	+=	 | x += 3	  | x = x + 3	
	-=	 | x -= 3	  | x = x - 3	
	*=	 | x *= 3	  | x = x * 3	
	/=	 | x /= 3	  | x = x / 3	
	%=	 | x %= 3	  | x = x % 3	
	&=	 | x &= 3	  | x = x & 3	
	|=	 | x |= 3	  | x = x | 3	
	^=	 | x ^= 3	  | x = x ^ 3	
	>>=	 | x >>= 3	| x = x >> 3	
	<<=	 | x <<= 3	| x = x << 3

	*/
	var x = 10
  x +=5
  fmt.Println(x) // ==> 15

}