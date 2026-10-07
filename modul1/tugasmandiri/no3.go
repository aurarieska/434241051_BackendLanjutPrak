package main

func swapbypointer(a, b *int){ //*int artinya alamat dari variabel bertipe int
// jadi parameter function ini harus berupa alamatnya
	var temp1 = *a 
	*a=*b 
	*b=temp1
}

func swapbyvalue(a, b int){
	var temp2 = a
	a=b 
	b=temp2
}

func updateslice(s *[]string, newItem string){
	*s = append(*s, newItem)
}



