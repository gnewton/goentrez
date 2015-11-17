package goentrez

type ESearchRoot struct {
	EeSearchResult *EeSearchResult `xml:" eSearchResult,omitempty" json:"eSearchResult,omitempty"`
}

type EeSearchResult struct {
	ECount            *ECount            `xml:" Count,omitempty" json:"Count,omitempty"`
	EIdList           *EIdList           `xml:" IdList,omitempty" json:"IdList,omitempty"`
	EWebEnv           *EWebEnv           `xml:"WebEnv,omitempty" json:"WebEnv,omitempty"`
	EQueryTranslation *EQueryTranslation `xml:" QueryTranslation,omitempty" json:"QueryTranslation,omitempty"`
	ERetMax           *ERetMax           `xml:" RetMax,omitempty" json:"RetMax,omitempty"`
	ERetStart         *ERetStart         `xml:" RetStart,omitempty" json:"RetStart,omitempty"`
	ETranslationSet   *ETranslationSet   `xml:" TranslationSet,omitempty" json:"TranslationSet,omitempty"`
	ETranslationStack *ETranslationStack `xml:" TranslationStack,omitempty" json:"TranslationStack,omitempty"`
}

type EWebEnv struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ERetMax struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ERetStart struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EIdList struct {
	EId []*EId `xml:" Id,omitempty" json:"Id,omitempty"`
}

type EId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ETranslationSet struct {
	ETranslation *ETranslation `xml:" Translation,omitempty" json:"Translation,omitempty"`
}

type ETranslation struct {
	EFrom *EFrom `xml:" From,omitempty" json:"From,omitempty"`
	ETo   *ETo   `xml:" To,omitempty" json:"To,omitempty"`
}

type EFrom struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ETo struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ETranslationStack struct {
	EOP      []*EOP      `xml:" OP,omitempty" json:"OP,omitempty"`
	ETermSet []*ETermSet `xml:" TermSet,omitempty" json:"TermSet,omitempty"`
}

type EOP struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ETermSet struct {
	ECount   *ECount   `xml:" Count,omitempty" json:"Count,omitempty"`
	EExplode *EExplode `xml:" Explode,omitempty" json:"Explode,omitempty"`
	EField   *EField   `xml:" Field,omitempty" json:"Field,omitempty"`
	ETerm    *ETerm    `xml:" Term,omitempty" json:"Term,omitempty"`
}

type ETerm struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EField struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ECount struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EExplode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EQueryTranslation struct {
	Text string `xml:",chardata" json:",omitempty"`
}
