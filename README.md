# lwdx
Lightweight document transformations.
* Accepts and processes content authored in any of multiple formats (XML/XHTML/Markdown)
* Documents comprise _maps_ that transclude modular, structured _topics_ (this is from DITA) 
* Implementation is based on emerging [OASIS standard](https://github.com/oasis-open/dita-lightweight) (but compliance is very much _not_ guaranteed) 
* Roundtrip conversions among all three formats for authoring+publishing freedom+flexibility 
* Straightforward handling of XML mixed content... 
* Avoiding the horrors of Golang's record-oriented XML annotations !
## Contains 
* API
* CLI commands 
## Features
* Relies on Golang's `encoder/xml` package for parsing & tokenization
* Hides as much XML & DITA complexity as possible (It's the NJ Style: _It is slightly better to be simple than correct._)
* Formats input files opinionatedly (like `gofmt`), which simplifies diff'ing for changes under version control 
* Can write out parse trees generated by [`xmlx`](https://github.com/jteeuwen/go-pkg-xmlx), [`etree`](https://github.com/beevik/etree), [`mxj`](https://github.com/clbanning/mxj), [`x2j`](https://github.com/clbanning/mxj/tree/master/x2j), [Go `html`](https://godoc.org/golang.org/x/net/html)
## License
* [__LGPL 3.0__](https://www.gnu.org/licenses/lgpl-3.0.en.html)
* [Discussion (#1)](https://www.whitesourcesoftware.com/whitesource-blog/top-10-gpl-questions-answered/): <i><small>The LGPL and GPL licenses differ with one major exception; with LGPL the requirement that you open up the source code to your own extensions to the software is removed. You are only obliged to subject your modifications to the original free library to the LGPL. Since the free library is always subject to the LGPL, it must be possible for any user of your software to modify, recompile or replace the free LGPL library and use its modified version with your software.</small></i>
* [Discussion (#2)](https://en.wikipedia.org/wiki/GNU_Lesser_General_Public_License): <i><small>The main difference between the GPL and the LGPL is that the latter allows the work to be linked with (in the case of a library, "used by") a  non-(L)GPLed program, regardless of whether it is free software or proprietary software. The non-(L)GPLed program can then be distributed under any terms if it is not a derivative work.</small></i>
* [Discussion (#3)](https://www.gnu.org/licenses/why-not-lgpl.html): <i><small>Using the ordinary GPL for a library gives free software developers an advantage over proprietary developers: a library that they can use, while proprietary developers cannot use it. Using the ordinary GPL is not advantageous for every library. There are reasons that can make it better to use the Lesser GPL in certain cases. The most common case is when a free library's features are readily available for proprietary software through other libraries. In that case, the library cannot give free software any particular advantage, so it is better to use the Lesser GPL for that library.</small></i>
* (Future:_TBS_) Dual licensing to provide an option for unencumbered commercial use 
## Missing
* Tests =:-O 
* XML DTD-style entity substitutions (but didn't you get the memo? yer sposta use DITA `conref`)
* Support for `ditaval` files for conditional publishing 
* Details of comment handling in Markdown, and how it can support roundtripping for complex LwDITA tags (note that LwDITA relies on CommonMark with extensions) 
* Speczn
## TODOs & Roadmap
* Implement character entities  
* Publishing! (this is the first step to usefulness, and will be done by using the DITA-OT's preprocessing stage as a guide) 
* A schema-driven WYSIWYG IBE (in-browser editor, probably oughta be based on [CodeMirror](https://codemirror.net/demo/xmlcomplete.html)) 
* Some level of integration with _DITA for Small Teams_ [D4ST](http://www.dita-for-small-teams.org/)
* Enhanced support for common DITA mechanisms like conref files and glossary files 
## Contributing
* Contact me at fbaube at welho daaht com 
