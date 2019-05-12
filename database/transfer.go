package database

import "github.com/fasmide/schttp/packer"

// Transfer is both a Sink and a Source
type Transfer interface {
	// TransferTo must accept a PackerCloser and send files to it
	TransferTo(packer.PackerCloser) error

	// TakeTransfer must return a PackerCloser (typically just it self - if it supports receiving files)
	TakeTransfer() (packer.PackerCloser, error)
}
