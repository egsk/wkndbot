package cnf

import (
	"encoding/json"
	"io/ioutil"
)

type Credentials struct {
	Token string
}

type Users struct {
	Allowed []string
}

type Config struct {
	Credentials Credentials
	Users       Users
	Commands    map[string]Command
}

type Command struct {
	Provider string
	Command []string
}

var config Config
var initStatus bool

func loadCredentials() Credentials {
	var credentials Credentials
	file, err := ioutil.ReadFile("_cnf/credentials.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &credentials)
	if err != nil {
		panic(err)
	}

	return credentials
}

func loadUsers() Users {
	var users Users
	file, err := ioutil.ReadFile("_cnf/users.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &users)
	if err != nil {
		panic(err)
	}

	return users
}

func loadCommands() map[string]Command {
	var commands map[string]Command
	file, err := ioutil.ReadFile("_cnf/commands.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &commands)
	if err != nil {
		panic(err)
	}

	return commands
}

func Get() Config {
	if !initStatus {
		config.Credentials = loadCredentials()
		config.Users = loadUsers()
		config.Commands = loadCommands()
		initStatus = true
	}

	return config
}
