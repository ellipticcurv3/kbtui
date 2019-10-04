// +build !rm_basic_commands allcommands deletecmd

package main

import (
	"fmt"
	"strconv"
)

func init() {
	command := Command{
		Cmd:         []string{"delete", "del"},
		Description: "$messageId - Delete a message by $messageId",
		Help:        "",
		Exec:        cmdDelete,
	}

	RegisterCommand(command)
}
func cmdDelete(cmd []string) {
	messageID, _ := strconv.Atoi(cmd[1])
	chat := k.NewChat(channel)
	_, err := chat.Delete(messageID)
	if err != nil {
		printToView("Feed", fmt.Sprintf("There was an error deleting your message."))
	}

}
