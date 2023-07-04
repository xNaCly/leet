package main;import(."fmt"
."os"
."time");func main(){for _,a:=range Args[1:]{v,_:=Parse("2006-2-1",a);Println(v.Weekday())}}
