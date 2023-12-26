package models

type AdminJournalCreateReq struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Isbn        string `json:"isbn,omitempty"`
	Author      string `json:"author,omitempty"`
	Email       string `json:"email,omitempty"`
}

type AdminJournalCreateRes struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Isbn        string `json:"isbn,omitempty"`
	Author      string `json:"author,omitempty"`
	Email       string `json:"email,omitempty"`
}

type AdminJournalUpdateReq struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Isbn        string `json:"isbn,omitempty"`
	Author      string `json:"author,omitempty"`
	Email       string `json:"email,omitempty"`
}

type AdminJournalUpdateRes struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Isbn        string `json:"isbn,omitempty"`
	Author      string `json:"author,omitempty"`
	Email       string `json:"email,omitempty"`
}

type JournalUpdateReq struct {
	Id                        string        `json:"id,omitempty"`
	CoverPhoto                string        `json:"cover_photo,omitempty"`
	Title                     string        `json:"title,omitempty"`
	Description               string        `json:"description,omitempty"`
	JournalData               []JournalData `json:"journal_data,omitempty"`
	AcceptanceRate            string        `json:"acceptance_rate,omitempty"`
	SubmissionToFinalDecision string        `json:"submission_to_final_decision,omitempty"`
	AcceptanceToPublication   string        `json:"acceptance_to_publication,omitempty"`
	CitationIndicator         string        `json:"citation_indicator,omitempty"`
	ImpactFactor              string        `json:"impact_factor,omitempty"`
	Subjects                  []Subject     `json:"subjects,omitempty"`
}

type JournalData struct {
	JournalId string `json:"journal_id,omitempty"`
	Text      string `json:"text,omitempty"`
	Type      string `json:"type,omitempty"`
	ShortText string `json:"short_text,omitempty"`
}

type Subject struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
