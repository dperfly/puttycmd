package putty

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var Output = widget.NewMultiLineEntry()

func Run() {
	app := app.New()
	nw := app.NewWindow("puttyCmd")
	nw.SetContent(container.NewVBox(
		widget.NewLabel("Privacy mode change"),
		container.NewGridWithColumns(4,
			widget.NewButton("Incognito", Incognito),
			widget.NewButton("Tracking", Tracking),
			widget.NewButton("Location", Location),
			widget.NewButton("Personal", Personal)),
		widget.NewLabel("disable services"),
		container.NewGridWithColumns(3,
			widget.NewLabel("rhonk"),
			widget.NewButton("recover", EnableRhonk),
			widget.NewButton("disable", DisableRhonk)),
		container.NewGridWithColumns(3,
			widget.NewLabel("carfinder"),
			widget.NewButton("recover", EnableCarFinder),
			widget.NewButton("disable", DisableCarFinder)),
		container.NewGridWithColumns(3,
			widget.NewLabel("vehiclehealth"),
			widget.NewButton("recover", EnableVehicleHealth),
			widget.NewButton("disable", DisableVehicleHealth)),
		container.NewGridWithColumns(3,
			widget.NewLabel("statusreport"),
			widget.NewButton("recover", EnableStatusReport),
			widget.NewButton("disable", DisableStatusReport)),
		widget.NewLabel("download ocu/cns service list"),
		container.NewGridWithColumns(2,
			widget.NewButton("ocu", OCU),
			widget.NewButton("cns", CNS)),
		Output,
	))
	nw.Resize(fyne.Size{Width: 800, Height: 400})

	loginPage := app.NewWindow("login")
	hostInput := widget.NewEntry()
	hostInput.SetText(host)
	portInput := widget.NewEntry()
	portInput.SetText(port)
	userInput := widget.NewEntry()
	userInput.SetText(username)
	passInput := widget.NewEntry()
	passInput.Password = true
	passInput.SetText(password)

	// install curl
	installCurl := func() {
		sftpClient, err := SftpConnect()
		if err != nil {
			dial := dialog.NewError(err, loginPage)
			dial.Resize(fyne.NewSize(300, 100))
			dial.Show()
		}
		remoteFile := "/mnt/"
		fileName := "curl"
		dstFile, err := sftpClient.Create(remoteFile + fileName)
		if err != nil {
			dial := dialog.NewError(err, loginPage)
			dial.Resize(fyne.NewSize(300, 100))
			dial.Show()
		}
		defer dstFile.Close()
		dstFile.Write(resourceCurl.StaticContent)
		//set chmod +x curl
		session, err := GetSession()
		defer session.Close()
		if err != nil {
			dial := dialog.NewError(err, loginPage)
			dial.Resize(fyne.NewSize(300, 100))
			dial.Show()
		}
		_, err = session.Output("cd /mnt; chmod +x curl")
		if err != nil {
			dial := dialog.NewError(err, loginPage)
			dial.Resize(fyne.NewSize(300, 100))
			dial.Show()
		}

	}

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
			installCurl()

			loginPage.Hide()
			nw.Show()
		}),
	))
	loginPage.Resize(fyne.Size{Width: 400, Height: 300})
	loginPage.Show()

	app.Run()
}
