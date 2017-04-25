# lwdx
Lightweight document transformations. 
* LwDITA-based structured modular authoring in XML/XHTML/Markdown 
* Roundtrip transformations among all three for authoring freedom 
* Tailored to correctly handle XML mixed content 
* Avoids the horrors of Golang XML annotations 
## Contains 
* CLI commands
* API
## Features
* Relies on Golang encoder/xml package for parsing & tokenization
* Avoids as much XML & DITA complexity as possible
* Can write out parse trees generated by [`xmlx`](https://github.com/jteeuwen/go-pkg-xmlx), [`etree`](https://github.com/beevik/etree), [`mxj`](https://github.com/clbanning/mxj), [`x2j`](https://github.com/clbanning/mxj/tree/master/x2j), [Go `html`](https://godoc.org/golang.org/x/net/html)
## Licensing
* LGPL 3.0 
* [Discussion](https://www.whitesourcesoftware.com/whitesource-blog/top-10-gpl-questions-answered/): <i>The LGPL and GPL licenses differ with one major exception; with LGPL the requirement that you open up the source code to your own extensions to the software is removed. You are only obliged to subject your modifications to the original free library to the LGPL. Since the free library is always subject to the LGPL, it must be possible for any user of your software to modify, recompile or replace the free LGPL library and use its modified version with your software.</i>
* (Future) Dual licensing to provide an option for unencumbered commercial use 
## Missing
* Tests 
* XML entity substitutions (but didn't you get the memo? you're sposta use DITA `conref`)
* Schema-driven WYSIWYG IBE (in-browser editor) (but this is roadmapped) 
