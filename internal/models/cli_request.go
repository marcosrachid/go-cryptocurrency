package models

type CLIRequest struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}
