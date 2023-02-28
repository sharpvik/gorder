package main

import (
	"io"
	"os"

	"github.com/sharpvik/gorder/order"
)

type Gorder struct {
	*Config
	fileName           string
	orderedPackageFile *order.File
}

type Config struct {
	dryRun bool
}

func NewGorder(config *Config) *Gorder {
	return &Gorder{
		Config: config,
	}
}

func (app *Gorder) Order(fileName string) (err error) {
	app.fileName = fileName
	app.orderedPackageFile, err = order.ReadFile(app.fileName)
	if err != nil {
		return err
	}
	if app.dryRun {
		return app.runDry()
	}
	return app.runForReal()
}

func (app *Gorder) runForReal() error {
	file, err := os.Create(app.fileName) //! Overrides the original file.
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, app.orderedPackageFile.Pretty())
	return err
}

func (app *Gorder) runDry() error {
	_, err := io.Copy(os.Stdout, app.orderedPackageFile.Pretty())
	return err
}
