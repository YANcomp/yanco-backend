package domain

type Banner struct {
	ID          int64  `json:"ID" db:"id"`
	Header      string `json:"header,omitempty" db:"header"`
	Description string `json:"description,omitempty" db:"description"`
	TargetType  string `json:"targetType,omitempty" db:"targettype"`
	Target      string `json:"target,omitempty" db:"target"`
	Image       string `json:"image,omitempty" db:"image"`
	ImageMobile string `json:"imageMobile,omitempty" db:"imagemobile"`
	IsCatalog   bool   `json:"isCatalog,omitempty" db:"iscatalog"`
}
