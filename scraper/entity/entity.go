package entity

type JobPosting struct {
	PostURL     string
	Company     string
	ContentText string
	EasyApply   bool
	ApplyURL    string
}

func (j JobPosting) ToCSV() []string {
	row := make([]string, 5)
	row[0] = j.PostURL
	row[1] = j.Company
	row[2] = j.ContentText

	if j.EasyApply {
		row[3] = "1"
	} else {
		row[3] = "0"
	}

	row[4] = j.ApplyURL

	return row
}
