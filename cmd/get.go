// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	RootCmd.AddCommand(GetCmd)

	RootCmd.PersistentFlags().String("format", "table", "Output format for data (defaults to table)")
}

func outputList(cmd *cobra.Command, resources interface{}) error {
	dataFormat, err := cmd.Flags().GetString("format")
	if err != nil {
		return err
	}

	switch dataFormat {
	case "table":
		return outputTable(false, cmd.OutOrStdout(), resources)
	case "json":
		return outputJSON(cmd.OutOrStdout(), resources)
	}

	return fmt.Errorf("Unknown data format %v", dataFormat)
}

func outputJSON(out io.Writer, resources interface{}) error {
	j, err := json.Marshal(resources)
	if err != nil {
		return err
	}

	_, err = out.Write(j)
	return err
}

func outputTable(dataOnly bool, out io.Writer, resources interface{}) error {
	if !dataOnly {
		fmt.Fprint(out, "\n")
	}

	table := tablewriter.NewWriter(out)
	table.SetBorder(false)
	table.SetHeaderLine(false)
	table.SetColumnSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	fieldNames, err := extractFieldNames(resources)
	if err != nil {
		return err
	}

	if !dataOnly {
		table.SetHeader(fieldNames)
	}

	data, err := formatTableData(resources, fieldNames)
	if err != nil {
		return err
	}
	table.AppendBulk(data)

	table.Render()

	if !dataOnly {
		fmt.Fprintf(out, "\n%d records returned\n", len(data))
	}

	return nil
}

func extractFieldNames(resource interface{}) ([]string, error) {
	itemType := reflect.TypeOf(resource).Elem()
	fieldNames := make([]string, itemType.NumField())

	for i := 0; i < itemType.NumField(); i++ {
		fieldNames[i] = itemType.Field(i).Name
	}

	return fieldNames, nil
}

func formatTableData(resource interface{}, fieldNames []string) ([][]string, error) {
	values := reflect.ValueOf(resource)
	data := make([][]string, values.Len())

	for i := 0; i < values.Len(); i++ {
		data[i] = make([]string, len(fieldNames))
		value := values.Index(i)
		for j, fieldName := range fieldNames {
			rawFieldValue := value.FieldByName(fieldName).Interface()
			data[i][j] = fmt.Sprintf("%+v", rawFieldValue)
		}
	}

	return data, nil
}
