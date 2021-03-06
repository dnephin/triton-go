//
// Copyright (c) 2018, Joyent, Inc. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//

package list

import (
	"github.com/joyent/triton-go/cmd/agent/account"
	cfg "github.com/joyent/triton-go/cmd/config"
	"github.com/joyent/triton-go/cmd/internal/command"
	"github.com/olekukonko/tablewriter"
	"github.com/sean-/conswriter"
	"github.com/spf13/cobra"
)

var Cmd = &command.Command{
	Cobra: &cobra.Command{
		Args:         cobra.NoArgs,
		Use:          "list",
		Short:        "list Triton SSH Keys",
		Aliases:      []string{"ls"},
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cons := conswriter.GetTerminal()

			c, err := cfg.NewTritonConfig()
			if err != nil {
				return err
			}

			a, err := account.NewAccountClient(c)
			if err != nil {
				return err
			}

			keys, err := a.ListKeys()
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(cons)
			table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
			table.SetHeaderLine(false)
			table.SetAutoFormatHeaders(true)

			table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
			table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
			table.SetCenterSeparator("")
			table.SetColumnSeparator("")
			table.SetRowSeparator("")

			table.SetHeader([]string{"FINGERPRINT", "NAME"})

			for _, key := range keys {
				table.Append([]string{key.Fingerprint, key.Name})
			}

			table.Render()

			return nil
		},
	},
	Setup: func(parent *command.Command) error {
		return nil
	},
}
