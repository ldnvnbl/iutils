package main

import (
	"fmt"
	"os"
	"strings"

	uuidpkg "github.com/google/uuid"
)

func main() {
	uuid := uuidpkg.New().String()
	if len(os.Args) > 1 {
		uuid = strings.ToLower(os.Args[1])
	}
	if len(uuid) == 36 && strings.Contains(uuid, "-") {
		fmt.Println(dashUUIDToNormal(uuid))
		fmt.Println(dashUUIDToHex(uuid))
		fmt.Println(uuid)
	} else if len(uuid) == 34 && strings.HasPrefix(uuid, "0x") {
		fmt.Println(hexUUIDToNormal(uuid))
		fmt.Println(hexUUIDToDash(uuid))
		fmt.Println(uuid)
	} else if len(uuid) == 32 {
		fmt.Println(normalToDash(uuid))
		fmt.Println(normalToHex(uuid))
		fmt.Println(uuid)
	}
}

func dashUUIDToNormal(uuid string) string {
	return strings.Replace(uuid, "-", "", -1)
}

func dashUUIDToHex(uuid string) string {
	return "0x" + dashUUIDToNormal(uuid)
}

func hexUUIDToNormal(uuid string) string {
	return uuid[2:]
}

func hexUUIDToDash(uuid string) string {
	normalUUID := hexUUIDToNormal(uuid)
	return fmt.Sprintf("%s-%s-%s-%s-%s", normalUUID[:8], normalUUID[8:12], normalUUID[12:16], normalUUID[16:20], normalUUID[20:])
}

func normalToHex(uuid string) string {
	return "0x" + uuid
}

func normalToDash(uuid string) string {
	normalUUID := uuid
	return fmt.Sprintf("%s-%s-%s-%s-%s", normalUUID[:8], normalUUID[8:12], normalUUID[12:16], normalUUID[16:20], normalUUID[20:])
}
