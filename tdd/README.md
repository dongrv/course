

```
func Distance(x1, y1, x2, y2 float64) float64 {
    v := math.Abs(x1 - x2)
    o := math.Abs(y1 - y2)
    return math.Sqrt(v*v + o*o)
}
```