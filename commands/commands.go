package commands

import (
	"errors"

	"github.com/nursakn/pokedexcli/config"
)

type Command struct {
	name string
	description string
	callback func(*config.Config, []string) error
}

type CommandsMap map[string]Command

func (commandsMap CommandsMap) Send(commandName string, cfg *config.Config, args []string) error {
command, ok := commandsMap[commandName]

	if !ok {
		return errors.New("command not found")
	}

	return command.callback(cfg, args)
}

