package loops

import "fmt"

func LoopsTest(){
	// Loops
	
	// Kind of loop in rust or While true in other languajes
	for {
		fmt.Print("Close? (y/n): ")
		var c rune;
		
		fmt.Scanf("%c\n", &c)

		if c == 'n' || c == 'N'{
			continue
		}
		if c == 'y' || c == 'Y'{
			break
		}
		fmt.Println("Not recognized")
	}

	fmt.Println("First loop closed")

	var c rune;
	for c != 'y' && c != 'Y'{
		fmt.Print("Close? (y/n): ")
		fmt.Scanf("%c\n", &c)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}