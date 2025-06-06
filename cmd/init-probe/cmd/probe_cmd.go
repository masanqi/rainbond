// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/goodrain/rainbond/cmd/init-probe/healthy"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// GetCmds get init probe cmds
func GetCmds() cli.Commands {
	var commands cli.Commands
	commands = []cli.Command{
		{
			Name:  "decoupling_probe",
			Usage: "decoupling probe depend service endpoint whether ready",
			Action: func(ctx *cli.Context) error {
				logrus.Info("start decoupling probe all depend service")
				controller, err := healthy.NewDecouplingDependServiceHealthController()
				if err != nil {
					logrus.Fatalf("new decoupling probe controller failure %s", err.Error())
					return err
				}
				controller.Check()
				return nil
			},
		},
	}
	return commands
}
