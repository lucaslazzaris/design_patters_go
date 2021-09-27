package main

import "fmt"

type Point struct {
	fromProxy bool
}

type Event struct {
}

type Graphic interface {
	draw(Point) string
	handleMouse(Event) string

	getExtent() Point
	
	load(string)
	save(string)
}

// Implements Graphic
type Image struct {
	fileName string
	extent Point
}


func (image *Image) getExtent() Point {
	return image.extent
}

func (image *Image) draw(point Point) string {
	return "Drawing point on image"
}

func (image *Image) handleMouse(event Event) string {
	return "Handling mouse on image"
}

func (image *Image) save(to string) {
	fmt.Println("Saving to", to)
}

func (image *Image) load(from string) {
	fmt.Println("Loading from", from)
}

type ImageProxy struct {
	image *Image
	extent Point
	fileName string
}

func newImageProxy(fileName string) *ImageProxy {
	return &ImageProxy{
		fileName: fileName,
		extent: Point{fromProxy: true},
		image: nil,
	}
}

func (proxy *ImageProxy) getImage() *Image {
	if (proxy.image == nil){
		proxy.image = &Image{fileName: proxy.fileName}
	}
	return proxy.image
}

func (proxy *ImageProxy) getExtent() Point {
	if (proxy.extent.fromProxy) {
		proxy.extent = proxy.image.extent
	}
	return proxy.extent
}

func (proxy *ImageProxy) draw(point Point) string {
	return proxy.image.draw(point) + " Drawing on proxy"
}

func (proxy *ImageProxy) handleMouse(event Event) string {
	return proxy.image.handleMouse(event) + " Handling mouse on proxy"
}

func (proxy *ImageProxy) save(to string) {
	fmt.Println("Saving to", to)
}

func (proxy *ImageProxy) load(from string) {
	fmt.Println("Loading from", from)
}