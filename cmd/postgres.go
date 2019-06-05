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
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var Host string
var Port int
var User string
var Password string
var Database string

// postgresCmd represents the postgres command
var postgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		err := checkPostgres(Host, Port, User, Password, Database)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func checkPostgres(host string, port int, user string, pass string, database string) error {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pass,
		database,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(postgresCmd)
	postgresCmd.Flags().StringVarP(&Host, "host", "H", "localhost", "The hostname of the postgres server")
	postgresCmd.Flags().IntVarP(&Port, "port", "p", 5432, "The postgres server port")
	postgresCmd.Flags().StringVarP(&User, "user", "U", "postgres", "The postgres username")
	postgresCmd.Flags().StringVarP(&Password, "password", "P", "postgres", "The postgres password")
	postgresCmd.Flags().StringVarP(&Database, "database", "d", "postgres", "The postgres database name")
}
