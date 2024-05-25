package main

type DBReader interface {
	Read() error
}
