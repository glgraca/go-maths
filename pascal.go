package main

import (
  "fmt"
  "os"
  "strconv"
)

func pascal(n uint64) <-chan uint64 {
  out:=make(chan uint64, n+1)
  if n==0 {
    go func() {
      out<-1
      close(out)
    }()
  } else {
    in:=pascal(n-1)
    go func() {
      a:=<-in
      out<-1
      for b:=range in {
        out<-a+b
        a=b
      }
      out<-1
      close(out)
    }()
  }
  return out;
}

func main() {
  n,_:=strconv.ParseUint(os.Args[1], 10, 64);
  line:=pascal(n);
  for n:=range line {
    fmt.Printf("%d ", n);
  }
}
