package main

func main() {
     // START1 OMIT
     for i:= 0; i < 1000; i++ {
       go func() {  for {}  } ()
     }

     fmt.Scanf("Wait For It:%d", &i)
     // END1 OMIT
}