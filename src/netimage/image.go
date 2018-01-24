package netimage

import (
	"gopkg.in/gographics/imagick.v3/imagick"
	"log"
)

type Image struct {
	File   string
	Twidth uint32
	Thight uint32
}

func (i *Image) Resize() []byte {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	var err error

	err = mw.ReadImage(i.File)
	if err != nil {
		log.Panic("Error opening image!", err)
	}

	mw.ResizeImage(uint(i.Twidth), uint(i.Thight), imagick.FILTER_LANCZOS)
	if err != nil {
		log.Panic("Error resizing image!", err)
	}

	mw.SetImageCompressionQuality(95)
	return mw.GetImageBlob()
}
