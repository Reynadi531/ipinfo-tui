package main

import (
	"ipinfo-tui/pkg"
	"os"
	"time"

	"github.com/pterm/pterm"
)

func init() {
	if len(os.Args) < 2 {
		pterm.Error.Println("Please provide an IP address")
		os.Exit(0)
	}
}

func GreenBgBlackFg(s string) string {
	return pterm.NewStyle(pterm.BgLightGreen, pterm.FgBlack, pterm.Bold).Sprintf(s)
}

func main() {
	ipArg := os.Args[1:][0]

	pterm.DefaultBigText.WithLetters(pterm.NewLettersFromStringWithStyle("IPInfo TUI", pterm.NewStyle(pterm.FgLightGreen))).Render()

	pterm.DefaultBasicText.Printfln("Requested on %s\n", time.Now().Local().Format("Mon 2 Jan 2006 15:04:05 MST"))

	pterm.DefaultSpinner.RemoveWhenDone = true
	pterm.DefaultSpinner.Start()

	res, err := pkg.CallAPI(ipArg)
	pterm.DefaultSpinner.Stop()

	if err != nil {
		pterm.Error.Println(err)
		os.Exit(0)
	}

	panels := pterm.Panels{
		{{Data: GreenBgBlackFg("IP Address : ")}, {Data: res.IP}},
		{{Data: GreenBgBlackFg("Hostname : ")}, {Data: res.Hostname}},
		{{Data: GreenBgBlackFg("City : ")}, {Data: res.City}},
		{{Data: GreenBgBlackFg("Country : ")}, {Data: res.Country}},
		{{Data: GreenBgBlackFg("Loc : ")}, {Data: res.Loc}},
		{{Data: GreenBgBlackFg("Org : ")}, {Data: res.Org}},
		{{Data: GreenBgBlackFg("Postal : ")}, {Data: res.Postal}},
		{{Data: GreenBgBlackFg("Timezone : ")}, {Data: res.Timezone}},
	}

	pterm.DefaultPanel.WithPanels(panels).WithSameColumnWidth(true).Render()

}
