package goentrez

const PARAM_DATETYPE = "datetype"
const PARAM_DB = "db"
const PARAM_FIELD = "field"
const PARAM_MAXDATE = "maxdate"
const PARAM_MINDATE = "mindate"
const PARAM_QUERY_KEY = "query_key"
const PARAM_RELDATE = "reldate"
const PARAM_RETMAX = "retmax"
const PARAM_RETMODE = "retmode" // not used
const PARAM_RETSTART = "retstart"
const PARAM_RETTYPE = "rettype"
const PARAM_SORT = "sort"
const PARAM_TERM = "term"
const PARAM_USE_HISTORY = "usehistory"
const PARAM_WEB_ENV = "WebEnv"

var VALID_DATABASES = map[string]bool{
	"bioproject":      true,
	"biosample":       true,
	"biosystems":      true,
	"books":           true,
	"cdd":             true,
	"gap":             true,
	"dbvar":           true,
	"epigenomics":     true,
	"nucest":          true,
	"gene":            true,
	"genome":          true,
	"gds":             true,
	"geoprofiles":     true,
	"nucgss":          true,
	"homologene":      true,
	"mesh":            true,
	"toolkit":         true,
	"ncbisearch":      true,
	"nlmcatalog":      true,
	"nuccore":         true,
	"omia":            true,
	"popset":          true,
	"probe":           true,
	"protein":         true,
	"proteinclusters": true,
	"pcassay":         true,
	"pccompound":      true,
	"pcsubstance":     true,
	"pubmed":          true,
	"pmc":             true,
	"snp":             true,
	"sra":             true,
	"structure":       true,
	"taxonomy":        true,
	"unigene":         true,
	"unists":          true,
}
