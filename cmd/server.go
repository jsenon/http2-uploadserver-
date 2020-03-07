/*
Copyright © 2020 Julien SENON

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/jsenon/http2-uploadserver/configs"
	mylog "github.com/jsenon/http2-uploadserver/internal/log"
	"github.com/jsenon/http2-uploadserver/internal/web"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command when called without any subcommands
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  `Stat a server that will test an upload`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.With().Str("Service", configs.Service).Logger()
		log.Logger = log.With().Str("Version", configs.Version).Logger()

		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		if loglevel {
			err := mylog.SetDebug()
			if err != nil {
				log.Error().Msgf("Could not set loglevel to debug: %v", err)
			}
			log.Debug().Msg("Log level set to Debug")
		}
		web.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	cobra.OnInitialize(initConfig)
}
