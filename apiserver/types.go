package main

import (
	"time"
)

type VCode struct {
	PhoneNum string
	VCode    string
	Time     time.Time
}

type Tenant struct {
	Name             string
	Org              string
	PhoneNum         string
	Email            string
	Password         string
	Secret           string
	IsPremiumAccount bool
	PremiumUntil     time.Time
}
