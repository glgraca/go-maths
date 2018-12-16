package main

import (
  "fmt"
)

func choose(n int, k int, vals []interface{}, ch chan []interface{}, indices []int) {
  if k>0 {
    for m:=k; m<=n; m++ {
      choose(m-1, k-1, vals, ch, append([]int{m},indices...)); 
      if k==1 {
        line:=make([]interface{},len(indices)+1)
        line[0]=vals[m-1]
        for i:=0; i<len(indices); i++ {
          line[i+1]=vals[indices[i]-1]
        }
        ch<-line
      }
    }
  } 
}

func Choose(k int, vals []interface{}) <-chan []interface{} {
  ch:=make(chan []interface{});  
  go func() {
    choose(len(vals), k, vals, ch, nil);
    close(ch)
  }()
  return ch
}

func main() {
  names:=[]interface{}{"A", "B", "C", "D"};

  threes:= Choose(3, names);

  for line:=range threes {
    fmt.Println(line)
  }
}
