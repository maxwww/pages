package telegram

type Responses struct {
	Start          string `mapstructure:"start"`
	UnknownCommand string `mapstructure:"unknown_command"`
	Description    string `mapstructure:"description"`
}

type Errors struct {
	Default string `mapstructure:"default"`
}
