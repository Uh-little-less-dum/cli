// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    ulldPluginConfig, err := UnmarshalUlldPluginConfig(bytes)
//    bytes, err = ulldPluginConfig.Marshal()

package schemas_ulld_plugin_config

import "bytes"
import "errors"

import "encoding/json"

func UnmarshalUlldPluginConfig(data []byte) (UlldPluginConfig, error) {
	var r UlldPluginConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UlldPluginConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UlldPluginConfig struct {
	AdditionalImports                                                           *AdditionalImports      `json:"additionalImports,omitempty"`
	CommandPalette                                                              []CommandPaletteElement `json:"commandPalette,omitempty"`
	Components                                                                  []Component             `json:"components,omitempty"`
	// An optional export of an mdx file that describes the use of your plugin.                         
	Documentation                                                               *string                 `json:"documentation,omitempty"`
	Events                                                                      *Events                 `json:"events,omitempty"`
	// For display purposes. Does not need to match npm the way pluginName does.                        
	Label                                                                       string                  `json:"label"`
	NavigationLinks                                                             []NavigationLink        `json:"navigationLinks,omitempty"`
	Pages                                                                       []Page                  `json:"pages,omitempty"`
	Parsers                                                                     *Parsers                `json:"parsers,omitempty"`
	PluginID                                                                    *string                 `json:"pluginId,omitempty"`
	PluginName                                                                  string                  `json:"pluginName"`
	Settings                                                                    *Settings               `json:"settings,omitempty"`
	Slot                                                                        *Slot                   `json:"slot,omitempty"`
	Styles                                                                      *Styles                 `json:"styles,omitempty"`
	Tailwind                                                                    *Tailwind               `json:"tailwind,omitempty"`
	Trpc                                                                        *Trpc                   `json:"trpc,omitempty"`
}

type AdditionalImports struct {
	// Similar to additionalImports.root, but only applied to pages with rendered mdx content.           
	Mdx                                                                                         []string `json:"mdx,omitempty"`
	// Additional imports to add to the RootLayout. These will not be added to the dom, but may          
	// be used to import files that don't need to be executed like css or scss files. Must match         
	// export in package.json                                                                            
	Root                                                                                        []string `json:"root,omitempty"`
}

type CommandPaletteElement struct {
	Href                                                                                       *string `json:"href,omitempty"`
	// An optional export field in your package.json that resolves to a file with a default            
	// export of a synchronous function that accepts the current pathname and returns a boolean        
	// to indicate if this item should be shown in the command palette. This method is called          
	// just before the item is to be shown.                                                            
	IsAvailable                                                                                *string `json:"isAvailable,omitempty"`
	Label                                                                                      string  `json:"label"`
}

type Component struct {
	ComponentID                                                                                 *string          `json:"componentId,omitempty"`
	// Must start with a capital letter.                                                                         
	ComponentName                                                                               string           `json:"componentName"`
	// The package.json export that points to a .md or .mdx file that provides a quick reference                 
	// for the component. This should be simple enough to be opened in the command palette in                    
	// split view. If only docsExport or fullDocsExport is provided, this shorter docsExport is                  
	// heavily preferred.                                                                                        
	DocsExport                                                                                  *string          `json:"docsExport,omitempty"`
	// This can be an array to apply aliases to the same component. The component won't be                       
	// imported twice.                                                                                           
	Embeddable                                                                                  *EmbeddableUnion `json:"embeddable"`
	// The path that this component is exported as in your package.json.                                         
	Export                                                                                      string           `json:"export"`
	ExportedPropsName                                                                           *string          `json:"exportedPropsName,omitempty"`
	// Similar to docsExport, but will be shown when users redirect to an entire page. This                      
	// markdown can be more complex and include examples, but the components used must only                      
	// consist of core ULLD components and your plugin, since we don't know what plugins other                   
	// user's will be using.                                                                                     
	FullDocsExport                                                                              *string          `json:"fullDocsExport,omitempty"`
	// This only applys if the component is meant to override an existing slot. All 'slots' in                   
	// the appConfigSchema exported from @ulld/configschema/main/zod at the slots key will have                  
	// an accompanying 'subslot' schema exported from @ulld/configschema/subslots/<name of that                  
	// slot>. This schema will be a record of a specific set of keys unique to that slot, and a                  
	// path to a component. All of these paths are initially populated by internally developed                   
	// components, but can be overridden if your plugin defines the initial slot at the                          
	// developerConfigSchema.slot path, and then overrides one or more subslots unique to that                   
	// slot. This componentConfigSchema.slot path will be that subslot if it applies. Most                       
	// components that are embedded in a user's notes, and don't modify the app as a whole do                    
	// not occupy slots.                                                                                         
	Slot                                                                                        *string          `json:"slot,omitempty"`
	// Help user's find your component both before they install it, and while searching for                      
	// documentation.                                                                                            
	Tags                                                                                        []string         `json:"tags,omitempty"`
}

type EmbeddableElement struct {
	// An object key that matches your regex. 99% of the time, this will just be the                   
	// regexToInclude property without the leading '<'. It will default to that, but if your           
	// regexToInclude property is more specific and includes other special characters, you             
	// should provide this label yourself.                                                             
	Label                                                                                      *string `json:"label,omitempty"`
	// String passed to new RegExp(<regexToInclude>) to determine if a component should be             
	// imported. The raw content of a mdx file will be tested using this regex, and imported if        
	// a match is found. Due to the nature of jsx, 99% of the time, the componentName property         
	// can be used with a prefix of < to give '<MyComponentName'                                       
	RegexToInclude                                                                             string  `json:"regexToInclude"`
}

type Events struct {
	// The package.json export of a function that will run each time a user backs up their notes        
	// to a json file. This function should return a single json object.                                
	OnBackup                                                                                    *string `json:"onBackup,omitempty"`
	// The package.json export of a function that will run once during the build process. This          
	// can be used for copying files over to the public directory, but most other copying of            
	// files can be handled by setting up the pluginConfig appropriately.                               
	OnBuild                                                                                     *string `json:"onBuild,omitempty"`
	// The package.json export of a function that will be used to restore data from a previous          
	// backup. The argument will be an optional object of an identical type to what was stored          
	// using the onBackup event.                                                                        
	OnRestore                                                                                   *string `json:"onRestore,omitempty"`
	// The package.json export of a function that will run each time a user syncs their notes           
	// with their database.                                                                             
	OnSync                                                                                      *string `json:"onSync,omitempty"`
}

type NavigationLink struct {
	Category   *Category `json:"category,omitempty"`
	Href       string    `json:"href"`
	Icon       *string   `json:"icon,omitempty"`
	Label      string    `json:"label"`
	PluginName *string   `json:"pluginName,omitempty"`
}

type Page struct {
	// The export in your package.json that matches a single component that will act as this          
	// page. The component must be the default export, and will be passed all props that the          
	// page receives, like params and searchParams. This can include intercepted modal routes         
	// with the @modal/(.)... syntax, nested routes and parameter based routes with the               
	// [someParam] syntax.                                                                            
	Export                                                                                    string  `json:"export"`
	// Whether or not the file at the export field exports a type named PageProps to apply to         
	// the parent page properties. This type should include a params property and/or a                
	// searchParams property as they apply to the page itself. This is mostly for the sake of         
	// completeness, but can help other developers to work in a bug free environment if they          
	// choose to extend their own app.                                                                
	ExportsPageProps                                                                          *bool   `json:"exportsPageProps,omitempty"`
	// The optional slot key that matches a corresponding slot of type page.                          
	Slot                                                                                      *string `json:"slot,omitempty"`
	// The target URL to place this page at. This is synonomous with a file path from the root        
	// of the app directory, including intercepted routes. An intercepted modal route for             
	// example should appear as `@modal/(.)myPath/...` even though `@modal` doesn't appear in         
	// the URL. This must be unique, as if it overwrites an existing route it will not be             
	// applied.                                                                                       
	TargetURL                                                                                 *string `json:"targetUrl,omitempty"`
}

type Parsers struct {
	Mdx *Mdx `json:"mdx,omitempty"`
}

type Mdx struct {
	Conditions                                                                               *Conditions `json:"conditions,omitempty"`
	// The export in your package.json that exports the parser function. The parsing function            
	// must be the file's default export.                                                                
	Export                                                                                   string      `json:"export"`
}

type Conditions struct {
	FrontMatter *string `json:"frontMatter,omitempty"`
}

type Settings struct {
	// The export of a function that handles the saving of your settings specific to your         
	// plugin, most likely using the DJT model.                                                   
	OnSettingsSave                                                                        string  `json:"onSettingsSave"`
	// Export of a page that will be included in the user's settings page in it's own tab.        
	SettingPageExport                                                                     string  `json:"settingPageExport"`
	Subtitle                                                                              *string `json:"subtitle,omitempty"`
	// The label applied to the tab of the user's setting page for this plugins settings.         
	// Defaults to the developerConfigSchema.pluginName property.                                 
	Title                                                                                 *string `json:"title,omitempty"`
}

type Styles struct {
	// Optional export of a scss file that should be imported to pages with mdx content only.        
	Mdx                                                                                      *string `json:"mdx,omitempty"`
	// Optional export of a scss file that should be imported to all pages.                          
	Root                                                                                     *string `json:"root,omitempty"`
}

type Tailwind struct {
}

type Trpc struct {
	// The optional export path of a trpc router that will be appended to the main router at the       
	// devloperConfigSchema.trpc.routerName property.                                                  
	RouterExport                                                                                string `json:"routerExport"`
	RouterName                                                                                  string `json:"routerName"`
}

type Category string

const (
	AI               Category = "AI"
	Academic         Category = "academic"
	Calendar         Category = "calendar"
	CategoryMath     Category = "math"
	CategorySnippets Category = "snippets"
	Code             Category = "code"
	Database         Category = "database"
	General          Category = "general"
	Ml               Category = "ML"
	Notebooks        Category = "notebooks"
	Organization     Category = "organization"
	ProjectPlanning  Category = "project-planning"
	Research         Category = "research"
	School           Category = "school"
	Search           Category = "search"
	TaskManagement   Category = "task-management"
	Work             Category = "work"
	Writing          Category = "writing"
)

type Slot string

const (
	Bibliography   Slot = "bibliography"
	CommandPalette Slot = "commandPalette"
	Dashboard      Slot = "dashboard"
	Editor         Slot = "editor"
	Form           Slot = "form"
	Navigation     Slot = "navigation"
	PDF            Slot = "pdf"
	SlotMath       Slot = "math"
	SlotSnippets   Slot = "snippets"
	TaskManager    Slot = "taskManager"
	UI             Slot = "UI"
)

// This can be an array to apply aliases to the same component. The component won't be
// imported twice.
type EmbeddableUnion struct {
	EmbeddableElement      *EmbeddableElement
	EmbeddableElementArray []EmbeddableElement
}

func (x *EmbeddableUnion) UnmarshalJSON(data []byte) error {
	x.EmbeddableElementArray = nil
	x.EmbeddableElement = nil
	var c EmbeddableElement
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.EmbeddableElementArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.EmbeddableElement = &c
	}
	return nil
}

func (x *EmbeddableUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.EmbeddableElementArray != nil, x.EmbeddableElementArray, x.EmbeddableElement != nil, x.EmbeddableElement, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
			*pi = nil
	}
	if pf != nil {
			*pf = nil
	}
	if pb != nil {
			*pb = nil
	}
	if ps != nil {
			*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
			return false, err
	}

	switch v := tok.(type) {
	case json.Number:
			if pi != nil {
					i, err := v.Int64()
					if err == nil {
							*pi = &i
							return false, nil
					}
			}
			if pf != nil {
					f, err := v.Float64()
					if err == nil {
							*pf = &f
							return false, nil
					}
					return false, errors.New("Unparsable number")
			}
			return false, errors.New("Union does not contain number")
	case float64:
			return false, errors.New("Decoder should not return float64")
	case bool:
			if pb != nil {
					*pb = &v
					return false, nil
			}
			return false, errors.New("Union does not contain bool")
	case string:
			if haveEnum {
					return false, json.Unmarshal(data, pe)
			}
			if ps != nil {
					*ps = &v
					return false, nil
			}
			return false, errors.New("Union does not contain string")
	case nil:
			if nullable {
					return false, nil
			}
			return false, errors.New("Union does not contain null")
	case json.Delim:
			if v == '{' {
					if haveObject {
							return true, json.Unmarshal(data, pc)
					}
					if haveMap {
							return false, json.Unmarshal(data, pm)
					}
					return false, errors.New("Union does not contain object")
			}
			if v == '[' {
					if haveArray {
							return false, json.Unmarshal(data, pa)
					}
					return false, errors.New("Union does not contain array")
			}
			return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
			return json.Marshal(*pi)
	}
	if pf != nil {
			return json.Marshal(*pf)
	}
	if pb != nil {
			return json.Marshal(*pb)
	}
	if ps != nil {
			return json.Marshal(*ps)
	}
	if haveArray {
			return json.Marshal(pa)
	}
	if haveObject {
			return json.Marshal(pc)
	}
	if haveMap {
			return json.Marshal(pm)
	}
	if haveEnum {
			return json.Marshal(pe)
	}
	if nullable {
			return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
