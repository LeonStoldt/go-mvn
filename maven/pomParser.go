package maven

import (
	"encoding/xml"
	"github.com/creekorful/mvnparser"
	"log"
)

func ParsePom(pom []byte) {
	var project mvnparser.MavenProject
	if err := xml.Unmarshal(pom, &project); err != nil {
		log.Fatalf("unable to unmarshal pom file. Reason: %s", err)
	}

	log.Print(project.Parent)

	log.Print(project.GroupId)    // -> com.example
	log.Print(project.ArtifactId) // -> my-app
	log.Print(project.Version)    // -> 1.0.0-SNAPSHOT

	for _, dep := range project.Dependencies {
		log.Print(dep.GroupId)
		log.Print(dep.ArtifactId)
		log.Print(dep.Version)
	}

	/* what we need:

	Project
		GroupId
		ArtifactId
		Version

	Parent:
		GroupId
		ArtifactId
		Version

	Properties (map[string]string)

	Dependencies
		GroupId
		ArtifactId
		Version
		(Exclusions)

	... tbc

	*/
}
