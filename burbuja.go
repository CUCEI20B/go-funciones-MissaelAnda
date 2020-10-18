package main

func Burbuja(s []int64)  {
  var aux int64
	var i int = 0
	for i < len(s)-1 {
		if s[i+1] < s[i] {
			aux = s[i]
			s[i] = s[i+1]
			s[i+1] = aux

			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
}

func main()  {
	
}