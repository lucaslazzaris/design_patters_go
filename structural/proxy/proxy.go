package main

import "fmt"

type Point struct {
	fromProxy bool
}

type Event struct {
}

type Graphic interface {
	draw(Point)
	handleMouse(Event)

	getExtent() Point
	
	load(string)
	save(string)
}

// Implements Graphic
type Image struct {
	fileName string
	extent Point
}

// Implements Graphic
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

func (proxy *ImageProxy) draw(point Point) {
	proxy.image.draw(point)
}

func (proxy *ImageProxy) handleMouse(event Event) {
	proxy.image.handleMouse(event)
}

func (proxy *ImageProxy) save(to string) {
	fmt.Println("Saving to", to)
}

func (proxy *ImageProxy) load(from string) {
	fmt.Println("Loading from", from)
}