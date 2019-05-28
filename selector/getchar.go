package selector

// #include <stdio.h>
import "C"

func getchar() {
	C.getchar()
}
