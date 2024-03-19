package note

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type Database struct {
	inMemory map[int]Note
}

func NewDB() *Database {
	db := &Database{}
	db.inMemory = make(map[int]Note)

	return db
}

func (d *Database) AddNote(n Note) {
	d.inMemory[n.Id] = n
}

func (d *Database) RemoveNote(id int) {
	delete(d.inMemory, id)
}

func (d *Database) UpdateNote(n Note) {
	d.inMemory[n.Id] = n
}

func (d *Database) GetNote(id int) (Note, bool) {
	note, ok := d.inMemory[id]
	return note, ok
}

func (d *Database) GetNotes() []Note {
	notes := []Note{}
	for _, note := range d.inMemory {
		notes = append(notes, note)
	}
	return notes
}
