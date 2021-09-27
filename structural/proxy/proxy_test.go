package main

import "testing"

func TestProxy(t *testing.T){
	t.Run("Creates image from proxy", func(t *testing.T) {
		proxy := newImageProxy("fileProxy.png")
		image := proxy.getImage()

		actualDrawing := proxy.draw(Point{})
		actualMouseHandling := proxy.handleMouse(Event{})
		
		expectedDrawing := "Drawing point on image Drawing on proxy"
		expectedMouseHandling := "Handling mouse on image Handling mouse on proxy"

		if image == nil {
			t.Errorf("Image is nil")
		}


		if actualDrawing != expectedDrawing {
			t.Errorf("Actual: %s Expected: %s", actualDrawing, expectedDrawing)
		}

		if actualMouseHandling != expectedMouseHandling {
			t.Errorf("Actual: %s Expected: %s", actualMouseHandling, expectedMouseHandling)
		}
	})
		t.Run("Creates image from proxy", func(t *testing.T) {
		image := Image{fileName: "file.png", extent: Point{fromProxy: false}}
		proxy := newImageProxy("fileProxy.png")
		proxy.image = &image
		actualDrawing := proxy.draw(Point{})
		actualMouseHandling := proxy.handleMouse(Event{})
		
		expectedDrawing := "Drawing point on image Drawing on proxy"
		expectedMouseHandling := "Handling mouse on image Handling mouse on proxy"

		if actualDrawing != expectedDrawing {
			t.Errorf("Actual: %s Expected: %s", actualDrawing, expectedDrawing)
		}

		if actualMouseHandling != expectedMouseHandling {
			t.Errorf("Actual: %s Expected: %s", actualMouseHandling, expectedMouseHandling)
		}
	})
}