package lwdx

// Reference:
// Pandoc:
// https://hackage.haskell.org/package/pandoc-types-1.23/docs/Text-Pandoc-Definition.html#t:Block
// Notion:
// https://www.notion.so/help/writing-and-editing-basics#types-of-content

// CTType specifies a "basic" (common) content tag, divided into
// Block and Inline, and maybe some that are neither. We start
// with a basic set from Pandoc and then add to it as necessary.
type CTType string

const (
	CT_type_ERROR CTType = "ERR" // ERROR

	/* Notion
		Basic blocks
		Text: Just your regular old plain text!
		Page: Adds a sub-page inside your page. Can nest.
		To-do list: Checkboxes for tasks and such. Shortcut [].
		Heading 1: The largest heading, add easily with shortcut /h1.
		Heading 2: The medium  heading, add easily with shortcut /h2.
		Heading 3: The smallest heading, add easily with shortcut /h3.
		Table: Creates a simple set of columns and rows.
		Bulleted list: Bullets. Shortcut - + space.
		Numbered list: Indents your list and auto generates the next number.
		Toggle List:   Toggles to nest/open/hide content. Shortcut > + space.
		Quote: Larger text to break quotes out from the rest of your document.
		Divider: Creates a thin gray line to break up text. Shortcut ---.
		Link to page: A block that links to another page in your workspace.
		Callout: Boxed text for tips, warnings, disclaimers, etc. With emoji!
		Databases
	        Added in-line.
		Table - Inline: Adds a table on your page.
		Board - Inline: Adds a kanban board on your page.
		Gallery - Inline: Adds a gallery on your page.
		List - Inline: Adds a list more minimal than a table to your page.
	*/

	// Pandoc BLOCK elements:

	// ?  Plain [Inline] - Plain text (not para)
	// Y  Para  [Inline] - Paragraph
	CT_Blk_Para = "bPara"
	// ?  LineBlock [[Inline]] - Multiple non-breaking lines
	// Y  CodeBlock Attr Text  - Code block (literal) w attrs
	CT_Blk_CdBlk = "bCode"
	// ?  RawBlock Format Text - Raw block
	// Y  BlockQuote [Block]   - Block quote (list of blocks)
	CT_Blk_Quote = "bQuot"
	// Y  OrderedList ListAttrs [[Block]] -
	//        Attrs + list of items, each a list of blocks
	CT_Blk_OList = "bOList"
	// Y  BulletList [[Block]] - List of items, each a list of blocks
	CT_Blk_UList = "bUList"
	// Y  DefinitionList [([Inline], [[Block]])] -
	//         Each list item is a pair of a term (list of inlines)
	w //         and one or more definitions (each a list of blocks)
	// Y  Header Int Attr [Inline]  Level (integer) and text (inlines)
	CT_Blk_Hedr = "bHedr"
	// Y  HorizontalRule - Horizontal rule
	CT_Blk_HRule = "bHRule"
	// Y  Table Attr Caption [ColSpec] TableHead [TableBody] TableFoot -
	//          Table, with attrs, caption, opt'l short caption,
	//          column alignments & widths (req'd), table head,
	//          table bodies, and table foot
	CT_Blk_Table = "bTable"
	// Y  Figure Attr Caption [Block] -
	//           Figure, with attrs, caption, and content (list of blocks)
	CT_Blk_Figure = "bFigure"
	// Y  Div Attr [Block] - Generic block container with attrs
	CT_Blk_Div = "bDiv"

	// NON-Pandoc BLOCK elements:

	// Pandoc INLINE elements:

	// Y  Str Text           - Text (string)
	CT_Inl_Text = "iText"
	// Y  Emph [Inline]      - Emphasized text (list of inlines)
	CT_Inl_Emph = "iEmph"
	// Y  Underline [Inline] - Underlined text (list of inlines)
	CT_Inl_Undln = "iUndln"
	// Y  Strong [Inline]    - Strongly emphasized text (list of inlines)
	CT_Inl_Strng = "iStrng"
	// Y  Strikeout [Inline] - Strikeout text (list of inlines)
	CT_Inl_Strike = "iStrike"
	// Y  Sprscript [Inline] - Superscripted text (list of inlines)
	CT_Inl_Super = "iSuper"
	// Y  Subscript [Inline] - Subscripted text (list of inlines)
	CT_Inl_Sub = "iSub"
	// ?  SmallCaps [Inline] - Small caps text (list of inlines)
	// Y  Quoted QuoteType [Inline] - Quoted text (list of inlines)
	CT_Inl_Quote = "iQuote"
	// ?  Cite [Citation] [Inline]  - Citation (list of inlines)
	CT_Inl_Citatn = "iCitatn"
	// Y  Code Attr Text - Inline code (literal)
	CT_Inl_Code = "iCode"
	// ?  Space     - Inter-word space
	// ?  SoftBreak - Soft line break
	// Y  LineBreak - Hard line break
	CT_Inl_LnBrk = "iLnBrk" // Not a Block break but a simple line break
	// ?  Math MathType Text    - TeX math (literal)
	// ?  RawInline Format Text - Raw inline
	// Y  Link Attr [Inline] Target -
	//         Hyperlink: alt text (list of inlines), target
	CT_Inl_Link = "iLink"
	// Y  Image Attr [Inline] Target -
	//          Image: alt text (list of inlines), target
	CT_Inl_Image = "iImage"
	// ?  Note [Block] - Footnote or endnote
	// Y  Span Attr [Inline] - Generic inline container with attrs
	CT_Inl_Span = "iSpan"

	// NON-Pandoc INLINE elements:

	CT_Inl_Bold   = "iBold"
	CT_Inl_Italic = "iTalic"

	// NON-Pandoc NEITHER elements:

)

func (CT CTType) LongForm() string {
	switch CT {
	case CT_Inl_LnBrk:
		return "Start-Tag"
	}
	return string(CT)
}

// ListAttrs: (Int, ListNumberStyle, ListNumberDelim)
// ListNumberStyle: DefaultStyle Example Decimal
//     LowerRoman UpperRoman LowerAlpha UpperAlpha
// ListNumberDelim: DefaultDelim Period OneParen TwoParens
