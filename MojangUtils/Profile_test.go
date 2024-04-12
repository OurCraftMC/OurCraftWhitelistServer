package MojangUtils

import (
	"fmt"
	"testing"
)

func Test_GetProfileByUserName(t *testing.T) {
	uuid, err := GetProfileByUserName("asfiusdfhsafgaogfaos8df")
	fmt.Println(uuid)
	if err != nil {
		fmt.Println(err)
	}
	uuid, err = GetProfileByUserName("Notch")
	fmt.Println(uuid)
	if err != nil {
		fmt.Println(err)
	}
}
