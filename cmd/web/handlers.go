package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	s := Assumptions{
		System: []SysAssumption{
			SysAssumption{Label: "Number of Geneos Environments (Dev, Prod)", VarName: "numEnv"},
			SysAssumption{Label: "Gateways per Team", VarName: "gwPerTeam"},
			SysAssumption{Label: "Number of process per server", VarName: "numProcPerServer"},
			SysAssumption{Label: "Number of logfiles per server", VarName: "numLogPerServer"},
			SysAssumption{Label: "Percent of servers Level 3 applies to", VarName: "percentL3"},
			SysAssumption{Label: "Percent of servers Level 4 applies to", VarName: "percentL4"},
			SysAssumption{Label: "Percent of servers Level 5 applies to", VarName: "percentL5"},
		},
	}
	app.render(w, r, "home.page.tmpl", &templateData{Sys: s})
}
