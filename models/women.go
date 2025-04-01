package models

type Women struct {
	Man
	IsMother bool
}

//func (w *Women) Breathe()    { w.breathing = true }
func (w *Women) Think()      { w.thinking = true }
//func (w *Women) Eat()        { w.eating = true }
func (w *Women) Sex() string { return "Mujer" }
