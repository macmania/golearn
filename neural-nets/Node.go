package Node

import(
    "log"
    "fmt"
)

//A representation of each Node
type Node struct{
  weight double //the weight
  isActivated bool
  activatedFunction func(double, double) double
}
