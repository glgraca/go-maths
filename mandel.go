package main

import (
  "image"
  "image/color"
  "image/png"
  "os"
  "math"
  "strconv"
)

func mandel(x, y, power float64) (c int) {
  var r=0.0;
  var i=0.0;
  var MAXITER=255;
  for r*r+i*i<4 && c<MAXITER {
    var t=math.Abs(power)*math.Atan2(i,r)
    var d=math.Pow(math.Sqrt(r*r+i*i),math.Abs(power))
    var nr=d*math.Cos(t)+x
    var ni=d*math.Sin(t)+y
    if(power<0.0) {
      r=nr/(nr*nr+ni*ni)
      i=-ni/(nr*nr+ni*ni)
    } else {
      r=nr
      i=ni
    }
    c++
  }
  return c;
}

func main() {
  power,_:=strconv.ParseFloat(os.Args[1], 64);
 
  mandelbrot:=image.NewRGBA(image.Rect(0,0,1023,1023));
  var minx float64=-2.0;
  var miny float64=-2.0;
  var maxx float64=2.0;
  var maxy float64=2.0;
  deltax:=(maxx-minx)/1024.0;
  deltay:=(maxy-miny)/1024.0;
  
  for x:=0; x<1024; x++ {
    for y:=0; y<1024; y++ {
      var c=mandel(minx+float64(x)*deltax, miny+float64(y)*deltay, power);
      red:=uint8((c*8)%255);
      green:=uint8((c*9)%255);
      blue:=uint8((255-c*4)%255);
      mandelbrot.Set(x, y, color.RGBA{red, green, blue, 0xff});
    }
  }
  
  outputFile,_:=os.Create("mandelbrot.png")
  png.Encode(outputFile, mandelbrot)
  outputFile.Close()
}