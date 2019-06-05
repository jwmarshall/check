// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

var Url string
var Status int

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		err := checkHTTP(Url, Status)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func checkHTTP(u string, s int) error {
	// Force client to not follow redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return err
	}
	resp, err := client.Get(u)
	if err != nil {
		return err
	}
	if resp.StatusCode != s {
		return fmt.Errorf("Status codes do not match: %d != %d", resp.StatusCode, s)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&Url, "url", "u", "http://localhost", "URL to be checked")
	httpCmd.Flags().IntVarP(&Status, "status", "s", 200, "The expected HTTP status code")
}
