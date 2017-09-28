# lwdx
Lightweight document transformations & assemblies, based on familiar markups and simple tools.
* Accepts and processes content authored in three interoperable input formats: XDITA _xml_, HDITA _html5_, MDITA _markdown_ (Note: this interop is a fundamental premise of LwDITA) 
* Documents are assembled from _maps_ that transclude modular, structured _topics_ (This feature is straight from DITA 1.X) 
* This implementation is based on the emerging [OASIS standard](https://github.com/oasis-open/dita-lightweight) (but full compliance is _not_ guaranteed) 
* The goal is roundtrip conversions among all three formats for freedom+flexibility in authoring+publishing  
* We desire straightforward handling of XML mixed content... 
* So `lwdx` chooses a DIY approach over the horrors of Golang's [XML annotations](https://godoc.org/encoding/xml#Marshal). 
## Contains 
* Golang API
* CLI commands 
## Features
* Relies on Golang's [`encoder/xml`](https://godoc.org/encoding/xml) package for parsing & tokenization
* Hides as much XML & DITA complexity as possible (Note that this is the Bell Labs UNIX "NJ Style": _It is slightly better to be simple than correct._)
* Formats input files opinionatedly (like `gofmt`), which simplifies diff'ing for changes under version control 
* For purposes of comparison and exploration, can write out parse trees generated by [`xmlx`](https://github.com/jteeuwen/go-pkg-xmlx), [`etree`](https://github.com/beevik/etree), [`mxj`](https://github.com/clbanning/mxj), [`x2j`](https://github.com/clbanning/mxj/tree/master/x2j), [Go `html`](https://godoc.org/golang.org/x/net/html)
## License
* [__LGPL 3.0__](https://www.gnu.org/licenses/lgpl-3.0.en.html)
* [Discussion (#1)](https://www.whitesourcesoftware.com/whitesource-blog/top-10-gpl-questions-answered/): <i><small>The LGPL and GPL licenses differ with one major exception; with LGPL the requirement that you open up the source code to your own extensions to the software is removed. You are only obliged to subject your modifications to the original free library to the LGPL. Since the free library is always subject to the LGPL, it must be possible for any user of your software to modify, recompile or replace the free LGPL library and use its modified version with your software.</small></i>
* [Discussion (#2)](https://en.wikipedia.org/wiki/GNU_Lesser_General_Public_License): <i><small>The main difference between the GPL and the LGPL is that the latter allows the work to be linked with (in the case of a library, "used by") a  non-(L)GPLed program, regardless of whether it is free software or proprietary software. The non-(L)GPLed program can then be distributed under any terms if it is not a derivative work.</small></i>
* (Future:_TBS_) Dual licensing to provide an option for unencumbered commercial use 
## Missing
* Tests =:-O 
* XML DTD-style entity substitutions (but didn't you get the memo? yer sposta use DITA `conref`)
* Details of comment handling in Markdown, and how it can support roundtripping for complex LwDITA tags (note that LwDITA relies on CommonMark with extensions) 
## TODOs & Roadmap
* Implement character entities, partly to enhance the input parser comparisons 
* Publishing! (this is the first step to usefulness, and will be done along the lines of the DITA-OT's preprocessing stage) 
* A schema-driven WYSIWYG IBE (in-browser editor, which probably oughta be based on [CodeMirror](https://codemirror.net/demo/xmlcomplete.html)) 
* Some level of integration with _DITA for Small Teams_ [_D4ST_](http://www.dita-for-small-teams.org/)
* Enhanced support for common DITA mechanisms like conref files and glossary files 
* Support for LwDITA-style specialization-by-example 
* _Not_ roadmapped: Support for `ditaval` files for conditional publishing 
## Contributing
* Contact me at fbaube at welho daaht com 
