package blabla 
        
import (
    "testing"
)

func TestUpperFunc(t *testing.Testing){

}

func TestBla(t *testing.Testing){

}

func ExampleUpperFunc(){

}

func ExampleBla(){

}

func BenchmarkUpperFunc(b *testing.B){
	for i := 0; i < b.N; i++ {
		UpperFunc()	//Enter the values that your function needs between the parentheses
	}
}
func BenchmarkBla(b *testing.B){
	for i := 0; i < b.N; i++ {
		Bla()	//Enter the values that your function needs between the parentheses
	}
} function needs between the parentheses
	}
}