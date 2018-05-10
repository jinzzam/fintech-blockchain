package command

import (
	"blockchain/log"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// CommandSlice is a type of slice of command.
type CommandSlice []Command

// FlagSlice is a type of slice of flags.
type FlagSlice []Flag

// Handler is a type to handle an executable function.
type Handler interface{}

// InputType is a type identifying command or flag.
type InputType int

const (
	// CommandType denotes that a string indicates a command.
	CommandType InputType = iota
	// FlagType denotes that a string indicates a flag.
	FlagType
)

// Command describes a console command. Name is a keyword on console
// to execute this command. Description will be printed on Usage.
// Commands contains sevral sub-commands. Run is a executable function
// invoked by calling this command.
type Command struct {
	Name          string
	ShortName     string
	Description   string
	Commands      CommandSlice
	Flags         FlagSlice
	DefaultParams []interface{}
	Run           Handler
}

var globalCommands CommandSlice

// CheckInputType determines whether the type of input parameter is
// a command or a flag.
func CheckInputType(in string) InputType {
	if strings.HasPrefix(in, "-") {
		return FlagType
	}

	return CommandType
}

// Usage print help messages of a command.
func (cc CommandSlice) Usage() {
	fmt.Printf("%s", cc.recursiveUsage(0))
}

func (cc CommandSlice) recursiveUsage(depth int) string {
	res := bytes.NewBuffer([]byte{})

	for _, c := range cc {
		fmt.Fprintf(res, "%s", c.recursiveUsage(depth))
	}

	return res.String()
}

// StartCommandWithFlags executes handler function of Usage() of a command.
func (cc CommandSlice) StartCommandWithFlags(cmd string) error {
	log.Debug("StartComrmandWithFlags: " + cmd)
	ss := strings.Split(cmd, " ")
	if len(ss) <= 0 {
		return nil
	}

	curr := cc.FindCommand(ss[0])

	if curr == nil {
		log.Error(ss[0] + " command cannot be found.")
		cc.Usage()
		return errors.New(ss[0] + " command cannot be found.")
	}

	log.Debug("Found command: " + curr.Name)

	for i := 1; i < len(ss); i++ {
		switch CheckInputType(ss[i]) {
		case CommandType:
			prev := curr
			log.Debug(curr)
			curr = curr.FindSubCommand(ss[i])
			log.Debug(curr)

			if curr == nil {
				if prev.Run == nil {
					prev.Usage()
				}

				if err := prev.RunHandler(ss[i:]...); err != nil {
					log.Error(err)

					return err
				}
				return nil
			}
		case FlagType:
			if ss[i] == "--help" {
				log.Debug("StartCommandWithFlags: " + curr.Name + " help")
				curr.Usage()
				return errors.New("The flag cannot be found")
			}
			if len(ss) < i+2 {
				log.Error("StartCommandWithFlags: Invalid command coposition")
				return errors.New("Invalid flags")
			}
			_ = curr.SetFlagValue(ss[i], ss[i+1])
			i++
		}
	}

	if curr.Run == nil {
		curr.Usage()
	}

	if err := curr.RunHandler(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// FindFullCommand finds a command for composited string.
func (cc CommandSlice) FindFullCommand(cmd string) *Command {
	log.Debug("FindFullCommand: " + cmd)
	ss := strings.Split(cmd, " ")

	if len(ss) <= 0 {
		return nil
	}

	curr := cc.FindCommand(ss[0])

	if curr == nil {
		log.Error("FindFullCommand: Invalid the first command.")
		return nil
	}

	for _, s := range ss[1:] {
		curr = curr.FindSubCommand(s)
		if curr == nil {
			log.Error("FindFullCommand: cannot find a sub command.")
			return nil
		}
	}

	log.Debug("FindFullCommand(Found): " + curr.Name)
	return curr
}

// FindCommand finds a command structure for a command string
func (cc CommandSlice) FindCommand(cmd string) *Command {
	for i, s := range cc {
		if s.Name == cmd || s.ShortName == cmd {
			return &cc[i]
		}
	}

	return nil
}

// FindSubCommand finds a sub-command from a command structure
// for a command string.
func (c *Command) FindSubCommand(cmd string) *Command {
	return c.Commands.FindCommand(cmd)
}

// AddSubCommand adds a sub-command to a parent command.
func (c *Command) AddSubCommand(cmd Command) error {
	log.Debug("AddSubCommand: " + cmd.Name + " to " + c.Name)
	if c.Commands == nil {
		c.Commands = make([]Command, 0)
	}
	if c.Commands.FindCommand(cmd.Name) != nil {
		log.Error("AddSubCommand: already exist.")
		return errors.New(cmd.Name + " command already exist.")
	}

	c.Commands = append(c.Commands, cmd)
	log.Debug("AddSubCommand(complted): " + cmd.Name + " to " + c.Name)
	return nil
}

// AddFlag adds a flag to a parent command.
func (c *Command) AddFlag(flag Flag) error {
	if flag.LongName == "" {
		return errors.New("a longname of the flag is mandotory")
	}

	if c.Flags.FindFlag(flag.LongName) != nil {
		return errors.New(flag.LongName + " flag already exist.")
	}

	if flag.Name != "" && c.Flags.FindFlag(flag.Name) != nil {
		return errors.New(flag.Name + " flag already exist.")
	}

	c.Flags = append(c.Flags, flag)

	return nil
}

// SetFlagValue assigns a value to a flag.
func (c *Command) SetFlagValue(flag string, val string) error {
	log.Debug("SetFlagValue: " + flag)
	f := c.Flags.FindFlag(flag)

	if f == nil {
		log.Error("SetFlagValue: Invalid flag")
		return errors.New("Invalid flag")
	}

	f.Value = val

	return nil
}

// Usage print the detail help messages for a command.
func (c Command) Usage() {
	fmt.Println(c.Name + " Usage")
	fmt.Println("-----------------------------------------------")
	if c.Commands != nil && len(c.Commands) > 0 {
		fmt.Println("Commands : ")
		for _, c := range c.Commands {
			fmt.Println("\t" + c.Name + "(" + c.ShortName + ")" + " " + c.Description)
		}
	}
	if c.Flags == nil || len(c.Flags) > 0 {
		fmt.Println("")
		fmt.Println("Flags : ")
		for _, c := range c.Flags {
			fmt.Println("\t" + c.Name + " " + c.Description + "(" + c.Value + ")")
		}
	}

}

func (c Command) recursiveUsage(depth int) string {
	res := bytes.NewBuffer([]byte{})

	line := ""
	for i := 0; i < depth; i++ {
		line += "\t"
	}

	if c.ShortName != "" {
		line += c.Name + "(" + c.ShortName + ")" + " " + c.Description
	} else {
		line += c.Name + " " + c.Description
	}

	fmt.Fprintf(res, "%s\n", line)
	if c.Commands != nil && len(c.Commands) > 0 {
		fmt.Fprintf(res, "%s", c.Commands.recursiveUsage(depth+1))
	}

	return res.String()
}

// RunHandler execute a registered handler.
func (c *Command) RunHandler(args ...string) error {
	if c.Run == nil {
		return errors.New("No handler")
	}

	typ := reflect.TypeOf(c.Run)

	if typ.Kind() != reflect.Func || typ.NumOut() != 1 {
		return errors.New("Invalid handler")
	}
	tnum := typ.NumIn()
	anum := len(args)
	dnum := len(c.DefaultParams)

	if tnum > anum+dnum {
		return fmt.Errorf("Incorrect arguments: %v arguments should be passed and %v default parameters have been set", tnum, dnum)
	}

	in := make([]reflect.Value, 0)

	for i := 0; i < anum; i++ {
		t := typ.In(i)

		switch tn := t.Kind(); tn {
		case reflect.Int:
			log.Debug(tn)
			v, err := strconv.Atoi(args[i])
			if err != nil {
				return err
			}
			in = append(in, reflect.ValueOf(v))
		case reflect.String:
			log.Debug(tn)
			in = append(in, reflect.ValueOf(args[i]))
		case reflect.Uint64:
			log.Debug(tn)
			v, err := strconv.ParseUint(args[i], 10, 64)
			if err != nil {
				return err
			}
			in = append(in, reflect.ValueOf(v))
		case reflect.Int64:
			log.Debug(tn)
			v, err := strconv.ParseInt(args[i], 10, 64)
			if err != nil {
				return err
			}
			in = append(in, reflect.ValueOf(v))
		case reflect.Slice:
			switch ts := t.Elem().Kind(); ts {
			case reflect.Uint8:
				log.Debug(tn)
				v, err := hex.DecodeString(args[i])
				if err != nil {
					return err
				}
				in = append(in, reflect.ValueOf(v))
			default:
				log.Error(ts)
				in = append(in, reflect.ValueOf(args[i]))
			}
		default:
			log.Error(tn)
			in = append(in, reflect.ValueOf(args[i]))
		}
	}

	if rem := tnum - anum; rem != 0 {
		for i := dnum - rem; i < dnum; i++ {
			in = append(in, reflect.ValueOf(c.DefaultParams[i]))
		}
	}

	var err error
	var res []reflect.Value

	f := reflect.ValueOf(c.Run)
	ch := make(chan int)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				v, _ := e.(string)
				err = errors.New(v)
			}
			ch <- 1
		}()

		res = f.Call(in)
	}()
	<-ch

	if err != nil {
		return err
	}

	val, ok := res[0].Interface().(error)

	if val != nil && !ok {
		switch v := val.(type) {
		default:
			log.Error(v)
		}

		return errors.New("It is unexpected return type of the handler")
	}

	return val
}
