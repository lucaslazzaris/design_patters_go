package main

import "fmt"
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
	drawRect(Point, Point) string
	drawPolygon([]Point)
	drawText(string, Point)

	getWindowImp() WindowImplementer
	getView() *View
}

type WindowImplementer interface {
	impTop()
	impBottom()
	impSetExtent(Point)
	impSetOrigin(Point)
	deviceRect(Coord, Coord, Coord, Coord) string
	deviceText(string, Coord, Coord)
	deviceBitmap(string, Coord, Coord)
}

// ApplicationWindow should implement Windower 
type ApplicationWindow struct {
	windowImp *ApplicationWindowImp
	view *View	
}

func(window *ApplicationWindow) open() {
	fmt.Println("Opening")
}

func(window *ApplicationWindow) close() {
	fmt.Println("Closing")
}

func(window *ApplicationWindow) iconify() {
	fmt.Println("Icon")
}

func(window *ApplicationWindow) deiconify() {
	fmt.Println("DeIcon")
}

func(window *ApplicationWindow) setOrigin(point Point) {
	fmt.Println("Origin", point)
}

func(window *ApplicationWindow) setExtent(point Point) {
	fmt.Println("Extent", point)
}

func(window *ApplicationWindow) raise() {
	fmt.Println("Raising")
}

func(window *ApplicationWindow) lower() {
	fmt.Println("Lower")
}

func(window *ApplicationWindow) drawLine(p1 Point, p2 Point) {
	fmt.Println("Drawing Line", p1, p2)
}

func(window *ApplicationWindow) drawRect(point1 Point, point2 Point) string {
	return window.windowImp.deviceRect(*point1.coord, *point1.coord, *point2.coord, *point2.coord)
}

func(window *ApplicationWindow) drawPolygon(points []Point) {
	fmt.Println("Drawing Polygon")
}

func(window *ApplicationWindow) drawText(text string, point Point) {
	fmt.Println("Drawing Text", text, point)
}

func (window *ApplicationWindow) drawContents() {
	window.view.drawOn(window)
}

func(window *ApplicationWindow) getWindowImp() WindowImplementer {
	return window.windowImp
}

func(window *ApplicationWindow) getView() *View {
	return window.view
}


// ApplicationWindow should implement WindowImplementer
type ApplicationWindowImp struct {

}

func (imp *ApplicationWindowImp) impTop() {
	fmt.Println("Top")
}

func (imp *ApplicationWindowImp) impBottom() {
	fmt.Println("Bottom")
}

func (imp *ApplicationWindowImp) impSetExtent(point Point) {
	fmt.Println("Extent on imp", point)
}

func (imp *ApplicationWindowImp) impSetOrigin(point Point) {
	fmt.Println("Origin on imp", point)
}

func (imp *ApplicationWindowImp) deviceRect(c1 Coord, c2 Coord, c3 Coord, c4 Coord) string {
	return fmt.Sprintf("Rect on imp %.1f %.1f %.1f %.1f", c1.x, c2.y, c3.x, c4.y)
}


func (imp *ApplicationWindowImp) deviceText(text string, c1 Coord, c2 Coord) {
	fmt.Println("Text on imp", text, c1, c2)
}

func (imp *ApplicationWindowImp) deviceBitmap(text string, c1 Coord, c2 Coord) {
	fmt.Println("Bitmap on imp", text, c1, c2)
}

// type IconWindow struct {
// 	windowImp *IconWindowImp
// 	view *View	
// 	bitmapName string
// }

// func(window IconWindow) drawContents() {
// 	if (window.windowImp != nil){
// 		window.windowImp.deviceBitmap(window.bitmapName, 0.0, 0.0)
// 	}
// }

// type IconWindowImp struct {

// }

type View struct {
}

func (view *View) drawOn(window Windower) {
	fmt.Println("Drawing on view", window)
}

