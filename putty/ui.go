package putty

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var Outputlog = widget.NewMultiLineEntry()

func Run() {

	var myapp = app.New()

	Outputlog.SetMinRowsVisible(10)
	Outputlog.Wrapping = fyne.TextWrapBreak

	loadPage := myapp.NewWindow("")
	loadPage.Resize(fyne.NewSize(300, 200))
	loadPage.SetContent(container.NewVBox(
		widget.NewLabel(""),
		container.NewGridWithColumns(1, widget.NewLabel("Please wait while initializing the environment")),
		container.NewGridWithColumns(1, widget.NewProgressBarInfinite()),
	))
	loadPage.CenterOnScreen()

	puttyCmdPage := myapp.NewWindow("puttyCmd")
	puttyCmdPage.CenterOnScreen()
	puttyCmdPage.SetContent(container.NewVBox(
		widget.NewLabel("Privacy mode change"),
		container.NewGridWithColumns(4,
			widget.NewButton("Incognito", Incognito),
			widget.NewButton("Tracking", Tracking),
			widget.NewButton("Location", Location),
			widget.NewButton("Personal", Personal)),
		widget.NewLabel("disable services"),
		container.NewGridWithColumns(3,
			widget.NewLabel("rhonk"),
			widget.NewButton("recover", RecoverRhonk),
			widget.NewButton("disable", DisableRhonk)),
		container.NewGridWithColumns(3,
			widget.NewLabel("carfinder"),
			widget.NewButton("recover", RecoverCarFinder),
			widget.NewButton("disable", DisableCarFinder)),
		container.NewGridWithColumns(3,
			widget.NewLabel("vehiclehealth"),
			widget.NewButton("recover", RecoverVehicleHealth),
			widget.NewButton("disable", DisableVehicleHealth)),
		container.NewGridWithColumns(3,
			widget.NewLabel("statusreport"),
			widget.NewButton("recover", RecoverStatusReport),
			widget.NewButton("disable", DisableStatusReport)),
		widget.NewLabel("download ocu/cns service list"),
		container.NewGridWithColumns(2,
			widget.NewButton("ocu", OCU),
			widget.NewButton("cns", CNS)),
		Outputlog,
	))
	puttyCmdPage.Resize(fyne.Size{Width: 800, Height: 600})

	loginPage := myapp.NewWindow("login")
	loginPage.CenterOnScreen()
	hostInput := widget.NewEntry()
	hostInput.SetText(host)
	portInput := widget.NewEntry()
	portInput.SetText(port)
	userInput := widget.NewEntry()
	userInput.SetText(username)
	passInput := widget.NewEntry()
	passInput.Password = true
	passInput.SetText(password)

	loginPage.SetContent(container.NewVBox(
		container.NewGridWithColumns(2, widget.NewLabel("host"), hostInput),
		container.NewGridWithColumns(2, widget.NewLabel("port"), portInput),
		container.NewGridWithColumns(2, widget.NewLabel("username"), userInput),
		container.NewGridWithColumns(2, widget.NewLabel("password"), passInput),
		widget.NewButton("login", func() {
			if hostInput.Text == "" {
				dial := dialog.NewError(errors.New("host input err"), loginPage)
				dial.Resize(fyne.NewSize(300, 100))
				dial.Show()
				return
			}

			if portInput.Text == "" {
				dial := dialog.NewError(errors.New("port input err"), loginPage)
				dial.Resize(fyne.NewSize(300, 100))
				dial.Show()
				return
			}

			if userInput.Text == "" {
				dial := dialog.NewError(errors.New("username input err"), loginPage)
				dial.Resize(fyne.NewSize(300, 100))
				dial.Show()
				return
			}
			if passInput.Text == "" {
				dial := dialog.NewError(errors.New("password input err"), loginPage)
				dial.Resize(fyne.NewSize(300, 100))
				dial.Show()
				return
			}
			username = userInput.Text
			password = passInput.Text
			host = hostInput.Text
			port = portInput.Text
			session, err := SSH(hostInput.Text, portInput.Text, userInput.Text, passInput.Text)
			if err != nil {
				dial := dialog.NewError(err, loginPage)
				dial.Resize(fyne.NewSize(300, 100))
				dial.Show()
				return
			}
			isOk := CheckSSH(session)
			if !isOk {
				dial := dialog.NewError(err, loginPage)
				dial.Resize(fyne.NewSize(300, 100))
				dial.Show()
				return
			}
			loginPage.Hide()
			loadPage.Show()
			// install curl
			installCurl := func() error {
				sftpClient, err := SftpConnect()
				if err != nil {
					return err
				}
				dstFile, err := sftpClient.Create(fmt.Sprintf("%s%s", remoteFile, fileName))
				if err != nil {
					return err
				}
				defer dstFile.Close()
				dstFile.Write(resourceCurl.StaticContent)
				//set chmod +x curl
				session, err := GetSession()
				defer session.Close()
				if err != nil {
					return err
				}
				_, err = session.Output("cd /mnt; chmod +x curl")
				if err != nil {
					return err
				}
				return nil
			}
			err = installCurl()
			if err != nil {
				dial := dialog.NewError(err, loadPage)
				dial.Resize(fyne.NewSize(300, 100))
				dial.Show()
				dial.SetOnClosed(func() {
					loadPage.Hide()
					loginPage.Show()
				})

			} else {
				loadPage.Hide()
				puttyCmdPage.Show()
			}

		}),
	))
	loginPage.Resize(fyne.Size{Width: 400, Height: 200})
	loginPage.Show()

	myapp.Run()
}
