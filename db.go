package main

import "sync"

type DB struct {
	names []name
	emails []email
	ssns []ssn
	billings []billing
	dobs []dob
	mux    sync.Mutex
}

func (d *DB) Save(n name, e email, d dob, s ssn, b billing) {
	d.mux.Lock()
	d.names = append(d.names, n)
	d.emails = append(d.emails, e)
	d.dobs = append(d.dobs, d)
	d.ssns = append(d.ssns, s)
	d.billings = append(d.billings, b)
	d.mux.Unlock()
}

func (d *DB) List() []email {
	return d.emails
}
