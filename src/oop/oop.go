package main

import (
	"fmt"
)

// This playground is meant to demonstrate and understand the oop (rather: type and value-oriented programming) of Go

// So, no classes, huh? What we do get are custom types of structs:

// Point2D is just that. A point.
type Point2D struct {
	x float64
	y float64
}

// ok now we have something like properties. but what about methods?
// turns out types and methods are separate things in Go. so you don't define them in the struct but
// rather "attach" them via a "receiver" type:

// Move moves a point by deltaX and deltaY
func (p *Point2D) Move(deltaX, deltaY float64) {
	// if you want to modify the "object" (i.e. the value) you need to pass a pointer
	// otherwise you would only get a copy (by-value)

	// this is actually short-hand for (*p).x and (*p).y. Go does magic dereferencing on struct pointers.
	p.x += deltaX
	p.y += deltaY
}

// So now by magic a Point2D has a "Move" method:
func movePoint(p *Point2D, dx, dy float64) {
	p.Move(dx, dy)
}

// pretty simple so far. but what about inheritance?
// surprise: there is none! Go rather encourages (enforces?) composition.
// inheritance is achieved (simulated?) via something called "embedding":

// Point3D pretends to subclass Point2D by "embedding" a Point2D type
type Point3D struct {
	Point2D
	z float64
}

func makePoint3D(x, y, z float64) *Point3D {
	// p := Point3D{x, y, z} // this, although intuitive, does not work

	// we initialize by sub-initializing the "embedded" type
	p := Point3D{Point2D{x, y}, z}

	return &p
}

// ok, weird enough. I see where this is going. but what about methods?
// I need a method to move a Point3D now...

// Move now moves a Point3D
func (p *Point3D) Move(deltaX, deltaY, deltaZ float64) {
	// Can we leverage existing code?
	// Turns out we can! This basically translates to super(deltaX, deltaY)
	p.Point2D.Move(deltaX, deltaY)

	//Now only the 3d part
	p.z += deltaZ
}

// The "Move" method is now shadowed (overloaded?) by its 3d-version:
// However, you can not really override a function in the same block, even if it has different arguments:
// func movePoint(p *Point3D, dx, dy, dz float64) {  <-- this does not work, redefinition of "movePoint"

func movePoint3D(p *Point3D, dx, dy, dz float64) {
	p.Move(dx, dy, dz)

	// we could move only the 2D-part of this point by calling
	// p.Point2D.Move(dx, dy)
}

// Could we somehow get around this "can't override functions" problem? It is really ugly, tho...

type Moveable interface {
	Move(deltas ...float64)
}

type Mover interface {
	Move(x, y float64)
}

func moveSomePoint(p Moveable, deltas ...float64) {
	p.Move(deltas...)
}

func main() {
	p2 := Point2D{12.5, 48.3}
	p3 := Point3D{Point2D{25.34, 42.34}, 4.8}

	fmt.Println("Made a 2d point at", p2)
	fmt.Println("Made a 3d point at", p3)

	movePoint(&p2, 12.5, -8.3)
	fmt.Println("Moved 2d point now at", p2)

	movePoint3D(&p3, 4.66, -2.34, 4.2)
	fmt.Println("Moved 3d point now at", p3)

	//moveSomePoint(&p2, -12.5, 8.3)  //does not work :(

}
