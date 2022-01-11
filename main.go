package main

import (
	"fmt"
	"github.com/LeonStoldt/go-mvn/fileexplorer"
	"github.com/LeonStoldt/go-mvn/maven"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	projectDir := "example/single-hierarchy-project"

	poms := fileexplorer.FindRecursivePomPaths(projectDir)

	for i := range poms {
		pom := poms[i]
		if strings.EqualFold(pom, projectDir+"/pom.xml") {
			log.Println("PARENT_POM:\n" + pom)
		} else {
			log.Println("CHILD_POM:\n" + pom)
		}
		printPomDetails(pom)
	}

	//mvnCommand := "dependency:list"
	//executeCommand("mvn", "-f", projectDir, mvnCommand)
}

func printPomDetails(pom string) {
	xmlFile, err := ioutil.ReadFile(pom)
	if err != nil {
		log.Fatal(err)
	}
	maven.ParsePom(xmlFile)
}

func executeCommand(command string, args ...string) string {
	fmt.Printf("DEBUG: execute %s with args: %s \n", command, args)
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("error executing command %s: %s \n", command, err)
		return ""
	}
	fmt.Printf("\n=> %s", output)
	return string(output[:])
}
