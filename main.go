package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writer(output string, body string)  {
	c := []byte(string(body))
	err := ioutil.WriteFile(output,c, 0644)
	check(err)
}

func parser(format string ,input string, output string) {
	var note Note
	var config Config

	inputFile, err := os.Open(input)
	if err != nil {
		check(err)
	}
	defer inputFile.Close()

	inputContent, _ := ioutil.ReadAll(inputFile)
	switch filepath.Ext(input) {
	case ".cfg":
		toml.Unmarshal(inputContent, &config)
	case ".xml":
		xml.Unmarshal(inputContent,&note)
	case ".yml":
		yaml.Unmarshal(inputContent,&note)
	case ".json":
		json.Unmarshal(inputContent,&note)
	default:
		log.Fatal("Input file format not recognized!")
	}

	switch format {
	case "yml":
		if !strings.Contains(output,".yml") {
			output = output + ".yml"
		}
		yml, _ := yaml.Marshal(&config)
		writer(output, string(yml))
	case "json":
		if !strings.Contains(output,".json") {
			output = output + ".json"
		}
		json, _ := json.Marshal(&note)
		writer(output, string(json))
	case "xml":
		if !strings.Contains(output,".xml") {
			output = output + ".xml"
		}
		xml, _ := xml.Marshal(&note)
		writer(output, string(xml))
	}

}

func main(){
	var input, output string

	app := cli.NewApp()
	app.Name = "midas"
	app.HelpName = "midas"
	app.Usage = "Parse files to many formats: xml, yml, json..."
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Copyright = "(c) 2019 BBVA"

	app.Commands = []cli.Command {
		{
			Name: "yml",
			Aliases: []string{"y"},
			Usage: "input -> YML",
			Description: "Parse input to `YAML` format",
			ArgsUsage: "-i `FILE` -o `YML`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "input, i",
					Usage: "Load xml from `PATH`",
					Destination: &input,
				},
				cli.StringFlag{
					Name: "output, o",
					Usage: "Write yml to `PATH`",
					Destination: &output,
				},
			},
			SkipFlagParsing:false,
			HideHelp:false,
			Hidden:false,
			HelpName:"yml",
			BashComplete: func(c *cli.Context) {
				fmt.Fprintf(c.App.Writer, "--better\n")
			},
			Action: func(c *cli.Context) error {
				log.Printf("Parsing %s to %s format...\n",input,c.Command.FullName())
				parser(c.Command.FullName(),input,output)
				return nil
			},
		},
		{
			Name: "json",
			Aliases: []string{"j"},
			Usage: "input -> JSON",
			Description: "Parse input to `JSON` format",
			ArgsUsage: "-i `FILE` -o `JSON`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "input, i",
					Usage: "Load xml from `PATH`",
					Destination: &input,
				},
				cli.StringFlag{
					Name: "output, o",
					Usage: "Write json to `PATH`",
					Destination: &output,
				},
			},
			SkipFlagParsing:false,
			HideHelp:false,
			Hidden:false,
			HelpName:"yml",
			BashComplete: func(c *cli.Context) {
				fmt.Fprintf(c.App.Writer, "--better\n")
			},
			Action: func(c *cli.Context) error {
				log.Printf("Parsing %s to %s format...\n",input,c.Command.FullName())
				parser(c.Command.FullName(),input,output)
				return nil
			},
		},
		{
			Name: "xml",
			Aliases: []string{"x"},
			Usage: "input -> XML",
			Description: "Parse input to `XML` format",
			ArgsUsage: "-i `FILE` -o `XML`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "input, i",
					Usage: "Load file from `PATH`",
					Destination: &input,
				},
				cli.StringFlag{
					Name: "output, o",
					Usage: "Write xml to `PATH`",
					Destination: &output,
				},
			},
			SkipFlagParsing:false,
			HideHelp:false,
			Hidden:false,
			HelpName:"yml",
			BashComplete: func(c *cli.Context) {
				fmt.Fprintf(c.App.Writer, "--better\n")
			},
			Action: func(c *cli.Context) error {
				log.Printf("Parsing %s to %s format...\n",input,c.Command.FullName())
				parser(c.Command.FullName(),input,output)
				return nil
			},
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "%q is not implemented yet\n", command)
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
