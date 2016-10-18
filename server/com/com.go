package com

import "fmt"

//MewErr
func NewErr(desc string, err error) error {
	if err == nil {
		return fmt.Errorf("%s", desc)
	}
	return fmt.Errorf("%s,%s", desc, err.Error())
}
