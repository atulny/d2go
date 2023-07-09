package main

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"os"
    "fmt"
 
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/textmeasure"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/canvas"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}  
 
func findfile() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
func showFilePicker(w fyne.Window) string{
 

    onChosen := func(f fyne.URIReadCloser, err error) {
        if err != nil {
            fmt.Println(err)
            return 
        }
        if f == nil {
            return 
        }
        fmt.Printf("chosen: %v", f.URI())
		save_dir =  f.URI().Path()// here value of save_dir shall be updated!
	 
		image := canvas.NewImageFromFile(save_dir)
		// image := canvas.NewImageFromURI(uri)
		// image := canvas.NewImageFromImage(src)
		// image := canvas.NewImageFromReader(reader, name)
		// image := canvas.NewImageFromFile(fileName)
		w.Content().Hide()
		image.FillMode = canvas.ImageFillOriginal
		w.SetContent( image)
        //_ = uri.Set(f.URI().String())
    }
    dialog.ShowFileOpen(onChosen, w)
	return save_dir
}
var save_dir string = "NoPathYet!"

func chooseDirectory(w fyne.Window) string {
    dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
        if err != nil {
            dialog.ShowError(err, w)
            return
        }
        if dir != nil {

            fmt.Println(dir.Path())
            save_dir = dir.Path() // here value of save_dir shall be updated!

        }
        fmt.Println(save_dir)

    }, w)
    return save_dir
}

func main1() {
    a := app.New()
    w := a.NewWindow("FileDialogTest")

    hello := widget.NewLabel("Hello!")
	 
    w.SetContent(container.NewVBox(
        hello,
        widget.NewButton("Go Get File!", func() {
            hello.SetText(showFilePicker(w)) 
        }),
    ))
    w.Resize(fyne.NewSize(500, 500))
    w.ShowAndRun()
}
// Remember to add if err != nil checks in production.
func build() {
	ruler, _ := textmeasure.NewRuler()
	defaultLayout := func(ctx context.Context, g *d2graph.Graph) error {
		return d2dagrelayout.Layout(ctx, g, nil)
	}
	dat, err := os.ReadFile("./test.d2")
	check(err)
	diagram, _, _ := d2lib.Compile(context.Background(), string(dat), &d2lib.CompileOptions{
		Layout: defaultLayout,
		Ruler:  ruler,
	})
	out, _ := d2svg.Render(diagram, &d2svg.RenderOpts{
		Pad:     d2svg.DEFAULT_PADDING,
		ThemeID: d2themescatalog.GrapeSoda.ID,
	})
	_ = ioutil.WriteFile(filepath.Join("out.svg"), out, 0600)
}
func main(){
	main1()
}