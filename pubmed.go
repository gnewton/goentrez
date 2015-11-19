package goentrez

// Generated using chidly https://github.com/gnewton/chidley

//<eFetchResult><ERROR>Supplied id parameter is empty.</ERROR></eFetchResult>

type errorRoot struct {
	FetchError *EFetchResultError `xml:" eFetchResult,omitempty"`
}
type EFetchResultError struct {
	Error *FetchError `xml:" ERROR,omitempty"`
}

type FetchError struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_root struct {
	Chi_PubmedArticleSet *Chi_PubmedArticleSet `xml:" PubmedArticleSet,omitempty" json:"PubmedArticleSet,omitempty"`
}

type Chi_PubmedArticleSet struct {
	Chi_PubmedArticle     []*Chi_PubmedArticle     `xml:" PubmedArticle,omitempty" json:"PubmedArticle,omitempty"`
	Chi_PubmedBookArticle []*Chi_PubmedBookArticle `xml:" PubmedBookArticle,omitempty" json:"PubmedBookArticle,omitempty"`
}

type Chi_PubmedArticle struct {
	Chi_MedlineCitation *Chi_MedlineCitation `xml:" MedlineCitation,omitempty" json:"MedlineCitation,omitempty"`
	Chi_PubmedData      *Chi_PubmedData      `xml:" PubmedData,omitempty" json:"PubmedData,omitempty"`
}

type Chi_MedlineCitation struct {
	Attr_Owner                  string                       `xml:" Owner,attr"  json:",omitempty"`
	Attr_Status                 string                       `xml:" Status,attr"  json:",omitempty"`
	Attr_VersionDate            string                       `xml:" VersionDate,attr"  json:",omitempty"`
	Attr_VersionID              string                       `xml:" VersionID,attr"  json:",omitempty"`
	Chi_Article                 *Chi_Article                 `xml:" Article,omitempty" json:"Article,omitempty"`
	Chi_ChemicalList            *Chi_ChemicalList            `xml:" ChemicalList,omitempty" json:"ChemicalList,omitempty"`
	Chi_CitationSubset          []*Chi_CitationSubset        `xml:" CitationSubset,omitempty" json:"CitationSubset,omitempty"`
	Chi_CommentsCorrectionsList *Chi_CommentsCorrectionsList `xml:" CommentsCorrectionsList,omitempty" json:"CommentsCorrectionsList,omitempty"`
	Chi_DateCompleted           *Chi_DateCompleted           `xml:" DateCompleted,omitempty" json:"DateCompleted,omitempty"`
	Chi_DateCreated             *Chi_DateCreated             `xml:" DateCreated,omitempty" json:"DateCreated,omitempty"`
	Chi_DateRevised             *Chi_DateRevised             `xml:" DateRevised,omitempty" json:"DateRevised,omitempty"`
	Chi_GeneSymbolList          *Chi_GeneSymbolList          `xml:" GeneSymbolList,omitempty" json:"GeneSymbolList,omitempty"`
	Chi_GeneralNote             []*Chi_GeneralNote           `xml:" GeneralNote,omitempty" json:"GeneralNote,omitempty"`
	Chi_InvestigatorList        *Chi_InvestigatorList        `xml:" InvestigatorList,omitempty" json:"InvestigatorList,omitempty"`
	Chi_KeywordList             *Chi_KeywordList             `xml:" KeywordList,omitempty" json:"KeywordList,omitempty"`
	Chi_MedlineJournalInfo      *Chi_MedlineJournalInfo      `xml:" MedlineJournalInfo,omitempty" json:"MedlineJournalInfo,omitempty"`
	Chi_MeshHeadingList         *Chi_MeshHeadingList         `xml:" MeshHeadingList,omitempty" json:"MeshHeadingList,omitempty"`
	Chi_NumberOfReferences      *Chi_NumberOfReferences      `xml:" NumberOfReferences,omitempty" json:"NumberOfReferences,omitempty"`
	Chi_OtherAbstract           *Chi_OtherAbstract           `xml:" OtherAbstract,omitempty" json:"OtherAbstract,omitempty"`
	Chi_OtherID                 []*Chi_OtherID               `xml:" OtherID,omitempty" json:"OtherID,omitempty"`
	Chi_PMID                    *Chi_PMID                    `xml:" PMID,omitempty" json:"PMID,omitempty"`
	Chi_PersonalNameSubjectList *Chi_PersonalNameSubjectList `xml:" PersonalNameSubjectList,omitempty" json:"PersonalNameSubjectList,omitempty"`
	Chi_SpaceFlightMission      []*Chi_SpaceFlightMission    `xml:" SpaceFlightMission,omitempty" json:"SpaceFlightMission,omitempty"`
	Chi_SupplMeshList           *Chi_SupplMeshList           `xml:" SupplMeshList,omitempty" json:"SupplMeshList,omitempty"`
}

type Chi_PMID struct {
	Attr_Version string `xml:" Version,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

type Chi_DateCreated struct {
	Chi_Day   *Chi_Day   `xml:" Day,omitempty" json:"Day,omitempty"`
	Chi_Month *Chi_Month `xml:" Month,omitempty" json:"Month,omitempty"`
	Chi_Year  *Chi_Year  `xml:" Year,omitempty" json:"Year,omitempty"`
}

type Chi_Day struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Year struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Month struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_DateRevised struct {
	Chi_Day   *Chi_Day   `xml:" Day,omitempty" json:"Day,omitempty"`
	Chi_Month *Chi_Month `xml:" Month,omitempty" json:"Month,omitempty"`
	Chi_Year  *Chi_Year  `xml:" Year,omitempty" json:"Year,omitempty"`
}

type Chi_MedlineJournalInfo struct {
	Chi_Country     *Chi_Country     `xml:" Country,omitempty" json:"Country,omitempty"`
	Chi_ISSNLinking *Chi_ISSNLinking `xml:" ISSNLinking,omitempty" json:"ISSNLinking,omitempty"`
	Chi_MedlineTA   *Chi_MedlineTA   `xml:" MedlineTA,omitempty" json:"MedlineTA,omitempty"`
	Chi_NlmUniqueID *Chi_NlmUniqueID `xml:" NlmUniqueID,omitempty" json:"NlmUniqueID,omitempty"`
}

type Chi_Country struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_MedlineTA struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_NlmUniqueID struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_ISSNLinking struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_CitationSubset struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_NumberOfReferences struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_KeywordList struct {
	Attr_Owner  string         `xml:" Owner,attr"  json:",omitempty"`
	Chi_Keyword []*Chi_Keyword `xml:" Keyword,omitempty" json:"Keyword,omitempty"`
}

type Chi_Keyword struct {
	Attr_MajorTopicYN string `xml:" MajorTopicYN,attr"  json:",omitempty"`
	Text              string `xml:",chardata" json:",omitempty"`
}

type Chi_OtherAbstract struct {
	Attr_Type                string                    `xml:" Type,attr"  json:",omitempty"`
	Attr_Language            string                    `xml:" Language,attr"  json:",omitempty"`
	Chi_AbstractText         []*Chi_AbstractText       `xml:" AbstractText,omitempty" json:"AbstractText,omitempty"`
	Chi_CopyrightInformation *Chi_CopyrightInformation `xml:" CopyrightInformation,omitempty" json:"CopyrightInformation,omitempty"`
}

type Chi_AbstractText struct {
	Attr_Label       string `xml:" Label,attr"  json:",omitempty"`
	Attr_NlmCategory string `xml:" NlmCategory,attr"  json:",omitempty"`
	Text             string `xml:",chardata" json:",omitempty"`
}

type Chi_CopyrightInformation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_PersonalNameSubjectList struct {
	Chi_PersonalNameSubject []*Chi_PersonalNameSubject `xml:" PersonalNameSubject,omitempty" json:"PersonalNameSubject,omitempty"`
}

type Chi_PersonalNameSubject struct {
	Chi_ForeName *Chi_ForeName `xml:" ForeName,omitempty" json:"ForeName,omitempty"`
	Chi_Initials *Chi_Initials `xml:" Initials,omitempty" json:"Initials,omitempty"`
	Chi_LastName *Chi_LastName `xml:" LastName,omitempty" json:"LastName,omitempty"`
	Chi_Suffix   *Chi_Suffix   `xml:" Suffix,omitempty" json:"Suffix,omitempty"`
}

type Chi_LastName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_ForeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Initials struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Suffix struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_GeneSymbolList struct {
	Chi_GeneSymbol []*Chi_GeneSymbol `xml:" GeneSymbol,omitempty" json:"GeneSymbol,omitempty"`
}

type Chi_GeneSymbol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_SupplMeshList struct {
	Chi_SupplMeshName []*Chi_SupplMeshName `xml:" SupplMeshName,omitempty" json:"SupplMeshName,omitempty"`
}

type Chi_SupplMeshName struct {
	Attr_Type string `xml:" Type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Chi_DateCompleted struct {
	Chi_Day   *Chi_Day   `xml:" Day,omitempty" json:"Day,omitempty"`
	Chi_Month *Chi_Month `xml:" Month,omitempty" json:"Month,omitempty"`
	Chi_Year  *Chi_Year  `xml:" Year,omitempty" json:"Year,omitempty"`
}

type Chi_Article struct {
	Attr_PubModel           string                   `xml:" PubModel,attr"  json:",omitempty"`
	Chi_Abstract            *Chi_Abstract            `xml:" Abstract,omitempty" json:"Abstract,omitempty"`
	Chi_ArticleDate         *Chi_ArticleDate         `xml:" ArticleDate,omitempty" json:"ArticleDate,omitempty"`
	Chi_ArticleTitle        *Chi_ArticleTitle        `xml:" ArticleTitle,omitempty" json:"ArticleTitle,omitempty"`
	Chi_AuthorList          []*Chi_AuthorList        `xml:" AuthorList,omitempty" json:"AuthorList,omitempty"`
	Chi_DataBankList        *Chi_DataBankList        `xml:" DataBankList,omitempty" json:"DataBankList,omitempty"`
	Chi_ELocationID         []*Chi_ELocationID       `xml:" ELocationID,omitempty" json:"ELocationID,omitempty"`
	Chi_GrantList           *Chi_GrantList           `xml:" GrantList,omitempty" json:"GrantList,omitempty"`
	Chi_Journal             *Chi_Journal             `xml:" Journal,omitempty" json:"Journal,omitempty"`
	Chi_Language            []*Chi_Language          `xml:" Language,omitempty" json:"Language,omitempty"`
	Chi_Pagination          *Chi_Pagination          `xml:" Pagination,omitempty" json:"Pagination,omitempty"`
	Chi_PublicationTypeList *Chi_PublicationTypeList `xml:" PublicationTypeList,omitempty" json:"PublicationTypeList,omitempty"`
	Chi_VernacularTitle     *Chi_VernacularTitle     `xml:" VernacularTitle,omitempty" json:"VernacularTitle,omitempty"`
}

type Chi_Pagination struct {
	Chi_MedlinePgn *Chi_MedlinePgn `xml:" MedlinePgn,omitempty" json:"MedlinePgn,omitempty"`
}

type Chi_MedlinePgn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_AuthorList struct {
	Attr_CompleteYN string        `xml:" CompleteYN,attr"  json:",omitempty"`
	Attr_Type       string        `xml:" Type,attr"  json:",omitempty"`
	Chi_Author      []*Chi_Author `xml:" Author,omitempty" json:"Author,omitempty"`
}

type Chi_Author struct {
	Attr_ValidYN       string              `xml:" ValidYN,attr"  json:",omitempty"`
	Chi_Affiliation    *Chi_Affiliation    `xml:" Affiliation,omitempty" json:"Affiliation,omitempty"`
	Chi_CollectiveName *Chi_CollectiveName `xml:" CollectiveName,omitempty" json:"CollectiveName,omitempty"`
	Chi_ForeName       *Chi_ForeName       `xml:" ForeName,omitempty" json:"ForeName,omitempty"`
	Chi_Identifier     *Chi_Identifier     `xml:" Identifier,omitempty" json:"Identifier,omitempty"`
	Chi_Initials       *Chi_Initials       `xml:" Initials,omitempty" json:"Initials,omitempty"`
	Chi_LastName       *Chi_LastName       `xml:" LastName,omitempty" json:"LastName,omitempty"`
	Chi_Suffix         *Chi_Suffix         `xml:" Suffix,omitempty" json:"Suffix,omitempty"`
}

type Chi_Affiliation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_CollectiveName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Identifier struct {
	Attr_Source string `xml:" Source,attr"  json:",omitempty"`
	Text        string `xml:",chardata" json:",omitempty"`
}

type Chi_Language struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Journal struct {
	Chi_ISOAbbreviation *Chi_ISOAbbreviation `xml:" ISOAbbreviation,omitempty" json:"ISOAbbreviation,omitempty"`
	Chi_ISSN            *Chi_ISSN            `xml:" ISSN,omitempty" json:"ISSN,omitempty"`
	Chi_JournalIssue    *Chi_JournalIssue    `xml:" JournalIssue,omitempty" json:"JournalIssue,omitempty"`
	Chi_Title           *Chi_Title           `xml:" Title,omitempty" json:"Title,omitempty"`
}

type Chi_ISSN struct {
	Attr_IssnType string `xml:" IssnType,attr"  json:",omitempty"`
	Text          string `xml:",chardata" json:",omitempty"`
}

type Chi_JournalIssue struct {
	Attr_CitedMedium string       `xml:" CitedMedium,attr"  json:",omitempty"`
	Chi_Issue        *Chi_Issue   `xml:" Issue,omitempty" json:"Issue,omitempty"`
	Chi_PubDate      *Chi_PubDate `xml:" PubDate,omitempty" json:"PubDate,omitempty"`
	Chi_Volume       *Chi_Volume  `xml:" Volume,omitempty" json:"Volume,omitempty"`
}

type Chi_Volume struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Issue struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_PubDate struct {
	Chi_Day         *Chi_Day         `xml:" Day,omitempty" json:"Day,omitempty"`
	Chi_MedlineDate *Chi_MedlineDate `xml:" MedlineDate,omitempty" json:"MedlineDate,omitempty"`
	Chi_Month       *Chi_Month       `xml:" Month,omitempty" json:"Month,omitempty"`
	Chi_Season      *Chi_Season      `xml:" Season,omitempty" json:"Season,omitempty"`
	Chi_Year        *Chi_Year        `xml:" Year,omitempty" json:"Year,omitempty"`
}

type Chi_MedlineDate struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Season struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Title struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_ISOAbbreviation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_ArticleTitle struct {
	Attr_book string `xml:" book,attr"  json:",omitempty"`
	Attr_part string `xml:" part,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Chi_Abstract struct {
	Chi_AbstractText         []*Chi_AbstractText       `xml:" AbstractText,omitempty" json:"AbstractText,omitempty"`
	Chi_CopyrightInformation *Chi_CopyrightInformation `xml:" CopyrightInformation,omitempty" json:"CopyrightInformation,omitempty"`
}

type Chi_PublicationTypeList struct {
	Chi_PublicationType []*Chi_PublicationType `xml:" PublicationType,omitempty" json:"PublicationType,omitempty"`
}

type Chi_PublicationType struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_VernacularTitle struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_ELocationID struct {
	Attr_EIdType string `xml:" EIdType,attr"  json:",omitempty"`
	Attr_ValidYN string `xml:" ValidYN,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

type Chi_DataBankList struct {
	Attr_CompleteYN string          `xml:" CompleteYN,attr"  json:",omitempty"`
	Chi_DataBank    []*Chi_DataBank `xml:" DataBank,omitempty" json:"DataBank,omitempty"`
}

type Chi_DataBank struct {
	Chi_AccessionNumberList *Chi_AccessionNumberList `xml:" AccessionNumberList,omitempty" json:"AccessionNumberList,omitempty"`
	Chi_DataBankName        *Chi_DataBankName        `xml:" DataBankName,omitempty" json:"DataBankName,omitempty"`
}

type Chi_DataBankName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_AccessionNumberList struct {
	Chi_AccessionNumber []*Chi_AccessionNumber `xml:" AccessionNumber,omitempty" json:"AccessionNumber,omitempty"`
}

type Chi_AccessionNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_ArticleDate struct {
	Attr_DateType string     `xml:" DateType,attr"  json:",omitempty"`
	Chi_Day       *Chi_Day   `xml:" Day,omitempty" json:"Day,omitempty"`
	Chi_Month     *Chi_Month `xml:" Month,omitempty" json:"Month,omitempty"`
	Chi_Year      *Chi_Year  `xml:" Year,omitempty" json:"Year,omitempty"`
}

type Chi_GrantList struct {
	Attr_CompleteYN string       `xml:" CompleteYN,attr"  json:",omitempty"`
	Chi_Grant       []*Chi_Grant `xml:" Grant,omitempty" json:"Grant,omitempty"`
}

type Chi_Grant struct {
	Chi_Acronym *Chi_Acronym `xml:" Acronym,omitempty" json:"Acronym,omitempty"`
	Chi_Agency  *Chi_Agency  `xml:" Agency,omitempty" json:"Agency,omitempty"`
	Chi_Country *Chi_Country `xml:" Country,omitempty" json:"Country,omitempty"`
	Chi_GrantID *Chi_GrantID `xml:" GrantID,omitempty" json:"GrantID,omitempty"`
}

type Chi_GrantID struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Acronym struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Agency struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_MeshHeadingList struct {
	Chi_MeshHeading []*Chi_MeshHeading `xml:" MeshHeading,omitempty" json:"MeshHeading,omitempty"`
}

type Chi_MeshHeading struct {
	Chi_DescriptorName *Chi_DescriptorName  `xml:" DescriptorName,omitempty" json:"DescriptorName,omitempty"`
	Chi_QualifierName  []*Chi_QualifierName `xml:" QualifierName,omitempty" json:"QualifierName,omitempty"`
}

type Chi_DescriptorName struct {
	Attr_MajorTopicYN string `xml:" MajorTopicYN,attr"  json:",omitempty"`
	Attr_Type         string `xml:" Type,attr"  json:",omitempty"`
	Text              string `xml:",chardata" json:",omitempty"`
}

type Chi_QualifierName struct {
	Attr_MajorTopicYN string `xml:" MajorTopicYN,attr"  json:",omitempty"`
	Text              string `xml:",chardata" json:",omitempty"`
}

type Chi_ChemicalList struct {
	Chi_Chemical []*Chi_Chemical `xml:" Chemical,omitempty" json:"Chemical,omitempty"`
}

type Chi_Chemical struct {
	Chi_NameOfSubstance *Chi_NameOfSubstance `xml:" NameOfSubstance,omitempty" json:"NameOfSubstance,omitempty"`
	Chi_RegistryNumber  *Chi_RegistryNumber  `xml:" RegistryNumber,omitempty" json:"RegistryNumber,omitempty"`
}

type Chi_RegistryNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_NameOfSubstance struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_CommentsCorrectionsList struct {
	Chi_CommentsCorrections []*Chi_CommentsCorrections `xml:" CommentsCorrections,omitempty" json:"CommentsCorrections,omitempty"`
}

type Chi_CommentsCorrections struct {
	Attr_RefType  string         `xml:" RefType,attr"  json:",omitempty"`
	Chi_Note      *Chi_Note      `xml:" Note,omitempty" json:"Note,omitempty"`
	Chi_PMID      *Chi_PMID      `xml:" PMID,omitempty" json:"PMID,omitempty"`
	Chi_RefSource *Chi_RefSource `xml:" RefSource,omitempty" json:"RefSource,omitempty"`
}

type Chi_RefSource struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Note struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_OtherID struct {
	Attr_Source string `xml:" Source,attr"  json:",omitempty"`
	Text        string `xml:",chardata" json:",omitempty"`
}

type Chi_GeneralNote struct {
	Attr_Owner string `xml:" Owner,attr"  json:",omitempty"`
	Text       string `xml:",chardata" json:",omitempty"`
}

type Chi_InvestigatorList struct {
	Chi_Investigator []*Chi_Investigator `xml:" Investigator,omitempty" json:"Investigator,omitempty"`
}

type Chi_Investigator struct {
	Attr_ValidYN    string           `xml:" ValidYN,attr"  json:",omitempty"`
	Chi_Affiliation *Chi_Affiliation `xml:" Affiliation,omitempty" json:"Affiliation,omitempty"`
	Chi_ForeName    *Chi_ForeName    `xml:" ForeName,omitempty" json:"ForeName,omitempty"`
	Chi_Initials    *Chi_Initials    `xml:" Initials,omitempty" json:"Initials,omitempty"`
	Chi_LastName    *Chi_LastName    `xml:" LastName,omitempty" json:"LastName,omitempty"`
	Chi_Suffix      *Chi_Suffix      `xml:" Suffix,omitempty" json:"Suffix,omitempty"`
}

type Chi_SpaceFlightMission struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_PubmedData struct {
	Chi_ArticleIdList     *Chi_ArticleIdList     `xml:" ArticleIdList,omitempty" json:"ArticleIdList,omitempty"`
	Chi_History           *Chi_History           `xml:" History,omitempty" json:"History,omitempty"`
	Chi_PublicationStatus *Chi_PublicationStatus `xml:" PublicationStatus,omitempty" json:"PublicationStatus,omitempty"`
}

type Chi_History struct {
	Chi_PubMedPubDate []*Chi_PubMedPubDate `xml:" PubMedPubDate,omitempty" json:"PubMedPubDate,omitempty"`
}

type Chi_PubMedPubDate struct {
	Attr_PubStatus string      `xml:" PubStatus,attr"  json:",omitempty"`
	Chi_Day        *Chi_Day    `xml:" Day,omitempty" json:"Day,omitempty"`
	Chi_Hour       *Chi_Hour   `xml:" Hour,omitempty" json:"Hour,omitempty"`
	Chi_Minute     *Chi_Minute `xml:" Minute,omitempty" json:"Minute,omitempty"`
	Chi_Month      *Chi_Month  `xml:" Month,omitempty" json:"Month,omitempty"`
	Chi_Year       *Chi_Year   `xml:" Year,omitempty" json:"Year,omitempty"`
}

type Chi_Hour struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Minute struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_PublicationStatus struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_ArticleIdList struct {
	Chi_ArticleId []*Chi_ArticleId `xml:" ArticleId,omitempty" json:"ArticleId,omitempty"`
}

type Chi_ArticleId struct {
	Attr_IdType string `xml:" IdType,attr"  json:",omitempty"`
	Text        string `xml:",chardata" json:",omitempty"`
}

type Chi_PubmedBookArticle struct {
	Chi_BookDocument   *Chi_BookDocument   `xml:" BookDocument,omitempty" json:"BookDocument,omitempty"`
	Chi_PubmedBookData *Chi_PubmedBookData `xml:" PubmedBookData,omitempty" json:"PubmedBookData,omitempty"`
}

type Chi_BookDocument struct {
	Chi_Abstract         *Chi_Abstract         `xml:" Abstract,omitempty" json:"Abstract,omitempty"`
	Chi_ArticleIdList    *Chi_ArticleIdList    `xml:" ArticleIdList,omitempty" json:"ArticleIdList,omitempty"`
	Chi_ArticleTitle     *Chi_ArticleTitle     `xml:" ArticleTitle,omitempty" json:"ArticleTitle,omitempty"`
	Chi_AuthorList       []*Chi_AuthorList     `xml:" AuthorList,omitempty" json:"AuthorList,omitempty"`
	Chi_Book             *Chi_Book             `xml:" Book,omitempty" json:"Book,omitempty"`
	Chi_ContributionDate *Chi_ContributionDate `xml:" ContributionDate,omitempty" json:"ContributionDate,omitempty"`
	Chi_DateRevised      *Chi_DateRevised      `xml:" DateRevised,omitempty" json:"DateRevised,omitempty"`
	Chi_ItemList         []*Chi_ItemList       `xml:" ItemList,omitempty" json:"ItemList,omitempty"`
	Chi_KeywordList      *Chi_KeywordList      `xml:" KeywordList,omitempty" json:"KeywordList,omitempty"`
	Chi_Language         []*Chi_Language       `xml:" Language,omitempty" json:"Language,omitempty"`
	Chi_LocationLabel    []*Chi_LocationLabel  `xml:" LocationLabel,omitempty" json:"LocationLabel,omitempty"`
	Chi_PMID             *Chi_PMID             `xml:" PMID,omitempty" json:"PMID,omitempty"`
	Chi_Sections         *Chi_Sections         `xml:" Sections,omitempty" json:"Sections,omitempty"`
}

type Chi_LocationLabel struct {
	Attr_Type string `xml:" Type,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Chi_Sections struct {
	Chi_Section []*Chi_Section `xml:" Section,omitempty" json:"Section,omitempty"`
}

type Chi_Section struct {
	Chi_LocationLabel []*Chi_LocationLabel `xml:" LocationLabel,omitempty" json:"LocationLabel,omitempty"`
	Chi_Section       []*Chi_Section       `xml:" Section,omitempty" json:"Section,omitempty"`
	Chi_SectionTitle  []*Chi_SectionTitle  `xml:" SectionTitle,omitempty" json:"SectionTitle,omitempty"`
}

type Chi_SectionTitle struct {
	Attr_book string `xml:" book,attr"  json:",omitempty"`
	Attr_part string `xml:" part,attr"  json:",omitempty"`
	Attr_sec  string `xml:" sec,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Chi_Book struct {
	Chi_AuthorList      []*Chi_AuthorList    `xml:" AuthorList,omitempty" json:"AuthorList,omitempty"`
	Chi_BeginningDate   *Chi_BeginningDate   `xml:" BeginningDate,omitempty" json:"BeginningDate,omitempty"`
	Chi_BookTitle       *Chi_BookTitle       `xml:" BookTitle,omitempty" json:"BookTitle,omitempty"`
	Chi_CollectionTitle *Chi_CollectionTitle `xml:" CollectionTitle,omitempty" json:"CollectionTitle,omitempty"`
	Chi_Edition         *Chi_Edition         `xml:" Edition,omitempty" json:"Edition,omitempty"`
	Chi_EndingDate      *Chi_EndingDate      `xml:" EndingDate,omitempty" json:"EndingDate,omitempty"`
	Chi_Isbn            []*Chi_Isbn          `xml:" Isbn,omitempty" json:"Isbn,omitempty"`
	Chi_Medium          *Chi_Medium          `xml:" Medium,omitempty" json:"Medium,omitempty"`
	Chi_PubDate         *Chi_PubDate         `xml:" PubDate,omitempty" json:"PubDate,omitempty"`
	Chi_Publisher       *Chi_Publisher       `xml:" Publisher,omitempty" json:"Publisher,omitempty"`
}

type Chi_BookTitle struct {
	Attr_book string `xml:" book,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Chi_BeginningDate struct {
	Chi_Year *Chi_Year `xml:" Year,omitempty" json:"Year,omitempty"`
}

type Chi_Edition struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Publisher struct {
	Chi_PublisherLocation *Chi_PublisherLocation `xml:" PublisherLocation,omitempty" json:"PublisherLocation,omitempty"`
	Chi_PublisherName     *Chi_PublisherName     `xml:" PublisherName,omitempty" json:"PublisherName,omitempty"`
}

type Chi_PublisherName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_PublisherLocation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_EndingDate struct {
	Chi_Year *Chi_Year `xml:" Year,omitempty" json:"Year,omitempty"`
}

type Chi_Medium struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_Isbn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_CollectionTitle struct {
	Attr_book string `xml:" book,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Chi_ContributionDate struct {
	Chi_Day   *Chi_Day   `xml:" Day,omitempty" json:"Day,omitempty"`
	Chi_Month *Chi_Month `xml:" Month,omitempty" json:"Month,omitempty"`
	Chi_Year  *Chi_Year  `xml:" Year,omitempty" json:"Year,omitempty"`
}

type Chi_ItemList struct {
	Attr_ListType string    `xml:" ListType,attr"  json:",omitempty"`
	Chi_Item      *Chi_Item `xml:" Item,omitempty" json:"Item,omitempty"`
}

type Chi_Item struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Chi_PubmedBookData struct {
	Chi_ArticleIdList     *Chi_ArticleIdList     `xml:" ArticleIdList,omitempty" json:"ArticleIdList,omitempty"`
	Chi_History           *Chi_History           `xml:" History,omitempty" json:"History,omitempty"`
	Chi_PublicationStatus *Chi_PublicationStatus `xml:" PublicationStatus,omitempty" json:"PublicationStatus,omitempty"`
}
