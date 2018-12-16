package main

import (
  "fmt"
)

func factoradics(n int, ch chan []int, indices []int) {
  if n==1 {
    indices=append(indices, 0)
    ch<-indices
  } else {
    for i:=0; i<n; i++ {
      factoradics(n-1, ch, append(indices, i))
    }
  }
}

func Factoradics(n int) <-chan []int {
  ch:=make(chan []int)
  go func() {
    factoradics(n, ch, nil)
    close(ch)
  }()
  return ch
}

func Permutations(vals []interface{}) <- chan []interface{} {
  ch:=make(chan []interface{})
  fact:=Factoradics(len(vals))

  go func() {
    for r:=range fact {
      tmp:=make([]interface{}, len(r))
      copy(tmp, vals)
      permutation:=make([]interface{}, len(r))
      for i:=0; i<len(r); i++ {
        v:=tmp[r[i]]
        tmp=append(tmp[:r[i]], tmp[r[i]+1:]...)
        permutation[i]=v
      }
      ch<-permutation
    }
    close(ch)
  }()

  return ch
}

func main() {
  ch:=Permutations([]interface{}{"A", "B", "C", "D"})

  for r:=range ch {
    fmt.Println(r)
  }  
}
