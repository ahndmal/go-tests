package main


func Compose(f, g func(x float) float)
                  func(x float) float {
     return func(x float) float {
        return f(g(x))
    }
}


func main() {
   print(Compose(sin, cos)(0.5))
}



