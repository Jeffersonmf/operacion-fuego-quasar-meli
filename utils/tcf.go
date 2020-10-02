/*
Support to the Try Catch Finally blocks. THIS IS EXPERIMENTAL CODE.
Author: Jefferson Marchetti Ferreira
*/
package utils

import "fmt"

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func TestTryCatchFinally() {

	Block{
		Try: func() {
			Throw("Generic Exception")
		},
		Catch: func(e Exception) {
			fmt.Println("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finally...")
		},
	}.Do()
}
