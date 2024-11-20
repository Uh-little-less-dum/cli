// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appConfig, err := UnmarshalAppConfig(bytes)
//    bytes, err = appConfig.Marshal()

package schemas_app_config

import "bytes"
import "errors"

import "encoding/json"

func UnmarshalAppConfig(data []byte) (AppConfig, error) {
	var r AppConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppConfig struct {
	// Always read directly from the file system, except in cases where use of a database is                       
	// required. This will negatively impact performance and portability, but might be necessary                   
	// for specific use cases.                                                                                     
	AlwaysPreferFS                                                                              *bool              `json:"alwaysPreferFs,omitempty"`
	AutoSubject                                                                                 []AutoSubject      `json:"autoSubject,omitempty"`
	AutoTag                                                                                     []AutoTag          `json:"autoTag,omitempty"`
	AutoTopic                                                                                   []AutoTopic        `json:"autoTopic,omitempty"`
	// The root relative path to the .bib file to be integrated within your app.                                   
	BibPath                                                                                     *string            `json:"bibPath,omitempty"`
	Build                                                                                       *Build             `json:"build,omitempty"`
	Code                                                                                        *Code              `json:"code,omitempty"`
	Credentials                                                                                 *Credentials       `json:"credentials,omitempty"`
	// The root relative path to the .csl file to be used for styling citations. An incredibly                     
	// robust repo of different csl styles can be found at                                                         
	// https://github.com/citation-style-language/styles                                                           
	CslPath                                                                                     *string            `json:"cslPath,omitempty"`
	Database                                                                                    *AppConfigDatabase `json:"database,omitempty"`
	DateHandling                                                                                *DateHandling      `json:"dateHandling,omitempty"`
	// The path to the root of the folder which holds your notes.                                                  
	FSRoot                                                                                      string             `json:"fsRoot"`
	// An ordered list of different parsable file types. This is only required when a file name                    
	// exists with multiple parsable file extensions in the same directory. In most cases the                      
	// default is adequate.                                                                                        
	FileTypePriority                                                                            []FileTypePriority `json:"fileTypePriority,omitempty"`
	// fsRoot relative path to the directory for automatically generated content. In almost all                    
	// use cases this can be left to it's default value.                                                           
	GeneratedDir                                                                                *string            `json:"generatedDir,omitempty"`
	// File paths within the root directory which should be completely ignored by ULLD.                            
	IgnoreFilepaths                                                                             []interface{}      `json:"ignoreFilepaths,omitempty"`
	// An array of either glob strings or RegExp's with which to test file paths. Those                            
	// evaluating to true will always be rendered from the database, regardless of other global                    
	// settings.                                                                                                   
	IgnorePreferFSExtensions                                                                    []interface{}      `json:"ignorePreferFsExtensions,omitempty"`
	Jupyter                                                                                     *Jupyter           `json:"jupyter,omitempty"`
	// A map of key value pairs of commonly referenced urls.                                                       
	LinkAliases                                                                                 map[string]string  `json:"linkAliases,omitempty"`
	Math                                                                                        *Math              `json:"math,omitempty"`
	Meta                                                                                        *Meta              `json:"meta,omitempty"`
	Navigation                                                                                  *Navigation        `json:"navigation,omitempty"`
	// This is the main location to describe the structure of your notes. Break up your note                       
	// directory into as many categories as you like, but this app is designed to allow for                        
	// increasingly refined searching and filtering. Categories of 'math', 'physics' and                           
	// 'chemistry' would likely fit most users better than 'calc1', 'calc2', 'linearAlgebra',                      
	// etc. For use cases such as those, please look at the 'autoTag', 'autoSubject', and                          
	// 'autoTopic' feature.                                                                                        
	NoteTypes                                                                                   []NoteType         `json:"noteTypes,omitempty"`
	Performance                                                                                 *Performance       `json:"performance,omitempty"`
	Plotting                                                                                    *Plotting          `json:"plotting,omitempty"`
	Plugins                                                                                     *PluginsUnion      `json:"plugins"`
	Slots                                                                                       *Slots             `json:"slots,omitempty"`
	// fsRoot relative path to the temporary directory. In almost all use cases this can be left                   
	// to it's default value.                                                                                      
	TempDir                                                                                     *string            `json:"tempDir,omitempty"`
	Terminal                                                                                    *Terminal          `json:"terminal,omitempty"`
	UI                                                                                          *AppConfigUI       `json:"UI,omitempty"`
}

type AutoSubject struct {
	// The glob pattern to test paths against.       
	Path                                      string `json:"path"`
	Subject                                   string `json:"subject"`
}

type AutoTag struct {
	// The glob pattern to test paths against.       
	Path                                      string `json:"path"`
	Tag                                       string `json:"tag"`
}

type AutoTopic struct {
	// The glob pattern to test paths against.       
	Path                                      string `json:"path"`
	Topic                                     string `json:"topic"`
}

type Build struct {
	AdditionalUserContent *AdditionalUserContent `json:"additionalUserContent,omitempty"`
	Database              *BuildDatabase         `json:"database,omitempty"`
}

type AdditionalUserContent struct {
	CSS     *string `json:"css,omitempty"`
	Favicon *string `json:"favicon,omitempty"`
	Logo    *string `json:"logo,omitempty"`
}

type BuildDatabase struct {
	Postgres   *PostgresClass `json:"postgres,omitempty"`
	Prioritize *Prioritize    `json:"prioritize,omitempty"`
	Type       *Type          `json:"type,omitempty"`
}

type PostgresClass struct {
	ConnectionURI *string `json:"connectionURI,omitempty"`
	DBName        *string `json:"dbName,omitempty"`
	Port          *int64  `json:"port,omitempty"`
}

type Code struct {
	Editor             *Editor             `json:"editor,omitempty"`
	SyntaxHighlighting *SyntaxHighlighting `json:"syntaxHighlighting,omitempty"`
	Theme              *ThemeClass         `json:"theme,omitempty"`
}

type Editor struct {
	DefaultLangauge *string `json:"defaultLangauge,omitempty"`
	UseVimMode      *bool   `json:"useVimMode,omitempty"`
}

type SyntaxHighlighting struct {
	DefaultLanguage *DefaultLanguageUnion `json:"defaultLanguage"`
	Transformers    *Transformers         `json:"transformers,omitempty"`
}

type DefaultLanguageClass struct {
	Block  Block `json:"block"`
	Inline Block `json:"inline"`
}

type Transformers struct {
	// shiki#transformerNotationDiff            
	LineDiff                              *bool `json:"lineDiff,omitempty"`
	// shiki#transformerNotationErrorLevel      
	LineErrorLevel                        *bool `json:"lineErrorLevel,omitempty"`
	// shiki#transformerNotationFocus           
	LineFocus                             *bool `json:"lineFocus,omitempty"`
	// shiki#transformerMetaHighlight           
	LineHighlight                         *bool `json:"lineHighlight,omitempty"`
	// shiki#transformerMetaWordHighlight       
	RegexHighlight                        *bool `json:"regexHighlight,omitempty"`
}

type ThemeClass struct {
	// Syntax highlighting theme to be used when in dark mode. Can be overriden with the           
	// ?shikiHlDark=dracula search param.                                                          
	Dark                                                                                 *DarkEnum `json:"dark,omitempty"`
	// Syntax highlighting theme to be used when in light mode. Can be overriden with the          
	// ?shikiHlLight=nord search param.                                                            
	Light                                                                                *DarkEnum `json:"light,omitempty"`
}

type Credentials struct {
	// Path to the service account credentials file to enable calendar integration.        
	GoogleServiceAccountJSONPath                                                   *string `json:"googleServiceAccountJsonPath,omitempty"`
}

type AppConfigDatabase struct {
	// Whether or not to remove notes from the database if they are no longer found in the local      
	// file system. This is best set to true if storing all notes in one location, but can be         
	// set to false to avoid losing notes that may have been lost elsewhere.                          
	RemoveIfNotPresentInFS                                                                      *bool `json:"removeIfNotPresentInFs,omitempty"`
	// Whether or not to store the formatted content along with the raw content This will             
	// improve performance and load times when rendering content  not in 'preferFs' mode, but         
	// will increase the storage size of each note. Recommended: true if running and storing          
	// notes locally where network speeds and monthly fees aren't an issue.                           
	StoreFormatted                                                                              *bool `json:"storeFormatted,omitempty"`
}

type DateHandling struct {
	DefaultTimeDisplayType                                                                    *DefaultTimeDisplayType `json:"defaultTimeDisplayType,omitempty"`
	// The time to be used to override the value found locally. In most cases this should be                          
	// left empty to allow the timezone to be derived naturally, but in cases of some mangaged                        
	// devices this may not be accurate. To see the timezone that your browser would provide,                         
	// open the developer console and type 'Intl.DateTimeFormat().resolvedOptions().timeZone'.                        
	// If that property does not match what is expected, set this value accordingly.                                  
	DefaultTimeZone                                                                           *string                 `json:"defaultTimeZone,omitempty"`
	// Whether or not to enable format strings that follow the syntax described by dayjs's                            
	// advancedFormat plugin.                                                                                         
	EnableAdvancedFormat                                                                      *bool                   `json:"enableAdvancedFormat,omitempty"`
	Format                                                                                    *FormatUnion            `json:"format"`
}

type FormatClass struct {
	// The format string passed to dayjs to format dates for display fields where size is less         
	// of a concern. Formatting using the advancedFormat plugin is enabled by default. Please          
	// see the documentation of dayjs for more information about the date formatting syntax,           
	// both with and without the advanced format plugin.                                               
	Long                                                                                       *string `json:"long,omitempty"`
	// The format string passed to dayjs to format dates for short or small display fields.            
	// Formatting using the advancedFormat plugin is enabled by default. Please see the                
	// documentation of dayjs for more information about the date formatting syntax, both with         
	// and without the advanced format plugin.                                                         
	Short                                                                                      *string `json:"short,omitempty"`
	// The format string passed to dayjs to format dates for fields where the time is relevant.        
	// Formatting using the advancedFormat plugin is enabled by default. Please see the                
	// documentation of dayjs for more information about the date formatting syntax, both with         
	// and without the advanced format plugin.                                                         
	TimeOnly                                                                                   *string `json:"timeOnly,omitempty"`
	// The format string passed to dayjs to format dates for fields where the time is relevant.        
	// Formatting using the advancedFormat plugin is enabled by default. Please see the                
	// documentation of dayjs for more information about the date formatting syntax, both with         
	// and without the advanced format plugin.                                                         
	WithTime                                                                                   *string `json:"withTime,omitempty"`
}

type Jupyter struct {
	CellInputWrappers                                                                           map[string]*CellInputWrapperValue `json:"cellInputWrappers,omitempty"`
	// The *absolute* path to the python environment with which to open Jupyter cells and                                         
	// notebooks.                                                                                                                 
	Environment                                                                                 *string                           `json:"environment,omitempty"`
	// Whether or not to execute notebook cells immediately after loading.                                                        
	Execute                                                                                     *bool                             `json:"execute,omitempty"`
	// Whether or not to initally fold jupyter input cells that are embedded within mdx notes.                                    
	InitiallyFoldCells                                                                          *bool                             `json:"initiallyFoldCells,omitempty"`
	JupyterNotebookProps                                                                        *JupyterNotebookProps             `json:"jupyterNotebookProps,omitempty"`
	// Port on which JupyterServer instance is running. This value must match the value in your                                   
	// local jupyter server config.                                                                                               
	JupyterPort                                                                                 *int64                            `json:"jupyterPort,omitempty"`
	JupyterReactProps                                                                           *JupyterReactProps                `json:"jupyterReactProps,omitempty"`
	// A secure token with which to connect to the jupyter server instance. This token must also                                  
	// be present in the jupyter_server_config.py file related to that environment. Any 64                                        
	// character random string of alpha-numeric characters will work.                                                             
	JupyterToken                                                                                *string                           `json:"jupyterToken,omitempty"`
	// Kernel name to use. Can be overriden with the search param ?kernel=someKernelName                                          
	Kernel                                                                                      *string                           `json:"kernel,omitempty"`
	NbConvert                                                                                   *NbConvert                        `json:"nbConvert,omitempty"`
	SyntaxHighlightTheme                                                                        *string                           `json:"syntaxHighlightTheme,omitempty"`
}

type CellInputWrapperClass struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type JupyterNotebookProps struct {
	// iPython widgets which to include in the notebook environment. This is most likely                            
	// unnecessary, but is made available for those with Jupyter experience and more advanced                       
	// use cases.                                                                                                   
	BundledIPyWidgets                                                                           []BundledIPyWidget  `json:"bundledIPyWidgets,omitempty"`
	// Can be set on a per-use basis with the ?cellMeta=true or ?cellMeta=false search param.                       
	CellMetadataPanel                                                                           *bool               `json:"cellMetadataPanel,omitempty"`
	// The margin of each cell that should be left for the sidebar menu. This is often best left                    
	// to the default or to be populated by plugins that might alter the default layout.                            
	CellSidebarMargin                                                                           *float64            `json:"cellSidebarMargin,omitempty"`
	// iPython widgets which to include in the notebook environment. This is most likely                            
	// unnecessary, but is made available for those with Jupyter experience and more advanced                       
	// use cases.                                                                                                   
	ExternalIPyWidgets                                                                          []ExternalIPyWidget `json:"externalIPyWidgets,omitempty"`
	// A css-in-tsx compatible string to be used for the notebook's height. This is made                            
	// accessible to combat possible bugs, but in most use cases this should remain the default                     
	// value.                                                                                                       
	Height                                                                                      *string             `json:"height,omitempty"`
	// A css-in-tsx compatible string to be used for the notebook's maximum height. This is made                    
	// accessible to combat possible bugs, but in most use cases this should remain the default                     
	// value.                                                                                                       
	MaxHeight                                                                                   *string             `json:"maxHeight,omitempty"`
	// Whether or not to load the nbgrader tool with the notebook. This can be overridden with                      
	// the ?nbgrader=true or ?nbgrader=false search param.                                                          
	Nbgrader                                                                                    *bool               `json:"nbgrader,omitempty"`
	// Open the notebook without the ability to write to it. This can be overridden with the                        
	// search param ?nbReadOnly=true or ?nbReadOnly=false.                                                          
	ReadOnly                                                                                    *bool               `json:"readOnly,omitempty"`
}

type BundledIPyWidget struct {
	Module  *Module `json:"module"`
	Name    string  `json:"name"`
	Version string  `json:"version"`
}

type ExternalIPyWidget struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type JupyterReactProps struct {
	// Whether the notebook is collaborative or not. This can also be set as needed by using the            
	// search param ?colab=true                                                                             
	Collaborative                                                                               *bool       `json:"collaborative,omitempty"`
	// Whether or not to use jupyterLite.                                                                   
	Lite                                                                                        *bool       `json:"lite,omitempty"`
	ServerUrls                                                                                  *ServerUrls `json:"serverUrls,omitempty"`
}

type ServerUrls struct {
	// The url with which to reach the jupyter server. Under the majority of use cases these            
	// values should be left to autopopulate based on the jupyter.serverPort value. Applying            
	// your own serverUrl should be reserved for those with experience connecting to a remote           
	// jupyter instance.                                                                                
	BaseURL                                                                                     *string `json:"baseUrl,omitempty"`
	// The websocket url with which to reach the jupyter server. Under the majority of use cases        
	// these values should be left to autopopulate based on the jupyter.serverPort value.               
	// Applying your own serverUrl should be reserved for those with experience connecting to a         
	// remote jupyter instance.                                                                         
	WsURL                                                                                       *string `json:"wsUrl,omitempty"`
}

type NbConvert struct {
	// The timeout in seconds with which to allow Jupyter notebooks to be converted to other                           
	// formats. This option will likely be deprecated as the ability to output notebooks in                            
	// alternative formats such as a pdf is being pushed back.                                                         
	ConversionTimeout                                                                        *float64                  `json:"conversionTimeout,omitempty"`
	CustomConversionFunction                                                                 *CustomConversionFunction `json:"customConversionFunction,omitempty"`
	// Whether or not to execute notebooks before converting via nbConvert. Can be overriden                           
	// with the ?nbConvertExec=true of ?nbConvertExec=false search param.                                              
	Execute                                                                                  *bool                     `json:"execute,omitempty"`
	// The template option that is passed to nbConvert. Can be overridden by the search param                          
	// ?nbConvertTemp=lab or ?nbConvertTemp=classic                                                                    
	NbConvertTemplate                                                                        *string                   `json:"nbConvertTemplate,omitempty"`
	// The path to nbConvert. This option will likely be deprecated.                                                   
	NbconvertPath                                                                            *string                   `json:"nbconvertPath,omitempty"`
	// The fsRoot relative file path where converted notebooks should be saved.                                        
	NotebookOutputDir                                                                        *string                   `json:"notebookOutputDir,omitempty"`
}

type CustomConversionFunction struct {
}

type Math struct {
	Constants                                                                                  map[string]float64  `json:"constants,omitempty"`
	// URL of the latex font to be used. This should be left to it's default value unless an                       
	// additional collection of .woff files are included in the build.additionalPublicResources                    
	// field.                                                                                                      
	LatexFontURL                                                                               *string             `json:"latexFontUrl,omitempty"`
	LatexPackages                                                                              *LatexPackagesUnion `json:"latexPackages"`
}

type Meta struct {
	Desc  *string `json:"desc,omitempty"`
	Title *string `json:"title,omitempty"`
}

type Navigation struct {
	// The default length of search results to be retrieved for each search result type. Default                     
	// values will likely work well for the default layout, but as more layouts become available                     
	// this setting may be adjusted to fit different search result layouts.                                          
	MaxResultLength                                                                             *MaxResultLength     `json:"maxResultLength,omitempty"`
	NavbarBreakpoint                                                                            *NavbarBreakpoint    `json:"navbarBreakpoint,omitempty"`
	// Either the document type id or the internalLink id to be to included in the navbar.                           
	NavbarLinks                                                                                 []NavbarLinkElement  `json:"navbarLinks,omitempty"`
	// Either the document type id or the internalLink id to be to included in the navbar.                           
	SidebarLinks                                                                                []SidebarLinkElement `json:"sidebarLinks,omitempty"`
}

// The default length of search results to be retrieved for each search result type. Default
// values will likely work well for the default layout, but as more layouts become available
// this setting may be adjusted to fit different search result layouts.
type MaxResultLength struct {
	Categories *float64 `json:"categories,omitempty"`
	Equations  *float64 `json:"equations,omitempty"`
	SearchAll  *float64 `json:"searchAll,omitempty"`
	Snippets   *float64 `json:"snippets,omitempty"`
}

type NavbarBreakpoint struct {
	// Pixel width at which the navbar will display all of it's contents.                               
	Full                                                                                       *float64 `json:"full,omitempty"`
	// Pixel width at which the navbar will display a minimal bit of content to render properly         
	// on more narrow screens. All pixel widths below this value will render a navbar with only         
	// required content. With many layouts, this means that pixel widths below this value will          
	// render only a search bar.                                                                        
	Minimal                                                                                    *float64 `json:"minimal,omitempty"`
}

type NavbarLinkClass struct {
	Href  *string `json:"href,omitempty"`
	Icon  *string `json:"icon,omitempty"`
	Label string  `json:"label"`
}

type SidebarLinkClass struct {
	Href  *string `json:"href,omitempty"`
	Icon  string  `json:"icon"`
	Label *string `json:"label,omitempty"`
}

type NoteType struct {
	// Automatically append these subjects to all notes of this document type. This can also be                      
	// done interactively through the app's UI after it is built.                                                    
	AutoSubject                                                                                 []string             `json:"autoSubject,omitempty"`
	// Automatically append these tags to all notes of this document type. This can also be done                     
	// interactively through the app's UI after it is built.                                                         
	AutoTag                                                                                     []string             `json:"autoTag,omitempty"`
	// Automatically append these topics to all notes of this document type. This can also be                        
	// done interactively through the app's UI after it is built.                                                    
	AutoTopic                                                                                   []string             `json:"autoTopic,omitempty"`
	// A unique key which describes the nature of this document type: 'MathNote', 'Journal',                         
	// 'References', etc...                                                                                          
	DocType                                                                                     *string              `json:"docType,omitempty"`
	// The path to the root of the directory which holds this document type. This path must be                       
	// both relative to the appConfig.fsRoot folder and a child of it. For example, a root                           
	// directory at '/Users/steve/Desktop/notes' might have folders within it of                                     
	// /Users/steve/Desktop/notes/math and /Users/steve/Desktop/notes/physics. These                                 
	// appConfig.noteTypes[0].fs should be '/math' and the latter should be '/physics'.                              
	FS                                                                                          string               `json:"fs"`
	// A glob style string to test a file path for this note type. Should return true if a given                     
	// file is of this note type.                                                                                    
	FilePathPattern                                                                             *string              `json:"filePathPattern,omitempty"`
	// A unique key to be used internally to deal with this doctype.                                                 
	ID                                                                                          *string              `json:"id,omitempty"`
	Icon                                                                                        *string              `json:"icon,omitempty"`
	InNavbar                                                                                    *bool                `json:"inNavbar,omitempty"`
	InSidebar                                                                                   *bool                `json:"inSidebar,omitempty"`
	// Keywords to help with search results and command sorting related to this document type.                       
	Keywords                                                                                    []string             `json:"keywords,omitempty"`
	// The label to be used for this document type where automatically generated links, commands                     
	// and buttons are referencing it.                                                                               
	Label                                                                                       string               `json:"label"`
	// An extra weight between 0 and 100 to apply to matches. This can be very important when                        
	// the location of one document type exists as a child of another, in which case an                              
	// increased weight should be applied to the child document type. Default: 50                                    
	MatchWeight                                                                                 *float64             `json:"matchWeight,omitempty"`
	// Replace references to the 'subject' tag with this label. This is useful for things like                       
	// managing freelance work, where 'subjects' might better be described as 'Jobs' or                              
	// 'Clients'.                                                                                                    
	SubjectLabel                                                                                *string              `json:"subjectLabel,omitempty"`
	// Replace references to the 'topic' tag with this label. This is useful for things like                         
	// managing freelance work, where 'topics' might better be described as 'Jobs' or 'Clients'.                     
	TopicLabel                                                                                  *string              `json:"topicLabel,omitempty"`
	UI                                                                                          *NoteTypeUI          `json:"UI,omitempty"`
	// The url at which this note should be displayed.                                                               
	URL                                                                                         *string              `json:"url,omitempty"`
	// Url search paramters to be appended to generated buttons and links for this doc type in                       
	// some cases. Useful for things like populating an initial list or opening with certain                         
	// default override-able settings.                                                                               
	URLQuery                                                                                    map[string]*URLQuery `json:"urlQuery,omitempty"`
}

type NoteTypeUI struct {
	Styles *Styles `json:"styles,omitempty"`
}

type Styles struct {
	Dark  *DarkClass `json:"dark,omitempty"`
	Light *DarkClass `json:"light,omitempty"`
}

type DarkClass struct {
	// Css classes to be appended to UI specific to this doc type. Tailwind classes will work.        
	Bg                                                                                        *string `json:"bg,omitempty"`
	// Css classes to be appended to UI specific to this doc type. Tailwind classes will work.        
	Fg                                                                                        *string `json:"fg,omitempty"`
}

type Performance struct {
	// The period to wait in milliseconds between parsing latex content that is expected to         
	// change rapidly.                                                                              
	LatexParsingDebounceTimeout                                                            *float64 `json:"latexParsingDebounceTimeout,omitempty"`
	// The period to wait in milliseconds between parsing markdown and mdx content that is          
	// expected to change rapidly.                                                                  
	MdxParsingDebounceTimeout                                                              *float64 `json:"mdxParsingDebounceTimeout,omitempty"`
}

type Plotting struct {
	PlotColorCycleMethod *PlotColorCycleMethod `json:"plotColorCycleMethod,omitempty"`
	PlotColorList        *PlotColorListUnion   `json:"plotColorList"`
}

type PlotColorListClass struct {
	Dark  []string `json:"dark"`
	Light []string `json:"light"`
}

type PluginsClass struct {
	Name        string   `json:"name"`
	ParserIndex *float64 `json:"parserIndex,omitempty"`
	Version     *string  `json:"version,omitempty"`
}

type Slots struct {
	Bibliography   *Bibliography `json:"bibliography"`
	CommandPalette *Bibliography `json:"commandPalette"`
	Dashboard      *Bibliography `json:"dashboard"`
	Editor         *Bibliography `json:"editor"`
	Form           *Bibliography `json:"form"`
	Math           *Bibliography `json:"math"`
	Navigation     *Bibliography `json:"navigation"`
	PDF            *Bibliography `json:"pdf"`
	Snippets       *Bibliography `json:"snippets"`
	TaskManager    *Bibliography `json:"taskManager"`
	UI             *Bibliography `json:"UI"`
}

type AppConfi struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}

type Terminal struct {
	LogLevel *LogLevel `json:"logLevel,omitempty"`
}

type AppConfigUI struct {
	AutoApplyMdxTitles *bool                  `json:"autoApplyMdxTitles,omitempty"`
	Colors             map[string]interface{} `json:"colors,omitempty"`
	Media              *Media                 `json:"media,omitempty"`
	Table              *Table                 `json:"table,omitempty"`
	Text               *Text                  `json:"text,omitempty"`
	Theme              *ThemeEnum             `json:"theme,omitempty"`
}

type Media struct {
	ImageMap                                                                                   map[string]string `json:"imageMap,omitempty"`
	ImageRemoteTest                                                                            []interface{}     `json:"imageRemoteTest,omitempty"`
	// Whether or not to include a default image map of light weight svg images that respond to                  
	// the theme's colors to be easily embedded by their alias. Can be disabled to minimize the                  
	// build size if they are unlikely to be used.                                                               
	IncludeDefaultImageMap                                                                     *bool             `json:"includeDefaultImageMap,omitempty"`
}

type Table struct {
	// The default height at which to limit tables. This can be overridden for each not                    
	// individually in that note's front matter. If the value is a string, it must be able to be           
	// interpretted by css-in-tsx syntax. Numbers will automatically be interpreted as pixels.             
	MaxHeight                                                                                   *MaxHeight `json:"maxHeight"`
}

type Text struct {
	// Whether or not to italicize block quote text.                                  
	BlockQuoteItalic                                                  *bool           `json:"blockQuoteItalic,omitempty"`
	// An array of font config objects to be bundled during the build.                
	FontPaths                                                         *FontPathsUnion `json:"fontPaths"`
}

type FontPath struct {
	// Should be an absolute path. As this is only required at build time, it is not necessary       
	// for this path to be a child of the fsRoot directory.                                          
	Path                                                                                      string `json:"path"`
	Style                                                                                     Style  `json:"style"`
	Weight                                                                                    Weight `json:"weight"`
}

type Prioritize string

const (
	Size  Prioritize = "size"
	Speed Prioritize = "speed"
)

type Type string

const (
	Postgres Type = "postgres"
	Sqlite   Type = "sqlite"
)

type Block string

const (
	APL              Block = "apl"
	ASM              Block = "asm"
	AWK              Block = "awk"
	Abap             Block = "abap"
	Actionscript3    Block = "actionscript-3"
	Ada              Block = "ada"
	AngularHTML      Block = "angular-html"
	AngularTs        Block = "angular-ts"
	Apache           Block = "apache"
	Apex             Block = "apex"
	Applescript      Block = "applescript"
	Ara              Block = "ara"
	Astro            Block = "astro"
	Ballerina        Block = "ballerina"
	Bash             Block = "bash"
	Bat              Block = "bat"
	Batch            Block = "batch"
	Be               Block = "be"
	Beancount        Block = "beancount"
	Berry            Block = "berry"
	Bibtex           Block = "bibtex"
	Bicep            Block = "bicep"
	Blade            Block = "blade"
	BlockC           Block = "c++"
	BlockF           Block = "f#"
	C                Block = "c"
	COBOL            Block = "cobol"
	CS               Block = "cs"
	CSS              Block = "css"
	CSV              Block = "csv"
	Cadence          Block = "cadence"
	Cdc              Block = "cdc"
	Clarity          Block = "clarity"
	Clj              Block = "clj"
	Clojure          Block = "clojure"
	Cmake            Block = "cmake"
	Cmd              Block = "cmd"
	Codeql           Block = "codeql"
	Coffee           Block = "coffee"
	Coffeescript     Block = "coffeescript"
	Console          Block = "console"
	Cpp              Block = "cpp"
	Cql              Block = "cql"
	Crystal          Block = "crystal"
	Csharp           Block = "csharp"
	Cue              Block = "cue"
	Cypher           Block = "cypher"
	D                Block = "d"
	Dart             Block = "dart"
	Dax              Block = "dax"
	Diff             Block = "diff"
	Docker           Block = "docker"
	Dockerfile       Block = "dockerfile"
	DreamMaker       Block = "dream-maker"
	Elixir           Block = "elixir"
	Elm              Block = "elm"
	Erb              Block = "erb"
	Erl              Block = "erl"
	Erlang           Block = "erlang"
	F                Block = "f"
	F03              Block = "f03"
	F08              Block = "f08"
	F18              Block = "f18"
	F77              Block = "f77"
	F90              Block = "f90"
	F95              Block = "f95"
	FS               Block = "fs"
	Fish             Block = "fish"
	For              Block = "for"
	FortranFixedForm Block = "fortran-fixed-form"
	FortranFreeForm  Block = "fortran-free-form"
	Fsharp           Block = "fsharp"
	Fsl              Block = "fsl"
	Gdresource       Block = "gdresource"
	Gdscript         Block = "gdscript"
	Gdshader         Block = "gdshader"
	Gherkin          Block = "gherkin"
	GitCommit        Block = "git-commit"
	GitRebase        Block = "git-rebase"
	Gjs              Block = "gjs"
	GlimmerJS        Block = "glimmer-js"
	GlimmerTs        Block = "glimmer-ts"
	Glsl             Block = "glsl"
	Gnuplot          Block = "gnuplot"
	Go               Block = "go"
	Gql              Block = "gql"
	Graphql          Block = "graphql"
	Groovy           Block = "groovy"
	Gts              Block = "gts"
	HCL              Block = "hcl"
	HTML             Block = "html"
	HTMLDerivative   Block = "html-derivative"
	HTTP             Block = "http"
	Hack             Block = "hack"
	Haml             Block = "haml"
	Handlebars       Block = "handlebars"
	Haskell          Block = "haskell"
	Hbs              Block = "hbs"
	Hjson            Block = "hjson"
	Hlsl             Block = "hlsl"
	Hs               Block = "hs"
	Imba             Block = "imba"
	Ini              Block = "ini"
	JS               Block = "js"
	JSON             Block = "json"
	Jade             Block = "jade"
	Java             Block = "java"
	Javascript       Block = "javascript"
	Jinja            Block = "jinja"
	Jison            Block = "jison"
	Json5            Block = "json5"
	Jsonc            Block = "jsonc"
	Jsonl            Block = "jsonl"
	Jsonnet          Block = "jsonnet"
	Jssm             Block = "jssm"
	Jsx              Block = "jsx"
	Julia            Block = "julia"
	Kotlin           Block = "kotlin"
	Kql              Block = "kql"
	Kt               Block = "kt"
	Kts              Block = "kts"
	Kusto            Block = "kusto"
	LISP             Block = "lisp"
	Latex            Block = "latex"
	Less             Block = "less"
	Liquid           Block = "liquid"
	Logo             Block = "logo"
	Lua              Block = "lua"
	Make             Block = "make"
	Makefile         Block = "makefile"
	Markdown         Block = "markdown"
	Marko            Block = "marko"
	Matlab           Block = "matlab"
	Md               Block = "md"
	Mdc              Block = "mdc"
	Mdx              Block = "mdx"
	Mermaid          Block = "mermaid"
	Mojo             Block = "mojo"
	Move             Block = "move"
	NIM              Block = "nim"
	Nar              Block = "nar"
	Narrat           Block = "narrat"
	Nextflow         Block = "nextflow"
	Nf               Block = "nf"
	Nginx            Block = "nginx"
	Nix              Block = "nix"
	Nu               Block = "nu"
	Nushell          Block = "nushell"
	Objc             Block = "objc"
	ObjectiveC       Block = "objective-c"
	ObjectiveCpp     Block = "objective-cpp"
	Ocaml            Block = "ocaml"
	PERL             Block = "perl"
	PHP              Block = "php"
	PS               Block = "ps"
	Pascal           Block = "pascal"
	Perl6            Block = "perl6"
	Plsql            Block = "plsql"
	Postcss          Block = "postcss"
	Powerquery       Block = "powerquery"
	Powershell       Block = "powershell"
	Prisma           Block = "prisma"
	Prolog           Block = "prolog"
	Properties       Block = "properties"
	Proto            Block = "proto"
	Ps1              Block = "ps1"
	Pug              Block = "pug"
	Puppet           Block = "puppet"
	Purescript       Block = "purescript"
	PurpleC          Block = "c#"
	Py               Block = "py"
	Python           Block = "python"
	Ql               Block = "ql"
	R                Block = "r"
	Raku             Block = "raku"
	Razor            Block = "razor"
	Rb               Block = "rb"
	Reg              Block = "reg"
	Rel              Block = "rel"
	Riscv            Block = "riscv"
	Rs               Block = "rs"
	Rst              Block = "rst"
	Ruby             Block = "ruby"
	Rust             Block = "rust"
	SAS              Block = "sas"
	SQL              Block = "sql"
	SSHConfig        Block = "ssh-config"
	Sass             Block = "sass"
	Scala            Block = "scala"
	Scheme           Block = "scheme"
	Scss             Block = "scss"
	Sh               Block = "sh"
	Shader           Block = "shader"
	Shaderlab        Block = "shaderlab"
	Shell            Block = "shell"
	Shellscript      Block = "shellscript"
	Shellsession     Block = "shellsession"
	Smalltalk        Block = "smalltalk"
	Solidity         Block = "solidity"
	Sparql           Block = "sparql"
	Spl              Block = "spl"
	Splunk           Block = "splunk"
	Stata            Block = "stata"
	Styl             Block = "styl"
	Stylus           Block = "stylus"
	Svelte           Block = "svelte"
	Swift            Block = "swift"
	SystemVerilog    Block = "system-verilog"
	Tasl             Block = "tasl"
	Tcl              Block = "tcl"
	Terraform        Block = "terraform"
	Tex              Block = "tex"
	Tf               Block = "tf"
	Tfvars           Block = "tfvars"
	Toml             Block = "toml"
	Ts               Block = "ts"
	Tsx              Block = "tsx"
	Turtle           Block = "turtle"
	Twig             Block = "twig"
	Typ              Block = "typ"
	Typescript       Block = "typescript"
	Typst            Block = "typst"
	V                Block = "v"
	VB               Block = "vb"
	Verilog          Block = "verilog"
	Vhdl             Block = "vhdl"
	Vim              Block = "vim"
	Viml             Block = "viml"
	Vimscript        Block = "vimscript"
	Vue              Block = "vue"
	VueHTML          Block = "vue-html"
	Vy               Block = "vy"
	Vyper            Block = "vyper"
	WASM             Block = "wasm"
	Wenyan           Block = "wenyan"
	Wgsl             Block = "wgsl"
	Wl               Block = "wl"
	Wolfram          Block = "wolfram"
	XML              Block = "xml"
	XSL              Block = "xsl"
	YAML             Block = "yaml"
	Yml              Block = "yml"
	Zenscript        Block = "zenscript"
	Zig              Block = "zig"
	Zsh              Block = "zsh"
	文言               Block = "文言"
)

// Syntax highlighting theme to be used when in dark mode. Can be overriden with the
// ?shikiHlDark=dracula search param.
//
// Syntax highlighting theme to be used when in light mode. Can be overriden with the
// ?shikiHlLight=nord search param.
type DarkEnum string

const (
	Andromeeda             DarkEnum = "andromeeda"
	AuroraX                DarkEnum = "aurora-x"
	AyuDark                DarkEnum = "ayu-dark"
	CatppuccinFrappe       DarkEnum = "catppuccin-frappe"
	CatppuccinLatte        DarkEnum = "catppuccin-latte"
	CatppuccinMacchiato    DarkEnum = "catppuccin-macchiato"
	CatppuccinMocha        DarkEnum = "catppuccin-mocha"
	DarkPlus               DarkEnum = "dark-plus"
	DarkRed                DarkEnum = "red"
	Dracula                DarkEnum = "dracula"
	DraculaSoft            DarkEnum = "dracula-soft"
	GithubDark             DarkEnum = "github-dark"
	GithubDarkDimmed       DarkEnum = "github-dark-dimmed"
	GithubLight            DarkEnum = "github-light"
	LightPlus              DarkEnum = "light-plus"
	MaterialTheme          DarkEnum = "material-theme"
	MaterialThemeDarker    DarkEnum = "material-theme-darker"
	MaterialThemeLighter   DarkEnum = "material-theme-lighter"
	MaterialThemeOcean     DarkEnum = "material-theme-ocean"
	MaterialThemePalenight DarkEnum = "material-theme-palenight"
	MinDark                DarkEnum = "min-dark"
	MinLight               DarkEnum = "min-light"
	Monokai                DarkEnum = "monokai"
	NightOwl               DarkEnum = "night-owl"
	Nord                   DarkEnum = "nord"
	OneDarkPro             DarkEnum = "one-dark-pro"
	Poimandres             DarkEnum = "poimandres"
	RosePine               DarkEnum = "rose-pine"
	RosePineDawn           DarkEnum = "rose-pine-dawn"
	RosePineMoon           DarkEnum = "rose-pine-moon"
	SlackDark              DarkEnum = "slack-dark"
	SlackOchin             DarkEnum = "slack-ochin"
	SolarizedDark          DarkEnum = "solarized-dark"
	SolarizedLight         DarkEnum = "solarized-light"
	Synthwave84            DarkEnum = "synthwave-84"
	TokyoNight             DarkEnum = "tokyo-night"
	Vesper                 DarkEnum = "vesper"
	VitesseBlack           DarkEnum = "vitesse-black"
	VitesseDark            DarkEnum = "vitesse-dark"
	VitesseLight           DarkEnum = "vitesse-light"
)

// Similar to a digital clock.
//
// A short but more natural description of time. This requires the
// dateParseFormat.enableAdvancedFormat option.
//
// Time with relative times. ie: '2 days ago', 'tomorrow', etc. This requires the
// dateParseFormat.enableAdvancedFormat option. Default: summarized.
type DefaultTimeDisplayType string

const (
	Analog      DefaultTimeDisplayType = "analog"
	Descriptive DefaultTimeDisplayType = "descriptive"
	Summarized  DefaultTimeDisplayType = "summarized"
)

type FileTypePriority string

const (
	DB                   FileTypePriority = ".db"
	Excel                FileTypePriority = ".excel"
	FileTypePriorityCSV  FileTypePriority = ".csv"
	FileTypePriorityHTML FileTypePriority = ".html"
	FileTypePriorityJSON FileTypePriority = ".json"
	FileTypePriorityMd   FileTypePriority = ".md"
	FileTypePriorityMdx  FileTypePriority = ".mdx"
	FileTypePrioritySQL  FileTypePriority = ".sql"
	FileTypePriorityTex  FileTypePriority = ".tex"
	Hdf5                 FileTypePriority = ".hdf5"
	Ipynb                FileTypePriority = ".ipynb"
	Numpy                FileTypePriority = ".numpy"
	PDF                  FileTypePriority = ".pdf"
	Pickle               FileTypePriority = ".pickle"
	Tsv                  FileTypePriority = ".tsv"
)

type LatexPackagesEnum string

const (
	All LatexPackagesEnum = "all"
)

type PlotColorCycleMethod string

const (
	InOrder PlotColorCycleMethod = "inOrder"
	Random  PlotColorCycleMethod = "random"
)

type LogLevel string

const (
	Debug   LogLevel = "debug"
	Info    LogLevel = "info"
	None    LogLevel = "none"
	Verbose LogLevel = "verbose"
)

type Style string

const (
	Bold   Style = "bold"
	Italic Style = "italic"
	Normal Style = "normal"
)

type Weight string

const (
	The100 Weight = "100"
	The200 Weight = "200"
	The300 Weight = "300"
	The400 Weight = "400"
	The500 Weight = "500"
	The600 Weight = "600"
	The700 Weight = "700"
	The800 Weight = "800"
	The900 Weight = "900"
)

type FontPathsEnum string

const (
	Default FontPathsEnum = "default"
)

type ThemeEnum string

const (
	Blue     ThemeEnum = "blue"
	Gray     ThemeEnum = "gray"
	Green    ThemeEnum = "green"
	Neutral  ThemeEnum = "neutral"
	Orange   ThemeEnum = "orange"
	Rose     ThemeEnum = "rose"
	Slate    ThemeEnum = "slate"
	Stone    ThemeEnum = "stone"
	ThemeRed ThemeEnum = "red"
	Ulld     ThemeEnum = "ulld"
	Violet   ThemeEnum = "violet"
	Yellow   ThemeEnum = "yellow"
	Zinc     ThemeEnum = "zinc"
)

type DefaultLanguageUnion struct {
	DefaultLanguageClass *DefaultLanguageClass
	Enum                 *Block
}

func (x *DefaultLanguageUnion) UnmarshalJSON(data []byte) error {
	x.DefaultLanguageClass = nil
	x.Enum = nil
	var c DefaultLanguageClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.DefaultLanguageClass = &c
	}
	return nil
}

func (x *DefaultLanguageUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.DefaultLanguageClass != nil, x.DefaultLanguageClass, false, nil, x.Enum != nil, x.Enum, false)
}

type FormatUnion struct {
	FormatClass *FormatClass
	String      *string
}

func (x *FormatUnion) UnmarshalJSON(data []byte) error {
	x.FormatClass = nil
	var c FormatClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FormatClass = &c
	}
	return nil
}

func (x *FormatUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FormatClass != nil, x.FormatClass, false, nil, false, nil, false)
}

type CellInputWrapperValue struct {
	CellInputWrapperClass *CellInputWrapperClass
	String                *string
}

func (x *CellInputWrapperValue) UnmarshalJSON(data []byte) error {
	x.CellInputWrapperClass = nil
	var c CellInputWrapperClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CellInputWrapperClass = &c
	}
	return nil
}

func (x *CellInputWrapperValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.CellInputWrapperClass != nil, x.CellInputWrapperClass, false, nil, false, nil, false)
}

type Module struct {
	String    *string
	StringMap map[string]string
}

func (x *Module) UnmarshalJSON(data []byte) error {
	x.StringMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, false, nil, true, &x.StringMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Module) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, false, nil, x.StringMap != nil, x.StringMap, false, nil, false)
}

type LatexPackagesUnion struct {
	Enum        *LatexPackagesEnum
	StringArray []string
}

func (x *LatexPackagesUnion) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *LatexPackagesUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

type NavbarLinkElement struct {
	NavbarLinkClass *NavbarLinkClass
	String          *string
}

func (x *NavbarLinkElement) UnmarshalJSON(data []byte) error {
	x.NavbarLinkClass = nil
	var c NavbarLinkClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.NavbarLinkClass = &c
	}
	return nil
}

func (x *NavbarLinkElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.NavbarLinkClass != nil, x.NavbarLinkClass, false, nil, false, nil, false)
}

type SidebarLinkElement struct {
	SidebarLinkClass *SidebarLinkClass
	String           *string
}

func (x *SidebarLinkElement) UnmarshalJSON(data []byte) error {
	x.SidebarLinkClass = nil
	var c SidebarLinkClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.SidebarLinkClass = &c
	}
	return nil
}

func (x *SidebarLinkElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.SidebarLinkClass != nil, x.SidebarLinkClass, false, nil, false, nil, false)
}

type URLQuery struct {
	Double     *float64
	String     *string
	UnionArray []MaxHeight
}

func (x *URLQuery) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, &x.String, true, &x.UnionArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *URLQuery) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, x.String, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, false)
}

// The default height at which to limit tables. This can be overridden for each not
// individually in that note's front matter. If the value is a string, it must be able to be
// interpretted by css-in-tsx syntax. Numbers will automatically be interpreted as pixels.
type MaxHeight struct {
	Double *float64
	String *string
}

func (x *MaxHeight) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, &x.Double, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *MaxHeight) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type PlotColorListUnion struct {
	PlotColorListClass *PlotColorListClass
	StringArray        []string
}

func (x *PlotColorListUnion) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.PlotColorListClass = nil
	var c PlotColorListClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PlotColorListClass = &c
	}
	return nil
}

func (x *PlotColorListUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, x.PlotColorListClass != nil, x.PlotColorListClass, false, nil, false, nil, false)
}

type PluginsUnion struct {
	PluginsClass *PluginsClass
	String       *string
	UnionArray   []Plugin
}

func (x *PluginsUnion) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	x.PluginsClass = nil
	var c PluginsClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PluginsClass = &c
	}
	return nil
}

func (x *PluginsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArray != nil, x.UnionArray, x.PluginsClass != nil, x.PluginsClass, false, nil, false, nil, false)
}

type Plugin struct {
	PluginsClass *PluginsClass
	String       *string
}

func (x *Plugin) UnmarshalJSON(data []byte) error {
	x.PluginsClass = nil
	var c PluginsClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PluginsClass = &c
	}
	return nil
}

func (x *Plugin) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PluginsClass != nil, x.PluginsClass, false, nil, false, nil, false)
}

type Bibliography struct {
	AppConfi   *AppConfi
	String     *string
	UnionArray []UI
}

func (x *Bibliography) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	x.AppConfi = nil
	var c AppConfi
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.AppConfi = &c
	}
	return nil
}

func (x *Bibliography) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArray != nil, x.UnionArray, x.AppConfi != nil, x.AppConfi, false, nil, false, nil, false)
}

type UI struct {
	AppConfi *AppConfi
	String   *string
}

func (x *UI) UnmarshalJSON(data []byte) error {
	x.AppConfi = nil
	var c AppConfi
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.AppConfi = &c
	}
	return nil
}

func (x *UI) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.AppConfi != nil, x.AppConfi, false, nil, false, nil, false)
}

// An array of font config objects to be bundled during the build.
type FontPathsUnion struct {
	Enum          *FontPathsEnum
	FontPathArray []FontPath
}

func (x *FontPathsUnion) UnmarshalJSON(data []byte) error {
	x.FontPathArray = nil
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.FontPathArray, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *FontPathsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.FontPathArray != nil, x.FontPathArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
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
