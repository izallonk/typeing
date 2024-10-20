package engine


type Rect struct{
    X,Y,Width,Height float64
}

func NewRect(x,y,width,height float64) Rect{
  return Rect{
    X: x,
    Y: y,
    Width :width,
    Height: height,
    
  }
}
func (r *Rect) Right() float64{
  return r.X + r.Width
}
func (r *Rect) Left() float64{
  return r.X
}
func (r *Rect) Top() float64{
  return r.Y 
}
func (r *Rect) Bottom() float64{
  return r.Y  + r.Height
}

func (r Rect) Colision(other Rect) bool{
  return  r.Right() >= other.X &&
          r.X <= other.Right() &&
          r.Bottom() >= other.Y &&
          r.Y <= other.Bottom()
  
}

func (r *Rect) IsInside(x,y float64) bool{
    return x > r.X && x<r.Right() && y> r.Y && y<r.Bottom()
}

