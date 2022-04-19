package model

import "fmt"

func CreateTodo(name, todo string) error {
	fmt.Println("hello i'm here")
	insertQ, err := con.Query("INSERT INTO TODO VALUES (?, ?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO WHERE NAME=?", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
