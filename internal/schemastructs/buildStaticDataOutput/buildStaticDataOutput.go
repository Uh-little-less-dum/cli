// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    buildStaticDataOutput, err := UnmarshalBuildStaticDataOutput(bytes)
//    bytes, err = buildStaticDataOutput.Marshal()

package schemas_build_static_data_output

import "encoding/json"

func UnmarshalBuildStaticDataOutput(data []byte) (BuildStaticDataOutput, error) {
	var r BuildStaticDataOutput
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BuildStaticDataOutput) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BuildStaticDataOutput struct {
	ComponentDocs     []ComponentDoc   `json:"componentDocs,omitempty"`
	FSRoot            string           `json:"fsRoot"`
	NavigationLinks   []NavigationLink `json:"navigationLinks,omitempty"`
	SettingPages      []SettingPage    `json:"settingPages,omitempty"`
	TailwindSources   []string         `json:"tailwindSources,omitempty"`
	TranspilePackages []string         `json:"transpilePackages,omitempty"`
}

type ComponentDoc struct {
	ComponentID                                                                             string    `json:"componentId"`
	ComponentName                                                                           string    `json:"componentName"`
	// The syntax for the component as it appears in JSX. Synonymous with the regexToInclude          
	// property in the developer config.                                                              
	EmbeddableSyntax                                                                        []string  `json:"embeddableSyntax,omitempty"`
	FilePaths                                                                               FilePaths `json:"filePaths"`
	PluginName                                                                              string    `json:"pluginName"`
	Tags                                                                                    []string  `json:"tags,omitempty"`
	Urls                                                                                    Urls      `json:"urls"`
}

type FilePaths struct {
	Full  *string `json:"full,omitempty"`
	Short *string `json:"short,omitempty"`
}

type Urls struct {
	Full  *string `json:"full,omitempty"`
	Short *string `json:"short,omitempty"`
}

type NavigationLink struct {
	Category   *Category `json:"category,omitempty"`
	Href       string    `json:"href"`
	Icon       *string   `json:"icon,omitempty"`
	Label      string    `json:"label"`
	PluginName *string   `json:"pluginName,omitempty"`
}

type SettingPage struct {
	Href       string  `json:"href"`
	PluginName string  `json:"pluginName"`
	Subtitle   *string `json:"subtitle,omitempty"`
	Title      *string `json:"title,omitempty"`
}

type Category string

const (
	AI              Category = "AI"
	Academic        Category = "academic"
	Calendar        Category = "calendar"
	Code            Category = "code"
	Database        Category = "database"
	General         Category = "general"
	Math            Category = "math"
	Ml              Category = "ML"
	Notebooks       Category = "notebooks"
	Organization    Category = "organization"
	ProjectPlanning Category = "project-planning"
	Research        Category = "research"
	School          Category = "school"
	Search          Category = "search"
	Snippets        Category = "snippets"
	TaskManagement  Category = "task-management"
	Work            Category = "work"
	Writing         Category = "writing"
)
