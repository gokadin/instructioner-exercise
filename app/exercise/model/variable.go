package model

type Variable struct {
    Name string `json:"name"`
    Type string `json:"type"`
    RangeStart string `json:"rangeStart"`
    RangeEnd string `json:"rangeEnd"`
    Default string `json:"default"`
}
