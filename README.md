# Spinner
spinner is a small ui loader for fyne
## how to use?
to create a spinner use this code:
```go
a := app.New()
w := a.NewWindow()
loader := spinner.New()
mainWindow := container.NewVBox(loader)
w.SetContent(mainWindow)
```
