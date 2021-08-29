package main

type Coord struct {
	x float32
	y float32
}

type Point struct {
	coord *Coord
}

type Windower interface {
	drawContents()
	open()
	close()
	iconify()
	deiconify()
	setOrigin(Point)
	setExtent(Point)
	raise()
	lower()
	drawLine(Point, Point)
	drawRect(Point, Point)
	drawPoligon([]Point)
	drawText(string, Point)

	getWindowImp() WindowImplementer
	getView() *View
}

type WindowImplementer interface {
	impTop()
	impBottom()
	ImpSetExtent(Point)
	impSetOrigin(Point)
	deviceRect(Coord, Coord, Coord, Coord)
	deviceText(string, Coord, Coord)
	deviceBitmap(string, Coord, Coord)
}

// ApplicationWindow should implement Windower 
type ApplicationWindow struct {
	windowImp *ApplicationWindowImp
	view *View	
}

func(window ApplicationWindow) drawRect(point1 Point, point2 Point) {
	window.windowImp.deviceRect(point1.coord, point2.coord)
}

func (window ApplicationWindow) drawContents() {
	window.view.drawOn(window)
}

func(window ApplicationWindow) getWindowImp() *ApplicationWindowImp {
	return window.windowImp
}

type IconWindow struct {
	windowImp *IconWindowImp
	view *View	
	bitmapName string
}

func(window IconWindow) drawContents() {
	if (window.windowImp != nil){
		window.windowImp.deviceBitmap(window.bitmapName, 0.0, 0.0)
	}
}

// ApplicationWindow should implement WindowImplementer
type ApplicationWindowImp struct {

}

type IconWindowImp struct {

}

type View struct {

}