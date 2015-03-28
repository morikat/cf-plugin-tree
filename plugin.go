package main

import (
	"encoding/json"
  "errors"
  "fmt"
	"github.com/cloudfoundry/cli/cf/configuration/config_helpers"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/plugin"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type AppSearchResoures struct {
	Metadata AppSearchMetaData `json:"metadata"`
}

type AppSearchMetaData struct {
	Guid string `json:"guid"`
	Url  string `json:"url"`
}

type AppSearchResults struct {
	Resources []AppSearchResoures `json:"resources"`
}

type Tree struct{}

func (t *Tree) GetMetadata() plugin.PluginMetadata {
	versionParts := strings.Split(string(VERSION), ".")
	major, _ := strconv.Atoi(versionParts[0])
	minor, _ := strconv.Atoi(versionParts[1])
	patch, _ := strconv.Atoi(strings.TrimSpace(versionParts[2]))

	return plugin.PluginMetadata{
		Name: "tree",
		Version: plugin.VersionType{
			Major: major,
			Minor: minor,
			Build: patch,
		},
		Commands: []plugin.Command{
			{
				Name:     "tree",
				HelpText: "like a tree command for an cf application",
				UsageDetails: plugin.Usage{
					Usage: "cf tree <app-name>",
				},
			},
		},
	}
}

func fatalIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stdout, "error:", err)
		os.Exit(1)
	}
}

func main() {
	plugin.Start(new(Tree))
}

type TreeRepo struct {
	conn     plugin.CliConnection
	appGuid  string
	appName  string
	fileTree FileTree
}

type FileTree []string

func NewTreeRepo(conn plugin.CliConnection) *TreeRepo {
	return &TreeRepo{
		conn:     conn,
		fileTree: make([]string, 0),
	}
}

func (repo *TreeRepo) findAppGuid(spaceGuid string, appName string) {

	appQuery := fmt.Sprintf("/v2/spaces/%v/apps?q=name:%v&inline-relations-depth=1", spaceGuid, appName)
	cmd := []string{"curl", appQuery}
	output, _ := repo.conn.CliCommandWithoutTerminalOutput(cmd...)
	res := &AppSearchResults{}
	json.Unmarshal([]byte(strings.Join(output, "")), &res)
	repo.appGuid = res.Resources[0].Metadata.Guid
}

func (repo *TreeRepo) getAppFilePath(filePath string) []string {
	appQuery := fmt.Sprintf("/v2/apps/%v/instances/0/files/%v", repo.appGuid, filePath)
	cmd := []string{"curl", appQuery}
	output, err := repo.conn.CliCommandWithoutTerminalOutput(cmd...)
	if err != nil {
		fmt.Println("PLUGIN ERROR: Error from TreeCommand: ", err)
		return []string{}
	}
	return output

}
func (repo *TreeRepo) printTree(path string) {
	re := regexp.MustCompile("\\/$")
	replaceWord := ""
	path = re.ReplaceAllString(path, replaceWord)

	re = regexp.MustCompile("\\/(\\w|\\z|\\.|\\-|_)+")
	repl := []byte("|  ")

	repCount := strings.Count(path, "/") - 1

	i := 0
	treePath := re.ReplaceAllFunc([]byte(path), func(s []byte) []byte {
		if i < repCount {
			i += 1
			return repl
		}
		return s
	})
	fmt.Println(strings.Replace(string(treePath), "/", "|--", -1))
}

func (repo *TreeRepo) getFilesRecursive(filePath string) {
	for _, val := range repo.getAppFilePath(filePath) {
		for _, line := range strings.Split(val, "\n") {
			fileName := strings.Split(line, " ")[0]
			fileFullPath := filePath + fileName
			if fileName != "" {
				repo.fileTree = append(repo.fileTree, fileFullPath)
				if m, _ := regexp.MatchString(".*/", fileName); m {
					repo.getFilesRecursive(fileFullPath)
				} else {
					//fmt.Println(fileName + " is file")
					//repo.fileTree = append(repo.fileTree, fileFullPath)
					//fmt.Println(fileFullPath)
				}
			}
		}
	}
}

func (f FileTree) Len() int {
	return len(f)
}

func (f FileTree) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f FileTree) Less(i, j int) bool {
	return f[i] < f[j]
}

func (t *Tree) Run(cliConnection plugin.CliConnection, args []string) {
	appName := args[1]
	if appName == "" {
		fmt.Println("error")
		os.Exit(1)

	}
	treeRepo := NewTreeRepo(cliConnection)

  confRepo := core_config.NewRepositoryFromFilepath(config_helpers.DefaultFilePath(), fatalIf)
	spaceGuid := confRepo.SpaceFields().Guid
	treeRepo.findAppGuid(spaceGuid, appName)
	treeRepo.getFilesRecursive("/")

	sort.Sort(treeRepo.fileTree)
	for _, val := range treeRepo.fileTree {
		treeRepo.printTree(val)
	}
}

func checkArgs(cliConnection plugin.CliConnection, args []string) error {
  if len(args) < 2 {
    if args[0] == "tree" {
      cliConnection.CliCommand(args[0], "-h")
      return errors.New("Appname is needed")
    }
  }
  return nil
}
