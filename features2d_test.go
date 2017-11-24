package gocv

import (
	"testing"
)

func TestAgastFeatureDetector(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in AgastFeatureDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	ad := NewAgastFeatureDetector()
	defer ad.Close()

	kp := ad.Detect(img)
	if len(kp) < 2800 {
		t.Errorf("Invalid KeyPoint array in AgastFeatureDetector test: %d", len(kp))
	}
}

func TestFastFeatureDetector(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in FastFeatureDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	fd := NewFastFeatureDetector()
	defer fd.Close()

	kp := fd.Detect(img)
	if len(kp) < 2690 {
		t.Errorf("Invalid KeyPoint array in FastFeatureDetector test: %d", len(kp))
	}
}

func TestORB(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in AgastFeatureDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	od := NewORB()
	defer od.Close()

	kp := od.Detect(img)
	if len(kp) != 500 {
		t.Errorf("Invalid KeyPoint array in ORB test: %d", len(kp))
	}

	mask := NewMat()
	defer mask.Close()

	kp2, desc := od.DetectAndCompute(img, mask)
	if len(kp2) != 500 {
		t.Errorf("Invalid KeyPoint array in ORB DetectAndCompute: %d", len(kp2))
	}

	if desc.Empty() {
		t.Error("Invalid Mat desc in ORB DetectAndCompute")
	}
}

func TestSimpleBlobDetector(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in SimpleBlobDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	bd := NewSimpleBlobDetector()
	defer bd.Close()

	kp := bd.Detect(img)
	if len(kp) != 2 {
		t.Errorf("Invalid KeyPoint array in SimpleBlobDetector test: %d", len(kp))
	}
}
