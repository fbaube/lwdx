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
* Can make parse trees generated [xmlx](https://github.com/jteeuwen/go-pkg-xmlx), [etree](https://github.com/beevik/etree), [mxj](https://github.com/clbanning/mxj), [x2j](https://github.com/clbanning/mxj/tree/master/x2j), [Go html](https://godoc.org/golang.org/x/net/html)
## Licensing
* LGPL 3.0 
* Dual licensing for unencumbered commercial use in the future (maybe)
## Lacks
* Tests 
* XML entity substitutions (but you were sposta use DITA conref's)
* Schema-driven WYSIWYG IBE (in-browser editor) (but this is roadmapped) 
