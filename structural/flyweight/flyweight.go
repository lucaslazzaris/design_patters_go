package main

type Window struct {
}

type GlyphContext struct {
}

type BTree struct {

}

type Font struct {
	index int
	fonts BTree
}

type Glyph interface {
	draw(*Window, GlyphContext)

	setFont(*Font, GlyphContext)
	getFont(GlyphContext) *Font

	first(GlyphContext)
	next(GlyphContext)
	isDone(GlyphContext) bool
	current(GlyphContext) *Glyph

	insert(*Glyph, GlyphContext)
	remove(GlyphContext)
}

// Should implement Glyph interface
type Character struct {
	charcode string
}

type Row struct {
}

type Column struct {	
}

type GlyphFactory struct {
	characters map[string]Character
}

func (factory *GlyphFactory) createCharacter(char string) *Character {
	character, ok := factory.characters[char]
	if(!ok){
		factory.characters[char] = Character { charcode: char }
		character = factory.characters[char]
	}
	return &character
}

func (factory *GlyphFactory) createRow(char string) *Row {
	return new(Row)
}

func (factory *GlyphFactory) createColumn(char string) *Column {
	return new(Column)
}